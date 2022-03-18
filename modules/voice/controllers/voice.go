/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:18:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:27:33
 * @Description: 接收数据
 */
package controllers

import (
	"gin-genshin-voice/library/logit"
	"gin-genshin-voice/library/response"

	voiceModel "gin-genshin-voice/modules/voice/model"
	voiceService "gin-genshin-voice/modules/voice/services"

	"github.com/gin-gonic/gin"
)

func GetVoice(ctx *gin.Context) {
	inputs, hasError := getGetVoicePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
	}
	err := voiceService.GetVoice(ctx, &inputs)
	if err != nil {
		logit.Logger.Error(err.Error())
		response.StdFailed(ctx, err.Error())
	}
	// response.StdSuccess(ctx)
}

func getGetVoicePars(ctx *gin.Context) (voiceModel.CharacterPars, bool) {
	var inputs voiceModel.CharacterPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
