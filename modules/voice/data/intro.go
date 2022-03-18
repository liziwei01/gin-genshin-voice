/*
 * @Author: liziwei01
 * @Date: 2022-03-18 20:22:17
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 22:41:26
 * @Description: file content
 */
package data

import (
	"fmt"
	voiceModel "gin-genshin-voice/modules/voice/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

const (
	AddressGenshinBBS = "https://bbs.mihoyo.com/ys"
	VoiceContent      = "/obc/content/"
	DetailTail        = "/detail?bbs_presentation_style=no_header"

	PatternCharCode = "https://uploadstatic.mihoyo.com/ys-obc/[0-9]{3,4}/[0-9]{1,2}/[0-9]{1,2}/[0-9]{6,12}/[0-9A-Za-z_]{40,60}.mp3"
)

func GetVoiceURL(ctx *gin.Context, pars *voiceModel.CharacterPars) error {
	url := getIntroURL(pars.CharCode)
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(response.StatusCode)
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}
	bodyStr := gconv.String(body)
	RegCharCode := regexp.MustCompile(PatternCharCode)
	urlList := RegCharCode.FindAllString(bodyStr, -1)
	if len(urlList) == 0 {
		ctx.Status(http.StatusNotFound)
		return fmt.Errorf("no voice url found")
	}
	pars.VoiceURL = getRandomVoice(urlList)
	return nil
}

func getIntroURL(charCode string) string {
	return AddressGenshinBBS + VoiceContent + charCode + DetailTail
}

func getRandomVoice(charURLList []string) string {
	listLen := len(charURLList)
	randNum := rand.Intn(listLen)
	return charURLList[randNum]
}
