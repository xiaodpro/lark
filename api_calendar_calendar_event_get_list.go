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

// GetCalendarEventList 该接口用于以当前身份（应用 / 用户）获取日历下的日程列表。
//
// 身份由 Header Authorization 的 Token 类型决定。
// - 当前身份必须对日历有reader、writer或owner权限才会返回日程详细信息（调用[获取日历](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, role字段可查看权限）。
// - 仅支持primary、shared和resource类型的日历获取日程列表。
// - 调用时首先使用 page_token 分页拉取存量数据, 之后使用 sync_token 增量同步变更数据。
// - 为了确保调用方日程同步数据的一致性, 在使用sync_token时, 不能同时使用start_time和end_time, 否则可能造成日程数据缺失。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list
func (r *CalendarService) GetCalendarEventList(ctx context.Context, request *GetCalendarEventListReq, options ...MethodOptionFunc) (*GetCalendarEventListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventList mock CalendarGetCalendarEventList method
func (r *Mock) MockCalendarGetCalendarEventList(f func(ctx context.Context, request *GetCalendarEventListReq, options ...MethodOptionFunc) (*GetCalendarEventListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventList = f
}

// UnMockCalendarGetCalendarEventList un-mock CalendarGetCalendarEventList method
func (r *Mock) UnMockCalendarGetCalendarEventList() {
	r.mockCalendarGetCalendarEventList = nil
}

// GetCalendarEventListReq ...
type GetCalendarEventListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"`  // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	PageSize   *int64  `query:"page_size" json:"-"`   // 一次请求要求返回最大数量, 默认500, 取值范围为[50, 1000], 示例值: 50, 默认值: `500`, 取值范围: `50` ～ `1000`
	AnchorTime *string `query:"anchor_time" json:"-"` // 拉取anchor_time之后的日程, 为timestamp, 示例值: "1609430400"
	PageToken  *string `query:"page_token" json:"-"`  // 上次请求Response返回的分页标记, 首次请求时为空, 示例值: "ListCalendarsPageToken_1632452910_1632539310"
	SyncToken  *string `query:"sync_token" json:"-"`  // 上次请求Response返回的增量同步标记, 分页请求未结束时为空, 示例值: "ListCalendarsSyncToken_1632452910"
	StartTime  *string `query:"start_time" json:"-"`  // 日程开始Unix时间戳, 单位为秒, 示例值: "1631777271"
	EndTime    *string `query:"end_time" json:"-"`    // 日程结束Unix时间戳, 单位为秒, 示例值: "1631777271"
}

// GetCalendarEventListResp ...
type GetCalendarEventListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否有下一页数据
	PageToken string                          `json:"page_token,omitempty"` // 下次请求需要带上的分页标记, 90 天有效期
	SyncToken string                          `json:"sync_token,omitempty"` // 下次请求需要带上的增量同步标记, 90 天有效期
	Items     []*GetCalendarEventListRespItem `json:"items,omitempty"`      // 日程列表
}

