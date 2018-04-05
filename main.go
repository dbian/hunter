package main

import (
	"os"

	"bytes"
	"io/ioutil"

	"fmt"

	"github.com/anaskhan96/soup"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

type info struct {
	magLink string
}

func getMovieInfo(url string) *info {
	m := &info{}
	s := getHtmlSoap(url)
	r := s.FindAll("td", "bgcolor", "#fdfddf")

	if len(r) < 2 {
		return nil
	}

	m.magLink = r[1].Find("a").Attrs()["href"]
	return m
}

func getHtmlSoap(url string) soup.Root {
	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println("can not fetch url", url)
		os.Exit(1)
	}
	utf, err := GbkToUtf8([]byte(resp))
	utfStr := string(utf)
	doc := soup.HTMLParse(utfStr)
	//fmt.Println(utfStr)
	return doc
}

func fetchLatestMagLinks() {
	h := "http://www.ygdy8.net/"
	doc := getHtmlSoap(h + "html/gndy/oumei/index.html")
	links := doc.FindAll("table", "class", "tbspan")
	for _, link := range links {
		name := link.Find("a").FindNextElementSibling()
		url := h + name.Attrs()["href"]
		//fmt.Println("parse link:", url)
		if m := getMovieInfo(url); m != nil {
			fmt.Println(name.Text())
			fmt.Println(m.magLink)
		}
	}
}

func main() {
	fetchLatestMagLinks()
}
