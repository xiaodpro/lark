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

// DeleteWikiSpaceMember 此接口用于删除知识空间成员或管理员。
//
// 知识空间具有[类型](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)和[可见性](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)的概念。不同的类型或可见性可以对本操作做出限制:
// - 可见性限制: 公开知识空间（visibility为public）对租户所有用户可见, 因此不支持再删除成员, 但可以删除管理员。
// - 类型限制: 个人知识空间 （type为person）为个人管理的知识空间, 不支持删除管理员。但可以删除成员。
// 知识空间权限要求, 当前用户或应用:
// - 为知识空间管理员
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/delete
func (r *DriveService) DeleteWikiSpaceMember(ctx context.Context, request *DeleteWikiSpaceMemberReq, options ...MethodOptionFunc) (*DeleteWikiSpaceMemberResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteWikiSpaceMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteWikiSpaceMember mock enable")
		return r.cli.mock.mockDriveDeleteWikiSpaceMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteWikiSpaceMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteWikiSpaceMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteWikiSpaceMember mock DriveDeleteWikiSpaceMember method
func (r *Mock) MockDriveDeleteWikiSpaceMember(f func(ctx context.Context, request *DeleteWikiSpaceMemberReq, options ...MethodOptionFunc) (*DeleteWikiSpaceMemberResp, *Response, error)) {
	r.mockDriveDeleteWikiSpaceMember = f
}

// UnMockDriveDeleteWikiSpaceMember un-mock DriveDeleteWikiSpaceMember method
func (r *Mock) UnMockDriveDeleteWikiSpaceMember() {
	r.mockDriveDeleteWikiSpaceMember = nil
}

// DeleteWikiSpaceMemberReq ...
type DeleteWikiSpaceMemberReq struct {
	SpaceID    string `path:"space_id" json:"-"`     // 知识空间id, 示例值: "7008061636015554580"
	MemberID   string `path:"member_id" json:"-"`    // 成员id, 值的类型由请求体的 member_type 参数决定, 示例值: "g64fb7g7"
	MemberType string `json:"member_type,omitempty"` // “openchat” - 群id, “userid” - 用户id, “email” - 邮箱, “opendepartmentid” - 部门id, “openid” - 应用openid, “unionid” - [unionid](/:ssltoken/home/user-identity-introduction/union-id, ), 示例值: "userid"
	MemberRole string `json:"member_role,omitempty"` // 角色: “admin” - 管理员, “member” - 成员, 示例值: "admin"
}

// DeleteWikiSpaceMemberResp ...
type DeleteWikiSpaceMemberResp struct {
	Member *DeleteWikiSpaceMemberRespMember `json:"member,omitempty"` // 成员信息
}

// DeleteWikiSpaceMemberRespMember ...
type DeleteWikiSpaceMemberRespMember struct {
	MemberType string `json:"member_type,omitempty"` // “openchat” - 群id, “userid” - 用户id, “email” - 邮箱, “opendepartmentid” - 部门id, “openid” - 应用openid, “unionid” - [unionid](/:ssltoken/home/user-identity-introduction/union-id, )
	MemberID   string `json:"member_id,omitempty"`   // 用户id, 值的类型由上面的 member_type 参数决定
	MemberRole string `json:"member_role,omitempty"` // 角色: “admin” - 管理员, “member” - 成员
}

// deleteWikiSpaceMemberResp ...
type deleteWikiSpaceMemberResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteWikiSpaceMemberResp `json:"data,omitempty"`
}
