package voice

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type VoiceStore struct {
	IP              string
	Port            string
	TranslateMe     string
	Translated      string
	LastTranslated  string
	LastTranslateMe string
	SpeakMe         string
	Client          *http.Client
	ChanSpeakMe     chan string
	Terminatate     chan bool
}

func Create() (v VoiceStore) {
	v.Client = &http.Client{
		Timeout: 1 * time.Second,
	}

	v.ChanSpeakMe = make(chan string)
	v.Terminatate = make(chan bool)

	return v
}

func (v *VoiceStore) Start() (err error) {
	out, err := v.Requset("get_engine", `{"Text": ""}`)
	if err != nil {
		return err
	}
	fmt.Println(out)

	// com.google.android.tts com.acapelagroup.android.tts
	out, err = v.Requset("set_engine", `{"Text": "com.google.android.tts"}`)
	if err != nil {
		return err
	}
	fmt.Println(out)

	out, err = v.Requset("set_speech_rate", `{"SpeechRate": 3}`)
	if err != nil {
		return err
	}
	fmt.Println(out)

	err = v.SpeekLoop()

	return
}

func (v *VoiceStore) SpeekLoop() (err error) {
FOR0:
	for {
		select {
		case v.SpeakMe = <-v.ChanSpeakMe:
		case <-v.Terminatate:
			break FOR0
		}

		str := fmt.Sprintf(`{"Text": "%s"}`, v.SpeakMe)
		out, err := v.Requset("play_on_android", str)
		if err != nil {
			return err
		}
		fmt.Println(out)
		time.Sleep(time.Second * 1)
	}
	return
}

func (v *VoiceStore) Stop() {
}

func (v *VoiceStore) Requset(method, input string) (out string, err error) {
	url := fmt.Sprintf("http://%s/%s", v.IP, method)
	data := []byte(input)
	r := bytes.NewReader(data)

	resp, err := v.Client.Post(url, "application/json", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	out = string(bodyBytes)
	return
}
