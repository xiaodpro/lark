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

// GetCalendarACLList 该接口用于以当前身份（应用 / 用户）获取日历的控制权限列表。
//
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限, 并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/list
func (r *CalendarService) GetCalendarACLList(ctx context.Context, request *GetCalendarACLListReq, options ...MethodOptionFunc) (*GetCalendarACLListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarACLList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarACLList mock enable")
		return r.cli.mock.mockCalendarGetCalendarACLList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarACLList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/acls",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarACLListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarACLList mock CalendarGetCalendarACLList method
func (r *Mock) MockCalendarGetCalendarACLList(f func(ctx context.Context, request *GetCalendarACLListReq, options ...MethodOptionFunc) (*GetCalendarACLListResp, *Response, error)) {
	r.mockCalendarGetCalendarACLList = f
}

// UnMockCalendarGetCalendarACLList un-mock CalendarGetCalendarACLList method
func (r *Mock) UnMockCalendarGetCalendarACLList() {
	r.mockCalendarGetCalendarACLList = nil
}

// GetCalendarACLListReq ...
type GetCalendarACLListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 小于10取10, 最大值: `50`
}

// GetCalendarACLListResp ...
type GetCalendarACLListResp struct {
	Acls      []*GetCalendarACLListRespACL `json:"acls,omitempty"`       // 入参日历对应的acl列表
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCalendarACLListRespACL ...
type GetCalendarACLListRespACL struct {
	ACLID string                          `json:"acl_id,omitempty"` // acl资源ID。参见[ACL ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/introduction)
	Role  CalendarRole                    `json:"role,omitempty"`   // 对日历的访问权限, 可选值有: unknown: 未知权限, free_busy_reader: 游客, 只能看到忙碌/空闲信息, reader: 订阅者, 查看所有日程详情, writer: 编辑者, 创建及修改日程, owner: 管理员, 管理日历及共享设置
	Scope *GetCalendarACLListRespACLScope `json:"scope,omitempty"`  // 权限范围
}

// GetCalendarACLListRespACLScope ...
type GetCalendarACLListRespACLScope struct {
	Type   string `json:"type,omitempty"`    // 权限类型, 当type为User时, 值为open_id/user_id/union_id, 可选值有: user: 用户
	UserID string `json:"user_id,omitempty"` // 用户ID, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
}

// getCalendarACLListResp ...
type getCalendarACLListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarACLListResp `json:"data,omitempty"`
}