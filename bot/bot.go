package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"os"
)

type Configuration struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func CreateTweet() string {
	puzzle := sudoku.CreatePuzzle(17) // 17 is the min givens for a unique solution
	var buffer bytes.Buffer
	for i := 0; i < len(puzzle); i += 9 {
		row := fmt.Sprintf("%v\n", string(puzzle[i:i+9]))
		buffer.WriteString(row)
	}
	buffer.WriteString("@MissSudoku")
	return buffer.String()
}

func GetTwitterApi() *anaconda.TwitterApi {
	file, _ := os.Open("settings.json")
	decoder := json.NewDecoder(file)
	settings := &Configuration{}
	decoder.Decode(&settings)

	anaconda.SetConsumerKey(settings.ConsumerKey)
	anaconda.SetConsumerSecret(settings.ConsumerSecret)
	api := anaconda.NewTwitterApi(settings.AccessToken, settings.AccessSecret)

	return api
}
