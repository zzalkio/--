package respository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var UsersLoginInfo map[string]User

// 初始化数据库
func Init(cfg *conf.MySQLConfig) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&interpolateParams=false", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=false"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	fmt.Println("数据库连接成功！")
	var user User
	UsersLoginInfo = make(map[string]User)
	UsersLoginInfo["0"] = user
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Video{})
	Db.AutoMigrate(&UserLike{})
	Db.AutoMigrate(&Comment{})
	Db.AutoMigrate(&FollowFollower{})
	return err
}

/*func CreatUserinfo() map[string]User {
	all, _ := NewUserDaoInstance().QueryAll()
	return all
}*/
