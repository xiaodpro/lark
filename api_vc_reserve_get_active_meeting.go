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

// GetVCReserveActiveMeeting 获取一个预约的当前活跃会议。
//
// 只能获取归属于自己的预约的活跃会议（一个预约最多有一个正在进行中的会议）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get_active_meeting
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/reserve/get_active_meeting
func (r *VCService) GetVCReserveActiveMeeting(ctx context.Context, request *GetVCReserveActiveMeetingReq, options ...MethodOptionFunc) (*GetVCReserveActiveMeetingResp, *Response, error) {
	if r.cli.mock.mockVCGetVCReserveActiveMeeting != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#GetVCReserveActiveMeeting mock enable")
		return r.cli.mock.mockVCGetVCReserveActiveMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCReserveActiveMeeting",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reserves/:reserve_id/get_active_meeting",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCReserveActiveMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCReserveActiveMeeting mock VCGetVCReserveActiveMeeting method
func (r *Mock) MockVCGetVCReserveActiveMeeting(f func(ctx context.Context, request *GetVCReserveActiveMeetingReq, options ...MethodOptionFunc) (*GetVCReserveActiveMeetingResp, *Response, error)) {
	r.mockVCGetVCReserveActiveMeeting = f
}

// UnMockVCGetVCReserveActiveMeeting un-mock VCGetVCReserveActiveMeeting method
func (r *Mock) UnMockVCGetVCReserveActiveMeeting() {
	r.mockVCGetVCReserveActiveMeeting = nil
}

// GetVCReserveActiveMeetingReq ...
type GetVCReserveActiveMeetingReq struct {
	ReserveID        string  `path:"reserve_id" json:"-"`         // 预约ID（预约的唯一标识）, 示例值: "6911188411932033028"
	WithParticipants *bool   `query:"with_participants" json:"-"` // 是否需要参会人列表, 默认为false, 示例值: false
	UserIDType       *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCReserveActiveMeetingResp ...
type GetVCReserveActiveMeetingResp struct {
	Meeting *GetVCReserveActiveMeetingRespMeeting `json:"meeting,omitempty"` // 会议数据
}

// GetVCReserveActiveMeetingRespMeeting ...
type GetVCReserveActiveMeetingRespMeeting struct {
	ID                          string                                             `json:"id,omitempty"`                            // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）
	Topic                       string                                             `json:"topic,omitempty"`                         // 会议主题
	URL                         string                                             `json:"url,omitempty"`                           // 会议链接（飞书用户可通过点击会议链接快捷入会）
	MeetingNo                   string                                             `json:"meeting_no,omitempty"`                    // 会议号
	Password                    string                                             `json:"password,omitempty"`                      // 会议密码
	CreateTime                  string                                             `json:"create_time,omitempty"`                   // 会议创建时间（unix时间, 单位sec）
	StartTime                   string                                             `json:"start_time,omitempty"`                    // 会议开始时间（unix时间, 单位sec）
	EndTime                     string                                             `json:"end_time,omitempty"`                      // 会议结束时间（unix时间, 单位sec）
	HostUser                    *GetVCReserveActiveMeetingRespMeetingHostUser      `json:"host_user,omitempty"`                     // 主持人
	Status                      int64                                              `json:"status,omitempty"`                        // 会议状态, 可选值有: 1: 会议呼叫中, 2: 会议进行中, 3: 会议已结束
	ParticipantCount            string                                             `json:"participant_count,omitempty"`             // 参会峰值人数
	ParticipantCountAccumulated string                                             `json:"participant_count_accumulated,omitempty"` // 累计参会人数
	Participants                []*GetVCReserveActiveMeetingRespMeetingParticipant `json:"participants,omitempty"`                  // 参会人列表
	Ability                     *GetVCReserveActiveMeetingRespMeetingAbility       `json:"ability,omitempty"`                       // 会中使用的能力
}

// GetVCReserveActiveMeetingRespMeetingAbility ...
type GetVCReserveActiveMeetingRespMeetingAbility struct {
	UseVideo        bool `json:"use_video,omitempty"`         // 是否使用视频
	UseAudio        bool `json:"use_audio,omitempty"`         // 是否使用音频
	UseShareScreen  bool `json:"use_share_screen,omitempty"`  // 是否使用共享屏幕
	UseFollowScreen bool `json:"use_follow_screen,omitempty"` // 是否使用妙享（magic share）
	UseRecording    bool `json:"use_recording,omitempty"`     // 是否使用录制
	UsePstn         bool `json:"use_pstn,omitempty"`          // 是否使用PSTN
}

// GetVCReserveActiveMeetingRespMeetingHostUser ...
type GetVCReserveActiveMeetingRespMeetingHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: 飞书用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// GetVCReserveActiveMeetingRespMeetingParticipant ...
type GetVCReserveActiveMeetingRespMeetingParticipant struct {
	ID                string `json:"id,omitempty"`                  // 用户ID
	FirstJoinTime     string `json:"first_join_time,omitempty"`     // 首次入会时间, 秒级Unix时间戳
	FinalLeaveTime    string `json:"final_leave_time,omitempty"`    // 最终离会时间, 秒级Unix时间戳
	InMeetingDuration string `json:"in_meeting_duration,omitempty"` // 累计在会中时间, 时间单位: 秒
	UserType          int64  `json:"user_type,omitempty"`           // 用户类型, 可选值有: 1: 飞书用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
	IsHost            bool   `json:"is_host,omitempty"`             // 是否为主持人
	IsCohost          bool   `json:"is_cohost,omitempty"`           // 是否为联席主持人
	IsExternal        bool   `json:"is_external,omitempty"`         // 是否为外部参会人
	Status            int64  `json:"status,omitempty"`              // 参会人状态, 可选值有: 1: 呼叫中, 2: 在会中, 3: 正在响铃, 4: 不在会中或已经离开会议
}

// getVCReserveActiveMeetingResp ...
type getVCReserveActiveMeetingResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *GetVCReserveActiveMeetingResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
