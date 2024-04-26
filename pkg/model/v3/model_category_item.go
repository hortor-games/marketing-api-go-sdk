/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 商品类目
type CategoryItem struct {
	CategoryId       *int64               `json:"category_id,omitempty"`
	CategoryName     *string              `json:"category_name,omitempty"`
	ParentCategoryId *int64               `json:"parent_category_id,omitempty"`
	Level            *int64               `json:"level,omitempty"`
	CategoryPath     *[]CategoryPathsItem `json:"category_path,omitempty"`
}