/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 轮播文案组件
type LivingDescStruct struct {
	LivingDescSwitch *bool     `json:"living_desc_switch,omitempty"`
	DescList         *[]string `json:"desc_list,omitempty"`
}