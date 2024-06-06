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

// GetCalendarEventInstanceList 调用该接口以当前身份（应用或用户）获取指定日历中的某一重复日程信息。
//
// 当前身份由 Header Authorization 的 Token 类型决定。tenant_access_token 指应用身份, user_access_token 指用户身份。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/instances
func (r *CalendarService) GetCalendarEventInstanceList(ctx context.Context, request *GetCalendarEventInstanceListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventInstanceList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventInstanceList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventInstanceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventInstanceList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/instances",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventInstanceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventInstanceList mock CalendarGetCalendarEventInstanceList method
func (r *Mock) MockCalendarGetCalendarEventInstanceList(f func(ctx context.Context, request *GetCalendarEventInstanceListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventInstanceList = f
}

// UnMockCalendarGetCalendarEventInstanceList un-mock CalendarGetCalendarEventInstanceList method
func (r *Mock) UnMockCalendarGetCalendarEventInstanceList() {
	r.mockCalendarGetCalendarEventInstanceList = nil
}

// GetCalendarEventInstanceListReq ...
type GetCalendarEventInstanceListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"` // 日历 ID。关于日历 ID 可参见[日历 ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_HF9U2MbibE8PPpjro6xjqa@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`    // 日程 ID, 创建日程时会返回日程 ID。你也可以调用以下接口获取某一日历的 ID, [获取日程列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list), [搜索日程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search), 示例值: "75d28f9b-e35c-4230-8a83-4a661497db54_0"
	StartTime  string  `query:"start_time" json:"-"` // 开始时间, Unix 时间戳, 单位为秒。该参数与 end_time 用于设置时间范围, 即重复日程的查询区间为 （start_time, end_time）, 注意: start_time 与 end_time 之间的时间区间不能超过 2年, 示例值: 1631777271
	EndTime    string  `query:"end_time" json:"-"`   // 结束时间, Unix 时间戳, 单位为秒。该参数与 start_time 用于设置时间范围, 即重复日程的查询区间为 （start_time, end_time）, 注意: start_time 与 end_time 之间的时间区间不能超过 2年, 示例值: 1631777271
	PageSize   *int64  `query:"page_size" json:"-"`  // 一次调用返回的日程数量上限, 示例值: 10, 默认值: `50`, 取值范围: `10` ～ `500`
	PageToken  *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0=
}

// GetCalendarEventInstanceListResp ...
type GetCalendarEventInstanceListResp struct {
	Items     []*GetCalendarEventInstanceListRespItem `json:"items,omitempty"`      // 重复日程的日程 instance 列表。
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetCalendarEventInstanceListRespItem ...
type GetCalendarEventInstanceListRespItem struct {
	EventID     string                                         `json:"event_id,omitempty"`     // 日程实例 ID, 注意: 重复日程实例的 ID 与其他日程 ID 不同, 其 ID 包含了实例原始时间（Original time）, 数据格式为秒级时间戳。例如: `2cf525f0-1e67-4b04-ad4d-30b7f003903c_1713168000`, 其中 `1713168000` 即为实例原始时间。
	Summary     string                                         `json:"summary,omitempty"`      // 日程主题。
	Description string                                         `json:"description,omitempty"`  // 日程描述。
	StartTime   *GetCalendarEventInstanceListRespItemStartTime `json:"start_time,omitempty"`   // 日程开始时间。
	EndTime     *GetCalendarEventInstanceListRespItemEndTime   `json:"end_time,omitempty"`     // 日程结束时间。
	Status      string                                         `json:"status,omitempty"`       // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException bool                                           `json:"is_exception,omitempty"` // 日程是否是重复日程的例外日程。了解例外日程, 可参见[例外日程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction#71c5ec78)。
	AppLink     string                                         `json:"app_link,omitempty"`     // 日程的 app_link, 用于跳转到具体的某个日程。
	Location    *GetCalendarEventInstanceListRespItemLocation  `json:"location,omitempty"`     // 日程地点。
}

// GetCalendarEventInstanceListRespItemEndTime ...
type GetCalendarEventInstanceListRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 结束时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 指日程具体的结束时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）。
	Timezone  string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准。
}

// GetCalendarEventInstanceListRespItemLocation ...
type GetCalendarEventInstanceListRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称。
	Address   string  `json:"address,omitempty"`   // 地点地址。
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准
}

// GetCalendarEventInstanceListRespItemStartTime ...
type GetCalendarEventInstanceListRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 开始时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 指日程具体的开始时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）。
	Timezone  string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准。
}

// getCalendarEventInstanceListResp ...
type getCalendarEventInstanceListResp struct {
	Code  int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                            `json:"msg,omitempty"`  // 错误描述
	Data  *GetCalendarEventInstanceListResp `json:"data,omitempty"`
	Error *ErrorDetail                      `json:"error,omitempty"`
}
