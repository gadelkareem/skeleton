package services

import (
    "errors"
    "fmt"
    "net/mail"

    "backend/kernel"
    "backend/models"
    "backend/queue"
    "backend/queue/workers"
    "github.com/go-gomail/gomail"
    "github.com/matcornic/hermes/v2"
)

type (
    EmailService struct {
        Dialer                *gomail.Dialer
        SenderCloser          gomail.SendCloser
        smtpUser, senderEmail string
        q                     *queue.QueManager
    }
)

func NewEmailService(d *gomail.Dialer, sc gomail.SendCloser, q *queue.QueManager) *EmailService {
    return &EmailService{Dialer: d,
        SenderCloser: sc,
        smtpUser:     kernel.App.Config.String("smtp::smtpUser"),
        senderEmail:  kernel.App.Config.String("smtp::senderEmail"),
        q:            q,
    }
}

// send sends the email
func (s *EmailService) Send(recipientEmail, subject, htmlBody, txtBody string) error {

    if recipientEmail == "" {
        return errors.New("no receiver emails configured")
    }
    if subject == "" {
        return errors.New("no subject specified")
    }

    from := mail.Address{
        Name:    s.smtpUser,
        Address: s.senderEmail,
    }

    m := gomail.NewMessage()
    m.SetHeader("From", from.String())
    m.SetHeader("To", recipientEmail)
    m.SetHeader("Subject", subject)

    m.SetBody("text/plain", txtBody)
    m.AddAlternative("text/html", htmlBody)

    sc := s.SenderCloser
    if sc == nil {
        var err error
        sc, err = s.Dialer.Dial()
        if err != nil {
            return err
        }
    }

    return gomail.Send(sc, m)
}

func (s *EmailService) WelcomeEmail(recipientName, recipientEmail string) (err error) {

    email := hermes.Email{
        Body: hermes.Body{
            Name:     recipientName,
            Greeting: "Welcome",
            Intros: []string{
                fmt.Sprintf("Welcome to %s! We're very excited to have you on board.", kernel.SiteName),
            },
            Outros: []string{
                "Need help, or have questions? Just reply to this email, we'd love to help.",
            },
        },
    }

    return s.generate(email, recipientEmail, fmt.Sprintf("Welcome to %s", kernel.SiteName))
}

func (s *EmailService) VerifyUserEmail(recipientName, recipientEmail, verifyLink string) (err error) {

    email := hermes.Email{
        Body: hermes.Body{
            Name: recipientName,
            Intros: []string{
                fmt.Sprintf("Thank you for checking out %s service.", kernel.SiteName),
            },
            Actions: []hermes.Action{
                {
                    Instructions: fmt.Sprintf("To get started with %s, please verify your email by clicking here:", kernel.SiteName),
                    Button: hermes.Button{
                        Color: "#22BC66", // Optional action button color
                        Text:  "Verify Your Email",
                        Link:  verifyLink,
                    },
                },
            },
            Outros: []string{
                "Need help, or have questions? Just reply to this email, we'd love to help.",
            },
        },
    }

    return s.generate(email, recipientEmail, fmt.Sprintf("Verify your %s email address", kernel.SiteName))
}

func (s *EmailService) ForgotPasswordEmail(m *models.User) (err error) {

    email := hermes.Email{
        Body: hermes.Body{
            Name: m.GetFullName(),
            Intros: []string{
                fmt.Sprintf("You have received this email because a password reset request for %s account was received.", kernel.SiteName),
            },
            Actions: []hermes.Action{
                {
                    Instructions: fmt.Sprintf("Your user name is %s", m.Username),
                },
                {
                    Instructions: "Click the button below to reset your password:",
                    Button: hermes.Button{
                        Color: "#DC4D2F",
                        Text:  "Reset your password",
                        Link:  m.GetResetPasswordURL(),
                    },
                },
            },
            Outros: []string{
                "Need help, or have questions? Just reply to this email, we'd love to help.",
            },
        },
    }

    return s.generate(email, m.Email, fmt.Sprintf("Reset your %s password", kernel.SiteName))
}

func (s *EmailService) generate(email hermes.Email, recipientEmail, subject string) (err error) {

    // Configure hermes by setting a theme and your product info
    hr := hermes.Hermes{
        // Optional Theme
        // Theme: new(Default)
        Product: hermes.Product{
            // Appears in header & footer of e-mails
            Name: kernel.SiteName,
            Link: kernel.App.FrontEndURL,
            // Optional product logo
            Logo:      "http://gadelkareem.github.io/images/skeleton.png",
            Copyright: "Copyright © 2020 Skeleton. All rights reserved.",
        },
    }

    // Generate an HTML email with the provided contents (for modern clients)
    html, err := hr.GenerateHTML(email)
    if err != nil {
        return err
    }

    // Generate the plaintext version of the e-mail (for clients that do not support xHTML)
    txt, err := hr.GeneratePlainText(email)
    if err != nil {
        return err
    }

    // // Optionally, preview the generated HTML e-mail by writing it to a local file
    // err = ioutil.WriteFile("preview.html", []byte(html), 0644)
    // if err != nil {
    // 	return err
    // }

    if s.q != nil {
        return s.enqueue(recipientEmail, subject, html, txt)
    }

    return s.Send(recipientEmail, subject, html, txt)
}

func (s *EmailService) enqueue(recipientEmail, subject, htmlBody, txtBody string) error {
    return s.q.Enqueue(workers.SendMailType, workers.SendMailReq{
        RecipientEmail: recipientEmail,
        Subject:        subject,
        HTML:           htmlBody,
        Text:           txtBody,
    })
}
