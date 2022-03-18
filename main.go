/*
 * @Author: liziwei01
 * @Date: 2022-03-18 15:45:07
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 23:31:13
 * @Description: file content
 */
package main

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gogf/gf/util/gconv"
)

func main() {
	// app, err := bootstrap.Setup()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // 注册接口路由
	// httpapi.InitRouters(app.Handler)

	// app.Start()
	charVoiceDisplayURL := "https://bbs.mihoyo.com/ys/obc/channel/map/80/84?bbs_presentation_style=no_header"

	// PatternCharVoice := `content_id:[0-9]{4},title:\"[^\x00-\xff]{1,20} 角色语音",ext`
	PatternCharVoice := `<a title="[^\x00-\xff]{1,20} 角色语音" href="/ys/obc/content/[0-9]{4}/detail\?bbs_presentation_style=no_header">`

	PatternContentID := `[0-9]{4}`
	PatternCharName := `[^\x00-\xff]{1,20}`

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
	urlList := []string{}
	for i := 0; i < len(urlList); i++ {
		urlList = append(urlList, charVoiceHRefList[i][len(charVoiceHRefList[i])-12:len(charVoiceHRefList[i])-1])
	}
}
