package workers

import (
    "encoding/json"
    "fmt"

    "backend/queue"
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
    if err := json.Unmarshal(j.Args, &r); err != nil {
        return fmt.Errorf("Unable to unmarshal job arguments into request: %s, err: %+v ", j.Args, err)
    }

    return w.s.Send(r.MobileNumber, r.Message)
}
