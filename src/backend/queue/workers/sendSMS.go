package workers

import (
    "encoding/json"
    "fmt"

    "backend/queue"
    "github.com/astaxie/beego/logs"
    "github.com/gadelkareem/que"
)

const SendSMSType = "sendSMS"

type smsSender interface {
    Send(number string, msg string) error
}

type sendSMS struct {
    s smsSender
}

type SendSMSReq struct {
    MobileNumber, Message string
}

func NewSendSMS(s smsSender) queue.Worker {
    return &sendSMS{s: s}
}

func (w sendSMS) Type() string {
    return SendSMSType
}

func (w sendSMS) Run(j *que.Job) error {
    var r SendSMSReq
    err := json.Unmarshal(j.Args, &r)
    if err != nil {
        return fmt.Errorf("Unable to unmarshal job arguments into request: %s, err: %+v ", j.Args, err)
    }

    err = w.s.Send(r.MobileNumber, r.Message)
    if err != nil {
        logs.Error("Worker %s failed! Error: %s", w.Type(), err)
    }
    return err
}
