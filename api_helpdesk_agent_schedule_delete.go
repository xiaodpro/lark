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

// DeleteHelpdeskAgentSchedule 该接口用于删除客服
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/delete
func (r *HelpdeskService) DeleteHelpdeskAgentSchedule(ctx context.Context, request *DeleteHelpdeskAgentScheduleReq, options ...MethodOptionFunc) (*DeleteHelpdeskAgentScheduleResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDeleteHelpdeskAgentSchedule != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteHelpdeskAgentSchedule mock enable")
		return r.cli.mock.mockHelpdeskDeleteHelpdeskAgentSchedule(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "DeleteHelpdeskAgentSchedule",
		Method:              "DELETE",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/agents/:agent_id/schedules",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(deleteHelpdeskAgentScheduleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskDeleteHelpdeskAgentSchedule mock HelpdeskDeleteHelpdeskAgentSchedule method
func (r *Mock) MockHelpdeskDeleteHelpdeskAgentSchedule(f func(ctx context.Context, request *DeleteHelpdeskAgentScheduleReq, options ...MethodOptionFunc) (*DeleteHelpdeskAgentScheduleResp, *Response, error)) {
	r.mockHelpdeskDeleteHelpdeskAgentSchedule = f
}

// UnMockHelpdeskDeleteHelpdeskAgentSchedule un-mock HelpdeskDeleteHelpdeskAgentSchedule method
func (r *Mock) UnMockHelpdeskDeleteHelpdeskAgentSchedule() {
	r.mockHelpdeskDeleteHelpdeskAgentSchedule = nil
}

// DeleteHelpdeskAgentScheduleReq ...
type DeleteHelpdeskAgentScheduleReq struct {
	AgentID string `path:"agent_id" json:"-"` // agent user id, 示例值: "12345"
}

// DeleteHelpdeskAgentScheduleResp ...
type DeleteHelpdeskAgentScheduleResp struct {
}

// deleteHelpdeskAgentScheduleResp ...
type deleteHelpdeskAgentScheduleResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteHelpdeskAgentScheduleResp `json:"data,omitempty"`
}