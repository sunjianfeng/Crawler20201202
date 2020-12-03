package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetMovie(url string) {
	fmt.Println(url)
	//new 一个 request，再设置其header
	req, _ := http.NewRequest("GET", url, nil)
	// 设置
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	//

	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		// name
		fmt.Println("name:" + s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text())
		// year
		fmt.Println("year:" + s.ChildrenFiltered(`.year`).Text())
	})

	// #info > span:nth-child(1) > span.attrs
	director := ""
	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		// 导演
		director += s.Text()
		//fmt.Println(s.Text())
	})
	fmt.Println("导演:" + director)
	//fmt.Println("\n")

	pl := ""
	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		pl += s.Text()
	})
	fmt.Println("编剧:" + pl)

	charactor := ""
	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		charactor += s.Text()
	})
	fmt.Println("主演:" + charactor)

	typeStr := ""
	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		typeStr += s.Text()
	})
	fmt.Println("类型:" + typeStr)
}

func GetToplist(url string) []string {
	var urls []string
	//new 一个 request，再设置其header
	req, _ := http.NewRequest("GET", url, nil)
	// 设置
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("===============================================================", resp.StatusCode)
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("//////////////////////////////////////", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("#content div div.article ol li div div.info div.hd a").
		Each(func(i int, s *goquery.Selection) {
			// year
			fmt.Printf("%v", s)
			herf, _ := s.Attr("href")
			urls = append(urls, herf)
		})
	return urls
}

func main() {
	url := "https://movie.douban.com/top250?start=0"
	var urls []string
	urls = GetToplist(url)
	fmt.Println("%v", urls)
	for _, url := range urls {
		GetMovie(url)
	}
}
