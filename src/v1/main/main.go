package v1

import (
	conf "GoWeb/src/v1/conf"
	controllers "GoWeb/src/v1/controllers"
	mysql "GoWeb/src/v1/db/mysql"
	redis "GoWeb/src/v1/db/redis"
	logm "GoWeb/src/v1/logm"
	"GoWeb/src/version"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	err := conf.ReadConf("./conf/web.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------start connect mysql--------")
	err = mysql.NewMysqlConn(conf.MySqlCfg.UserName, conf.MySqlCfg.PassWord, conf.MySqlCfg.Ip, int16(conf.MySqlCfg.Port), conf.MySqlCfg.DBName, conf.MySqlCfg.ConPool)
	fmt.Println("------start connect mysql--------:", err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----------start connect redis-----------")
	err = redis.NewRedisPool(conf.RedisCfg.Ip, conf.RedisCfg.Port, conf.RedisCfg.PassWord, 16, conf.RedisCfg.ConPool)
	if err != nil {
		fmt.Println(err)
		return
	}
	logm.NewLogRus()
	printVersion()
	startServer()
}

func printVersion() {
	fmt.Println(version.VerionStr)
	fmt.Println(version.VerionInt)
	fmt.Println(version.VerionGit)
}

func startServer() {
	r := gin.Default()
	initRouterV1(r)
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.WebCfg.Ip, conf.WebCfg.Port))
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	srv := &http.Server{
		Addr:    l.Addr().String(),
		Handler: r,
	}
	err = srv.Serve(l)
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Server error:", err)
		return
	}
}

func initRouterV1(r *gin.Engine) {
	version := "v1/"
	//user
	g1 := r.Group(version + "user/")
	g1.GET("add", controllers.AddUser)
	g1.GET("modify", controllers.ModifyUser)
	g1.GET("del", controllers.DelUser)
	g1.POST("add", controllers.AddUser)
	g1.POST("modify", controllers.ModifyUser)
	g1.POST("del", controllers.DelUser)
}
