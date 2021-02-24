package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/exitstop/speaker/internal/console"
	"github.com/exitstop/speaker/internal/google"
	"github.com/exitstop/speaker/internal/gtts"
	"github.com/exitstop/speaker/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	var nFlag string
	var nFlagTransel bool

	flag.StringVar(&nFlag, "ip", "192.168.0.133", "ip")
	flag.BoolVar(&nFlagTransel, "t", false, "translate")
	flag.Parse()

	{
		logger.InitLog("trace", "speaker", "/home/bg/go/src/github.com/exitstop/speaker/logs")

		logrus.SetLevel(logrus.InfoLevel)

		logrus.WithFields(logrus.Fields{
			"Keyboard": "ok",
		}).Info("Init")
	}
	gstore := google.Create()

	// Запускаем браузер
	if err := gstore.Start(); err != nil {
		log.Println(err)
		return
	}
	defer gstore.Stop()

	// voice google tts
	v := gtts.Create()
	v.IP = nFlag + ":8484"

	gstore.Terminatate = v.Terminatate

	//Signal := make(chan os.Signal)
	//signal.Notify(Signal, os.Interrupt)
	//signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	//// graceful shutdown
	//go func() {
	//<-Signal
	//gstore.Terminatate <- true
	//}()

	gstore.SendTranslateToSpeak = v.ChanSpeakMe

	go func() {
		//console.Keyboard()
		console.Add(&gstore, &v)
		console.Low()
		console.Event()

		gstore.Terminatate <- true
		gstore.Terminatate <- true
	}()

	// Переводчик сам будет слать в chan ChanSpeakMe, чтобы голос воспроизводился
	//go func() {
	//time.Sleep(3 * time.Second)
	//gstore.chantranslateme <- `You can also specify JSHandle as the property value if you want live objects to be passed into the event:`
	//time.Sleep(3 * time.Second)
	//gstore.ChanTranslateMe <- `Here you hand errorChannelWatch the errorList as a value.`

	//time.Sleep(3 * time.Second)
	//gstore.ChanTranslateMe <- `To remedy the situation, either hand a slice pointer to errorChannelWatch or rewrite it as a call to a closure, capturing errorList.`

	//gstore.Terminatate <- true
	//v.Terminatate <- true
	//}()

	go func() {
		// Обработка строк для перевода, посылаемых через ChanTranslateMe
		if err := gstore.LoopTransalate(); err != nil {
			log.Println(err)
			return
		}
	}()

	//go func() {
	//for i := 0; i < 20; i++ {
	//time.Sleep(time.Second * 1)
	//str := fmt.Sprintf("Привет мир %d", i)
	//v.ChanSpeakMe <- str
	//}
	//}()

	err := v.Start()
	if err != nil {
		gstore.Terminatate <- true
		gstore.Terminatate <- true
		fmt.Println(err)
		return
	}

	defer v.Stop()

	//if err := gstore.Stop(); err != nil {
	//log.Println(err)
	//return
	//}
}
