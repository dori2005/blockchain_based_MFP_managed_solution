package main

import ( //몰라 겁나 인포트

	"fmt"       //입출력
	"io/ioutil" //편한 입출력 유틸리티 지원
	"net/http"  //네트웍 프로그래밍을 위함
	"strconv"
	"strings" //

	"github.com/antchfx/htmlquery"
)

type PrinterInfo struct {
	Id        string `json:"id"` //프린터 ID
	Ip        string `json:"url"` //프린터 IP 데이터
	Black     int    `json:"black,string"`
	Cyan      int    `json:"cyan,string"`
	Magenta   int    `json:"magenta,string"`
	Yellow    int    `json:"yellow,string"`
	Drum      int    `json:"drum,string"`
	ErrorCode int    `json:"Errorcode,string`
	Paper     int    `json:"Paer,string`
}

func getHtml(Ip string, num int) (string, error) { //Ip를 받아 url로 변경하여 페이지 OPEN
	var url string
	if num == 1 { //stat
		url = "http://" + Ip + "/status/statgeneral.htm" //http 모듈을 이용하여 웹 반환 데이터 가져옴
	} else if num == 2 { //error
		url = "http://" + Ip + "/status/statevent.htm"
	} else if num == 3 { //toner
		url = "http://" + Ip + "/status/statsupl.htm"
	} else if num == 4 { //paper
		url = "http://" + Ip + "/setting/counter.htm"
	}

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close() //Get으로 요청하여 서버로부터 응답을 받으면 바디를 닫아줘야 함

	body, err := ioutil.ReadAll(res.Body) //반환된 바디를 분리

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil //분리된 바디를 반환
}

func numSlice(e string) int { //문자열 중 숫자만 추출하여, INT 형으로 변환
	if e == "Replace Soon" {
		return 0
	}
	if e[2] == ' ' {
		e = e[:2]
	} else if e[3] == ' ' {
		e = e[:3]
	}
	i, _ := strconv.Atoi(e)
	return i
}

func CrawlerStat(printerVal *PrinterInfo) bool {

	//
	// 상태 크롤링
	//
	html, err := getHtml(printerVal.Ip, 1) //stat
	if err != nil {
		return true
	}

	doc, err := htmlquery.Parse(strings.NewReader(html)) //html쿼리 파서로 분석

	list := htmlquery.Find(doc, "//tr/td/a[@href]")
	if strings.HasPrefix(htmlquery.InnerText(list[0]), "Ready to print") {
		return true
	} else {
		return false
	}
}

func CrawlerAll(printerVal *PrinterInfo) (string, error) {

	//
	// 토너량 크롤링
	//
	html, err := getHtml(printerVal.Ip, 3) //사이트 바디 분리해서 받아왔다 (스트링)

	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(html)) //html쿼리 파서로 분석

	list := htmlquery.Find(doc, "//td/table/tbody/tr/td")

	numSlice(htmlquery.InnerText(list[8]))
	numSlice(htmlquery.InnerText(list[12]))
	numSlice(htmlquery.InnerText(list[16]))
	numSlice(htmlquery.InnerText(list[20]))
	numSlice(htmlquery.InnerText(list[24]))

	printerVal.Cyan = numSlice(htmlquery.InnerText(list[8]))
	printerVal.Magenta = numSlice(htmlquery.InnerText(list[12]))
	printerVal.Yellow = numSlice(htmlquery.InnerText(list[16]))
	printerVal.Black = numSlice(htmlquery.InnerText(list[20]))
	printerVal.Drum = numSlice(htmlquery.InnerText(list[24]))

	//
	// 오류정보 크롤링
	//
	html, err = getHtml(printerVal.Ip, 2) //사이트 바디 분리해서 받아왔다 (스트링)

	if err != nil {
		return html, err
	}

	doc, err = htmlquery.Parse(strings.NewReader(html)) //html쿼리 파서로 분석

	list = htmlquery.Find(doc, "//tr/td/font[@size]")

	listLen := len(list)
	stat := 0
	for i := 4; i < listLen; i += 3 {
		text := htmlquery.InnerText(list[i])
		if strings.HasPrefix(text, "091-402") {
			stat = stat + 1
		}
		if strings.HasPrefix(text, "Out of paper.") {
			stat = stat + 2
		}
		if strings.HasPrefix(text, "Power Saver.") {
			stat = stat + 4
		}
	}
	printerVal.ErrorCode = stat

	//
	// 총 사용 종이량
	//
	html, err = getHtml(printerVal.Ip, 4) //사이트 바디 분리해서 받아왔다 (스트링)

	if err != nil {
		return html, err
	}

	doc, err = htmlquery.Parse(strings.NewReader(html)) //html쿼리 파서로 분석

	list = htmlquery.Find(doc, "//td/table/tbody/tr[1]/td[2]")
	text := htmlquery.InnerText(list[0])
	telen := len(text)
	printerVal.Paper, _ = strconv.Atoi(text[:telen-1])

	return html, err
}
