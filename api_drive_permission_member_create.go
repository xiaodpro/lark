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

// CreateDriveMemberPermission 该接口用于根据 filetoken 给用户增加文档的权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/create
func (r *DriveService) CreateDriveMemberPermission(ctx context.Context, request *CreateDriveMemberPermissionReq, options ...MethodOptionFunc) (*CreateDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveCreateDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveMemberPermission",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveMemberPermission mock DriveCreateDriveMemberPermission method
func (r *Mock) MockDriveCreateDriveMemberPermission(f func(ctx context.Context, request *CreateDriveMemberPermissionReq, options ...MethodOptionFunc) (*CreateDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveCreateDriveMemberPermission = f
}

// UnMockDriveCreateDriveMemberPermission un-mock DriveCreateDriveMemberPermission method
func (r *Mock) UnMockDriveCreateDriveMemberPermission() {
	r.mockDriveCreateDriveMemberPermission = nil
}

// CreateDriveMemberPermissionReq ...
type CreateDriveMemberPermissionReq struct {
	Token            string `path:"token" json:"-"`              // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type             string `query:"type" json:"-"`              // 文件类型, 放于query参数中, 如: `?type=doc`, 示例值: "doc", 可选值有: `doc`: 文档, `sheet`: 电子表格, `file`: 云空间文件, `wiki`: 知识库节点, `bitable`: 多维表格, `docx`: 新版文档, `folder`: 文件夹（未来支持）
	NeedNotification *bool  `query:"need_notification" json:"-"` // 添加权限后是否通知对方, 注意: 使用`tenant_access_token`访问不支持该参数, 示例值: false, 默认值: `false`
	MemberType       string `json:"member_type,omitempty"`       // 用户类型, 与请求体中的`member_id`要对应, 可选值有: `email`: 飞书邮箱, `openid`: [开放平台ID](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), `openchat`: [开放平台群组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), `opendepartmentid`:[开放平台部门ID](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview), `userid`: [用户自定义ID](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 注意: 使用`tenant_access_token`访问不支持添加`opendepartmentid`, 示例值: "openid"
	MemberID         string `json:"member_id,omitempty"`         // 用户类型下的值, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	Perm             string `json:"perm,omitempty"`              // 需要更新的权限, 可选值有: `view`: 可阅读, `edit`: 可编辑, `full_access`: 所有权限, 示例值: "view"
}

// CreateDriveMemberPermissionResp ...
type CreateDriveMemberPermissionResp struct {
	Member *CreateDriveMemberPermissionRespMember `json:"member,omitempty"` // 本次添加权限的用户信息
}

// CreateDriveMemberPermissionRespMember ...
type CreateDriveMemberPermissionRespMember struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型, 与请求体中的`member_id`要对应, 可选值有: `email`: 飞书邮箱, `openid`: [开放平台ID](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), `openchat`: [开放平台群组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), `opendepartmentid`:[开放平台部门ID](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview), `userid`: [用户自定义ID](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 注意: 使用`tenant_access_token`访问不支持添加`opendepartmentid`
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
	Perm       string `json:"perm,omitempty"`        // 需要更新的权限, 可选值有: `view`: 可阅读, `edit`: 可编辑, `full_access`: 所有权限
}

// createDriveMemberPermissionResp ...
type createDriveMemberPermissionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveMemberPermissionResp `json:"data,omitempty"`
}
