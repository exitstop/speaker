package google

import (
	"fmt"
	"time"

	"github.com/mxschmitt/playwright-go"
)

func Create() (gstore GStore) {
	return
}

func (s *GStore) Start() (err error) {
	s.Pw, err = playwright.Run()
	if err != nil {
		return
	}
	s.Browser, err = s.Pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
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

	//// mw.config.values is the JS object where Wikipedia stores wiki metadata
	//handle, err := page.EvaluateHandle("mw.config.values", struct{}{})
	//if err != nil {
	//log.Fatalf("could not acquire JSHandle: %v\n", err)
	//}
	//// mw.config.values.wgPageName is the name of the current page
	//pageName, err := handle.(playwright.JSHandle).GetProperty("wgPageName")
	//if err != nil {
	//log.Fatalf("could not get Wikipedia page name: %v\n", err)
	//}

	//fmt.Printf("Lots of type casting, brought to you by %s\n", pageName)

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

	if err = s.SetText("hello"); err != nil {
		err = fmt.Errorf("could not set text: %v\n", err)
		return
	}

	for i := 0; i < 15; i++ {
		text, err := s.GetTextGoogle()
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Println(text)
		break
	}

	return
}

func (s *GStore) GetTextGoogle() (text string, err error) {
	handle, err := s.Page.EvaluateHandle(JS_GET_TEXT_GOOGLE)
	if err != nil {
		err = fmt.Errorf("could not get text google: %v\n", err)
		return
	}

	text = handle.String()
	if text == "" || text == "JSHandle@" {
		err = fmt.Errorf("empty text: %v\n")
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
