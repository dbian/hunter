package main

import (
	"bytes"
	"io/ioutil"

	"fmt"

	"time"

	"strings"

	"github.com/anaskhan96/soup"
	"github.com/dbian/hunter/qtorrent"
	"github.com/dbian/hunter/queue"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/gomail.v2"
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
		return getHtmlSoap(url)
	}
	utf, err := GbkToUtf8([]byte(resp))
	utfStr := string(utf)
	doc := soup.HTMLParse(utfStr)
	//fmt.Println(utfStr)
	return doc
}

func emailNotify(title, links string) {
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("mail_from"))
	m.SetHeader("To", viper.GetString("mail_to"))
	m.SetHeader("Subject", "为您下载了新的科幻电影:"+title)
	m.SetBody("text/html", "链接: "+links)

	d := gomail.NewDialer(
		"smtp.qq.com", viper.GetInt("mail_fromport"),
		viper.GetString("mail_username"),
		viper.GetString("mail_password"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

var Queue = queue.NewLimitQueue()

func fetchLatestMagLinks() {
	h := "http://www.ygdy8.net/"
	doc := getHtmlSoap(h + "html/gndy/oumei/index.html")
	links := doc.FindAll("table", "class", "tbspan")
	for _, link := range links {
		name := link.Find("a").FindNextElementSibling()
		url := h + name.Attrs()["href"]
		//fmt.Println("parse link:", url)
		if m := getMovieInfo(url); m != nil {
			title := name.Text()
			if strings.Index(title, "科幻") == -1 {
				continue
			}
			if Queue.Push(queue.Data{Title: title, Mag: m.magLink}) {
				fmt.Println("downloading:", title)
				qtorrent.AddMagLink(
					m.magLink,
					viper.GetString("qtorrent_username"),
					viper.GetString("qtorrent_password"))
				emailNotify(title, url)
			}
		}
	}
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	fetchLatestMagLinks()

	t := time.NewTicker(time.Minute * 30)
	for range t.C {
		fetchLatestMagLinks()

	}
}

func init() {
	setupConfig()
	qtorrent.SetUrl(viper.GetString("qtorrent_url"))
}
