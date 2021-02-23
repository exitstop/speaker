package voice

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type VoiceStore struct {
	IP          string
	Port        string
	SpeakMe     string
	Client      *http.Client
	ChanSpeakMe chan string
	Terminatate chan bool
	ChanPause   chan bool
	Pause       bool
	SpeechSpeed float64
	NoTranslate bool // не переводить текст
}

func Create() (v VoiceStore) {
	v.Client = &http.Client{
		Timeout: 2 * time.Second,
	}

	v.ChanSpeakMe = make(chan string)
	v.Terminatate = make(chan bool)
	v.ChanPause = make(chan bool)

	v.SpeechSpeed = 3.7

	return v
}

func (v *VoiceStore) Start() (err error) {
	out, err := v.Requset("get_engine", `{"Text": ""}`)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"out": out,
	}).Info("get_engine")

	// com.google.android.tts com.acapelagroup.android.tts
	out, err = v.Requset("set_engine", `{"Text": "com.google.android.tts"}`)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"out": out,
	}).Info("set_engine")

	out, err = v.Requset("set_speech_rate", `{"SpeechRate": 3}`)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"out": out,
	}).Info("set_speech_rate")

	str := fmt.Sprintf(`{"Text": "Инициализация успешна"}`)
	out, err = v.Requset("play_on_android", str)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"out": out,
	}).Info("play_on_android")

	err = v.SpeekLoop()

	return
}

func (v *VoiceStore) SpeedSub() (out string, speed float64, err error) {
	v.SpeechSpeed -= 0.1
	strSpeed := fmt.Sprintf(`{"SpeechRate": %.2f}`, v.SpeechSpeed)
	out, err = v.Requset("set_speech_rate", strSpeed)
	out = fmt.Sprintf("%s %.1f", out, v.SpeechSpeed)
	speed = v.SpeechSpeed
	return
}

func (v *VoiceStore) SpeedAdd() (out string, speed float64, err error) {
	v.SpeechSpeed += 0.1
	strSpeed := fmt.Sprintf(`{"SpeechRate": %.2f}`, v.SpeechSpeed)
	out, err = v.Requset("set_speech_rate", strSpeed)
	out = fmt.Sprintf("%s %.1f", out, v.SpeechSpeed)
	speed = v.SpeechSpeed
	return
}

func (v *VoiceStore) SpeekLoop() (err error) {
FOR0:
	for {
		select {
		case v.SpeakMe = <-v.ChanSpeakMe:
		case <-v.Terminatate:
			break FOR0
		case v.Pause = <-v.ChanPause:
			v.Pause = <-v.ChanPause
			v.SpeakMe = "пауза снята"
		}

		str := fmt.Sprintf(`{"Text": "%s"}`, v.SpeakMe)
		out, err := v.Requset("play_on_android", str)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("SpeakLoop")
			continue
		}
		logrus.WithFields(logrus.Fields{
			"out": out,
		}).Info("play_on_android")

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
