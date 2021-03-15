package youtube

import (
	"fmt"
	"testing"
)

func TestGetYoutubeVideoInfo(t *testing.T) {
	// You google api key -> https://console.developers.google.com/apis/dashboard
	var key = "AIzaSyDaliXcM6Z6xTREYY8X2ExjPsQ2XOafk1Y"
	var token = ""

	// Youtube Id Or url -> https://www.youtube.com/watch?v=W_JLuPyI3iU
	var source = "https://www.youtube.com/watch?v=W_JLuPyI3iU"

	videoModel, err := GetYoutubeVideoInfo(source, key, token)
	if err != nil {
		panic(err)
	}

	fmt.Println("videoId:", videoModel.Items[0].Id)
}