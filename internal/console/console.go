package console

import (
	"fmt"
	"regexp"

	"github.com/FaceChainTeam/gocommonutil/logger"
	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
	"github.com/exitstop/speaker/internal/voice"
	"github.com/go-vgo/robotgo"
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

	logger.InitLog("trace", "proxy", "/var/log/facechain/")

	logrus.SetLevel(logrus.ErrorLevel)

	logrus.WithFields(logrus.Fields{
		"Keyboard": "ok",
	}).Info("keyboard")

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

func Add(event chan string, voice *voice.VoiceStore) {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	robotgo.EventHook(hook.KeyDown, []string{"-", "alt"}, func(e hook.Event) {
		fmt.Println("-", "alt")
		out, speed, err := voice.SpeedSub()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(out)
		voice.ChanSpeakMe <- fmt.Sprintf("%.1f", speed)
	})

	robotgo.EventHook(hook.KeyDown, []string{"+", "alt"}, func(e hook.Event) {
		fmt.Println("+", "alt")
		out, speed, err := voice.SpeedAdd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(out)
		voice.ChanSpeakMe <- fmt.Sprintf("%.1f", speed)
	})

	fmt.Println("--- Please press c---")
	robotgo.EventHook(hook.KeyDown, []string{"c"}, func(e hook.Event) {
		text, _ := clipboard.ReadAll()
		reg, err := regexp.Compile("[^a-zA-Z0-9 .,]+")
		if err != nil {
			fmt.Println(err)
			return
		}
		processedString := reg.ReplaceAllString(text, " ")

		logrus.WithFields(logrus.Fields{
			"SendoToGoole": processedString,
		}).Warn("google")

		event <- processedString
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
	ok := robotgo.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := robotgo.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
