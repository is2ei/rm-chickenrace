package main

import (
	"fmt"
	"path/filepath"
	"os"
	"net/url"
	"net/http"
	"flag"
	"strings"
	"io/ioutil"
	"github.com/ChimeraCoder/anaconda"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	msg := fmt.Sprintln(strings.Join(os.Args, " "))
	msg += GetTotalEnrtyInfoMessage(args)

	err := DeleteAll(args)
	if err != nil {
		msg += "\n\nError!!\n" + err.Error()
	}
	fmt.Println(msg)

	if Exists("/rm.typetalk.conf") {
		s := loadSetting("/rm.typetalk.conf")
		url := s[0]
		fmt.Println(url)
		PostTypetalk(msg, url)
	}
	if Exists("/rm.twitter.conf") {
		s := loadSetting("/rm.twitter.conf")
		consumerKey := s[0]
		consumerSecret := s[1]
		accessToken := s[2]
		accessTokenSecret := s[3]
		PostTwitter(msg, consumerKey, consumerSecret, accessToken, accessTokenSecret)
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func loadSetting(path string) []string {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return strings.Split(string(text),"\n")
}

type EntryInfo struct {
	directoryCount int64
	fileCount int64
	totalSize int64
}

func NewEntryInfo() *EntryInfo {
	return &EntryInfo{
		directoryCount: int64(0),
		fileCount: int64(0),
		totalSize: int64(0),
	}
}

func (e1 *EntryInfo) Add(e2 *EntryInfo) {
	e1.directoryCount += e2.directoryCount
	e1.fileCount += e2.fileCount
	e1.totalSize += e2.totalSize
}

func DeleteAll(args[] string) error {
	for _, a := range args {
		err := os.RemoveAll(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTotalEnrtyInfoMessage(args[] string) string {
	e := EntryInfo{}
	for _, a := range args {
		e.Add(GetEntryInfo(a))
	}
	return fmt.Sprintf("  directories: %d\n  files: %d\n  total size: %d byte", e.directoryCount, e.fileCount, e.totalSize)
}

func GetEntryInfo(root string) *EntryInfo {
	e := NewEntryInfo()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			e.directoryCount += 1
		} else {
			e.fileCount += 1
		}
		e.totalSize += info.Size()
		return nil
	})
	if err != nil {
		return NewEntryInfo()
	}
	return e
}

func PostTypetalk(msg string, apiUrl string) {
	values := url.Values{}
	values.Add("message", msg)
	resp, err := http.PostForm(apiUrl, values)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
}

func PostTwitter(msg string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) {
	hashtag := "#rmコマンドチキンレース"
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	text := msg + "\n" + hashtag
	tweet, err := api.PostTweet(text, nil)
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(tweet.Text)
}