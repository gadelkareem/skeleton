package workers

import (
    "encoding/json"
    "fmt"

    "backend/queue"
    "github.com/gadelkareem/que"
)

const SendMailType = "sendMail"

type mailSender interface {
    Send(recipientEmail, subject, htmlBody, txtBody string) error
}

type sendMail struct {
    s mailSender
}

type SendMailReq struct {
    RecipientEmail, Subject, HTML, Text string
}

func NewSendMail(s mailSender) queue.Worker {
    return &sendMail{s: s}
}

func (w sendMail) Type() string {
    return SendMailType
}

func (w sendMail) Run(j *que.Job) error {
    var r SendMailReq
    if err := json.Unmarshal(j.Args, &r); err != nil {
        return fmt.Errorf("Unable to unmarshal job arguments into request: %s, err: %+v ", j.Args, err)
    }

    return w.s.Send(r.RecipientEmail, r.Subject, r.HTML, r.Text)
}
