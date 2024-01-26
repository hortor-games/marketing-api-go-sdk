/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 落地页信息
type DynamicCreativePageSpec struct {
	PageId                  *int64                   `json:"page_id,omitempty"`
	PageUrl                 *string                  `json:"page_url,omitempty"`
	ChannelsShopProductSpec *ChannelsShopProductSpec `json:"channels_shop_product_spec,omitempty"`
}
