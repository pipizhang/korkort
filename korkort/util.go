package korkort

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func IsFile(filePath string) bool {
	var err error
	var f os.FileInfo

	f, err = os.Stat(filePath)
	if err == nil && !f.IsDir() {
		return true
	} else {
		return false
	}
}

func DownloadImage(imageURL string, localFile string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		log.Println(err)
	}

	useragent, _ := Cfg.GetValue("app", "useragent")
	req.Header.Set("User-Agent", useragent)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(localFile, contents, 0644)
	if err != nil {
		log.Println(err)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetFileExtension(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return filepath.Ext(u.Path)
}
