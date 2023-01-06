package main

import (
	"fmt"
	"github.com/GoldenLeeK/blog-service/global"
	"github.com/GoldenLeeK/blog-service/interal/model"
	"github.com/GoldenLeeK/blog-service/interal/routers"
	"github.com/GoldenLeeK/blog-service/pkg/logger"
	"github.com/GoldenLeeK/blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal(fmt.Sprintf("init.setupSetting err :%v", err))
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatal(fmt.Sprintf("init.setupDBEngine  err :%v", err))
	}

	err = setupLogger()
	if err != nil {
		log.Fatal(fmt.Sprintf("init.Logger  err :%v", err))
	}

}
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func setupSetting() error {
	settingConfig, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = settingConfig.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = settingConfig.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = settingConfig.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil

}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}