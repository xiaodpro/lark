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

// CreateDocxBlock 指定需要操作的块, 为其创建一批子块, 并插入到指定位置。如果操作成功, 接口将返回新创建子块的富文本内容。
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnby5Y0yoACL3PdfZqrJEm6f#doxcnm0ooUe0s20GwwVB3a05rtb), 了解相关规则及约束。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/create
func (r *DriveService) CreateDocxBlock(ctx context.Context, request *CreateDocxBlockReq, options ...MethodOptionFunc) (*CreateDocxBlockResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDocxBlock != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDocxBlock mock enable")
		return r.cli.mock.mockDriveCreateDocxBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDocxBlock",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDocxBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDocxBlock mock DriveCreateDocxBlock method
func (r *Mock) MockDriveCreateDocxBlock(f func(ctx context.Context, request *CreateDocxBlockReq, options ...MethodOptionFunc) (*CreateDocxBlockResp, *Response, error)) {
	r.mockDriveCreateDocxBlock = f
}

// UnMockDriveCreateDocxBlock un-mock DriveCreateDocxBlock method
func (r *Mock) UnMockDriveCreateDocxBlock() {
	r.mockDriveCreateDocxBlock = nil
}

// CreateDocxBlockReq ...
type CreateDocxBlockReq struct {
	DocumentID         string       `path:"document_id" json:"-"`           // 文档的唯一标识, 示例值: "doxcnePuYufKa49ISjhD8Ih0ikh"
	BlockID            string       `path:"block_id" json:"-"`              // Block 的唯一标识, 示例值: "doxcnO6UW6wAw2qIcYf4hZpFIth"
	DocumentRevisionID *int64       `query:"document_revision_id" json:"-"` // 操作的文档版本, 1表示文档最新版本。若此时操作的版本为文档最新版本, 则需要持有文档的阅读权限；若此时操作的版本为文档的历史版本, 则需要持有文档的编辑权限, 示例值:1, 默认值: `-1`, 最小值: `-1`
	ClientToken        *string      `query:"client_token" json:"-"`         // 操作的唯一标识, 与接口返回值的 client_token 相对应, 用于幂等的进行更新操作。此值为空表示将发起一次新的请求, 此值非空表示幂等的进行更新操作, 示例值: "fe599b60-450f-46ff-b2ef-9f6675625b97"
	UserIDType         *IDType      `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: <md-enum>, <md-enum-item key="open_id" >用户的 open id</md-enum-item>, <md-enum-item key="union_id" >用户的 union id</md-enum-item>, <md-enum-item key="user_id" >用户的 user id</md-enum-item>, </md-enum>, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Children           []*DocxBlock `json:"children,omitempty"`             // 添加的孩子列表, 长度范围: `1` ～ `50`
	Index              *int64       `json:"index,omitempty"`                // 当前 block 在 children 中的插入位置, 起始值为 0, 最大值为原 children 长度, 示例值: 0, 默认值: `-1`
}

// CreateDocxBlockResp ...
type CreateDocxBlockResp struct {
	Children           []*DocxBlock `json:"children,omitempty"`             // 所添加的孩子的 Block 信息
	DocumentRevisionID int64        `json:"document_revision_id,omitempty"` // 当前 block children 创建成功后文档的版本号
	ClientToken        string       `json:"client_token,omitempty"`         // 操作的唯一标识, 更新请求中使用此值表示幂等的进行此次更新
}

// createDocxBlockResp ...
type createDocxBlockResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *CreateDocxBlockResp `json:"data,omitempty"`
}
