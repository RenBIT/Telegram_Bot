package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gocolly/colly"
)

type Mp3 struct {
	Title string
	File  string
	Time  int64
}

var url_mp3 Mp3

func ParseSiteMp3() {

	c := colly.NewCollector()

	// Ищим ид и в нем скрипт
	c.OnHTML("#player script", func(e *colly.HTMLElement) {
		//	fmt.Println(e.Text)
		//Вытащим ссылку на mp3 из скрипта, регуляркой
		re, _ := regexp.Compile(`(?:https?:\/\/)?(?:[\w\.]+)\.(?:[a-z]{2,6}\.?)(?:\/[\w\.]*)*\/?`)
		res := re.FindAllString(e.Text, -1)

		url_mp3.File = res[0]
		//fmt.Println(mp3.File) // false
		//e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://zvukiknig.net/skazki/30098-vsex-prekrasnej-istoriya-korolevy.html")

}

// Проверим размер файла с сайта
func getSizeFile(mp3 string) int {
	if mp3 == "" {
		return 0
	}
	rsp, err := http.Get(mp3)
	if err != nil {
		fmt.Println("HEAD request failed", err)
		defer rsp.Body.Close()
		return 0
	} else {
		// по-хорошему, тут надо обработать статус запроса
		defer rsp.Body.Close()
		return int(rsp.ContentLength)
	}
}
