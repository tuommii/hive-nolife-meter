package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// User ...
type User struct {
	Name        string
	Level       float64
	LevelString string
}

// ...
var (
	List        []User
	Path        string
	Login       string
	Cookie      string
	ProfilePage string
	BaseURL     string
)

// Usernames ...
var Usernames []string

func init() {
	Path = os.Getenv("DATA_FOLDER")
	Login = os.Getenv("LOGIN")
	Cookie = os.Getenv("COOKIE")
	BaseURL = "https://profile.intra.42.fr/users/"
	ProfilePage = "https://profile.intra.42.fr/"
	readUsersFromFile()
}

func readUsersFromFile() {
	jsonFile, err := os.Open(os.Getenv("USERS"))
	if err != nil {
		log.Fatal("users.json missing!")
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("error while parsing users.json!")
	}
	json.Unmarshal(bytes, &Usernames)
}
