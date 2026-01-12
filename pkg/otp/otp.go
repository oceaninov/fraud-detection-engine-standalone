package otp

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/sirupsen/logrus"
)

var auth smtp.Auth

type otp struct {
	cfg Config
}

func (o otp) Email(email, url string) error {
	const emailLayout = `
<html>
    <head></head>
    <body>
        <div style="font-family: Helvetica,Arial,sans-serif;min-width:1000px;overflow:auto;line-height:2">
            <div style="margin:50px auto;width:70%;padding:20px 0">
                <div style="border-bottom:1px solid #eee">
                    <a href="#" style="font-size:1.4em;color: #00466a;text-decoration:none;font-weight:600">
                        Reset Password Ocean - FDS Dashboard
                    </a>
                </div>
                <p>Halo,  
                   <br> Berikut adalah link untuk mengubah password Ocean - FDS Dashboard anda
                   <br> Email&emsp;: {{.Email}} 
                   <br> Link&emsp;&emsp;&emsp;: {{.Url}}
                </p>
                <p style="font-size:0.9em;"><br /></p>
                <hr style="border:none;border-top:1px solid #eee" />
                <div style="float:right;padding:8px 0;color:#aaa;font-size:0.8em;line-height:1;font-weight:300">
                </div>
            </div>
        </div>
    </body>
</html>`

	// Parse the embedded HTML template
	t, err := template.New("email").Parse(emailLayout)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("template parsing error: %+v", err)
	}

	// Prepare the email body
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, struct {
		Email string
		Url   string
	}{
		Email: email,
		Url:   url,
	}); err != nil {
		logrus.Error(err)
		return fmt.Errorf("mailer send error: %+v", err)
	}
	body := buf.String()

	// Email headers and content
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Reset Password\n"
	msg := []byte(subject + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%s", o.cfg.SMTPHost, o.cfg.SMTPPort)

	// Send the email
	if err := smtp.SendMail(addr, auth, o.cfg.SMTPSender, []string{email}, msg); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func NewOTP(config Config) OTP {
	auth = smtp.PlainAuth("", config.SMTPAuth, config.SMTPAuthPass, config.SMTPHost)
	return &otp{
		cfg: config,
	}
}
