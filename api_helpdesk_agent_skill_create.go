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

// CreateHelpdeskAgentSkill 该接口用于创建客服技能
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/create
func (r *HelpdeskService) CreateHelpdeskAgentSkill(ctx context.Context, request *CreateHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*CreateHelpdeskAgentSkillResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateHelpdeskAgentSkill != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#CreateHelpdeskAgentSkill mock enable")
		return r.cli.mock.mockHelpdeskCreateHelpdeskAgentSkill(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CreateHelpdeskAgentSkill",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_skills",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createHelpdeskAgentSkillResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCreateHelpdeskAgentSkill mock HelpdeskCreateHelpdeskAgentSkill method
func (r *Mock) MockHelpdeskCreateHelpdeskAgentSkill(f func(ctx context.Context, request *CreateHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*CreateHelpdeskAgentSkillResp, *Response, error)) {
	r.mockHelpdeskCreateHelpdeskAgentSkill = f
}

// UnMockHelpdeskCreateHelpdeskAgentSkill un-mock HelpdeskCreateHelpdeskAgentSkill method
func (r *Mock) UnMockHelpdeskCreateHelpdeskAgentSkill() {
	r.mockHelpdeskCreateHelpdeskAgentSkill = nil
}

// CreateHelpdeskAgentSkillReq ...
type CreateHelpdeskAgentSkillReq struct {
	Name     *string                            `json:"name,omitempty"`      // 技能名, 示例值: "test-skill"
	Rules    []*CreateHelpdeskAgentSkillReqRule `json:"rules,omitempty"`     // 技能rules
	AgentIDs []string                           `json:"agent_ids,omitempty"` // 客服 ids, 示例值: ["客服ID"]
}

// CreateHelpdeskAgentSkillReqRule ...
type CreateHelpdeskAgentSkillReqRule struct {
	ID               *string `json:"id,omitempty"`                // rule id, 参考[获取客服技能rules](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list) 用于获取rules options, 示例值: "test-skill-id"
	SelectedOperator *int64  `json:"selected_operator,omitempty"` // 运算符比较, 参考[客服技能运算符选项](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options), 示例值: 8
	Operand          *string `json:"operand,omitempty"`           // rule 操作数的值, 示例值: "{, "selected_departments": [, {, "id": "部门ID", "name": "IT", }, ], }"
	Category         *int64  `json:"category,omitempty"`          // rule 类型, 1-知识库, 2-工单信息, 3-用户飞书信息, 示例值: 3
}

// CreateHelpdeskAgentSkillResp ...
type CreateHelpdeskAgentSkillResp struct {
	AgentSkillID string `json:"agent_skill_id,omitempty"` // 客服技能id
}

// createHelpdeskAgentSkillResp ...
type createHelpdeskAgentSkillResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *CreateHelpdeskAgentSkillResp `json:"data,omitempty"`
}
