package console

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
	"github.com/exitstop/robotgo"
	"github.com/exitstop/speaker/internal/basic"
	"github.com/exitstop/speaker/internal/google"
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

type model struct {
	LogLevel    int
	MaxLogLevel int
}

var mod = model{
	LogLevel:    0,
	MaxLogLevel: len(LogLevelString),
}

var LogLevelString = [...]string{
	"info",
	"trace",
	"debug",
	"warning",
	"error",
	"fatal",
}

func (m *model) LevelIntToString() {
	m.LogLevel++
	id := m.LogLevel % m.MaxLogLevel
	if m.LogLevel >= m.MaxLogLevel {
		m.LogLevel = 0
	}

	switch id {
	case 0: //nolint:goconst
		logrus.SetLevel(logrus.InfoLevel)
	case 1: //nolint:goconst
		logrus.SetLevel(logrus.TraceLevel)
	case 2: //nolint:goconst
		logrus.SetLevel(logrus.DebugLevel)
	case 3:
		logrus.SetLevel(logrus.WarnLevel)
	case 4: //nolint:goconst
		logrus.SetLevel(logrus.ErrorLevel)
	case 5: //nolint:goconst
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	fmt.Println("logLevel", LogLevelString[id])
}

func Keyboard() (err error) {
	if err = keyboard.Open(); err != nil {
		return
	}

	defer func() {
		_ = keyboard.Close()
	}()

	//logger.InitLog("trace", "proxy", "/var/log/facechain/")

	//logrus.SetLevel(logrus.ErrorLevel)

	//logrus.WithFields(logrus.Fields{
	//"Keyboard": "ok",
	//}).Info("keyboard")

FOR0:
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyCtrlC:
			break FOR0
		}

		switch char {
		case 'q':
			break FOR0
		case 'c':
			break FOR0
		case 'l':
			mod.LevelIntToString()
		}

		if key == keyboard.KeyEsc {
			break FOR0
		}
	}
	//os.Exit(0)
	return
}

//func Add(gstore *google.GStore, voice *voice.VoiceStore) {
func Add(gstore *google.GStore, voice basic.VoiceInterface) {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		//robotgo.EventEnd()

		voice.ChSpeakMe("завершение программы")

		time.Sleep(1 * time.Second)

		//voice.Terminatate <- true
		//voice.Terminatate <- true
		voice.Exit()
		voice.Exit()
	})

	robotgo.EventHook(hook.KeyDown, []string{"p", "ctlr", "alt"}, func(e hook.Event) {
		fmt.Println("ctrl-alt-p")
		//robotgo.EventEnd()

		if !voice.GetPause() {
			//voice.ChanSpeakMe <- "пауза"
			voice.ChSpeakMe("пауза")
		}

		time.Sleep(time.Millisecond * 1)
		//voice.ChanPause <- !voice.Pause
		voice.SetPause()
	})

	robotgo.EventHook(hook.KeyDown, []string{"t", "alt"}, func(e hook.Event) {
		fmt.Println("alt-t")
		//robotgo.EventEnd()

		//voice.NoTranslate = !voice.NoTranslate

		if voice.TanslateOrNot() {
			//voice.ChanSpeakMe <-
			voice.ChSpeakMe("без перевода")
		} else {
			//voice.ChanSpeakMe <- "переводить текст"
			voice.ChSpeakMe("переводить текст")
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"-", "alt"}, func(e hook.Event) {
		fmt.Println("-", "alt")
		out, speed, err := voice.SpeedSub()
		if err != nil {
			fmt.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			"out": out,
		}).Info("speed-")

		//voice.ChanSpeakMe <- fmt.Sprintf("%.1f", speed)
		str := fmt.Sprintf("%.1f", speed)
		voice.ChSpeakMe(str)
	})

	robotgo.EventHook(hook.KeyDown, []string{"+", "alt"}, func(e hook.Event) {
		fmt.Println("+", "alt")
		out, speed, err := voice.SpeedAdd()
		if err != nil {
			fmt.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			"out": out,
		}).Info("speed+")

		//voice.ChanSpeakMe <- fmt.Sprintf("%.1f", speed)
		str := fmt.Sprintf("%.1f", speed)
		voice.ChSpeakMe(str)
	})

	fmt.Println("--- Please press c---")
	robotgo.EventHook(hook.KeyDown, []string{"c", "ctrl"}, func(e hook.Event) {
		//if voice.Pause {
		if voice.GetPause() {
			return
		}

		time.Sleep(time.Millisecond * 50)
		text, err := clipboard.ReadAll()

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Warn("clipboard")

			//voice.ChanSpeakMe <- fmt.Sprintf()
			voice.ChSpeakMe("не скопировалось")
			return
		}

		processedString, err := RegexWork(text)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Warn("regexp")
			return
		}

		//if voice.NoTranslate {
		if voice.TanslateOrNot() {
			processedString, err := RegexWorkRu(text)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Warn("regexp")
				return
			}
			//voice.ChanSpeakMe <- processedString
			voice.ChSpeakMe(processedString)
		} else {

			select {
			case gstore.ChanTranslateMe <- processedString:
				logrus.WithFields(logrus.Fields{
					"SendoToGoole": processedString,
				}).Warn("google")
			default:
				logrus.WithFields(logrus.Fields{
					"SendoToGoole": processedString,
				}).Error("Skip text")
			}
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"r", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("r", "ctrl", "shift")
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func Low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		fmt.Println("hook: ", ev)
	}
}

func Event() {
	//ok := robotgo.AddEvents("q", "ctrl", "shift")
	//if ok {
	//fmt.Println("add events...")
	//}

	//keve := robotgo.AddEvent("k")
	//if keve {
	//fmt.Println("you press... ", "k")
	//}

	//mleft := robotgo.AddEvent("mleft")
	//if mleft {
	//fmt.Println("you press... ", "mouse left button")
	//}
}

func RegexWork(tt string) (out string, err error) {
	reg0, err := regexp.Compile("[^a-zA-Z0-9 .,]+")
	if err != nil {
		return
	}
	reg2, err := regexp.Compile(`([\p{L}])\.([\p{L}])`)
	if err != nil {
		return
	}
	reg3, err := regexp.Compile(`([[:lower:]])([[:upper:]])`)
	if err != nil {
		return
	}
	reg4, err := regexp.Compile(`(\b(\p{L}+)\b)`)
	if err != nil {
		return
	}
	tt = reg0.ReplaceAllString(tt, " ")
	tt = reg4.ReplaceAllString(tt, " $1 ")
	tt = reg3.ReplaceAllString(tt, "$1 $2")
	tt = reg2.ReplaceAllString(tt, "$1. $2")

	singleSpacePattern := regexp.MustCompile(`\s+`)
	tt = singleSpacePattern.ReplaceAllString(tt, " ")
	tt = strings.ReplaceAll(tt, " .", ".")
	tt = strings.ReplaceAll(tt, " ,", ",")

	tt = strings.TrimSpace(tt)
	return tt, err
}

func RegexWorkRu(tt string) (out string, err error) {
	reg0, err := regexp.Compile("[^а-яА-Яa-zA-Z0-9 .,]+")
	if err != nil {
		return
	}
	tt = reg0.ReplaceAllString(tt, " ")

	singleSpacePattern := regexp.MustCompile(`\s+`)
	tt = singleSpacePattern.ReplaceAllString(tt, " ")
	tt = strings.ReplaceAll(tt, " .", ".")
	tt = strings.ReplaceAll(tt, " ,", ",")

	tt = strings.TrimSpace(tt)
	return tt, err
}
