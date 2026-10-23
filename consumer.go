package main

import (
	"fmt"
	"log"
	"net/smtp"
	//"os/exec"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintln("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email,"Just testing the email function")
		// msg := []byte(formattedMsg)
		msg, err := ExecuteTemp(recipient)

		if err != nil {
			fmt.Errorf("Worker %d Error parsing template for %s", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: sending mail to %s \n",id,recipient.Email)
		err =smtp.SendMail(smtpHost + ":" + smtpPort, nil, "example1@gmail.com", []string{recipient.Email},[]byte(msg))
		if err != nil{
			fmt.Errorf("Worker %d: failed to send mail to %s \n",id,recipient.Email)
			log.Fatal(err)
		}
		time.Sleep(50*time.Millisecond)
		fmt.Printf("Worker %d: sent mail to %s \n",id,recipient.Email)
	}
	defer wg.Done()
}

