/*
 * @Author: liziwei01
 * @Date: 2022-03-18 20:07:03
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:41:43
 * @Description: file content
 */
package data

import (
	"context"
	voiceModel "gin-genshin-voice/modules/voice/model"
	"math/rand"
)

func GetCharCode(ctx context.Context, pars *voiceModel.CharacterPars) error {
	// 如果指定了角色，则不进行随机
	if pars.CharCode == "" && pars.CharName == "" {
		pars.CharName = GetRandomChar()
		pars.CharCode = mapCharNameCode(pars.CharName)
	} else if pars.CharCode == "" && pars.CharName != "" {
		pars.CharCode = mapCharNameCode(pars.CharName)
		if len(pars.CharCode) == 0 {
			pars.CharName = GetRandomChar()
			pars.CharCode = mapCharNameCode(pars.CharName)
		}
	}
	return nil
}

func mapCharNameCode(charName string) string {
	charNameCodeMap := map[string]string{
		"八重神子": "3584",
		"五郎":   "3635",
		"埃洛伊":  "3634",
		"菲谢尔":  "3583",
		"荒泷一斗": "3492",
		"云堇":   "3471",
	}
	return charNameCodeMap[charName]
}

func GetRandomChar() string {
	charCode := map[int]string{
		0: "八重神子",
		1: "五郎",
		2: "埃洛伊",
		3: "菲谢尔",
		4: "荒泷一斗",
		5: "云堇",
	}
	charNum := len(charCode)
	randNum := rand.Intn(charNum)
	return charCode[randNum]
}
