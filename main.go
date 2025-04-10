package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
)

//go:embed static/*
var f embed.FS

func main() {

	PORT := 8080
	ipStr := getIp()
	fmt.Println("IPv4:", ipStr)
	url := fmt.Sprintf("http://%s:%d/index.html", ipStr,PORT)
	if err := printQRCode(url); err != nil {
		fmt.Println("QR Error:", err)
	}

	router := gin.Default()

	router.GET("/index.html", func(c *gin.Context) {
		file,_ := f.ReadFile("static/index.html")
		c.Data(200, "text/html",file)
	})

	fileAssets, _ := fs.Sub(f, "static/assets")
	router.StaticFS("/assets", http.FS(fileAssets))

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "pong",
		})
	})

	// /api/a?times=10
	router.GET("/api/:action", func(c *gin.Context) {
		action := c.Param("action")
		queryTimes := c.Query("times")

		times, err := strconv.Atoi(queryTimes)
		if err != nil {
			times = 1
		}

		var function func()

		switch action {
		case "a":
			function = func() {
				robotgo.Move(710, 782)
			}
		case "b":
			function = func() {
				robotgo.Move(710, 840)
			}
		case "c":
			function = func() {
				robotgo.Move(710, 894)
			}
		case "d":
			function = func() {
				robotgo.Move(710, 955)
			}
		}

		for range times {
			function()
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(135, 135)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(1412, 135)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(1095, 1006)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)

			function()
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(1412, 135)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(135, 135)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
			robotgo.Move(1095, 1006)
			robotgo.Click()
			time.Sleep(time.Millisecond * 100)
		}

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  fmt.Sprintf("/api/%s?times=%d", action, times),
		})
	})

	router.Run(fmt.Sprintf(":%d", PORT))
}
