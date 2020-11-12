package timeline

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestTimeLine_SliceTweet(t *testing.T) {
	timeLine := TimeLine{}
	timeLine.Text = strings.Repeat("a", MaxLenTweet)
	result := timeLine.SliceTweet()
	if len(result[0]) != MaxLenTweet || len(result) != 1 {
		t.Fatal("Error slice time line.")
	}

	timeLine.Text = strings.Repeat("a", MaxLenTweet+1)
	result = timeLine.SliceTweet()
	if len(result[0]) != MaxLenTweet || len(result[1]) != 1 {
		t.Fatal("Error slice time line.")
	}

	timeLine.Text = strings.Repeat("a", MaxLenTweet*2)
	result = timeLine.SliceTweet()
	if len(result[0]) != MaxLenTweet || len(result[1]) != MaxLenTweet {
		t.Fatal("Error slice time line.")
	}

	rand.Seed(time.Now().UnixNano())
	magicNumber := rand.Intn(MaxLenTweet-1+1) + 1
	timeLine.Text = strings.Repeat("a", MaxLenTweet+magicNumber)
	result = timeLine.SliceTweet()
	if len(result[0]) != MaxLenTweet || len(result[1]) != magicNumber {
		t.Fatal("Error slice time line.")
	}
}
