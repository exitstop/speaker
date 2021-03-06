package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/exitstop/speaker/internal/browser"
	hook "github.com/robotn/gohook"
)

//var ip string = "192.168.0.133:8484"
var ip string = "192.168.88.50:8484"
var nFlagTransel bool
var last string

func main() {
	var nFlag string
	flag.StringVar(&nFlag, "ip", "192.168.0.133", "ip")
	flag.BoolVar(&nFlagTransel, "t", false, "translate")
	flag.Parse()
	ip = nFlag + ":8484"

	if nFlagTransel {
		browser.InitBrowser()
		browser.SetBrowserLang1("en-US")
		browser.SetBrowserLang2("ru-RU")
		transl := browser.GetTranslate("hello world")
		fmt.Println(transl)
	}

	clientLong := &http.Client{
		Timeout: 4 * time.Second,
	}

	{
		url := "http://" + ip + "/get_engine"
		data := []byte(`{"Text": "hello"}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if true {
		url := "http://" + ip + "/set_engine"
		data := []byte(`{"Text": "com.acapelagroup.android.tts"}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	time.Sleep(2 * time.Second)
	if true {
		url := "http://" + ip + "/get_voice"
		data := []byte(`{"Text": "com.acapelagroup.android.tts"}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	if false {
		url := "http://" + ip + "/set_voice"
		data := []byte(`{"Text": "eng-USA"}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("SetVoice", bodyString)
	}
	{
		url := "http://" + ip + "/set_language"
		data := []byte(`{"Text": "eng-USA"}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("SetVoice", bodyString)
	}

	if true {
		url := "http://" + ip + "/set_speech_rate"
		data := []byte(`{"SpeechRate": 2}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	if true {
		url := "http://" + ip + "/set_pitch"
		data := []byte(`{"SetPitch": 0.5}`)
		r := bytes.NewReader(data)

		resp, err := clientLong.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	time.Sleep(2 * time.Second)

	{
		bigText := "Hello world"
		url := "http://" + ip + "/play_on_android"

		type Text struct {
			Text string `json:"Text"`
		}
		user := &Text{Text: bigText}
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}
		r := bytes.NewReader(b)

		client := &http.Client{
			Timeout: 1 * time.Second,
		}

		resp, err := client.Post(url, "application/json", r)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	add()
	low()
	return

}

var translSave string

func SaveTransl(input string) {
	translSave = input
}

func GatLastTransl() string {
	return translSave
}

func add() {
	fmt.Println("--- Please press ctrl + c to stop hook ---")

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	hook.Register(hook.KeyDown, []string{"c", "ctrl"}, func(e hook.Event) {
		//fmt.Println("ctrl-c")
		actual, err := clipboard.ReadAll()
		//hook.End()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(actual)
		var transl string

		if last == actual {
			// Повторяем в слух уже переведенное ранее
			transl = GatLastTransl()
			fmt.Println("-----------------------")
			fmt.Println("----last: ", last)
			fmt.Println("----transl: ", transl)
			fmt.Println("-----------------------")
		} else {

			words := strings.Fields(actual)

			if len(words) == 1 {
				actual = "word, " + actual
			}

			if actual == "" || actual == " " {
				return
			}

			last = actual

			if nFlagTransel {
				// hello world
				transl = browser.GetTranslateGoogle(actual)
			} else {
				transl = actual
			}

			SaveTransl(transl)
		}
		fmt.Println("transl:", transl, actual)

		if transl != "0" {

			bigText := RegexWork(transl)
			url := "http://" + ip + "/play_on_android"

			type Text struct {
				Text string `json:"Text"`
			}
			user := &Text{Text: bigText}
			b, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err)
				return
			}

			//data := []byte(`{"text": "` + bigText + `"}`)
			r := bytes.NewReader(b)

			resp, err := client.Post(url, "application/json", r)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		}
	})

	//fmt.Println("--- Please press w---")
	//hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
	//fmt.Println("w")
	//})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	//EvChan := hook.Start()
	_ = hook.Start()
	defer hook.End()

	//for ev := range EvChan {
	//fmt.Println("hook: ", ev)
	//}
}

//func TextClear() {
//text = strings.ReplaceAll(text, "//", " ")
//text = strings.ReplaceAll(text, "|", " ")
//text = strings.ReplaceAll(text, "*", " ")
//}

//func Substr(mess string) (string, error) {
//re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
//re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
//mess = re_leadclose_whtsp.ReplaceAllString(mess, "")
//mess = re_inside_whtsp.ReplaceAllString(mess, " ")
//}

func RegexWork(tt string) string {
	reg, _ := regexp.Compile(`[\_\]\[\@\#\/\:]+`)
	reg2, _ := regexp.Compile(`([\p{L}])\.([\p{L}])`)
	reg3, _ := regexp.Compile(`([[:lower:]])([[:upper:]])`)
	reg4, _ := regexp.Compile(`(\b(\p{L}+)\b)`)
	tt = reg.ReplaceAllString(tt, " ")
	tt = reg2.ReplaceAllString(tt, "$1. $2")
	tt = reg3.ReplaceAllString(tt, "$1 $2")
	tt = reg4.ReplaceAllString(tt, " $1 ")

	tt = strings.TrimSpace(tt)
	return tt
}
