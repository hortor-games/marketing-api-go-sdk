/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// SysStatus : 朋友圈头像及昵称跳转页状态
type SysStatus string

// List of SysStatus
const (
	SysStatus_NORMAL            SysStatus = "AD_STATUS_NORMAL"
	SysStatus_PENDING           SysStatus = "AD_STATUS_PENDING"
	SysStatus_DENIED            SysStatus = "AD_STATUS_DENIED"
	SysStatus_FROZEN            SysStatus = "AD_STATUS_FROZEN"
	SysStatus_PARTIALLY_PENDING SysStatus = "AD_STATUS_PARTIALLY_PENDING"
	SysStatus_PARTIALLY_NORMAL  SysStatus = "AD_STATUS_PARTIALLY_NORMAL"
	SysStatus_PREPARE           SysStatus = "AD_STATUS_PREPARE"
	SysStatus_DELETED           SysStatus = "AD_STATUS_DELETED"
	SysStatus_INVALID           SysStatus = "AD_STATUS_INVALID"
)
