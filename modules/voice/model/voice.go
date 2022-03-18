/*
 * @Author: liziwei01
 * @Date: 2022-03-03 19:45:35
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-18 21:21:16
 * @Description: model
 */
package model

type CharacterPars struct {
	CharCode string `form:"char_code" json:"char_code"`
	CharName string `form:"char_name" json:"char_name"`
	VoiceURL string `form:"voice_url" json:"voice_url"`
}
