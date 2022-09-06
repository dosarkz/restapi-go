package email

import (
	"net/smtp"
	"os"
)

// Object defines email payload data
type Object struct {
	To      string
	Body    string
	Subject string
	From    string
}

// SendMail method to send email to user
func SendMail(subject, body, to, from string) bool {
	emptyTo := []string{to}
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))
	// Here we do it all: connect to our server, set up a message and send it
	err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
		auth, from, emptyTo, msg)
	if err != nil {
		//	fmt.Println(err)
		return false
	}

	//fmt.Println("Email sent successfully")
	return true
}
