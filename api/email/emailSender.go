package email

import (
	"bytes"
	"encoding/csv"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/env"
	"gopkg.in/gomail.v2"
	"os"
	"time"
)

func SendCSVViaGmail(data [][]string) error {

	sender := env.GetEnvVariable("SENDER_EMAIL")
	password := env.GetEnvVariable("SENDER_PASSWORD")
	receiver := env.GetEnvVariable("RECEIVER_EMAIL")

	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)

	date := time.Now().Format("2006-01-02")

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "Products CSV "+date)

	// Create a buffer to hold the CSV content
	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)

	// Write data to the buffer
	err := csvWriter.WriteAll(data) // calls Flush internally
	if err != nil {
		return err
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "products-*.csv")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name()) // Clean up the file after sending the email

	// Write the buffer to the temporary file
	if _, err := tmpFile.Write(csvBuffer.Bytes()); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}

	// Attach the temporary file
	m.Attach(tmpFile.Name())

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
