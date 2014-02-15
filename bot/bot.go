package bot

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/dkua/go-sudoku"
	"os"
)

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
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	return api
}
