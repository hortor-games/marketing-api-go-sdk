/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 元素审核结果
type ElementRejectDetailInfoListStruct struct {
	ElementName        *string                         `json:"element_name,omitempty"`
	ElementType        ReviewElementType               `json:"element_type,omitempty"`
	ElementValue       *string                         `json:"element_value,omitempty"`
	Reason             *string                         `json:"reason,omitempty"`
	ReviewStatus       ReviewResultStatus              `json:"review_status,omitempty"`
	RejectInfoLocation *[]RejectInfoLocationListStruct `json:"reject_info_location,omitempty"`
}
