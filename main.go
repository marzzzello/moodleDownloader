package moodleDownloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/marzzzello/moodleAPI"
)

type MoodleDownloader struct {
	baseURL string
	token   string

	log   MoodleLogger
	fetch LookupUrl
}

func NewMoodleDownloader(baseURL string, token string) *MoodleDownloader {
	return &MoodleDownloader{
		baseURL: baseURL,
		token:   token,
	}
}

func getUserID(api *moodleAPI.MoodleApi) (int64, error) {
	_, _, _, userID, err := api.GetSiteInfo()
	if err != nil {
		return 0, err
	} else {
		return userID, nil

	}
}

func printInfo(api *moodleAPI.MoodleApi) {
	sitename, firstname, lastname, userID, err := api.GetSiteInfo()
	if err != nil {
		//fmt.Println("Error getting site info:", err.Error())
		println(readError(err.Error()))
	} else {
		fmt.Println("Sitename:", sitename)
		fmt.Println("Firstname:", firstname)
		fmt.Println("Lastname:", lastname)
		fmt.Println("UserID:", userID)
	}
}

func getToken(settingsFile string) string {

	type Settings struct {
		BaseURL  string `json:"baseURL"`
		Username string `json:"username"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open(settingsFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened settings.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Settings array
	var settings Settings
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'settings' which we defined above
	json.Unmarshal(byteValue, &settings)
	fmt.Println("Settings", settings)

	if settings.Token != "" {
		//check if Token is valid
		//TODO
		return settings.Token
	}

	//TODO: Get new Token, and save into settings.json
	/*
		Get token with curl 'https://moodle.ruhr-uni-bochum.de/m/login/token.php' --data 'username=YOUR_USERNAME&password=YOUR_PASSWORD&service=moodle_mobile_app'
		normally the token is valid for one month
	*/
	fmt.Println("Token", settings.Token)
	fmt.Println("Settings", settings)

	return settings.Token
}
