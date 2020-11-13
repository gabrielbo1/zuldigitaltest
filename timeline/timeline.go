package timeline

import (
	"fmt"
	"math/rand"
	"time"
)

//MaxLenTweet - Max len tweet print.
const MaxLenTweet = 45

type TimeLine struct {
	Text string `json:"text"`
}

//SliceTweet - Slice to text string in parts of length size 45.
func (timeLine *TimeLine) SliceTweet() []string {
	textCopy := timeLine.Text
	var result []string

	for i := 0; len(textCopy) > MaxLenTweet; {
		result = append(result, textCopy[i:(i+MaxLenTweet)])
		textCopy = textCopy[(i + MaxLenTweet):len(textCopy)]
	}

	if len(textCopy) > 0 {
		result = append(result, textCopy)
	}
	return result
}

func (timeLine *TimeLine) PrintlnTweet() {
	tweets := timeLine.SliceTweet()
	for i := range tweets {
		fmt.Printf("Tweet #%d: %s\n", i+1, tweets[i])
	}
}

func GetRandTimeLine(timeLines []TimeLine) TimeLine {
	if len(timeLines) == 0 {
		return TimeLine{}
	}
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(len(timeLines)-1+1) + 1
	return timeLines[randNumber]
}
