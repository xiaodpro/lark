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

// BatchCreateBitableRecord 该接口用于在数据表中新增多条记录, 单次调用最多新增 500 条记录。
//
// 该接口支持调用频率上限为 10 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create
func (r *BitableService) BatchCreateBitableRecord(ctx context.Context, request *BatchCreateBitableRecordReq, options ...MethodOptionFunc) (*BatchCreateBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableBatchCreateBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateBitableRecord mock enable")
		return r.cli.mock.mockBitableBatchCreateBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchCreateBitableRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchCreateBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchCreateBitableRecord mock BitableBatchCreateBitableRecord method
func (r *Mock) MockBitableBatchCreateBitableRecord(f func(ctx context.Context, request *BatchCreateBitableRecordReq, options ...MethodOptionFunc) (*BatchCreateBitableRecordResp, *Response, error)) {
	r.mockBitableBatchCreateBitableRecord = f
}

// UnMockBitableBatchCreateBitableRecord un-mock BitableBatchCreateBitableRecord method
func (r *Mock) UnMockBitableBatchCreateBitableRecord() {
	r.mockBitableBatchCreateBitableRecord = nil
}

// BatchCreateBitableRecordReq ...
type BatchCreateBitableRecordReq struct {
	AppToken   string                               `path:"app_token" json:"-"`     // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string                               `path:"table_id" json:"-"`      // table id, 示例值: "tblsRc9GRRXKqhvW"
	UserIDType *IDType                              `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Records    []*BatchCreateBitableRecordReqRecord `json:"records,omitempty"`      // 记录
}

// BatchCreateBitableRecordReqRecord ...
type BatchCreateBitableRecordReqRecord struct {
	CreatedBy        *BatchCreateBitableRecordReqRecordCreatedBy      `json:"created_by,omitempty"`         // 创建人
	CreatedTime      *int64                                           `json:"created_time,omitempty"`       // 创建时间, 示例值: 1610281603
	LastModifiedBy   *BatchCreateBitableRecordReqRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 修改人
	LastModifiedTime *int64                                           `json:"last_modified_time,omitempty"` // 最近更新时间, 示例值: 1610281603
	Fields           map[string]interface{}                           `json:"fields,omitempty"`             // 记录字段, 关于支持新增的字段类型, 请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
}

// BatchCreateBitableRecordReqRecordCreatedBy ...
type BatchCreateBitableRecordReqRecordCreatedBy struct {
	ID     *string `json:"id,omitempty"`      // 人员Id, 示例值: "testesttest"
	Name   *string `json:"name,omitempty"`    // 中文姓名, 示例值: "黄泡泡"
	EnName *string `json:"en_name,omitempty"` // 英文姓名, 示例值: "Paopao Huang"
	Email  *string `json:"email,omitempty"`   // 邮箱, 示例值: "huangpaopao@feishu.cn"
}

// BatchCreateBitableRecordReqRecordLastModifiedBy ...
type BatchCreateBitableRecordReqRecordLastModifiedBy struct {
	ID     *string `json:"id,omitempty"`      // 人员Id, 示例值: "testesttest"
	Name   *string `json:"name,omitempty"`    // 中文姓名, 示例值: "黄泡泡"
	EnName *string `json:"en_name,omitempty"` // 英文姓名, 示例值: "Paopao Huang"
	Email  *string `json:"email,omitempty"`   // 邮箱, 示例值: "huangpaopao@feishu.cn"
}

// BatchCreateBitableRecordResp ...
type BatchCreateBitableRecordResp struct {
	Records []*BatchCreateBitableRecordRespRecord `json:"records,omitempty"` // 记录
}

// BatchCreateBitableRecordRespRecord ...
type BatchCreateBitableRecordRespRecord struct {
	RecordID         string                                            `json:"record_id,omitempty"`          // 记录 id
	CreatedBy        *BatchCreateBitableRecordRespRecordCreatedBy      `json:"created_by,omitempty"`         // 创建人
	CreatedTime      int64                                             `json:"created_time,omitempty"`       // 创建时间
	LastModifiedBy   *BatchCreateBitableRecordRespRecordLastModifiedBy `json:"last_modified_by,omitempty"`   // 修改人
	LastModifiedTime int64                                             `json:"last_modified_time,omitempty"` // 最近更新时间
	Fields           map[string]interface{}                            `json:"fields,omitempty"`             // 记录字段, 关于支持新增的字段类型, 请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
}

// BatchCreateBitableRecordRespRecordCreatedBy ...
type BatchCreateBitableRecordRespRecordCreatedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// BatchCreateBitableRecordRespRecordLastModifiedBy ...
type BatchCreateBitableRecordRespRecordLastModifiedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// batchCreateBitableRecordResp ...
type batchCreateBitableRecordResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateBitableRecordResp `json:"data,omitempty"`
}