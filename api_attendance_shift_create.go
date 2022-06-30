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

// CreateAttendanceShift 班次是描述一次考勤任务时间规则的统称, 比如一天打多少次卡, 每次卡的上下班时间, 晚到多长时间算迟到, 晚到多长时间算缺卡等。
//
// - 创建一个考勤组前, 必须先创建一个或者多个班次。
// - 一个公司内的班次是共享的, 你可以直接引用他人创建的班次, 但是需要注意的是, 若他人修改了班次, 会影响到你的考勤组及其考勤结果。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create
func (r *AttendanceService) CreateAttendanceShift(ctx context.Context, request *CreateAttendanceShiftReq, options ...MethodOptionFunc) (*CreateAttendanceShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateAttendanceShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateAttendanceShift mock enable")
		return r.cli.mock.mockAttendanceCreateAttendanceShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateAttendanceShift",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/shifts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createAttendanceShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceCreateAttendanceShift mock AttendanceCreateAttendanceShift method
func (r *Mock) MockAttendanceCreateAttendanceShift(f func(ctx context.Context, request *CreateAttendanceShiftReq, options ...MethodOptionFunc) (*CreateAttendanceShiftResp, *Response, error)) {
	r.mockAttendanceCreateAttendanceShift = f
}

// UnMockAttendanceCreateAttendanceShift un-mock AttendanceCreateAttendanceShift method
func (r *Mock) UnMockAttendanceCreateAttendanceShift() {
	r.mockAttendanceCreateAttendanceShift = nil
}

// CreateAttendanceShiftReq ...
type CreateAttendanceShiftReq struct {
	ShiftName         string                                       `json:"shift_name,omitempty"`            // 班次名称, 示例值: "早班"
	PunchTimes        int64                                        `json:"punch_times,omitempty"`           // 打卡次数, 示例值: 1
	IsFlexible        *bool                                        `json:"is_flexible,omitempty"`           // 是否弹性打卡, 示例值: false
	FlexibleMinutes   *int64                                       `json:"flexible_minutes,omitempty"`      // 弹性打卡的时间, 示例值: 60
	NoNeedOff         *bool                                        `json:"no_need_off,omitempty"`           // 不需要打下班卡, 示例值: true
	PunchTimeRule     []*CreateAttendanceShiftReqPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule []*CreateAttendanceShiftReqLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      []*CreateAttendanceShiftReqRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息规则
}

// CreateAttendanceShiftReqLateOffLateOnRule ...
type CreateAttendanceShiftReqLateOffLateOnRule struct {
	LateOffMinutes int64 `json:"late_off_minutes,omitempty"` // 晚走多久, 示例值: 60
	LateOnMinutes  int64 `json:"late_on_minutes,omitempty"`  // 晚到多久, 示例值: 30
}

// CreateAttendanceShiftReqPunchTimeRule ...
type CreateAttendanceShiftReqPunchTimeRule struct {
	OnTime              string `json:"on_time,omitempty"`                // 上班时间, 示例值: "9:00"
	OffTime             string `json:"off_time,omitempty"`               // 下班时间, 示例值: "18:00, 第二天凌晨2点, 26:00"
	LateMinutesAsLate   int64  `json:"late_minutes_as_late,omitempty"`   // 晚到多久记为迟到, 示例值: 30
	LateMinutesAsLack   int64  `json:"late_minutes_as_lack,omitempty"`   // 晚到多久记为缺卡, 示例值: 60
	OnAdvanceMinutes    int64  `json:"on_advance_minutes,omitempty"`     // 最早多久可打上班卡, 示例值: 60
	EarlyMinutesAsEarly int64  `json:"early_minutes_as_early,omitempty"` // 早退多久记为早退, 示例值: 30
	EarlyMinutesAsLack  int64  `json:"early_minutes_as_lack,omitempty"`  // 早退多久记为缺卡, 示例值: 60
	OffDelayMinutes     int64  `json:"off_delay_minutes,omitempty"`      // 最晚多久可打下班卡, 示例值: 60
}

// CreateAttendanceShiftReqRestTimeRule ...
type CreateAttendanceShiftReqRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始, 示例值: "13:00"
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束, 示例值: "14:00"
}

// CreateAttendanceShiftResp ...
type CreateAttendanceShiftResp struct {
	Shift *CreateAttendanceShiftRespShift `json:"shift,omitempty"` // 班次
}

// CreateAttendanceShiftRespShift ...
type CreateAttendanceShiftRespShift struct {
	ShiftID           string                                             `json:"shift_id,omitempty"`              // 班次 ID
	ShiftName         string                                             `json:"shift_name,omitempty"`            // 班次名称
	PunchTimes        int64                                              `json:"punch_times,omitempty"`           // 打卡次数
	IsFlexible        bool                                               `json:"is_flexible,omitempty"`           // 是否弹性打卡
	FlexibleMinutes   int64                                              `json:"flexible_minutes,omitempty"`      // 弹性打卡的时间
	NoNeedOff         bool                                               `json:"no_need_off,omitempty"`           // 不需要打下班卡
	PunchTimeRule     []*CreateAttendanceShiftRespShiftPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule []*CreateAttendanceShiftRespShiftLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      []*CreateAttendanceShiftRespShiftRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息规则
}

// CreateAttendanceShiftRespShiftLateOffLateOnRule ...
type CreateAttendanceShiftRespShiftLateOffLateOnRule struct {
	LateOffMinutes int64 `json:"late_off_minutes,omitempty"` // 晚走多久
	LateOnMinutes  int64 `json:"late_on_minutes,omitempty"`  // 晚到多久
}

// CreateAttendanceShiftRespShiftPunchTimeRule ...
type CreateAttendanceShiftRespShiftPunchTimeRule struct {
	OnTime              string `json:"on_time,omitempty"`                // 上班时间
	OffTime             string `json:"off_time,omitempty"`               // 下班时间
	LateMinutesAsLate   int64  `json:"late_minutes_as_late,omitempty"`   // 晚到多久记为迟到
	LateMinutesAsLack   int64  `json:"late_minutes_as_lack,omitempty"`   // 晚到多久记为缺卡
	OnAdvanceMinutes    int64  `json:"on_advance_minutes,omitempty"`     // 最早多久可打上班卡
	EarlyMinutesAsEarly int64  `json:"early_minutes_as_early,omitempty"` // 早退多久记为早退
	EarlyMinutesAsLack  int64  `json:"early_minutes_as_lack,omitempty"`  // 早退多久记为缺卡
	OffDelayMinutes     int64  `json:"off_delay_minutes,omitempty"`      // 最晚多久可打下班卡
}

// CreateAttendanceShiftRespShiftRestTimeRule ...
type CreateAttendanceShiftRespShiftRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束
}

// createAttendanceShiftResp ...
type createAttendanceShiftResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateAttendanceShiftResp `json:"data,omitempty"`
}
