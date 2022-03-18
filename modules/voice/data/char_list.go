/*
 * @Author: liziwei01
 * @Date: 2022-03-18 22:25:40
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 23:46:12
 * @Description: file content
 */
package data

import (
	voiceModel "gin-genshin-voice/modules/voice/model"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

var (
	charVoiceDisplayURL = "https://bbs.mihoyo.com/ys/obc/channel/map/80/84?bbs_presentation_style=no_header"
	PatternCharVoice    = `<a title="[^\x00-\xff]{1,20} 角色语音" href="/ys/obc/content/[0-9]{4}/detail\?bbs_presentation_style=no_header">`
	PatternContentID    = `[0-9]{4}`
	PatternCharName     = `[^\x00-\xff]{1,20}`
)

func GetCharMap(ctx *gin.Context, pars *voiceModel.CharacterPars) (map[string]string, error) {
	response, _ := http.Get(charVoiceDisplayURL)
	body, _ := ioutil.ReadAll(response.Body)
	bodyStr := gconv.String(body)
	RegCharCode := regexp.MustCompile(PatternCharVoice)
	charVoiceHRefList := RegCharCode.FindAllString(bodyStr, -1)
	ContentIDList := []string{}
	CharNameList := []string{}
	for i := 0; i < len(charVoiceHRefList); i++ {
		contentid := regexp.MustCompile(PatternContentID).FindString(charVoiceHRefList[i])
		charName := regexp.MustCompile(PatternCharName).FindString(charVoiceHRefList[i])
		ContentIDList = append(ContentIDList, contentid)
		CharNameList = append(CharNameList, charName)
	}
	pars.CharCode = getRandomVoice(ContentIDList)
	return nil, nil
}
