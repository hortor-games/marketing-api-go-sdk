/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type XijingPageDeleteRequest struct {
	AccountId  *int64    `json:"account_id,omitempty"`
	PageIdList *[]string `json:"page_id_list,omitempty"`
}
