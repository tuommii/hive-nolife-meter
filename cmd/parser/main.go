package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"../../user"
	"golang.org/x/net/html"
)

// LevelToFloat parses string which comes format like "level 5 - 2%" or "level 5 - 40%"
// to float. So 5.02 and 5.40
func LevelToFloat(str string) float64 {
	var nums int
	var res string
	var level float64

	cut := " level%"
	arr := strings.Split(str, "-")
	if len(arr) != 2 {
		return 0.00
	}
	arr[0] = strings.Trim(arr[0], cut)
	arr[1] = strings.Trim(arr[1], cut)
	for _, rune := range arr[1] {
		if unicode.IsNumber(rune) {
			nums++
		}
	}
	res = arr[0]
	res += "."
	if nums == 1 {
		res += "0"
		res += arr[1]
	}
	if nums == 2 {
		res += arr[1]
	}
	level, err := strconv.ParseFloat(res, 64)
	if err != nil {
		level = 0.00
	}
	return level
}

// GetLevelString returns string that contains user's level
func GetLevelString(r io.Reader) string {
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			fmt.Printf("Error while parsing HTML: %v\n", tokenizer.Err())
			return ""
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "div" {
				for _, a := range token.Attr {
					if a.Val == "on-progress" {
						tokenizer.Next()
						t := tokenizer.Token()
						return t.Data
					}
				}
			}
		}
	}
	return ""
}

func showUsernames() bool {
	if os.Getenv("SHOW_USERNAMES") == "TRUE" {
		return true
	}
	return false
}

// FormatLevel ...
func FormatLevel(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

type jee func(float64) string

func main() {
	for _, username := range user.Usernames {
		r, err := os.Open(user.Path + username)
		if err != nil {
			fmt.Println("Could not open", user.Path+username)
			continue
		}
		lvl := GetLevelString(r)
		num := LevelToFloat(lvl)
		usr := user.User{Name: username, Level: num, LevelString: ""}
		usr.LevelString = fmt.Sprintf("%.2f", num)
		user.List = append(user.List, usr)
	}
	sort.Slice(user.List, func(i, j int) bool { return user.List[i].Level > user.List[j].Level })
	for i, user := range user.List {
		fmt.Printf("%3d. %-9s %.2f\n", i+1, user.Name, user.Level)
	}
	templ := template.Must(template.ParseFiles(os.Getenv("TEMPLATE")))
	t := time.Now()
	data := struct {
		PageTitle     string
		UserList      []user.User
		Time          string
		ShowUsernames bool
	}{
		"Nolife-Meter",
		user.List,
		t.Format(time.RFC850),
		showUsernames(),
	}
	f, err := os.Create(os.Getenv("INDEX"))
	if err != nil {
		log.Fatal("Error while creating HTML-file!")
	}
	defer f.Close()
	templ.Templates()[0].Execute(f, data)
	if err != nil {
		fmt.Println(err)
	}
}
