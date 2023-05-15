package main

import (
	"GoWeb/src/conf"
	"GoWeb/src/controllers"
	"GoWeb/src/logm"
	"GoWeb/src/version"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := conf.ReadConf("./conf/web.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	startLogRus()
	printVersion()
	startServer()
}

func startLogRus() {
	logm.NewLogRus()
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
