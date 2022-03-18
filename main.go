/*
 * @Author: liziwei01
 * @Date: 2022-03-18 15:45:07
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 23:48:16
 * @Description: file content
 */
package main

import (
	"gin-genshin-voice/bootstrap"
	"gin-genshin-voice/httpapi"
	"log"
)

func main() {
	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}
