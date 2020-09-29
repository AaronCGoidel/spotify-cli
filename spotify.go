package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const token = "" // PUT TOKEN HERE
const baseURL = "https://api.spotify.com/v1/me/player"

type playerT struct {
	IsPlaying bool `json:"is_playing"`
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func makeRequest(reqType string, endpoint string) *http.Response {
	req, err := http.NewRequest(reqType, baseURL+endpoint, nil)
	handleErr(err)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	handleErr(err)

	return res
}

func getInfo() bool {
	res := makeRequest("GET", "")
	body, err := ioutil.ReadAll(res.Body)
	handleErr(err)

	player := playerT{}
	json.Unmarshal(body, &player)

	return player.IsPlaying
}

func displayUsage() {
	fmt.Println("Wrong")
}

func play() {
	makeRequest("PUT", "/play")
}

func pause() {
	makeRequest("PUT", "/pause")
}

func skip() {
	makeRequest("POST", "/next")
}

func main() {
	isPlaying := getInfo()
	if len(os.Args) < 2 {
		if isPlaying {
			pause()
		} else {
			play()
		}
	} else if os.Args[1] == "n" {
		skip()
	}
}
