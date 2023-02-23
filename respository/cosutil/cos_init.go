/******
** @创建时间 : 2022/6/21 14:43
** @作者 : MUGUAGAI
******/
package cosutil

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var CosClient *cos.Client

func CosInit() error {
	u, _ := url.Parse("https://15112782276-1312349127.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	CosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			//SecretID: os.Getenv("AKIDMYU9WeV6RWLtAAuNYT8WdS1JK4XjgMlu"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			//SecretKey: os.Getenv("tVqvjGWVDfqlQpxoOhjUxA7nCLDo8BxS"),
			SecretID:  "AKIDMYU9WeV6RWLtAAuNYT8WdS1JK4XjgMlu", // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: "tVqvjGWVDfqlQpxoOhjUxA7nCLDo8BxS",     // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	return nil
}
