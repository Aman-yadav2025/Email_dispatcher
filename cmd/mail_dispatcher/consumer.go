package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	for recipient := range ch {
		fmt.Println(id,recipient)
	}
}