package main

import (
	"sync"
	"time"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	recipient := make(chan Recipient)

	var wg sync.WaitGroup
	go loadRecipient("./dummy_emails_200.csv", recipient)

	workerCount := 5

	for i:= 1;i<=workerCount;i++{
		go emailWorker(i, recipient,&wg)
	}
}
