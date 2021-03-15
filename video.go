/**
 * @Author: tyz
 * @Date: 2021/3/15 14:05
 * @Name: video_model.go
 */

package youtube

type VideoModel struct {
	Kind     string    `json:"kind"`
	ETag     string    `json:"etag"`
	Items    []*Video  `json:"items,omitempty"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type Video struct {
	Kind    string        `json:"kind"`
	ETag    string        `json:"etag"`
	Id      string        `json:"id"`
	Snippet *VideoSnippet `json:"snippet,omitempty"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type VideoSnippet struct {
	PublishedAt          string      `json:"publishedAt"`
	ChannelId            string      `json:"channelId"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	Thumbnails           *Thumbnails `json:"thumbnails,omitempty"`
	ChannelTitle         string      `json:"channelTitle"`
	Tags                 []string    `json:"tags"`
	CategoryId           string      `json:"categoryId"`
	LiveBroadcastContent string      `json:"liveBroadcastContent"`
	Localized            *Localized  `json:"localized,omitempty"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnails struct {
	Default  *Thumbnail `json:"default"`
	Medium   *Thumbnail `json:"medium"`
	High     *Thumbnail `json:"high"`
	Standard *Thumbnail `json:"standard"`
	Maxres   *Thumbnail `json:"maxres"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
