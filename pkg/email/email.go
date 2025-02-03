package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type EmailServiceImpl struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

func NewEmailService(smtpHost string, smtpPort int, smtpUsername string, smtpPassword string) EmailService {
	return &EmailServiceImpl{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
	}
}

func (es *EmailServiceImpl) SendEmail(to string, subject string, body string) error {
	// Build email message
	from := es.smtpUsername
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body + "\r\n")

	// Handle multiple recipients
	toAddresses := strings.Split(to, ",")

	// Create TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // For testing only, prod should use valid certificate
		ServerName:         es.smtpHost,
	}

	// Connect to SMTP server with SSL
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", es.smtpHost, es.smtpPort), tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, es.smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer client.Close()

	// Authentication
	auth := smtp.PlainAuth(
		"",
		es.smtpUsername,
		es.smtpPassword, // Use SMTP authorization code here
		es.smtpHost,
	)

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	// Set sender
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	// Set recipients
	for _, addr := range toAddresses {
		if err = client.Rcpt(strings.TrimSpace(addr)); err != nil {
			return fmt.Errorf("failed to add recipient %s: %v", addr, err)
		}
	}

	// Send email body
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %v", err)
	}
	defer w.Close()

	if _, err = w.Write(msg); err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	return nil
}
