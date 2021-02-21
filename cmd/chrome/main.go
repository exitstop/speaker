package main

import (
	"log"
	"time"

	"github.com/exitstop/speaker/internal/google"
)

func main() {
	gstore := google.Create()

	// Запускаем браузер
	if err := gstore.Start(); err != nil {
		log.Println(err)
		return
	}

	// После того как запустился браузер пошлем строки для перевода и запустим gstore.LoopTransalate()
	go func() {
		time.Sleep(1 * time.Second)
		gstore.ChanTranslateMe <- `You can also specify JSHandle as the property value if you want live objects to be passed into the event:`
		time.Sleep(1 * time.Second)
		gstore.ChanTranslateMe <- `Here you hand errorChannelWatch the errorList as a value.`

		time.Sleep(1 * time.Second)
		gstore.ChanTranslateMe <- `To remedy the situation, either hand a slice pointer to errorChannelWatch or rewrite it as a call to a closure, capturing errorList.`

		gstore.Terminatate <- true
	}()

	// Обработка строк для перевода, посылаемых через ChanTranslateMe
	if err := gstore.LoopTransalate(); err != nil {
		log.Println(err)
		return
	}

	if err := gstore.Stop(); err != nil {
		log.Println(err)
		return
	}
}
