package browser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/exitstop/speaker/internal/browserlist"
	"github.com/sclevine/agouti"
	"golang.org/x/text/language"
)

var driverTransl *agouti.WebDriver
var page *agouti.Page
var allNamePageOpen []string
var Lang1 string = "en-US"
var Lang2 string = "ru-RU"

const TIME_WAIT_MILLISECOND = 5000
const TIME_CYCLE = 50
const COUNT_TRANSLATE_SITE = 0 // 0 == google // 1 == google, yandex// google, yandex, promt
var Mu sync.Mutex
var TranslGoogle sync.Mutex
var TranslYandex sync.Mutex
var GlobalId int = -1
var DelayTranslate int64 = 0
var GlobalScreenShot int = 0
var InitBrowserTranslate int = 0
var scriptChangeLangPromt []byte
var FlagYandexElementTransl bool = false

func InitBrowser() {
	Mu.Lock()
	if InitBrowserTranslate == 1 {
		return
	}
	var err error
	//crxBytes, _ := ioutil.ReadFile("Adblock-Plus-kostenloser-Adblocker_v3.5.2.crx")
	//crxBytes2, _ := ioutil.ReadFile("Video-Speed-Controller_v0.5.6.crx")
	//crxBytes3, _ := ioutil.ReadFile("Dark-Mode_v0.3.3.crx")
	//crxBytes4, _ := ioutil.ReadFile("Clever-Mute_v0.5.1.crx")
	//crxByte5, _ := ioutil.ReadFile("~/elementYandex/elementYandex.crx")

	//pathCrx := "scripts/elementYandex/elementYandex.crx"
	//crxByte5, _ := ioutil.ReadFile(pathCrx)
	//fmt.Println(pathCrx)

	// --user-data-dir
	//"user-data-dir": "~/.config/google-chrome",
	var flagHeadless bool = false
	var flagHeadProxy bool = false
	var strHeadless string = ""
	var strProxy string = ""

	if flagHeadProxy {
		strProxy = "--proxy-server=127.0.0.1:9999"
	} else {
		strProxy = "--lang=en-US"
	}

	if flagHeadless {
		strHeadless = "--headless"
	} else {
		strHeadless = "--lang=en-US"
	}

	//
	//"--window-position=1000,1000" ,
	// --window-size=1920x1080
	// --start-maximized

	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(len(browserlist.BrowserList))
	userAgent := browserlist.BrowserList[randInt]
	userAgent = browserlist.BrowserList[0]
	log.Println("userAgent = " + userAgent)

	argsChrome := agouti.ChromeOptions("args", []string{strHeadless,
		"--no-first-run", "--user-data-dir=chrome_config",
		"--disable-gpu", "--no-sandbox", "--lang=en-US",
		"user-agent='" + userAgent + "'",
		"--enable-tab-audio-muting", "--mute-audio", "--no-default-browser-check", "--disable-setuid-sandbox",
		"--detach=false", strProxy},
	)

	//argsChrome2 := agouti.ChromeOptions("extensions", [][]byte{ crxBytes, crxBytes2, crxBytes4 })
	//argsChrome2 := agouti.ChromeOptions("extensions", [][]byte{crxByte5})
	argsChrome3 := agouti.ChromeOptions(
		"binary", "/usr/bin/google-chrome",
		//"binary", "/usr/bin/chromium-browser",
	)

	// http://www.assertselenium.com/java/list-of-chrome-driver-command-line-arguments/
	//capabilities := agouti.Capabilities{ "chromeOptions": map[string]interface{}{ "extensions": [][]byte{ crxBytes, crxBytes2 }, "args": []string{"--headless", "--disable-gpu", "--no-sandbox"}}, }
	//capabilities := agouti.Capabilities{ "chromeOptions": map[string]interface{}{ "extensions": [][]byte{ crxBytes, crxBytes2, crxBytes4 } } }

	//driverTransl = agouti.ChromeDriver( agouti.Desired(capabilities), argsChrome )
	if flagHeadless == false {
		//driverTransl = agouti.ChromeDriver(argsChrome, argsChrome2, argsChrome3)
		driverTransl = agouti.ChromeDriver(argsChrome, argsChrome3)
	} else {
		driverTransl = agouti.ChromeDriver(argsChrome)
	}

	if err := driverTransl.Start(); err != nil {
		// google-chrome version 86.0.4240.111
		// https://chromedriver.chromium.org/downloads
		// wget https://chromedriver.storage.googleapis.com/86.0.4240.22/chromedriver_linux64.zip
		// sudo cp chromedriver /usr/local/bin/
		fmt.Println("Failed to start driver:", err)
	}

	page, err = driverTransl.NewPage()
	if err != nil {
		fmt.Println("Failed to open page:", err)
	}

	// Ограничим ожидание загрузки двумя секундами
	page.SetPageLoad(2000)

	if COUNT_TRANSLATE_SITE > -1 {
		CreateNewPage(page, "googleTranslate", "https://translate.google.com/?hl=en#view=home&op=translate&sl=auto&tl=ru")
		SetEventOnGoogle(page)
	}
	if COUNT_TRANSLATE_SITE > 0 {
		CreateNewPage(page, "yandexTranslate", "https://translate.yandex.ru/?lang=en-ru")
		SetEventOnYandex(page)
	}
	if COUNT_TRANSLATE_SITE > 1 {
		CreateNewPage(page, "promtTranslate", "https://www.translate.ru/")
		SetEventOnPromt(page)
		ChangeLangPromt()
	}
	//CreateNewPage(page, "cambridgeTranslate", "https://dictionary.cambridge.org/ru/translate/")
	//CreateNewPage(page, "microsoftTranslate", "https://www.bing.com/translator/?cc=ru")
	//CreateNewPage(page, "", "")
	InitBrowserTranslate = 1
	Mu.Unlock()
}

