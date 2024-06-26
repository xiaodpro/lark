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

// GetAttendanceShiftList 翻页获取所有班次列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/list
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/shift/list
func (r *AttendanceService) GetAttendanceShiftList(ctx context.Context, request *GetAttendanceShiftListReq, options ...MethodOptionFunc) (*GetAttendanceShiftListResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceShiftList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceShiftList mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceShiftList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceShiftList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/shifts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceShiftListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceShiftList mock AttendanceGetAttendanceShiftList method
func (r *Mock) MockAttendanceGetAttendanceShiftList(f func(ctx context.Context, request *GetAttendanceShiftListReq, options ...MethodOptionFunc) (*GetAttendanceShiftListResp, *Response, error)) {
	r.mockAttendanceGetAttendanceShiftList = f
}

// UnMockAttendanceGetAttendanceShiftList un-mock AttendanceGetAttendanceShiftList method
func (r *Mock) UnMockAttendanceGetAttendanceShiftList() {
	r.mockAttendanceGetAttendanceShiftList = nil
}

// GetAttendanceShiftListReq ...
type GetAttendanceShiftListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `10`, 取值范围: `1` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: YrkvQ1wGaPVta45tkxuGiQ[
}

// GetAttendanceShiftListResp ...
type GetAttendanceShiftListResp struct {
	ShiftList []*GetAttendanceShiftListRespShift `json:"shift_list,omitempty"` // 班次列表
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetAttendanceShiftListRespShift ...
type GetAttendanceShiftListRespShift struct {
	ShiftID           string                                              `json:"shift_id,omitempty"`              // 班次 ID
	ShiftName         string                                              `json:"shift_name,omitempty"`            // 班次名称
	PunchTimes        int64                                               `json:"punch_times,omitempty"`           // 打卡次数
	IsFlexible        bool                                                `json:"is_flexible,omitempty"`           // 是否弹性打卡
	FlexibleMinutes   int64                                               `json:"flexible_minutes,omitempty"`      // 弹性打卡时间, 设置[上班最多可晚到]与[下班最多可早走]时间, 如果不设置flexible_rule则生效
	FlexibleRule      []*GetAttendanceShiftListRespShiftFlexibleRule      `json:"flexible_rule,omitempty"`         // 弹性打卡时间设置
	NoNeedOff         bool                                                `json:"no_need_off,omitempty"`           // 不需要打下班卡
	PunchTimeRule     []*GetAttendanceShiftListRespShiftPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule []*GetAttendanceShiftListRespShiftLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      []*GetAttendanceShiftListRespShiftRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息规则
	OvertimeRule      []*GetAttendanceShiftListRespShiftOvertimeRule      `json:"overtime_rule,omitempty"`         // 打卡规则
}

// GetAttendanceShiftListRespShiftFlexibleRule ...
type GetAttendanceShiftListRespShiftFlexibleRule struct {
	FlexibleEarlyMinutes int64 `json:"flexible_early_minutes,omitempty"` // 下班最多可早走（上班早到几分钟, 下班可早走几分钟）
	FlexibleLateMinutes  int64 `json:"flexible_late_minutes,omitempty"`  // 上班最多可晚到（上班晚到几分钟, 下班须晚走几分钟）
}

// GetAttendanceShiftListRespShiftLateOffLateOnRule ...
type GetAttendanceShiftListRespShiftLateOffLateOnRule struct {
	LateOffMinutes int64 `json:"late_off_minutes,omitempty"` // 晚走多久
	LateOnMinutes  int64 `json:"late_on_minutes,omitempty"`  // 晚到多久
}

// GetAttendanceShiftListRespShiftOvertimeRule ...
type GetAttendanceShiftListRespShiftOvertimeRule struct {
	OnOvertime  string `json:"on_overtime,omitempty"`  // 上班时间
	OffOvertime string `json:"off_overtime,omitempty"` // 下班时间
}

// GetAttendanceShiftListRespShiftPunchTimeRule ...
type GetAttendanceShiftListRespShiftPunchTimeRule struct {
	OnTime                   string `json:"on_time,omitempty"`                      // 上班时间
	OffTime                  string `json:"off_time,omitempty"`                     // 下班时间
	LateMinutesAsLate        int64  `json:"late_minutes_as_late,omitempty"`         // 晚到多久记为迟到
	LateMinutesAsLack        int64  `json:"late_minutes_as_lack,omitempty"`         // 晚到多久记为缺卡
	OnAdvanceMinutes         int64  `json:"on_advance_minutes,omitempty"`           // 最早多久可打上班卡
	EarlyMinutesAsEarly      int64  `json:"early_minutes_as_early,omitempty"`       // 早退多久记为早退
	EarlyMinutesAsLack       int64  `json:"early_minutes_as_lack,omitempty"`        // 早退多久记为缺卡
	OffDelayMinutes          int64  `json:"off_delay_minutes,omitempty"`            // 最晚多久可打下班卡
	LateMinutesAsSeriousLate int64  `json:"late_minutes_as_serious_late,omitempty"` // 晚到多久记为严重迟到
}

// GetAttendanceShiftListRespShiftRestTimeRule ...
type GetAttendanceShiftListRespShiftRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束
}

// getAttendanceShiftListResp ...
type getAttendanceShiftListResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *GetAttendanceShiftListResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
