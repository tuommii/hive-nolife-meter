package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"../../user"
)

func setHeaders(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36")
	req.Header.Set("Cookie", user.Cookie)
}

func createURL(username string) string {
	url := user.BaseURL + username
	if username == user.Login {
		url = user.ProfilePage
	}
	return url
}

func getHTML(username string) []byte {
	req, err := http.NewRequest("GET", createURL(username), nil)
	if err != nil {
		log.Fatal(err)
	}
	setHeaders(req)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func saveToFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	delay, err := strconv.Atoi(os.Getenv("DELAY"))
	if err != nil || delay < 0 {
		delay = 200
	}
	if user.Path == "" {
		fmt.Println("Missing DATA_FOLDER!")
		return
	}
	for _, username := range user.Usernames {
		saveToFile(user.Path+username, getHTML(username))
		fmt.Println(username, "downloaded and saved!")
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	fmt.Println("Done.")
}
