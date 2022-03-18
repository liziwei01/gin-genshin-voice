/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:16:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:42:42
 * @Description: 路由
 */

package routers

import (
	voiceControllers "gin-genshin-voice/modules/voice/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.Engine) {
	voiceGroup := router.Group("/voice")
	{
		voiceGroup.GET("/rand", voiceControllers.GetVoice)
		voiceGroup.GET("/randChar", voiceControllers.GetChar)
	}
}
