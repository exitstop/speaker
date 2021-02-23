package google

import (
	"fmt"
	"time"

	"github.com/exitstop/speaker/internal/browser"
	"github.com/mxschmitt/playwright-go"
	"github.com/sirupsen/logrus"
)

func Create() (gstore GStore) {
	gstore.TimeoutWaitTranslate = 100 * time.Millisecond
	gstore.CountLoopWaitTranslate = 30
	gstore.ChanTranslateMe = make(chan string)
	gstore.Terminatate = make(chan bool)

	return
}

func (s *GStore) Start() (err error) {
	s.Pw, err = playwright.Run()
	if err != nil {
		return
	}
	s.Browser, err = s.Pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		//Headless: playwright.Bool(false),
		//ChromiumSandbox: playwright.Bool(false),
	})
	if err != nil {
		err = fmt.Errorf("could not launch browser: %v\n", err)
		return
	}
	s.Page, err = s.Browser.NewPage()
	if err != nil {
		err = fmt.Errorf("could not create page: %v\n", err)
		return
	}

	err = s.Google()
	if err != nil {
		err = fmt.Errorf("Goole: %v\n", err)
		return
	}

	return
}

func (s *GStore) Stop() (err error) {
	if err = s.Browser.Close(); err != nil {
		err = fmt.Errorf("could not close browser: %v\n", err)
		return
	}
	if err = s.Pw.Stop(); err != nil {
		err = fmt.Errorf("could not stop Playwright: %v\n", err)
		return
	}
	return
}

func (s *GStore) Google() (err error) {
	s.Url = "https://translate.google.com/?hl=en#view=home&op=translate&sl=auto&tl=ru"

	if _, err = s.Page.Goto(s.Url); err != nil {
		err = fmt.Errorf("could not goto: %v\n", err)
		return
	}

	if err = s.SetEventGoogle(); err != nil {
		err = fmt.Errorf("could not set text: %v\n", err)
		return
	}

	return
}

// Запускаем обработчик ожидающий строки для перевода
func (s *GStore) LoopTransalate() (err error) {

FOR0:
	for {
		select {
		case s.ToTranslete = <-s.ChanTranslateMe:
		case <-s.Terminatate:
			break FOR0
		}

		if s.LastTranslete == s.ToTranslete {
			if s.TranslatedText != "" {
				s.SendTranslateToSpeak <- s.TranslatedText
				fmt.Println("REPEATE LAST TRANSLATE")
				continue
			}
		}
		s.LastTranslete = s.ToTranslete

		s.ClearVar()

		if err = s.SetText(s.ToTranslete); err != nil {
			err = fmt.Errorf("could not set text: %v\n", err)
			break
		}

		s.TranslatedText, err = s.WaitTextTranslate()

		if err != nil {
			err = fmt.Errorf("не удалось перевести: %s", err.Error())

			//s.SendTranslateToSpeak <- "не удалось перевести"
			//strErr := fmt.Sprintf("не удалось перевести: %s", err.Error())
			s.SendTranslateToSpeak <- err.Error()
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Translate")

			continue
		} else {
			//fmt.Println(translateText)
			s.SendTranslateToSpeak <- s.TranslatedText
		}
	}
	return
}

// Ждем пока появится перевод
func (s *GStore) WaitTextTranslate() (parseText string, err error) {
	var text string
	for i := 0; i < s.CountLoopWaitTranslate; i++ {
		text, err = s.GetTextGoogle()
		if err != nil {
			time.Sleep(s.TimeoutWaitTranslate)
			continue
		}
		parseText, err = browser.ParseGoogle5(text)
		if err != nil {
			err = fmt.Errorf("перевод не распарсился")
			return parseText, err
		}
		break
	}
	if parseText == "" {
		err = fmt.Errorf("пустая строка")
	}
	return
}

// Забираем перевод
func (s *GStore) GetTextGoogle() (text string, err error) {
	handle, err := s.Page.EvaluateHandle(JS_GET_TEXT_GOOGLE)
	if err != nil {
		err = fmt.Errorf("could not get text google: %v\n", err)
		return
	}

	text = handle.String()
	if text == "" || text == "JSHandle@" {
		err = fmt.Errorf("empty text: \n")
		return
	}

	return
}

// Очищаем переменную js с переводом
func (s *GStore) ClearVar() (err error) {
	_, err = s.Page.Evaluate(JS_CLEAR_VAR)
	if err != nil {
		err = fmt.Errorf("could clear var: %v\n", err)
		return
	}
	return
}

func (s *GStore) SetEventGoogle() (err error) {
	_, err = s.Page.Evaluate(JS_SET_EVENT_GOOGLE)
	if err != nil {
		err = fmt.Errorf("could not set event google: %v\n", err)
		return
	}
	return
}

func (s *GStore) SetText(text string) (err error) {
	jsInput := fmt.Sprintf(JS_SET_TYPE, text)

	_, err = s.Page.Evaluate(jsInput)

	if err != nil {
		err = fmt.Errorf("could not acquire JSHandle: text %s %v\n", jsInput, err)
		return
	}
	return
}
