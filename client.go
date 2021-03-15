/**
 * @Author: tyz
 * @Date: 2021/3/15 14:13
 * @Name: client.go
 */

package youtube

import (
	"encoding/json"
	"errors"
	"github.com/smok-dev-team/ngx"
	"net/http"
	"strings"
)

const (
	k_YOUTUBE_API_URL       = "https://www.googleapis.com/youtube"
	K_YOUTUB_API_VIDEO_INFO = "/v3/videos"
)

type Client struct {
	key         string
	accessToken string
	apiDomain   string
}

func NewClient(key, accessToken string) (client *Client) {
	client = &Client{}
	client.key = key
	client.accessToken = accessToken
	client.apiDomain = k_YOUTUBE_API_URL
	return client
}

func (c *Client) BuildAPI(paths ...string) string {
	var path = c.apiDomain
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (c *Client) GetOneVideoInfoById(videoId string) (*VideoModel, error) {
	api := c.BuildAPI(K_YOUTUB_API_VIDEO_INFO)
	url := api + "?id=" + videoId + "&key=" + c.key + "&part=snippet"
	req := ngx.NewRequest("GET", url)

	if c.accessToken != "" {
		req.SetHeader("Authorization", "Bearer "+c.accessToken)
	}

	rsq := req.Exec()

	videoModel := &VideoModel{}
	if rsq.StatusCode() == http.StatusOK {
		err := json.Unmarshal(rsq.MustBytes(), videoModel)

		if err != nil {
			return nil, err
		}

		if len(videoModel.Items) == 0 {
			return nil, errors.New("video data is null")
		}

	} else {
		return nil, rsq.Error()
	}

	return videoModel, nil
}
