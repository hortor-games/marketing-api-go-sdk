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
type DpPageSpec struct {
	MiniProgramSpec *DpMiniProgramSpec `json:"mini_program_spec,omitempty"`
	AlitaPageSpec   *DpAlitaPageSpec   `json:"alita_page_spec,omitempty"`
}
