package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	recipient := make(chan Recipient)

	var wg sync.WaitGroup
	go loadRecipient("./Emails.csv", recipient)

	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipient, &wg)
	}

	wg.Wait()
}

func ExecuteTemp(r Recipient) (string, error) {
	t, err := template.ParseFiles("./email.tmpl")

	if err != nil {
		return "", nil
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)

	if err != nil {
		return "", nil
	}
	return tpl.String(), nil
}
