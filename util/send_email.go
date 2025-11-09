package util

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to string, subject string, resetLink string) error {
	from := "bookshop430@gmail.com"
	appPassword := "iewaxjyygwxxyzba"

	body := fmt.Sprintf(`
	<html>
		<body style="font-family: Arial, sans-serif; background-color:#f8f9fa; padding:20px;">
			<div style="max-width:600px; margin:auto; background-color:white; padding:20px; border-radius:10px; box-shadow:0 0 8px rgba(0,0,0,0.1);">
				<h2 style="color:#007BFF;">Password Reset Request</h2>
				<p>Hello,</p>
				<p>We received a request to reset your password for your BookShop account.</p>
				<p>Click the button below to set a new password:</p>
				<p style="text-align:center;">
					<a href="%s" style="display:inline-block; background-color:#007BFF; color:white; padding:10px 20px; text-decoration:none; border-radius:5px;">Reset Password</a>
				</p>
				<p>This link will expire in <b>15 minutes</b>. If you didn’t request this, you can safely ignore this email.</p>
				<br>
				<p style="font-size:12px; color:#666;">— The BookShop Team</p>
			</div>
		</body>
	</html>`, resetLink)

	msg := fmt.Sprintf("From: BookShop Support <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"+
		"%s", from, to, subject, body)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		smtp.PlainAuth("", from, appPassword, smtpHost),
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
