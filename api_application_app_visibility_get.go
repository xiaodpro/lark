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

// GetApplicationAppVisibility 该接口用于查询应用在该企业内可以被使用的范围, 只能被企业自建应用调用。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN
// new doc: https://open.feishu.cn/document/server-docs/application-v6/admin/obtain-the-app-availability-in-an-organization
func (r *ApplicationService) GetApplicationAppVisibility(ctx context.Context, request *GetApplicationAppVisibilityReq, options ...MethodOptionFunc) (*GetApplicationAppVisibilityResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationAppVisibility != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationAppVisibility mock enable")
		return r.cli.mock.mockApplicationGetApplicationAppVisibility(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationAppVisibility",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v2/app/visibility",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationAppVisibilityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationAppVisibility mock ApplicationGetApplicationAppVisibility method
func (r *Mock) MockApplicationGetApplicationAppVisibility(f func(ctx context.Context, request *GetApplicationAppVisibilityReq, options ...MethodOptionFunc) (*GetApplicationAppVisibilityResp, *Response, error)) {
	r.mockApplicationGetApplicationAppVisibility = f
}

// UnMockApplicationGetApplicationAppVisibility un-mock ApplicationGetApplicationAppVisibility method
func (r *Mock) UnMockApplicationGetApplicationAppVisibility() {
	r.mockApplicationGetApplicationAppVisibility = nil
}

// GetApplicationAppVisibilityReq ...
type GetApplicationAppVisibilityReq struct {
	AppID         string  `query:"app_id" json:"-"`          // 目标应用的 ID
	UserPageToken *string `query:"user_page_token" json:"-"` // 分页拉取用户列表起始位置标示, 不填表示从头开始
	UserPageSize  *int64  `query:"user_page_size" json:"-"`  // 本次拉取用户列表最大个数(最大值 1000, 0 自动最大个数 )
}

// GetApplicationAppVisibilityResp ...
type GetApplicationAppVisibilityResp struct {
	Departments          []*GetApplicationAppVisibilityRespDepartment     `json:"departments,omitempty"`           // 可用部门列表
	InvisibleDepartments []string                                         `json:"invisible_departments,omitempty"` // 禁用部门列表
	Users                []*GetApplicationAppVisibilityRespUser           `json:"users,omitempty"`                 // 可用用户列表（仅包含单独设置的用户, 可用部门、用户组中的用户未展开）
	InvisibleUsers       []*GetApplicationAppVisibilityRespInvisibleUser  `json:"invisible_users,omitempty"`       // 禁用用户列表（仅包含单独设置的用户, 可用部门、用户组中的用户未展开）
	Groups               []*GetApplicationAppVisibilityRespGroup          `json:"groups,omitempty"`                // 可用用户组列表
	InvisibleGroups      []*GetApplicationAppVisibilityRespInvisibleGroup `json:"invisible_groups,omitempty"`      // 禁用用户组列表
	IsVisibleToAll       int64                                            `json:"is_visible_to_all,omitempty"`     // 是否全员可见, 1: 是, 0: 否
	HasMoreUsers         int64                                            `json:"has_more_users,omitempty"`        // 是否还有更多可见用户, 1: 是, 0: 否
	UserPageToken        string                                           `json:"user_page_token,omitempty"`       // 拉取下一页用户列表时使用的 user_page_token
}

// GetApplicationAppVisibilityRespDepartment ...
type GetApplicationAppVisibilityRespDepartment struct {
	ID string `json:"id,omitempty"` // 自定义 department_id
}

// GetApplicationAppVisibilityRespGroup ...
type GetApplicationAppVisibilityRespGroup struct {
	ID string `json:"id,omitempty"` // 用户组 group_id
}

// GetApplicationAppVisibilityRespInvisibleGroup ...
type GetApplicationAppVisibilityRespInvisibleGroup struct {
	ID string `json:"id,omitempty"` // 用户组 group_id
}

// GetApplicationAppVisibilityRespInvisibleUser ...
type GetApplicationAppVisibilityRespInvisibleUser struct {
	UserID string `json:"user_id,omitempty"` // 用户的 user_id, 只返回给申请了 user_id 权限的企业自建应用
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}

// GetApplicationAppVisibilityRespUser ...
type GetApplicationAppVisibilityRespUser struct {
	UserID string `json:"user_id,omitempty"` // 用户的 user_id, 只返回给申请了 user_id 权限的企业自建应用
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
}

// getApplicationAppVisibilityResp ...
type getApplicationAppVisibilityResp struct {
	Code  int64                            `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg   string                           `json:"msg,omitempty"`  // 返回码的描述
	Data  *GetApplicationAppVisibilityResp `json:"data,omitempty"` // 返回的业务信息
	Error *ErrorDetail                     `json:"error,omitempty"`
}
