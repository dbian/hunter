package qtorrent

import (
	"fmt"

	"bytes"
	"mime/multipart"

	"github.com/franela/goreq"
)

var Home string

func SetUrl(u string) {
	Home = u
}

// Creates a new file upload http request with optional extra params
func newMultipartRequest(params map[string]string) (*bytes.Buffer, string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()
	return body, writer.FormDataContentType()
}

func AddMagLink(url, user, pwd string) {
	if !addMagLink(url, user, pwd) {
		addMagLink(url, user, pwd)
	}
}

func addMagLink(l string, username, password string) bool {
	r := goreq.Request{
		Uri:         Home + "/login",
		Method:      "POST",
		Body:        fmt.Sprintf(`username=%s&password=%s`, username, password),
		ContentType: "application/x-www-form-urlencoded",
	}
	r.AddHeader("Referer", Home)
	res, err := r.Do()
	if err != nil {
		fmt.Println(err)
		return false
	}

	v, _ := res.Body.ToString()
	if v != "Ok." {
		fmt.Println("server response invalid:", v)
		return false
	}

	r = goreq.Request{
		Uri:    Home + "/command/download",
		Method: "POST",
	}
	r.AddCookie(res.Cookies()[0])
	para := map[string]string{
		"urls":        l,
		"savepath":    "/home/pi/Downloads",
		"category":    "movies",
		"root_folder": "true",
	}
	r.Body, r.ContentType = newMultipartRequest(para)

	res, err = r.Do()
	if err != nil {
		fmt.Println("err:", err)
		return false
	}
	return true
}
