package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/FaceChainTeam/gocommonutil/logger"
	"github.com/exitstop/speaker/internal/console"
	"github.com/exitstop/speaker/internal/google"
	"github.com/exitstop/speaker/internal/voice"
	"github.com/sirupsen/logrus"
)

func main() {
	var nFlag string
	var nFlagTransel bool

	flag.StringVar(&nFlag, "ip", "192.168.0.133", "ip")
	flag.BoolVar(&nFlagTransel, "t", false, "translate")
	flag.Parse()

	{
		logger.InitLog("trace", "proxy", "/var/log/facechain/")

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

	// voice
	v := voice.Create()
	v.IP = nFlag + ":8484"

	gstore.SendTranslateToSpeak = v.ChanSpeakMe

	go func() {
		//console.Keyboard()
		console.Add(gstore.ChanTranslateMe, &v)
		console.Low()
		console.Event()

		gstore.Terminatate <- true
		v.Terminatate <- true
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
		fmt.Println(err)
		return
	}

	v.Stop()

	if err := gstore.Stop(); err != nil {
		log.Println(err)
		return
	}
}
