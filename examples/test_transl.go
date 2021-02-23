package main

import (
	"fmt"
	"strings"

	"github.com/exitstop/speaker/internal/browser"
)

func main() {
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
		`
)]}

[[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,34]\n]\n,[true]\n]\n]\n]\n,34]\n]\n,[[[null,\"Pozhaluysta, nazhmite ctrl + c, chtoby ostanovit' khuk\",null,null,null,[[\"Пожалуйста, нажмите ctrl + c, чтобы остановить хук\",[\"Пожалуйста, нажмите ctrl + c, чтобы остановить хук\",\"Пожалуйста, нажмите CTRL + C, чтобы остановить крюк\"]\n]\n]\n]\n]\n,\"ru\",1,\"auto\",[\"Please press ctrl + c to stop hook\",\"auto\",\"ru\",true]\n]\n,\"en\"]\n",null,null,null,"generic"]
,["di",24]
,["af.httprm",23,"-3626247769773888654",10]
,["e",4,null,null,670]
]]
`,
		`load(0): )]}'

898
[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,105]\n]\n,[true]\n]\n]\n]\n,105]\n]\n,[[[null,\"Vy takzhe mozhete ukazat' JSHandle v kachestve znacheniya svoystva, yesli khotite, chtoby zhivyye ob\\\"yekty peredavalis' v sobytiye:\",null,null,null,[[\"Вы также можете указать JSHandle в качестве значения свойства, если хотите, чтобы живые объекты передавались в событие:\",[\"Вы также можете указать JSHandle в качестве значения свойства, если хотите, чтобы живые объекты передавались в событие:\",\"Вы также можете указать JSHandle в качестве значения свойства, если вы хотите живые объекты должны быть переданы в случае:\"]\n]\n]\n]\n]\n,\"ru\",1,\"auto\",[\"You can also specify JSHandle as the property value if you want live objects to be passed into the event:\",\"auto\",\"ru\",true]\n]\n,\"en\"]\n",null,null,null,"generic"]
,["di",32]
,["af.httprm",31,"-3457224819668688087",266]
]
27
[["e",4,null,null,1214]
]`,
		`load(0): )]}'

3693
[["wrb.fr","MkEWBc","[[null,null,\"en\"]\n,[[[null,\"Nasledovaniye, khotya i polezno, imeyet svoi podvodnyye kamni. Eto chasto privodit k iyerarkhii klassov, a inogda povedeniye konechnogo ob\\\"yekta raspredelyayetsya po iyerarkhii. V iyerarkhii nasledovaniya superklassy chasto mogut byt' khrupkimi, potomu chto odno nebol'shoye izmeneniye superklassa mozhet povliyat' na mnogiye drugiye mesta v kode prilozheniya. Odna iz luchshikh veshchey, kotoryye mogut proizoyti, - eto skompilirovannyye yazyki s oshibkoy vremeni kompilyatsii, no deystvitel'no slozhnyye situatsii - eto te, gde net oshibok vremeni kompilyatsii, no tonkiye izmeneniya povedeniya privodyat k oshibkam v dopolnitel'nykh stsenariyakh. V kontse kontsov, takiye veshchi mozhet byt' ochen' slozhno otladit', v vashem kode nichego ne izmenilos'.Net prostogo sposoba ulovit' eto v takikh protsessakh, kak proverka koda, poskol'ku po zamyslu bazovyy klass i razrabotchiki, kotoryye podderzhivayut eto, ne zabotyatsya ili ne znayut o proizvodnykh klassakh. Al'ternativoy nasledovaniyu yavlyayetsya delegirovaniye povedeniya, takzhe nazyvayemogo kompozitsiyey. Vmesto is a eto imeyet otnosheniye. Eto otnositsya k ob\\\"yedineniyu prostykh tipov v boleye slozhnyye. Predydushchiye otnosheniya s zhivotnymi modeliruyutsya sleduyushchim obrazom:\",null,null,null,[[\"Наследование, хотя и полезно, имеет свои подводные камни. Это часто приводит к иерархии классов, а иногда поведение конечного объекта распределяется по иерархии. В иерархии наследования суперклассы часто могут быть хрупкими, потому что одно небольшое изменение суперкласса может повлиять на многие другие места в коде приложения. Одна из лучших вещей, которые могут произойти, - это скомпилированные языки с ошибкой времени компиляции, но действительно сложные ситуации - это те, где нет ошибок времени компиляции, но тонкие изменения поведения приводят к ошибкам в дополнительных сценариях. В конце концов, такие вещи может быть очень сложно отладить, в вашем коде ничего не изменилось.Нет простого способа уловить это в таких процессах, как проверка кода, поскольку по замыслу базовый класс и разработчики, которые поддерживают это, не заботятся или не знают о производных классах. Альтернативой наследованию является делегирование поведения, также называемого композицией. Вместо is a это имеет отношение. Это относится к объединению простых типов в более сложные. Предыдущие отношения с животными моделируются следующим образом:\"]\n]\n]\n]\n,\"ru\",1,\"auto\",[\"Inheritance , though useful , has its pitfalls . It often leads to a hierarchy of classes , and sometimes the behavior of the final object is spread across the hierarchy . In an inheritance hierarchy , super classes can often be fragile , because one little change to a superclass can ripple out and affect many other places in the application s code . One of the better things that can happen is a compile time error compiled languages , but the really tricky situations are those where there are no compile time errors , but , subtle behavior changes leading to errors bugs in fringe scenarios . Such things can be really hard to debug after all , nothing changed in your code There is no easy way to catch this in processes such as code review , since , by design , the base class and the developers who maintain that don t care or know about the derived classes . An alternative to inheritance is to delegate behavior , also called composition . Instead of an is a , this is a has a relationship . It refers to combining simple types to make more complex ones . The preceding Animal relationship is modeled as the following\",\"auto\",\"ru\",true]\n]\n,\"en\"]\n",null,null,null,"generic"]
]
57
[["di",31]
,["af.httprm",29,"5267913587000612704",10]
]
27
[["e",4,null,null,4737]
]
`,
		`load(0): )]}'

3053
[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,142]\n]\n,[true]\n]\n]\n,[1,[[[null,143]\n,[143,274]\n]\n,[false,true]\n]\n]\n,[2,[[[null,275]\n,[275,469]\n]\n,[false,true]\n]\n]\n]\n,469]\n]\n,[[[null,\"Chto takoye ugadayte khesh besplatnoy programmy gornodobyvayushchey kriptovalyuty, blagodarya kotoromu pol'zovatel' mozhet poluchit' neogranichennoye kolichestvo BTC iz odnogo besplatnogo lotereynogo bileta. Podrobneye Pol'zovateli Boleye krupnyye dzhekpoty My sozdayem basseyn premii Dzhekpota cherez sovmestnuyu rabotu vsekh videokart vsekh nashikh uchastnikov. Chtoby vyigrat' dzhekpot, vam nuzhno prosto dogadat'sya 4 poslednikh kolichestva sluchaynykh bitkoynovskoy tranzaktsii, i da, eto tak prostoye data rozygrysha 30.04.2021 Tekushchiy dzhekpot 110 000 Tekushchiye uchastniki № 1030\",null,null,null,[[\"Что такое угадайте хэш бесплатной программы горнодобывающей криптовалюты, благодаря которому пользователь может получить неограниченное количество BTC из одного бесплатного лотерейного билета.\",[\"Что такое угадайте хэш бесплатной программы горнодобывающей криптовалюты, благодаря которому пользователь может получить неограниченное количество BTC из одного бесплатного лотерейного билета.\",\"Что Guess Hash Программа добыча бесплатно криптовалюта, благодаря которому пользователь может получить неограниченное количество BTC от одного бесплатного лотерейного билета.\"]\n]\n,[\"Подробнее Пользователи Более крупные джекпоты Мы создаем бассейн премии Джекпота через совместную работу всех видеокарт всех наших участников.\",[\"Подробнее Пользователи Более крупные джекпоты Мы создаем бассейн премии Джекпота через совместную работу всех видеокарт всех наших участников.\",\"Больше пользователей превышают джекпоты Мы создаем призовой джекпот через совместную работу всех видеокарт все наши участников.\"]\n,true]\n,[\"Чтобы выиграть джекпот, вам нужно просто догадаться 4 последних количества случайных биткойновской транзакции, и да, это так простое дата розыгрыша 30.04.2021 Текущий джекпот 110 000 Текущие участники № 1030\",[\"Чтобы выиграть джекпот, вам нужно просто догадаться 4 последних количества случайных биткойновской транзакции, и да, это так простое дата розыгрыша 30.04.2021 Текущий джекпот 110 000 Текущие участники № 1030\",\"Для того, чтобы выиграть джек-пот вам нужно просто угадать 4 последних numders случайной сделки Bitcoin, и да, это так просто Дата розыгрыша 30.04.2021 текущий джекпот 110000 текущие участники номер 1030\"]\n,true]\n]\n]\n]\n,\"ru\",1,\"auto\",[\"What is Guess Hash A free cryptocurrency mining program , thanks to which user can receive unlimited BTC amount from one free lottery ticket . More users bigger Jackpots We are creating a jackpot prize pool through the joint work of all video cards of all our participants . To win the jackpot you need just to guess 4 last numders of a random bitcoin transaction , and yes , it is so simple draw date 30.04.2021 current jackpot 110,000 current participants number 1030\",\"auto\",\"ru\",true]\n]\n,\"en\"]\n",null,null,null,"generic"]
]
59
[["di",194]
,["af.httprm",193,"-634869799261197867",10]
]
27
[["e",4,null,null,4436]
]
`,
	}
	for _, it := range input {
		ret, _ := browser.ParseGoogle5(it)
		fmt.Println(ret)
		fmt.Println("--------------")
	}
	fmt.Println("################################")
	fmt.Println("################################")
	for _, it := range input {
		ret, _ := browser.ParseGoogle5(it)
		fmt.Println("FinalReusl: ", ret)
		fmt.Println("--------------")
	}
}