func Destroy() {
	page.Destroy()
}

func CreateNewPage(page *agouti.Page, name string, url string) {
	var value string
	var err error

	err = page.RunScript("window.open('','"+name+"')", nil, &value)
	if err != nil {
		log.Println("Error: "+name+" not open", err)
	} else {
		allNamePageOpen = append(allNamePageOpen, name)
		page.SwitchToWindow(name)
		if err := page.Navigate(url); err != nil {
			fmt.Println("Failed to navigate:", err)
		}
	}
}

func YandexElementGetTranslate(text string) string {
	if !FlagYandexElementTransl {
		return "0"
	}
	var value string
	var err error

	var localLang1 string = PasrseLangYandexElement(Lang1)
	if localLang1 == "none" {
		log.Println("Yandex element Language1: ", Lang1, ". Not found.")
		return "0"
	}

	var localLang2 string = PasrseLangYandexElement(Lang2)
	if localLang2 == "none" {
		log.Println("Yandex element Language2: ", Lang2, ". Not found.")
		return "0"
	}

	err = page.RunScript(`
window.MyGolobYandExep = 0;
	window.postMessage({action: 'YandexExtension::Translate', payload: TEXT, from: LANG_FROM, to: LANG_TO }, '*');
	`, map[string]interface{}{"TEXT": text, "LANG_FROM": localLang1, "LANG_TO": localLang2}, &value)
	if err != nil {
		fmt.Println("0Google RunScript: ", err)
	}
	for i := 0; i < (TIME_WAIT_MILLISECOND/TIME_CYCLE)/3; i++ {
		err = page.RunScript(`
		if(window.MyGolobYandExep == 1) {
			var retString = "" + window.GlobalText;
			window.MyGolobYandExep = 0;
			window.GlobalText ='0';
			return retString;
		}
		return "0"
		`, nil, &value)
		if err != nil {
			fmt.Println("Google3 RunScript: ", err)
			break
		}
		if value != "0" {
			log.Println("google(yandExept)(t:" + strconv.Itoa(i*TIME_CYCLE) + "): ")
			break
		}
		time.Sleep(TIME_CYCLE * time.Millisecond)
	}
	return value
}

