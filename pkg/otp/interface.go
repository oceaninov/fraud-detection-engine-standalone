package otp

type Config struct {
	SMTPHost     string
	SMTPPort     string
	SMTPAuthPass string
	SMTPAuth     string
	SMTPSender   string
}

type OTP interface {
	Email(email, url string) error
}
