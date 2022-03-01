package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/**
* @Author: liujunren
* @Date: 2022/3/1 10:54
 */
var defaultHttpClient = http.Client{Timeout: time.Second * 3}

func Get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client(request)
}

func Post(url string, header http.Header, data []byte) ([]byte, error) {

	url = strings.Trim(url, " ")
	newReader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, newReader)
	if err != nil {
		return nil, err
	}
	request.Header = header
	return client(request)
}

type File interface {
	GetFieldName() string
	GetName() string
	GetData() *os.File
}

func PostForm(url string, header http.Header, data url.Values) ([]byte, error) {
	if header == nil {
		header = http.Header{}
	}
	//header.Set("Content-Type","application/x-www-form-urlencoded")
	header.Set("Content-Type","application/x-www-form-urlencoded")
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header=header
	return client(req)
}
func PostFormFile(url string, header http.Header, data map[string]string, files ...File) ([]byte, error) {
	if header == nil {
		header = http.Header{}
	}
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	for _, file := range files {
		if file.GetData() != nil {
			formFile, err := writer.CreateFormFile(file.GetFieldName(), file.GetName())
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(formFile, file.GetData())
			if err != nil {
				return nil, err
			}
			header.Set("content-type", writer.FormDataContentType())
		}
	}

	for k, v := range data {
		err := writer.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}
	writer.Close()
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}
	request.Header = header

	return client(request)
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	defer bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(resp_body))

	return nil
}

func client(r *http.Request) ([]byte, error) {
	res, err := defaultHttpClient.Do(r)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request err,http code:%d", res.StatusCode)
	}

	defer res.Body.Close()
	all, err := ioutil.ReadAll(res.Body)
	return all, err
}
