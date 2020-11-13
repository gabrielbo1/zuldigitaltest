package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gabrielbo1/zuldigitaltest/config"
	"github.com/gabrielbo1/zuldigitaltest/timeline"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var token config.Token

func init() {
	config.FlagParse()
}

func parseBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getToken() {
	url := "https://n8e480hh63o547stou3ycz5lwz0958.herokuapp.com/" + config.GetValue(config.ApiVersion) + "/auth"
	resp, err := http.Post(url, "Application/Json", nil)
	if err == nil && resp.StatusCode == http.StatusOK {
		json.Unmarshal(parseBody(resp), &token)
	}
}

func getTimeLine() []timeline.TimeLine {
	if token.Token == "" {
		getToken()
	}
	var timeLines []timeline.TimeLine
	if token.Token != "" {
		client := http.Client{}
		req, _ := http.NewRequest(http.MethodGet, "https://n8e480hh63o547stou3ycz5lwz0958.herokuapp.com/"+
			config.GetValue(config.ApiVersion)+"/statuses/home_timeline.json", nil)
		req.Header.Set("Authorization", token.Token)
		resp, err := client.Do(req)

		if err == nil && resp.StatusCode == http.StatusOK {
			json.Unmarshal(parseBody(resp), &timeLines)
		}
	}
	return timeLines
}

func menu() {
	fmt.Println("\t----------------------------------------------------------")
	fmt.Println("\t* Prescione 1 para ver um novo tweet")
	fmt.Println("\t* Prescione 2 limpar")
	fmt.Println("\t* Qualquer outra tecla para sair")
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	go func() {
		for {
			getToken()
			time.Sleep(time.Second * 50)
		}
	}()

	menu()
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case '1':
			timeLines := getTimeLine()
			timeline := timeline.GetRandTimeLine(timeLines)
			clear()
			menu()
			timeline.PrintlnTweet()
			break
		case '2':
			clear()
			menu()
		default:
			os.Exit(0)
		}

	}
}
