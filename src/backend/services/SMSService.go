package services

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"

    "backend/kernel"
    "backend/queue"
    "backend/queue/workers"
    "github.com/astaxie/beego/context"
)

type (
    SMSService struct {
        c                                HttpClient
        accountSID, authToken, ownNumber string
        q                                *queue.QueManager
    }
    HttpClient interface {
        Do(req *http.Request) (*http.Response, error)
    }
)

func NewSMSService(c HttpClient, q *queue.QueManager) *SMSService {
    return &SMSService{
        c:          c,
        q:          q,
        accountSID: kernel.App.Config.String("sms::accountSID"),
        authToken:  kernel.App.Config.String("sms::authToken"),
        ownNumber:  kernel.App.Config.String("sms::ownNumber"),
    }
}

func (s *SMSService) Enqueue(number string, msg string) (err error) {
    if s.q != nil {
        return s.enqueue(number, msg)
    }
    return s.Send(number, msg)
}

func (s *SMSService) Send(number string, msg string) (err error) {

    d := url.Values{}
    d.Set("To", number)
    d.Set("From", s.ownNumber)
    d.Set("Body", msg)

    r, _ := http.NewRequest("POST", fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", s.accountSID), strings.NewReader(d.Encode()))
    r.SetBasicAuth(s.accountSID, s.authToken)
    r.Header.Add("Accept", context.ApplicationJSON)
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := s.c.Do(r)
    if resp.StatusCode > 300 {
        var b []byte
        b, err = ioutil.ReadAll(resp.Body)
        if err != nil {
            return
        }
        err = fmt.Errorf("twilio error[%d]: %s", resp.StatusCode, string(b))
    }
    return
}

func (s *SMSService) enqueue(number string, msg string) error {
    return s.q.Enqueue(workers.SendSMSType, workers.SendSMSReq{
        MobileNumber: number,
        Message:      msg,
    })
}