func GetTranslateGoogle(text string) string {
	var value string
	var ret string
	var err error

	page.SwitchToWindow(allNamePageOpen[0])

	ret = YandexElementGetTranslate(text)
	if ret != "0" && ret != "undefined" {
		return ret
	}

	fmt.Println("text: ", text)

	err = page.RunScript(`
	window.MyGolobalVar = 0;
	var obj = document.querySelector("#source");
	if (!obj) {
		obj = document.querySelector("#yDmH0d > c-wiz > div > div.WFnNle > c-wiz > div.OlSOob > c-wiz > div.ccvoYb > div.AxqVh > div.OPPzxe > c-wiz.rm1UF.UnxENd > span > span > div > textarea");
	}

	obj.value = TEXT;

	var event = new Event('input', {
		'bubbles': true,
		'cancelable': true
	});
	obj.dispatchEvent(event);
	`, map[string]interface{}{"TEXT": text}, &value)
	if err != nil {
		fmt.Println("1Google RunScript: ", err)
		return ""
	}

	for i := 0; i < TIME_WAIT_MILLISECOND/TIME_CYCLE; i++ {
		err = page.RunScript(`
if(window.MyGolobalVar == 2) {
			var retString = "" + window.GlobalText;
			window.MyGolobalVar = 0;
			window.GlobalText='';
			return retString;
		}
		return "0";
		`, nil, &value)
		if err != nil {
			fmt.Println("Google2 RunScript: ", err)
			break
		}
		if value != "0" {
			log.Println("google(t:" + strconv.Itoa(i*TIME_CYCLE) + "): ")
			break
		}
		time.Sleep(TIME_CYCLE * time.Millisecond)
	}
	fmt.Println("valueGoogle: ", value)
	if value != "0" && err == nil {
		fmt.Println("-----0000-1---", value)
		ret = ParseGoogle4(value)
	} else {
		ret = "0"
	}
	return ret
}

func SetEventYandexElement() {
	var value string
	var err error

	err = page.RunScript(`
addEventListener("message", function(event) {
		if (event.data.action == "TranslateMess") {
			console.log("content DONE: ", event.data.payload);
			window.GlobalText = event.data.payload;
			window.MyGolobYandExep = 1;
		}
	});
	return "0";
	`, nil, &value)
	if err != nil {
		log.Println("Error SetYandexElement: ", err)
	}
}

func SetEventOnGoogle(page *agouti.Page) {
	var value string
	var err error

	//SetEventYandexElement()

	err = page.RunScript(`
	window.MyGolobalVar = 0;
	window.GlobalText = "";

	function locationHashChanged() {
		console.log("----004-----");
		window.MyGolobalVar = 1;
	}

	window.onhashchange = locationHashChanged;

	console.log("----000-----")

	let oldXHROpen = window.XMLHttpRequest.prototype.open;
	window.XMLHttpRequest.prototype.open = function(method, url, async, user, password) {
		this.addEventListener('load', function() {
		console.log("----001-----")
			if (this.responseText.indexOf("generic") >= 0) {
			//if (window.MyGolobalVar == 1) {
			console.log("----002-----")
				if (this.responseText.indexOf("[[[") >= 0) {
				console.log("----003-----")
					window.GlobalText = '' + this.responseText
					console.log('load('+MyGolobalVar+'): ' + this.responseText);
					window.MyGolobalVar = 2
				}
			//}
			}
		});

		return oldXHROpen.apply(this, arguments);
	}

	function sleep(ms) {
		return new Promise(resolve => setTimeout(resolve, ms));
	}
	return "0"
	`, nil, &value)

	if err != nil {
		fmt.Println("SetEvent2Google RunScript: ", err)
	}
}

