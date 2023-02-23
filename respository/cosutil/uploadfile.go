/******
** @创建时间 : 2022/6/21 15:05
** @作者 : MUGUAGAI
******/
package cosutil

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"time"
)

func UploadVideo(fileName string, fileByte []byte) (url string, err error) {
	folderName := time.Now().Format("2006-01-02")
	key := filepath.Join("uploads"+folderName) + "/" + fileName
	_, err = CosClient.Object.Put(context.Background(), key, bytes.NewReader(fileByte), nil)
	objectURL := CosClient.Object.GetObjectURL(key)
	url = objectURL.String()
	fmt.Println(url)
	return url, err
}

func UploadCover(fileName string, fileByte []byte) (url string, err error) {
	folderName := time.Now().Format("2006-01-02")
	key := filepath.Join("cover"+folderName) + "/" + fileName
	_, err = CosClient.Object.Put(context.Background(), key, bytes.NewReader(fileByte), nil)
	objectURL := CosClient.Object.GetObjectURL(key)
	url = objectURL.String()
	fmt.Println(url)
	return url, err
}
