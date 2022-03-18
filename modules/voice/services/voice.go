/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:20:19
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 23:46:37
 * @Description: 完整服务
 */
package services

import (
	voiceData "gin-genshin-voice/modules/voice/data"
	voiceModel "gin-genshin-voice/modules/voice/model"

	"github.com/gin-gonic/gin"
)

func GetVoice(ctx *gin.Context, pars *voiceModel.CharacterPars) error {
	_, err := voiceData.GetCharMap(ctx, pars)
	err = voiceData.GetCharCode(ctx, pars)
	if err != nil {
		return err
	}
	err = voiceData.GetVoiceURL(ctx, pars)
	if err != nil {
		return err
	}
	return voiceData.GetVoice(ctx, pars)
}
