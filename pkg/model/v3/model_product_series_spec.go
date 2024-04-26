/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 商品系列详情
type ProductSeriesSpec struct {
	LogicOperator ProductSeriesSpecLogicOperator      `json:"logic_operator,omitempty"`
	Filters       *[]ProductSeriesSpecFilteringStruct `json:"filters,omitempty"`
}