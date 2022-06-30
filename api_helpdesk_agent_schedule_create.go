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

// CreateHelpdeskAgentSchedule 该接口用于创建客服
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/create
func (r *HelpdeskService) CreateHelpdeskAgentSchedule(ctx context.Context, request *CreateHelpdeskAgentScheduleReq, options ...MethodOptionFunc) (*CreateHelpdeskAgentScheduleResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateHelpdeskAgentSchedule != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#CreateHelpdeskAgentSchedule mock enable")
		return r.cli.mock.mockHelpdeskCreateHelpdeskAgentSchedule(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CreateHelpdeskAgentSchedule",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_schedules",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createHelpdeskAgentScheduleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCreateHelpdeskAgentSchedule mock HelpdeskCreateHelpdeskAgentSchedule method
func (r *Mock) MockHelpdeskCreateHelpdeskAgentSchedule(f func(ctx context.Context, request *CreateHelpdeskAgentScheduleReq, options ...MethodOptionFunc) (*CreateHelpdeskAgentScheduleResp, *Response, error)) {
	r.mockHelpdeskCreateHelpdeskAgentSchedule = f
}

// UnMockHelpdeskCreateHelpdeskAgentSchedule un-mock HelpdeskCreateHelpdeskAgentSchedule method
func (r *Mock) UnMockHelpdeskCreateHelpdeskAgentSchedule() {
	r.mockHelpdeskCreateHelpdeskAgentSchedule = nil
}

// CreateHelpdeskAgentScheduleReq ...
type CreateHelpdeskAgentScheduleReq struct {
	AgentSchedules []*CreateHelpdeskAgentScheduleReqAgentSchedule `json:"agent_schedules,omitempty"` // 新客服日程
}

// CreateHelpdeskAgentScheduleReqAgentSchedule ...
type CreateHelpdeskAgentScheduleReqAgentSchedule struct {
	AgentID       *string                                                `json:"agent_id,omitempty"`        // 客服id, [可以以普通用户身份在服务台发起工单, 从工单详情里面获取用户guest.id](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get), 示例值: "agent-id"
	Schedule      []*CreateHelpdeskAgentScheduleReqAgentScheduleSchedule `json:"schedule,omitempty"`        // 工作日程列表
	AgentSkillIDs []string                                               `json:"agent_skill_ids,omitempty"` // 客服技能 ids
}

// CreateHelpdeskAgentScheduleReqAgentScheduleSchedule ...
type CreateHelpdeskAgentScheduleReqAgentScheduleSchedule struct {
	StartTime *string `json:"start_time,omitempty"` // 开始时间, format 00:00 - 23:59, 示例值: "00:00"
	EndTime   *string `json:"end_time,omitempty"`   // 结束时间, format 00:00 - 23:59, 示例值: "24:00"
	Weekday   *int64  `json:"weekday,omitempty"`    // 星期几, 1 - Monday, 2 - Tuesday, 3 - Wednesday, 4 - Thursday, 5 - Friday, 6 - Saturday, 7 - Sunday, 9 - Everday, 10 - Weekday, 11 - Weekend, 示例值: 9
}

// CreateHelpdeskAgentScheduleResp ...
type CreateHelpdeskAgentScheduleResp struct {
}

// createHelpdeskAgentScheduleResp ...
type createHelpdeskAgentScheduleResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateHelpdeskAgentScheduleResp `json:"data,omitempty"`
}