func GetTranslateYandex(text string) string {
	var value string
	var ret string
	var err error
	page.SwitchToWindow(allNamePageOpen[1])

	ret = YandexElementGetTranslate(text)
	if ret != "0" && ret != "undefined" {
		return ret
	}

	err = page.RunScript(`
	window.MyGolobalVar = 0;

	//document.querySelector("#srcText").value = TEXT;
	document.querySelector("#textarea").value = TEXT;
	document.querySelector("#fakeArea").value = TEXT;
	var event = new Event('input', {
		'bubbles': true,
		'cancelable': true
	});
	document.querySelector("#textLayer").dispatchEvent(event);
	return "0";
	`, map[string]interface{}{"TEXT": text}, &value)
	if err != nil {
		fmt.Println("Yandex RunScript: ", err)
	}

	//ScreenShot()

	for i := 0; i < TIME_WAIT_MILLISECOND/TIME_CYCLE; i++ {
		err = page.RunScript(`
		if(window.MyGolobalVar == 1) {
			var retString = "" + window.GlobalText;
			window.GlobalText='';
			window.MyGolobalVar = 0;
			return retString;
		}
		return "0"
		`, nil, &value)
		if err != nil {
			fmt.Println("RunScript: ", err)
			break
		}
		if value != "0" {
			log.Println("yandex (t:" + strconv.Itoa(i*TIME_CYCLE) + "): ")
			break
		}
		time.Sleep(TIME_CYCLE * time.Millisecond)
	}
	ret = value

	fStr := `,"text":["`
	indexStart := strings.Index(value, fStr)
	if indexStart != -1 {
		value = value[indexStart+len(fStr):]
	} else {
		return "0"
	}
	indexEnd := strings.Index(value, `"]`)
	if indexEnd != -1 {
		ret = value[:indexEnd]
	} else {
		return "0"
	}
	return ret
}

func SetEventOnYandex(page *agouti.Page) {
	var value string
	var err error

	SetEventYandexElement()

	err = page.RunScript(`
	window.MyGolobalVar = 0
	window.GlobalText = ""

	let oldXHROpen = window.XMLHttpRequest.prototype.open;
	window.XMLHttpRequest.prototype.open = function(method, url, async, user, password) {
		this.addEventListener('load', function() {
			if (window.MyGolobalVar == 0) {
				if(  this.responseText.indexOf('{"align":')!=-1 ) {
					window.GlobalText = '' + this.responseText
					console.log('load(): ' + this.responseText);
					window.MyGolobalVar = 1
				}
			}
		});

		return oldXHROpen.apply(this, arguments);
	}

	function sleep(ms) {
		return new Promise(resolve => setTimeout(resolve, ms));
	}
	return "0"
	`, nil, &value)
	if err != nil {
		fmt.Println("SetEventYandex RunScript: ", err)
	}
}

func GetTranslatePromt(text string) string {
	return "Promt not implement"
}

func SetEventOnPromt(page *agouti.Page) {
	var value string
	var err error

	err = page.RunScript(`
	window.MyGolobalVar = 0
	window.GlobalText = ""

	return "0"
	`, nil, &value)
	if err != nil {
		fmt.Println("SetEventYandex RunScript: ", err)
	}
}

