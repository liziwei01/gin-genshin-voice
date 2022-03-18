/*
 * @Author: liziwei01
 * @Date: 2022-03-18 21:40:57
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:42:07
 * @Description: file content
 */
package services

import (
	"context"
	voiceData "gin-genshin-voice/modules/voice/data"
)

func GetChar(context.Context) (string, error) {
	return voiceData.GetRandomChar(), nil
}
