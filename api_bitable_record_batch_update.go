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

// BatchUpdateBitableRecord 该接口用于更新数据表中的多条记录, 单次调用最多更新 500 条记录。
//
// 该接口支持调用频率上限为 10 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update
func (r *BitableService) BatchUpdateBitableRecord(ctx context.Context, request *BatchUpdateBitableRecordReq, options ...MethodOptionFunc) (*BatchUpdateBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableBatchUpdateBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchUpdateBitableRecord mock enable")
		return r.cli.mock.mockBitableBatchUpdateBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchUpdateBitableRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchUpdateBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchUpdateBitableRecord mock BitableBatchUpdateBitableRecord method
func (r *Mock) MockBitableBatchUpdateBitableRecord(f func(ctx context.Context, request *BatchUpdateBitableRecordReq, options ...MethodOptionFunc) (*BatchUpdateBitableRecordResp, *Response, error)) {
	r.mockBitableBatchUpdateBitableRecord = f
}

// UnMockBitableBatchUpdateBitableRecord un-mock BitableBatchUpdateBitableRecord method
func (r *Mock) UnMockBitableBatchUpdateBitableRecord() {
	r.mockBitableBatchUpdateBitableRecord = nil
}

// BatchUpdateBitableRecordReq ...
type BatchUpdateBitableRecordReq struct {
	AppToken   string                               `path:"app_token" json:"-"`     // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string                               `path:"table_id" json:"-"`      // table id, 示例值: "tblsRc9GRRXKqhvW"
	UserIDType *IDType                              `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Records    []*BatchUpdateBitableRecordReqRecord `json:"records,omitempty"`      // 记录
}

// BatchUpdateBitableRecordReqRecord ...
type BatchUpdateBitableRecordReqRecord struct {
	RecordID         *string                                          `json:"record_id,omitempty"`          // 记录 id, 更新多条记录时必填, 示例值: "recqwIwhc6"
	CreatedBy        *BatchUpdateBitableRecordReqRecordCreatedBy      `json:"created_by,omitempty"`         // 创建人
	CreatedTime      *int64                                           `json:"created_time,omitempty"`       // 创建时间, 示例值: 1610281603
	LastModifiedBy   *BatchUpdateBitableRecordReqRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 修改人
	LastModifiedTime *int64                                           `json:"last_modified_time,omitempty"` // 最近更新时间, 示例值: 1610281603
	Fields           map[string]interface{}                           `json:"fields,omitempty"`             // 记录字段, 关于支持新增的字段类型, 请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
}

// BatchUpdateBitableRecordReqRecordCreatedBy ...
type BatchUpdateBitableRecordReqRecordCreatedBy struct {
	ID     *string `json:"id,omitempty"`      // 人员Id, 示例值: "testesttest"
	Name   *string `json:"name,omitempty"`    // 中文姓名, 示例值: "黄泡泡"
	EnName *string `json:"en_name,omitempty"` // 英文姓名, 示例值: "Paopao Huang"
	Email  *string `json:"email,omitempty"`   // 邮箱, 示例值: "huangpaopao@feishu.cn"
}

// BatchUpdateBitableRecordReqRecordLastModifiedBy ...
type BatchUpdateBitableRecordReqRecordLastModifiedBy struct {
	ID     *string `json:"id,omitempty"`      // 人员Id, 示例值: "testesttest"
	Name   *string `json:"name,omitempty"`    // 中文姓名, 示例值: "黄泡泡"
	EnName *string `json:"en_name,omitempty"` // 英文姓名, 示例值: "Paopao Huang"
	Email  *string `json:"email,omitempty"`   // 邮箱, 示例值: "huangpaopao@feishu.cn"
}

// BatchUpdateBitableRecordResp ...
type BatchUpdateBitableRecordResp struct {
	Records []*BatchUpdateBitableRecordRespRecord `json:"records,omitempty"` // 记录
}

// BatchUpdateBitableRecordRespRecord ...
type BatchUpdateBitableRecordRespRecord struct {
	RecordID         string                                            `json:"record_id,omitempty"`          // 记录 id, 更新多条记录时必填
	CreatedBy        *BatchUpdateBitableRecordRespRecordCreatedBy      `json:"created_by,omitempty"`         // 创建人
	CreatedTime      int64                                             `json:"created_time,omitempty"`       // 创建时间
	LastModifiedBy   *BatchUpdateBitableRecordRespRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 修改人
	LastModifiedTime int64                                             `json:"last_modified_time,omitempty"` // 最近更新时间
	Fields           map[string]interface{}                            `json:"fields,omitempty"`             // 记录字段, 关于支持新增的字段类型, 请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
}

// BatchUpdateBitableRecordRespRecordCreatedBy ...
type BatchUpdateBitableRecordRespRecordCreatedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// BatchUpdateBitableRecordRespRecordLastModifiedBy ...
type BatchUpdateBitableRecordRespRecordLastModifiedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// batchUpdateBitableRecordResp ...
type batchUpdateBitableRecordResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *BatchUpdateBitableRecordResp `json:"data,omitempty"`
}
