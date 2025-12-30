package main

import (
	"fmt"
	"net/smtp"
)

func SendEmailAlert(toEmail string, subject string, body string) error {
	// 1. Configuration (Best practice: Load these from env variables)
	// Example for Gmail. Outlook would be smtp.office365.com
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// WARNING: Never hardcode passwords. Use os.Getenv("EMAIL_PASSWORD")
	fromEmail := "youremail@gmail.com"
	password := "yourApppassword" // NOT your login password (see note below)

	// 2. Authentication
	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	// 3. Construct the Message (Headers + Body)
	// Note: The empty line "\r\n" separates headers from the body
	msg := []byte(
		"To: " + toEmail + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

	// 4. Send the email
	address := smtpHost + ":" + smtpPort
	err := smtp.SendMail(address, auth, fromEmail, []string{toEmail}, msg)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := SendEmailAlert("recepient@email.com", "Pipeline Alert", "The nightly build failed.")
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