func GetId() int {
	GlobalId++
	if GlobalId > COUNT_TRANSLATE_SITE {
		GlobalId = 0
	}
	return GlobalId
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTranslate(text string) string {
	var ret string
	text = strings.ReplaceAll(text, ",", "")

	Mu.Lock()
	//delta := MakeTimestamp() - DelayTranslate
	switch GetId() {
	case 0:
		ret = GetTranslateGoogle(text)
		break
	case 1:
		ret = GetTranslateYandex(text)
		break
	case 2:
		ret = GetTranslatePromt(text)
		break
	default:
		ret = GetTranslateGoogle(text)
	}
	//DelayTranslate = MakeTimestamp()
	Mu.Unlock()
	return ret
}

func ParseGoogle2(text string) string {
	in := []byte(text)
	LEN_ARR := len(in)
	var out []byte
	countBrackets := -3

	text = strings.ReplaceAll(text, ")]}'", "")
	fmt.Println(countBrackets)
	var flagOpenBrackets bool = false // flag Open brakets
	var indexOpenBrakets int = 0
	var indexCountBrakets int = 0

	for i := 0; i < LEN_ARR; i++ {
		c := in[i]
		if c == '[' {
			flagOpenBrackets = true
			indexOpenBrakets = i
		} else if c == ']' && flagOpenBrackets {
			flagOpenBrackets = false
			if indexCountBrakets == 2 {
				content := string(in[indexOpenBrakets+1 : i])
				content = strings.ReplaceAll(content, `\"`, "")
				contentArray := strings.Split(content, ",")
				fmt.Println("found [", indexCountBrakets, "]", contentArray[1])
				return contentArray[1]
			}
			indexOpenBrakets = i
			indexCountBrakets++
		}
	}
	return string(out)
}

func PasreGoogle(text string) string {
	in := []byte(text)
	LEN_ARR := len(in)
	var out []byte
	x := 0
	quotes := 0
	countBrackets := -3
	delta := 0
	for i := 0; i < LEN_ARR; i++ {
		c := in[i]
		if c == '[' || c == ']' {
			if c == '[' {
				delta++
			} else {
				delta--
			}
			countBrackets++
			if countBrackets > 3 && delta == 1 {
				break
			}
		}
	}
	countBrackets = countBrackets
	for i := 0; i < LEN_ARR; i++ {
		c := in[i]
		if c == '[' {
			countBrackets--
			x++
		} else if c == ']' {
			countBrackets--
			x--
		} else {
			if x == 3 {
				if c == '"' {
					quotes++
				} else {
					if quotes == 1 {
						out = append(out, c)
					}
				}
			} else {
				quotes = 0
			}
		}
		if countBrackets <= 0 {
			break
		}
	}
	return string(out)
}

func ScreenShot() {
	page.Screenshot("screenshot/" + strconv.Itoa(GlobalScreenShot) + ".jpg")
	GlobalScreenShot++
}

func DetectHeadless() {
	//CreateNewPage(page, "detect_headless", "https://infosimples.github.io/detect-headless/")

	//time.Sleep(1500 * time.Millisecond)
	//CreateNewPage(page, "ads", "https://checkadblock.ru/")
	//ScreenShot()
	CreateNewPage(page, "detect_headless", "https://intoli.com/blog/making-chrome-headless-undetectable/chrome-headless-test.html")
	//ScreenShot()
	CreateNewPage(page, "detect_headless2", "https://infosimples.github.io/detect-headless/")
	time.Sleep(1000 * time.Millisecond)
	//ScreenShot()
}

func SetBrowserLang1(str string) {
	Lang1 = language.MustParse(str).String()
	log.Println("lagparse = ", Lang1, " str = ", str)
}

func SetBrowserLang2(str string) {
	Lang2 = language.MustParse(str).String()
	log.Println("lagparse = ", Lang2, " str = ", str)
}

func RedirectPage() {
	var url string
	var err error

	log.Println("Lang1: ", Lang1, " Lang2: ", Lang2)

	page.SwitchToWindow("googleTranslate")
	url = fmt.Sprintf("https://translate.google.com/#view=home&op=translate&sl=%[1]s&tl=%[2]s",
		PasrseLangGoogle(Lang1), PasrseLangGoogle(Lang2))
	if err = page.Navigate(url); err != nil {
		fmt.Println("Failed google to navigate:", err)
	}
	SetEventOnGoogle(page)

	page.SwitchToWindow("yandexTranslate")
	url = fmt.Sprintf("https://translate.yandex.ru/?lang=%[1]s-%[2]s",
		PasrseLangYandex(Lang1), PasrseLangYandex(Lang2))
	if err = page.Navigate(url); err != nil {
		fmt.Println("Failed yandex to navigate:", err)
	}
	SetEventOnYandex(page)

}

func ChangeLangPromt() error {
	var err error
	var value string
	if scriptChangeLangPromt == nil {
		scriptChangeLangPromt, err = ioutil.ReadFile("./youtube/script_change_lang_promt.js")
		if err != nil {
			log.Println("Error change lang promt: ", err)
			return err
		}
	}

	pLang1 := PasrseLangPromt(Lang1)
	pLang2 := PasrseLangPromt(Lang2)

	if pLang1 != "none" && pLang2 != "none" {
		err = page.RunScript(string(scriptChangeLangPromt),
			map[string]interface{}{"LANG_1": pLang1, "LANG_2": pLang2}, &value)
		if err != nil {
			fmt.Println("3Google RunScript: ", err)
		}
	} else {
		err = errors.New("Promt not set lang pLang1: " + pLang1 + " pLang2: " + pLang2)
	}
	return err
}

func PasrseLangYandexElement(str string) string {
	switch str {
	case "en-US":
		return "en"
	case "ru-RU":
		return "ru"
	case "Bulgarian":
		return "bg"
	case "cs-CZ":
		return "cs"
	case "de-DE":
		return "de"
	case "es-ES":
		return "es"
	case "fr-FR":
		return "fr"
	case "it-IT":
		return "it"
	case "pl-PL":
		return "pl"
	case "ro-RO":
		return "ro"
	case "sr":
		return "sr"
	case "tr-TR":
		return "tr"
	case "uk-UA":
		return "uk"
	default:
	}
	return "none"
}

func PasrseLangGoogle(str string) string {
	switch str {
	case "ru-RU":
		return "ru"
	case "pl-PL":
		return "po"
	case "it-IT":
		return "it"
	case "Norwegian":
		return "no"
	case "Swedish":
		return "sv"
	case "en-US":
		return "en"
	case "fr-FR":
		return "fr"
	case "es-ES":
		return "es"
	case "de-DE":
		return "de"
	default:
	}
	return "none"
}

func PasrseLangYandex(str string) string {
	switch str {
	case "ru-RU":
		return "ru"
	case "pl-PL":
		return "pl"
	case "it-IT":
		return "it"
	case "Norwegian":
		return "no"
	case "Swedish":
		return "sv"
	case "en-US":
		return "en"
	case "fr-FR":
		return "fr"
	case "es-ES":
		return "es"
	case "de-DE":
		return "de"
	default:
	}
	return "none"
}

func PasrseLangPromt(str string) string {
	switch str {
	case "ru-RU":
		return "Русский"
	case "pl-PL":
		return "none"
	case "it-IT":
		return "Итальянский"
	case "Norwegian":
		return "none"
	case "Swedish":
		return "none"
	case "en-US":
		return "Английский"
	case "fr-FR":
		return "Французский"
	case "es-ES":
		return "Испанский"
	case "de-DE":
		return "Немецкий"
	default:
	}
	return "none"
}

/*
func PasrseLangYandexElement(str string) string {
	switch str {
		case "English": return "en"
		case "Russian": return "ru"
		case "Bulgarian": return "bg"
		case "Czech": return "cs"
		case "German": return "de"
		case "Spanish": return "es"
		case "French": return "fr"
		case "Italian": return "it"
		case "Polish": return "pl"
		case "Romanian": return "ro"
		case "Serbian": return "sr"
		case "Turkish": return "tr"
		case "Ukrainian": return "uk"
	default:
	}
	return "en"
}

func PasrseLangGoogle(str string) string {
	switch str {
		case "Russian": return "ru"
		case "Polish": return "po"
		case "Italian":  return "it"
		case "Norwegian": return "no"
		case "Swedish": return "sv"
		case "English": return "en"
		case "CanadianFrench": return "fr"
		case "Spanish": return "es"
		case "German": return "de"
	default:
	}
	return "en"
}

func PasrseLangYandex(str string) string {
	switch str {
		case "Russian": return "ru"
		case "Polish": return "pl"
		case "Italian":  return "it"
		case "Norwegian": return "no"
		case "Swedish": return "sv"
		case "English": return "en"
		case "CanadianFrench": return "fr"
		case "Spanish": return "es"
		case "German": return "de"
	default:
	}
	return "en"
}

func PasrseLangPromt(str string) string {
	switch str {
		case "Russian": return "Русский"
		case "Polish": return "none"
		case "Italian":  return "Итальянский"
		case "Norwegian": return "none"
		case "Swedish": return "none"
		case "English": return "Английский"
		case "CanadianFrench": return "Французский"
		case "Spanish": return "Испанский"
		case "German": return "Немецкий"
	default:
	}
	return "none"
}

case "bn_IN" : return "bn"
case "bs" : return "bs"
case "ca" : return "ca"
case "cs_CZ" : return "cs"
case "cy" : return "cy"
case "da_DK" : return "da"
case "de_DE" : return "de"
case "el_GR" : return "el"
case "en_AU" : return "en"
case "en_GB" : return "en"
case "en_IN" : return "en"
case "en_US" : return "en"
case "es_ES" : return "es"
case "es_US" : return "es"
case "et_EE" : return "et"
case "fi_FI" : return "fi"
case "fil_PH" : return "fi"
case "fr_CA" : return "fr"
case "fr_FR" : return "fr"
case "hi_IN" : return "hi"
case "hr" : return "hr"
case "hu_HU" : return "hu"
case "in_ID" : return "in"
case "it_IT" : return "it"
case "ja_JP" : return "ja"
case "jv_ID" : return "jv"
case "km_KH" : return "km"
case "ko_KR" : return "ko"
case "ku" : return "ku"
case "la" : return "la"
case "nb_NO" : return "nb"
case "ne_NP" : return "ne"
case "nl_NL" : return "nl"
case "pl_PL" : return "pl"
case "pt_BR" : return "pt"
case "pt_PT" : return "pt"
case "ro_RO" : return "ro"
case "ru_RU" : return "ru"
case "si_LK" : return "si"
case "sk_SK" : return "sk"
case "sq" : return "sq"
case "sr" : return "sr"
case "su_ID" : return "su"
case "sv_SE" : return "sv"
case "sw" : return "sw"
case "ta" : return "ta"
case "th_TH" : return "th"
case "tr_TR" : return "tr"
case "uk_UA" : return "uk"
case "vi_VN" : return "vi"
case "yue_HK" : return "yu"
case "zh_CN" : return "zh"
case "zh_TW" : return "zh"
case "bn_BD" : return "bn"
*/

func ParseGoogle3(text string) string {
	in := []byte(text)
	LEN_ARR := len(in)
	var out []byte
	countBrackets := -3

	text = strings.ReplaceAll(text, ")]}'", "")
	fmt.Println(countBrackets)
	var flagOpenBrackets bool = false // flag Open brakets
	var indexOpenBrakets int = 0
	var indexCountBrakets int = 0

	for i := 0; i < LEN_ARR; i++ {
		c := in[i]
		if c == '[' {
			flagOpenBrackets = true
			indexOpenBrakets = i
		} else if c == ']' && flagOpenBrackets {
			flagOpenBrackets = false
			//if indexCountBrakets == 2 {
			//content := string(in[indexOpenBrakets+1 : i])
			content := string(in[indexOpenBrakets:i])
			if strings.Contains(content, `\"`) {
				content = strings.ReplaceAll(content, `\"`, "")
				contentArray := strings.Split(content, ",")
				//fmt.Println("found [", indexCountBrakets, "]", contentArray[1])
				fmt.Println("found [", indexCountBrakets, "]", contentArray)
				//return contentArray[1]
				//}
			}
			indexOpenBrakets = i
			indexCountBrakets++
		}
	}
	return string(out)
}

func ParseGoogle4(text string) string {
	text = strings.ReplaceAll(text, ")]}'", "")
	indexRu := strings.Index(text, `\"ru\"`)

	// Находим секцию где есть ru
	contentArray := strings.Split(text, `\"ru\"`)
	if len(contentArray) != 2 {
		return ""
	}

	// возвращамемся назада и ищем [[
	braketsStart := strings.LastIndex(contentArray[0], "[[") + 2
	braketsEnd := strings.Index(contentArray[1], "]") + indexRu + 1

	textSplit := strings.Split(text[braketsStart:braketsEnd], `",[\"`)

	// Берем из каждой найденной только первые скобки
	var fullText string
	for _, it := range textSplit {
		braketsStart := strings.Index(it, `\"`)
		if braketsStart > 0 {
			fullText += it[0:braketsStart]
		}
	}

	return fullText
}
