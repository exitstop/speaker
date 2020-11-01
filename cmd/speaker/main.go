package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	clientLong := &http.Client{
		Timeout: 4 * time.Second,
	}

	{
		url := "http://192.168.0.133:8484/get_engine"
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

	{
		url := "http://192.168.0.133:8484/set_engine"
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

	if true {
		url := "http://192.168.0.133:8484/set_speech_rate"
		data := []byte(`{"SpeechRate": 3}`)
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
		bigText := "Вы считаете безопасно хранить my_key.jks в репозитории? Если злоумышленник получит свою копию ключа остается надеятся только на стойкость алгоритма шифрования. Я бы не стал подсаживаться на переменную github.run_number. Что делать если придется мигрировать на GitLab/TeamCity, или когда у внешнего сервиса собьется нумерация сборок? Вместо того, чтобы вручную докручивать образ ubuntu-latest, можно взять один из готовых от CircleCi. Совсем необязательно выносить работу keystore.properties на уровень всего скрипта сборки. Я предпочитаю иметь ограниченную область, в которой можно обратиться к RELEASE_STORE_PASSWORD. В репозитории на GitHub есть раздел с релизами, было бы круто автоматом отправлять туда артефакты сборки, в том числе от proguard/R8."
		url := "http://192.168.0.133:8484/play_on_android"
		data := []byte(`{"text": "` + bigText + `"}`)
		r := bytes.NewReader(data)

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
}
