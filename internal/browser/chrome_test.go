package browser

import (
	"fmt"
	"testing"

	"github.com/exitstop/speaker/internal/browser"
)

// go test internal/browser/chrome_test.go -run TestBrowser

func TestBrowser(t *testing.T) {
	input := []string{`
)]}'

[[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,14]\n]\n,[true]\n]\n]\n]\n,14]\n]\n,[[[null,\"poshel na khuy, suka\",null,null,null,[[\"пошел на хуй, сука\",[\"пошел на хуй, сука\",\"пошел на хуй сука\"]\n]\n]\n]\n]\n,\"ru\",1,\"auto\"]\n,\"en\"]\n",null,null,null,"generic"]
,["di",39]
,["af.httprm",38,"3520281249244905365",12]
,["e",4,null,null,409]
]]

`,
		`
)]}'

[[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,27]\n]\n,[true]\n]\n]\n]\n,27]\n]\n,[[[null,\"name \\u003d com.google.android.tts\",null,null,null,[[\"name \\u003d com.google.android.tts\",[\"name \\u003d com.google.android.tts\",\"Имя \\u003d com.google.android.tts\"]\n]\n]\n]\n]\n,\"ru\",1,\"auto\"]\n,\"en\"]\n",null,null,null,"generic"]
,["di",100]
,["af.httprm",98,"-5593638297484965670",20]
,["e",4,null,null,438]
]]
`,
		`
)]}'

[[["wrb.fr","QShL0","[[\"Splits the slices of s into all substrings separated by sep and returns a slice of the substrings between those separators.\",\"Split the slices are found in all substrings, split and return the sen with a slice of substrings between those delimiters.\"]\n]\n",null,null,null,"generic"]
,["di",133]
,["af.httprm",131,"6058031296088600526",17]
,["e",4,null,null,399]
]]
`,
		`
)]}'

[[["wrb.fr","MkEWBc","[[null,[[[null,\"Ok i found the answer, the javascript code was called on from inside a iframe and \\u003cb\\u003e\\u003ci\\u003eapparently\\u003c/i\\u003e\\u003c/b\\u003e, they act just like a separate html page. I used this code to solve it.\"]\n,null,2,[1]\n]\n]\n,\"en\",[[[0,[[[null,132]\n]\n,[true]\n]\n]\n,[1,[[[null,133]\n,[133,162]\n]\n,[false,true]\n]\n]\n]\n,162]\n]\n,[[[null,\"Khorosho, ya nashel otvet, kod javascript byl vyzvan iznutri iframe i, po suti, oni deystvuyut kak otdel'naya stranitsa html. YA ispol'zoval etot kod, chtoby reshit' etu problemu.\",null,null,null,[[\"Хорошо, я нашел ответ, код javascript был вызван изнутри iframe и, по сути, они действуют как отдельная страница html.\",[\"Хорошо, я нашел ответ, код javascript был вызван изнутри iframe и, по сути, они действуют как отдельная страница html.\",\"Ok я нашел ответ, Javascript код был вызван из внутри фрейма и appertly, они действуют так же, как отдельный HTML-страницы.\"]\n]\n,[\"Я использовал этот код, чтобы решить эту проблему.\",[\"Я использовал этот код, чтобы решить эту проблему.\"]\n,true]\n]\n]\n]\n,\"ru\",1,\"auto\"]\n,\"en\"]\n",null,null,null,"generic"]
,["di",155]
,["af.httprm",155,"-8985846247477363830",12]
,["e",4,null,null,1578]
]]
`,
	}
	//got := Abs(-1)
	//if got != 1 {
	//t.Errorf("Abs(-1) = %d; want 1", got)
	//}

	for _, it := range input {
		ret := browser.ParseGoogle4(it)
		fmt.Println(ret)
		fmt.Println("--------------")
	}
	t.Errorf("error")

	// Output: hello
	return
}
