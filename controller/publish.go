package controller

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/respository/cosutil"
	"github.com/RaymondCode/simple-demo/util/jwt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/RaymondCode/simple-demo/service"

	"github.com/RaymondCode/simple-demo/respository"

	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoListResponse struct {
	respository.Response
	VideoList []respository.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := c.PostForm("token")
	parseToken, err := jwt.ParseToken(token)
	username := parseToken.Username
	if _, exist := respository.UsersLoginInfo[username]; !exist {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// 获取上传文件信息
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	fileHandle, err := data.Open()
	fileByte, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videourl, err := cosutil.UploadVideo(filename, fileByte)
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user := respository.UsersLoginInfo[username]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	// 抽取视频封面 ffmpeg - start
	/*exec_ffmpeg_extract_cmd := "ffmpeg -i " + saveFile + " -ss 00:00:00 -frames:v 1 " + coverFile
	println("to tun:", exec_ffmpeg_extract_cmd)

	cmdArguments := []string{"-i", saveFile, "-ss", "00:00:00",
		"-frames:v", "1", coverFile}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	errFFMPEG := cmd.Run()
	if errFFMPEG != nil {
		log.Fatal(errFFMPEG)
	}
	fmt.Printf("command output: %q", out.String())*/
	// 抽取视频封面 ffmpeg - end
	u3 := uuid.New()
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	asJpeg, err := readFrameAsJpeg(videourl)
	coverurl, err := cosutil.UploadCover(coverPath, asJpeg)
	var video = respository.Video{
		Id:            worker.GetId(),
		Author:        user,
		PlayUrl:       videourl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateTime:    time.Now(),
		Title:         title,
	}

	if err := service.PublishVideo(video); err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, respository.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	parseToken, _ := jwt.ParseToken(token)
	username := parseToken.Username
	user := respository.UsersLoginInfo[username]
	c.JSON(http.StatusOK, VideoListResponse{
		Response: respository.Response{
			StatusCode: 0,
		},
		VideoList: respository.QueryVideosListByauthorid(user),
	})
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
