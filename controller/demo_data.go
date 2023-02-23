package controller

import "github.com/RaymondCode/simple-demo/respository"

var DemoVideos = []respository.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.5:8080/static/bear.mp4",
		CoverUrl:      "http://192.168.1.5:8080/static/屏幕截图 2021-02-16 163146.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.5:8080/static/bear.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []respository.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = respository.User{
	Id:            1,
	Name:          "666",
	FollowCount:   100,
	FollowerCount: 100,
	IsFollow:      false,
}
