/*
 * @Author: liziwei01
 * @Date: 2022-03-18 21:39:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:40:36
 * @Description: file content
 */
package controllers

import (
	"gin-genshin-voice/library/logit"
	"gin-genshin-voice/library/response"
	voiceService "gin-genshin-voice/modules/voice/services"

	"github.com/gin-gonic/gin"
)

func GetChar(ctx *gin.Context) {
	charName, err := voiceService.GetChar(ctx)
	if err != nil {
		logit.Logger.Error(err.Error())
		response.StdFailed(ctx, err.Error())
	}
	response.StdSuccess(ctx, charName)
}
