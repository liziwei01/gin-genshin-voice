/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:19:18
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:29:30
 * @Description: 数据处理
 */
package data

import (
	voiceModel "gin-genshin-voice/modules/voice/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	voice = "https://uploadstatic.mihoyo.com/ys-obc/2022/02/24/291696164/823198430724294c3fa85abe8ec90ddf_461883976801526917.mp3"
)

func GetVoice(ctx *gin.Context, pars *voiceModel.CharacterPars) error {
	response, err := http.Get(pars.VoiceURL)
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return err
	}
	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{
		"CharCode": pars.CharCode,
		"CharName": pars.CharName,
	}
	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	return nil
}
