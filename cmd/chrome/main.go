package main

import (
	"log"
	"time"

	"github.com/exitstop/speaker/internal/google"
)

func main() {
	gstore := google.Create()

	if err := gstore.Start(); err != nil {
		log.Println(err)
		return
	}

	time.Sleep(20 * time.Second)

	if err := gstore.Stop(); err != nil {
		log.Println(err)
		return
	}
}
