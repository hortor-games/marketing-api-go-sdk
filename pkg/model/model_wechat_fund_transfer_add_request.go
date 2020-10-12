/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatFundTransferAddRequest struct {
	AccountId    int64                 `json:"account_id,omitempty"`
	FundType     WechatAccountFundType `json:"fund_type,omitempty"`
	Amount       int64                 `json:"amount,omitempty"`
	TransferType TransferType          `json:"transfer_type,omitempty"`
	Operator     string                `json:"operator,omitempty"`
}