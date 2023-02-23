package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/RaymondCode/simple-demo/respository/cosutil"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"os"

	"github.com/RaymondCode/simple-demo/respository/redis"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/util"
)

func main() {
	c := cron.New()
	_, _ = c.AddFunc("@every 5s", redis.FavouriteToMysql)
	c.Start()
	if err := Init(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	r := gin.Default()
	initRouter(r)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
func Init() error {
	fmt.Println("初始化开始")
	if err := conf.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return err
	}
	if err := respository.Init(conf.Conf.MySQLConfig); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		return err
	}
	conf.GetAddress()
	fmt.Println(conf.Ip)
	if err := cosutil.CosInit(); err != nil {
		return err
	}
	fmt.Println("初始化完成")
	return nil
}
