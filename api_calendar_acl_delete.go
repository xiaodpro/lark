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

// DeleteCalendarACL 调用该接口以当前身份（应用或用户）删除指定日历内的某一访问控制, 即成员权限。
//
// - 当前身份由 Header Authorization 的 Token 类型决定。tenant_access_token 指应用身份, user_access_token 指用户身份。
// - 如果使用应用身份调用该接口, 则需要确保应用开启了[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 当前身份需要有日历的 owner 权限, 并且日历的类型只能为 primary 或 shared。你可以调用[查询日历信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, 获取日历类型以及当前身份对该日历的访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/calendar-acl/delete
func (r *CalendarService) DeleteCalendarACL(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarACL != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarACL mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarACL(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "DeleteCalendarACL",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteCalendarACLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarDeleteCalendarACL mock CalendarDeleteCalendarACL method
func (r *Mock) MockCalendarDeleteCalendarACL(f func(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error)) {
	r.mockCalendarDeleteCalendarACL = f
}

// UnMockCalendarDeleteCalendarACL un-mock CalendarDeleteCalendarACL method
func (r *Mock) UnMockCalendarDeleteCalendarACL() {
	r.mockCalendarDeleteCalendarACL = nil
}

// DeleteCalendarACLReq ...
type DeleteCalendarACLReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 需要删除访问控制的日历 ID, 创建共享日历时会返回日历 ID。你也可以调用以下接口获取某一日历的 ID, [查询主日历信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/primary), [查询日历列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/list), [搜索日历](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/search), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	ACLID      string `path:"acl_id" json:"-"`      // 访问控制 ID, 为日历创建访问控制时会返回访问控制 ID。你也可以调用[获取访问控制列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/list)接口, 获取指定日历内的访问控制信息, 示例值: "user_xxxxxx"
}

// DeleteCalendarACLResp ...
type DeleteCalendarACLResp struct {
}

// deleteCalendarACLResp ...
type deleteCalendarACLResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteCalendarACLResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
