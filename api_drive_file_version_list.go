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

// GetDriveFileVersionList 获取文档所有版本。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-version/list
func (r *DriveService) GetDriveFileVersionList(ctx context.Context, request *GetDriveFileVersionListReq, options ...MethodOptionFunc) (*GetDriveFileVersionListResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFileVersionList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFileVersionList mock enable")
		return r.cli.mock.mockDriveGetDriveFileVersionList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveFileVersionList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/versions",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveFileVersionListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveFileVersionList mock DriveGetDriveFileVersionList method
func (r *Mock) MockDriveGetDriveFileVersionList(f func(ctx context.Context, request *GetDriveFileVersionListReq, options ...MethodOptionFunc) (*GetDriveFileVersionListResp, *Response, error)) {
	r.mockDriveGetDriveFileVersionList = f
}

// UnMockDriveGetDriveFileVersionList un-mock DriveGetDriveFileVersionList method
func (r *Mock) UnMockDriveGetDriveFileVersionList() {
	r.mockDriveGetDriveFileVersionList = nil
}

// GetDriveFileVersionListReq ...
type GetDriveFileVersionListReq struct {
	FileToken  string  `path:"file_token" json:"-"`    // 源文档token, 示例值: "shtbcpM2mm3znrLfWnf4browTYp23"
	PageSize   int64   `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `100`
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "1665739388"
	ObjType    string  `query:"obj_type" json:"-"`     // 原文档类型, 示例值: "doc/docx/sheet/bitable", 可选值有: doc: doc文档, sheet: sheet文档, bitable: bitable文档, docx: docx文档
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDriveFileVersionListResp ...
type GetDriveFileVersionListResp struct {
	Items     []*GetDriveFileVersionListRespItem `json:"items,omitempty"`      // 版本文档列表
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetDriveFileVersionListRespItem ...
type GetDriveFileVersionListRespItem struct {
	Name        string `json:"name,omitempty"`         // 版本文档标题
	Version     string `json:"version,omitempty"`      // 版本文档版本号
	ParentToken string `json:"parent_token,omitempty"` // shtbcpM2mm3znrLfWnf4browTYp
	OwnerID     string `json:"owner_id,omitempty"`     // 版本文档所有者id
	CreatorID   string `json:"creator_id,omitempty"`   // 版本文档创建者id
	CreateTime  string `json:"create_time,omitempty"`  // 版本文档创建时间
	UpdateTime  string `json:"update_time,omitempty"`  // 版本文档更新时间
	Status      string `json:"status,omitempty"`       // 版本文档状态, 可选值有: 0: 正常状态, 1: 删除状态, 2: 回收站状态
	ObjType     string `json:"obj_type,omitempty"`     // 版本文档类型, 可选值有: doc: doc文档, sheet: sheet文档, bitable: bitable文档, docx: docx文档
	ParentType  string `json:"parent_type,omitempty"`  // 源文档类型, 可选值有: doc: doc文档, sheet: sheet文档, bitable: bitable文档, docx: docx文档
}

// getDriveFileVersionListResp ...
type getDriveFileVersionListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveFileVersionListResp `json:"data,omitempty"`
}