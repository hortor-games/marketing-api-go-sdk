/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 流量限制规则
type Rule struct {
	RuleType     *int64  `json:"RuleType,omitempty"`
	RuleTypeName *string `json:"RuleTypeName,omitempty"`
	RuleValue    *int64  `json:"RuleValue,omitempty"`
	RuleDesc     *string `json:"RuleDesc,omitempty"`
}
