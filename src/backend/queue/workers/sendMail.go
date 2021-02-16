package workers

import (
    "encoding/json"
    "fmt"

    "backend/queue"
    "github.com/astaxie/beego/logs"
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
    err := json.Unmarshal(j.Args, &r)
    if err != nil {
        return fmt.Errorf("Unable to unmarshal job arguments into request: %s, err: %+v ", j.Args, err)
    }

    err = w.s.Send(r.RecipientEmail, r.Subject, r.HTML, r.Text)
    if err != nil {
        logs.Error("Worker %s failed! Error: %s", w.Type(), err)
    }
    return err
}
