## Youtube
一个根据视频ID获取视频详细信息的工具库
#### GetYoutubeVideoInfo(source, key, token)

- source 视频ID或者链接
- key
- token


```
var key = "AIzaSyDaliXcM6Z6xTREYY8X2ExjPsQ2XOafk1Y"
var token = ""

// Youtube Id Or url -> https://www.youtube.com/watch?v=W_JLuPyI3iU
var source = "https://www.youtube.com/watch?v=W_JLuPyI3iU"

videoModel, err := GetYoutubeVideoInfo(source, key, token)
if err != nil {
    panic(err)
}
```