/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 礼包道具
type Gift struct {
	GiftId       *string     `json:"gift_id,omitempty"`
	GiftName     *string     `json:"gift_name,omitempty"`
	GiftType     *int64      `json:"gift_type,omitempty"`
	ResourceList *[]Resource `json:"resource_list,omitempty"`
}
