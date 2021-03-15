/**
 * @Author: tyz
 * @Date: 2021/3/15 11:41
 * @Name: youtube.go
 * @Description: https://developers.google.com/youtube/v3/docs/videos
 * @Google Key: https://console.developers.google.com/apis/dashboard
 */

package youtube

import (
	"errors"
	"net/url"
	"strings"
)

func GetYoutubeVideoInfo(source string, key, token string) (*VideoModel, error) {

	queryStr := source
	if strings.HasPrefix(source, "https://www.youtube.com/watch?") {
		queryStr = strings.Replace(source, "https://www.youtube.com/watch?", "", -1)
	} else if strings.HasPrefix(source, "http://www.youtube.com/watch?") {
		queryStr = strings.Replace(source, "http://www.youtube.com/watch?", "", -1)
	}

	param, err := url.ParseQuery(queryStr)
	if err != nil {
		return nil, err
	}

	var videoId = param.Get("v")
	if len(videoId) == 0 {
		videoId = source
	}

	if len(videoId) == 0 {
		return nil, errors.New("video id is not empty")
	}

	client := NewClient(key, token)
	return client.GetOneVideoInfoById(videoId)
}