/*

)]}'

[[["wrb.fr","MkEWBc","[[null,null,\"en\",[[[0,[[[null,14]\n]\n,[true]\n]\n]\n]\n,14]\n]\n,[[[null,\"poshel na khuy, suka\",null,null,null,[[\"пошел на хуй, сука\",[\"пошел на хуй, сука\",\"пошел на хуй сука\"]\n]\n]\n]\n]\n,\"ru\",1,\"auto\"]\n,\"en\"]\n",null,null,null,"generic"]
,["di",39]
,["af.httprm",38,"3520281249244905365",12]
,["e",4,null,null,409]
]]

)]}'

[[["wrb.fr","MkEWBc","[[null,[[[null,\"Ok i found the answer, the javascript code was called on from inside a iframe and \\u003cb\\u003e\\u003ci\\u003eapparently\\u003c/i\\u003e\\u003c/b\\u003e, they act just like a separate html page. I used this code to solve it.\"]\n,null,2,[1]\n]\n]\n,\"en\",[[[0,[[[null,132]\n]\n,[true]\n]\n]\n,[1,[[[null,133]\n,[133,162]\n]\n,[false,true]\n]\n]\n]\n,162]\n]\n,[[[null,\"Khorosho, ya nashel otvet, kod javascript byl vyzvan iznutri iframe i, po suti, oni deystvuyut kak otdel'naya stranitsa html. YA ispol'zoval etot kod, chtoby reshit' etu problemu.\",null,null,null,[[\"Хорошо, я нашел ответ, код javascript был вызван изнутри iframe и, по сути, они действуют как отдельная страница html.\",[\"Хорошо, я нашел ответ, код javascript был вызван изнутри iframe и, по сути, они действуют как отдельная страница html.\",\"Ok я нашел ответ, Javascript код был вызван из внутри фрейма и appertly, они действуют так же, как отдельный HTML-страницы.\"]\n]\n,[\"Я использовал этот код, чтобы решить эту проблему.\",[\"Я использовал этот код, чтобы решить эту проблему.\"]\n,true]\n]\n]\n]\n,\"ru\",1,\"auto\"]\n,\"en\"]\n",null,null,null,"generic"]
,["di",155]
,["af.httprm",155,"-8985846247477363830",12]
,["e",4,null,null,1578]
]]
*/

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

func ParseGoogle(text string) string {
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
