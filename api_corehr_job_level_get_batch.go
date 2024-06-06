// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// BatchGetCoreHRJobLevel 通过职级 ID 批量获取职级信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/job_level/batch_get
func (r *CoreHRService) BatchGetCoreHRJobLevel(ctx context.Context, request *BatchGetCoreHRJobLevelReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobLevelResp, *Response, error) {
	if r.cli.mock.mockCoreHRBatchGetCoreHRJobLevel != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#BatchGetCoreHRJobLevel mock enable")
		return r.cli.mock.mockCoreHRBatchGetCoreHRJobLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "BatchGetCoreHRJobLevel",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/job_levels/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCoreHRJobLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRBatchGetCoreHRJobLevel mock CoreHRBatchGetCoreHRJobLevel method
func (r *Mock) MockCoreHRBatchGetCoreHRJobLevel(f func(ctx context.Context, request *BatchGetCoreHRJobLevelReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobLevelResp, *Response, error)) {
	r.mockCoreHRBatchGetCoreHRJobLevel = f
}

// UnMockCoreHRBatchGetCoreHRJobLevel un-mock CoreHRBatchGetCoreHRJobLevel method
func (r *Mock) UnMockCoreHRBatchGetCoreHRJobLevel() {
	r.mockCoreHRBatchGetCoreHRJobLevel = nil
}

// BatchGetCoreHRJobLevelReq ...
type BatchGetCoreHRJobLevelReq struct {
	JobLevelIDs []string `json:"job_level_ids,omitempty"` // 职级 ID 列表, 示例值: ["1515"], 长度范围: `1` ～ `100`
}

// BatchGetCoreHRJobLevelResp ...
type BatchGetCoreHRJobLevelResp struct {
	Items []*BatchGetCoreHRJobLevelRespItem `json:"items,omitempty"` // 查询的职级信息
}

// BatchGetCoreHRJobLevelRespItem ...
type BatchGetCoreHRJobLevelRespItem struct {
	JobLevelID   string                                       `json:"job_level_id,omitempty"`  // 职级 ID
	LevelOrder   int64                                        `json:"level_order,omitempty"`   // 职级数值
	Code         string                                       `json:"code,omitempty"`          // 编码
	Name         []*BatchGetCoreHRJobLevelRespItemName        `json:"name,omitempty"`          // 名称
	Description  []*BatchGetCoreHRJobLevelRespItemDescription `json:"description,omitempty"`   // 描述
	Active       bool                                         `json:"active,omitempty"`        // 启用
	CustomFields []*BatchGetCoreHRJobLevelRespItemCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// BatchGetCoreHRJobLevelRespItemCustomField ...
type BatchGetCoreHRJobLevelRespItemCustomField struct {
	CustomApiName string                                         `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *BatchGetCoreHRJobLevelRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                          `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                         `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// BatchGetCoreHRJobLevelRespItemCustomFieldName ...
type BatchGetCoreHRJobLevelRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// BatchGetCoreHRJobLevelRespItemDescription ...
type BatchGetCoreHRJobLevelRespItemDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRJobLevelRespItemName ...
type BatchGetCoreHRJobLevelRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// batchGetCoreHRJobLevelResp ...
type batchGetCoreHRJobLevelResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *BatchGetCoreHRJobLevelResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
