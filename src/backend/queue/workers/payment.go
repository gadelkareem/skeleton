package workers

import (
	"encoding/json"
	"fmt"

	"backend/queue"
	"github.com/gadelkareem/que"
)

type paymentCommand int

const (
	PaymentType                                = "payment"
	CommandCustomerCreateUpdate paymentCommand = iota // 0
	CommandCustomer                                  // 1
)

type paymentService interface {
}
type userService interface {
	CreateOrUpdateCustomer(int64) error
}

type payment struct {
	s  paymentService
	us userService
}

type PaymentReq struct {
	Command paymentCommand
	UserID  int64
}

func NewPayment(s paymentService, us userService) queue.Worker {
	return &payment{s: s, us: us}
}

func (w payment) Type() string {
	return PaymentType
}

func (w payment) Run(j *que.Job) error {
	var r PaymentReq
	err := json.Unmarshal(j.Args, &r)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal job arguments into request: %s, err: %+v ", j.Args, err)
	}

	switch r.Command {
	case CommandCustomerCreateUpdate:
		return w.us.CreateOrUpdateCustomer(r.UserID)
	}
	// err = w.s.Send(r.RecipientEmail, r.Subject, r.HTML, r.Text)
	// if err != nil {
	// 	logs.Error("Worker %s failed! Error: %s", w.Type(), err)
	// }
	return err
}