// GetCalendarEventListRespItem ...
type GetCalendarEventListRespItem struct {
	EventID             string                                  `json:"event_id,omitempty"`              // 日程ID。参见[日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction)
	OrganizerCalendarID string                                  `json:"organizer_calendar_id,omitempty"` // 日程组织者日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction)
	Summary             string                                  `json:"summary,omitempty"`               // 日程标题
	Description         string                                  `json:"description,omitempty"`           // 日程描述；目前不支持编辑富文本描述, 如果日程描述通过客户端编辑过, 更新描述会导致富文本格式丢失
	StartTime           *GetCalendarEventListRespItemStartTime  `json:"start_time,omitempty"`            // 日程开始时间
	EndTime             *GetCalendarEventListRespItemEndTime    `json:"end_time,omitempty"`              // 日程结束时间
	Vchat               *GetCalendarEventListRespItemVchat      `json:"vchat,omitempty"`                 // 视频会议信息。
	Visibility          string                                  `json:"visibility,omitempty"`            // 日程公开范围, 新建日程默认为Default；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: default: 默认权限, 跟随日历权限, 默认仅向他人显示是否“忙碌”, public: 公开, 显示日程详情, private: 私密, 仅自己可见详情
	AttendeeAbility     string                                  `json:"attendee_ability,omitempty"`      // 参与人权限, 可选值有: none: 无法编辑日程、无法邀请其它参与人、无法查看参与人列表, can_see_others: 无法编辑日程、无法邀请其它参与人、可以查看参与人列表, can_invite_others: 无法编辑日程、可以邀请其它参与人、可以查看参与人列表, can_modify_event: 可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus      string                                  `json:"free_busy_status,omitempty"`      // 日程占用的忙闲状态, 新建日程默认为Busy；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: busy: 忙碌, free: 空闲
	Location            *GetCalendarEventListRespItemLocation   `json:"location,omitempty"`              // 日程地点
	Color               int64                                   `json:"color,omitempty"`                 // 日程颜色, 颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders           []*GetCalendarEventListRespItemReminder `json:"reminders,omitempty"`             // 日程提醒列表
	Recurrence          string                                  `json:"recurrence,omitempty"`            // 重复日程的重复性规则；参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)；, 不支持COUNT和UNTIL同时出现；, 预定会议室重复日程长度不得超过两年。
	Status              string                                  `json:"status,omitempty"`                // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException         bool                                    `json:"is_exception,omitempty"`          // 日程是否是一个重复日程的例外日程
	RecurringEventID    string                                  `json:"recurring_event_id,omitempty"`    // 例外日程的原重复日程的event_id
	Schemas             []*GetCalendarEventListRespItemSchema   `json:"schemas,omitempty"`               // 日程自定义信息；控制日程详情页的ui展示。
}

// GetCalendarEventListRespItemEndTime ...
type GetCalendarEventListRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// GetCalendarEventListRespItemLocation ...
type GetCalendarEventListRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称
	Address   string  `json:"address,omitempty"`   // 地点地址
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
}

// GetCalendarEventListRespItemReminder ...
type GetCalendarEventListRespItemReminder struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量, 正数时表示在日程开始前X分钟提醒, 负数时表示在日程开始后X分钟提醒, 新建或更新日程时传入该字段, 仅对当前身份生效
}

// GetCalendarEventListRespItemSchema ...
type GetCalendarEventListRespItemSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI名称。取值范围如下: ForwardIcon: 日程转发按钮, MeetingChatIcon: 会议群聊按钮, MeetingMinutesIcon: 会议纪要按钮, MeetingVideo: 视频会议区域, RSVP: 接受/拒绝/待定区域, Attendee: 参与者区域, OrganizerOrCreator: 组织者/创建者区域
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态。目前只支持hide, 可选值有: hide: 隐藏显示, readonly: 只读, editable: 可编辑, unknown: 未知UI项自定义状态, 仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接; 该字段暂不支持传入。
}

// GetCalendarEventListRespItemStartTime ...
type GetCalendarEventListRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// GetCalendarEventListRespItemVchat ...
type GetCalendarEventListRespItemVchat struct {
	VCType      string `json:"vc_type,omitempty"`     // 视频会议类型, 可选值有: vc: 飞书视频会议, 取该类型时, 其他字段无效。, third_party: 第三方链接视频会议, 取该类型时, icon_type、description、meeting_url字段生效。, no_meeting: 无视频会议, 取该类型时, 其他字段无效。, lark_live: 飞书直播, 内部类型, 飞书客户端使用, API不支持创建, 只读。, unknown: 未知类型, 做兼容使用, 飞书客户端使用, API不支持创建, 只读。
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空, 为空展示默认icon, 可选值有: vc: 飞书视频会议icon, live: 直播视频会议icon, default: 默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案, 可以为空, 为空展示默认文案
	MeetingURL  string `json:"meeting_url,omitempty"` // 视频会议URL
}

// getCalendarEventListResp ...
type getCalendarEventListResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventListResp `json:"data,omitempty"`
}