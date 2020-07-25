package tests

import (
    "bytes"
    "encoding/json"
    "errors"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "testing"
    "time"

    "backend/di"
    "backend/kernel"
    "backend/services"
    h "github.com/gadelkareem/go-helpers"
    "github.com/go-gomail/gomail"
    "github.com/stretchr/testify/assert"
)

var (
    mailbox = flag.Bool("mailbox", false, "Use mailhog/mailtrap instead of SMTP mock.")
)

func initEmailService(c *di.Container) {
    if *mailbox {
        c.EmailService = services.NewEmailService(kernel.SMTPDialer(), nil)
    } else {
        c.EmailService = services.NewEmailService(nil, newMockSenderCloser())
    }
}

type mockSender gomail.SendFunc

func (s mockSender) Send(from string, to []string, msg io.WriterTo) error {
    return s(from, to, msg)
}

type mockSendCloser struct {
    mockSender
    close func() error
}

func (s *mockSendCloser) Close() error {
    return s.close()
}

func newMockSenderCloser() *mockSendCloser {
    return &mockSendCloser{
        mockSender: func(from string, to []string, msg io.WriterTo) error { return nil },
        close:      func() error { return nil },
    }
}

func ExpectEmail(t *testing.T, wantFrom string, wantTo []string, wantSubject, wantBody string) {
    if wantFrom == "" {
        wantFrom = kernel.App.Config.String("smtp::senderEmail")
    }
    if *mailbox {
        return
    }
    s := &mockSendCloser{
        mockSender: stubSend(t, wantFrom, wantTo, wantSubject, wantBody),
        close: func() error {
            t.Error("Close() should not be called in Send()")
            return nil
        },
    }
    C.EmailService.SenderCloser = s
}

func ResetEmailService() {
    if !*mailbox {
        C.EmailService.SenderCloser = newMockSenderCloser()
    }
}

func CheckEmailRetry(t *testing.T, subject string) {
    if !*mailbox {
        return
    }
    rSubject := ""
    h.Retry(func() error {
        rSubject = CheckEmail(t)
        if subject != rSubject {
            time.Sleep(1 * time.Second)
            return errors.New("retry")
        }
        return nil
    }, 3)
    assert.Equal(t, subject, rSubject)
}

// check mailbox works with both mailhog and mailtrap services
func CheckEmail(t *testing.T) string {
    // check if email is sent
    mailServer := kernel.App.ConfigOrEnvVar("smtp::server", "MAIL_HOST")
    u := fmt.Sprintf("http://%s:8025/api/v1/messages", mailServer)
    if strings.Contains(mailServer, "mailtrap") {
        time.Sleep(2 * time.Second)
        u = fmt.Sprintf("https://mailtrap.io/api/v1/inboxes/%d/messages",
            kernel.App.Config.DefaultInt("smtp::inboxId", 0))
    }
    rq, err := http.NewRequest(http.MethodGet, u, nil)
    FailOnErr(t, err)
    rq.Header.Set("API-Token", kernel.App.Config.String("smtp::apiToken"))
    rs, err := http.DefaultClient.Do(rq)
    FailOnErr(t, err)
    defer rs.Body.Close()
    bd, err := ioutil.ReadAll(rs.Body)
    FailOnErr(t, err)
    var rbd []struct {
        Subject string
        Content struct {
            Headers map[string][]string
        }
    }
    err = json.Unmarshal(bd, &rbd)
    FailOnErr(t, err)

    s := rbd[0].Subject
    if s == "" {
        s = rbd[0].Content.Headers["Subject"][0]
    }
    return s
}

var subjRegex = regexp.MustCompile(`(?m)^Subject: (.+)\r\n`)

func stubSend(t *testing.T, wantFrom string, wantTo []string, wantSubject, wantBody string) mockSender {
    return func(from string, to []string, msg io.WriterTo) error {

        assert.Equal(t, from, wantFrom)
        assert.ElementsMatch(t, to, wantTo)

        buf := new(bytes.Buffer)
        _, err := msg.WriteTo(buf)
        if err != nil {
            t.Fatal(err)
        }
        b := buf.String()

        if wantSubject != "" {
            subject := subjRegex.FindStringSubmatch(b)
            assert.Equal(t, wantSubject, subject[1])

        }
        if wantBody != "" {
            compareBodies(t, b, wantBody)
        }
        return nil
    }
}

func compareBodies(t *testing.T, got, want string) {
    // We cannot do a simple comparison since the ordering of headers' fields
    // is random.
    gotLines := strings.Split(got, "\r\n")
    wantLines := strings.Split(want, "\r\n")

    // We only test for too many lines, missing lines are tested after
    if len(gotLines) > len(wantLines) {
        t.Fatalf("Message has too many lines, \ngot %d:\n%s\nwant %d:\n%s", len(gotLines), got, len(wantLines), want)
    }

    isInHeader := true
    headerStart := 0
    for i, line := range wantLines {
        if line == gotLines[i] {
            if line == "" {
                isInHeader = false
            } else if !isInHeader && len(line) > 2 && line[:2] == "--" {
                isInHeader = true
                headerStart = i + 1
            }
            continue
        }

        if !isInHeader {
            missingLine(t, line, got, want)
        }

        isMissing := true
        for j := headerStart; j < len(gotLines); j++ {
            if gotLines[j] == "" {
                break
            }
            if gotLines[j] == line {
                isMissing = false
                break
            }
        }
        if isMissing {
            missingLine(t, line, got, want)
        }
    }
}

func missingLine(t *testing.T, line, got, want string) {
    t.Fatalf("Missing line %q\ngot:\n%s\nwant:\n%s", line, got, want)
}
