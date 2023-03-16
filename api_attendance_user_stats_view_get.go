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

// GetAttendanceUserStatsView 查询开发者定制的日度统计或月度统计的统计报表表头设置信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query
func (r *AttendanceService) GetAttendanceUserStatsView(ctx context.Context, request *GetAttendanceUserStatsViewReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsViewResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserStatsView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserStatsView mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserStatsView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserStatsView",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_stats_views/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserStatsViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserStatsView mock AttendanceGetAttendanceUserStatsView method
func (r *Mock) MockAttendanceGetAttendanceUserStatsView(f func(ctx context.Context, request *GetAttendanceUserStatsViewReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsViewResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserStatsView = f
}

// UnMockAttendanceGetAttendanceUserStatsView un-mock AttendanceGetAttendanceUserStatsView method
func (r *Mock) UnMockAttendanceGetAttendanceUserStatsView() {
	r.mockAttendanceGetAttendanceUserStatsView = nil
}

// GetAttendanceUserStatsViewReq ...
type GetAttendanceUserStatsViewReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 响应体中的 user_id 的员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	Locale       string       `json:"locale,omitempty"`        // 语言类型, 示例值: "zh", 可选值有: en: 英语, ja: 日语, zh: 中文
	StatsType    string       `json:"stats_type,omitempty"`    // 统计类型, 示例值: "daily", 可选值有: daily: 日度统计, month: 月度统计
	UserID       *string      `json:"user_id,omitempty"`       // 操作者的用户id, * 必填字段(系统升级后, 新系统要求必填), 示例值: "dd31248a"
}

// GetAttendanceUserStatsViewResp ...
type GetAttendanceUserStatsViewResp struct {
	View *GetAttendanceUserStatsViewRespView `json:"view,omitempty"` // 统计视图
}

// GetAttendanceUserStatsViewRespView ...
type GetAttendanceUserStatsViewRespView struct {
	ViewID    string                                    `json:"view_id,omitempty"`    // 视图 ID
	StatsType string                                    `json:"stats_type,omitempty"` // 视图类型, 可选值有: daily: 日度统计, month: 月度统计
	UserID    string                                    `json:"user_id,omitempty"`    // 操作者的用户id
	Items     []*GetAttendanceUserStatsViewRespViewItem `json:"items,omitempty"`      // 用户设置字段
}

// GetAttendanceUserStatsViewRespViewItem ...
type GetAttendanceUserStatsViewRespViewItem struct {
	Code       string                                             `json:"code,omitempty"`        // 标题编号
	Title      string                                             `json:"title,omitempty"`       // 标题名称
	ChildItems []*GetAttendanceUserStatsViewRespViewItemChildItem `json:"child_items,omitempty"` // 子标题
}

// GetAttendanceUserStatsViewRespViewItemChildItem ...
type GetAttendanceUserStatsViewRespViewItemChildItem struct {
	Code       string `json:"code,omitempty"`        // 子标题编号
	Value      string `json:"value,omitempty"`       // 开关字段, 0: 关闭, 1: 开启（非开关字段场景: code = 51501 可选值为1-6）
	Title      string `json:"title,omitempty"`       // 子标题名称
	ColumnType int64  `json:"column_type,omitempty"` // 列类型
	ReadOnly   bool   `json:"read_only,omitempty"`   // 是否只读
	MinValue   string `json:"min_value,omitempty"`   // 最小值
	MaxValue   string `json:"max_value,omitempty"`   // 最大值
}

// getAttendanceUserStatsViewResp ...
type getAttendanceUserStatsViewResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserStatsViewResp `json:"data,omitempty"`
}
