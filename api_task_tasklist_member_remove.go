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

// RemoveTaskTasklistMember 移除清单的一个或多个协作成员。通过设置`members`字段表示要移除的成员信息。关于member的格式, 详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何表示任务和清单的成员？”章节。
//
// 清单中同一个成员只能有一个角色, 通过的member的id和type可以唯一确定一个成员, 因此请求参数中对于要删除的成员, 不需要填写"role"字段。
// 如果要移除的成员不在清单中, 则被自动忽略, 接口返回成功。
// 该接口不能用于移除清单所有者。如果要移除的成员是清单所有者, 则会被自动忽略。如要设置清单所有者, 需要调用[更新清单](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/patch)接口。
// 需要清单编辑权限。详情见[清单功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/overview)中的“清单是如何鉴权的？“章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/remove_members
func (r *TaskService) RemoveTaskTasklistMember(ctx context.Context, request *RemoveTaskTasklistMemberReq, options ...MethodOptionFunc) (*RemoveTaskTasklistMemberResp, *Response, error) {
	if r.cli.mock.mockTaskRemoveTaskTasklistMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#RemoveTaskTasklistMember mock enable")
		return r.cli.mock.mockTaskRemoveTaskTasklistMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "RemoveTaskTasklistMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasklists/:tasklist_guid/remove_members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(removeTaskTasklistMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskRemoveTaskTasklistMember mock TaskRemoveTaskTasklistMember method
func (r *Mock) MockTaskRemoveTaskTasklistMember(f func(ctx context.Context, request *RemoveTaskTasklistMemberReq, options ...MethodOptionFunc) (*RemoveTaskTasklistMemberResp, *Response, error)) {
	r.mockTaskRemoveTaskTasklistMember = f
}

// UnMockTaskRemoveTaskTasklistMember un-mock TaskRemoveTaskTasklistMember method
func (r *Mock) UnMockTaskRemoveTaskTasklistMember() {
	r.mockTaskRemoveTaskTasklistMember = nil
}

// RemoveTaskTasklistMemberReq ...
type RemoveTaskTasklistMemberReq struct {
	TasklistGuid string                               `path:"tasklist_guid" json:"-"` // 要移除协作人的清单全局唯一ID, 示例值: "d300a75f-c56a-4be9-80d1-e47653028ceb", 最大长度: `100` 字符
	UserIDType   *IDType                              `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	Members      []*RemoveTaskTasklistMemberReqMember `json:"members,omitempty"`      // 要移除的member列表。关于member的格式, 详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何表示任务和清单的成员？”章节, 长度范围: `1` ～ `500`
}

// RemoveTaskTasklistMemberReqMember ...
type RemoveTaskTasklistMemberReqMember struct {
	ID   *string `json:"id,omitempty"`   // 表示member的id, 示例值: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f", 最大长度: `100` 字符
	Type *string `json:"type,omitempty"` // 成员的类型, 支持: user: 普通用户, 此时member的id是一个表示用户的ID, 比如open_id。具体格式取决于user_id_type参数, chat: 群组, 此时member的id是一个Open Chat ID, app: 应用, 此时member的id是一个应用的ID, 示例值: "user", 默认值: `user`
	Role *string `json:"role,omitempty"` // 清单角色。移除清单成员时该字段不需要填写, 示例值: "editor", 最大长度: `20` 字符
}

// RemoveTaskTasklistMemberResp ...
type RemoveTaskTasklistMemberResp struct {
	Tasklist *RemoveTaskTasklistMemberRespTasklist `json:"tasklist,omitempty"` // 修改完成后的清单实体
}

// RemoveTaskTasklistMemberRespTasklist ...
type RemoveTaskTasklistMemberRespTasklist struct {
	Guid      string                                        `json:"guid,omitempty"`       // 清单的全局唯一ID
	Name      string                                        `json:"name,omitempty"`       // 清单名
	Creator   *RemoveTaskTasklistMemberRespTasklistCreator  `json:"creator,omitempty"`    // 清单创建者
	Owner     *RemoveTaskTasklistMemberRespTasklistOwner    `json:"owner,omitempty"`      // 清单所有者
	Members   []*RemoveTaskTasklistMemberRespTasklistMember `json:"members,omitempty"`    // 清单协作成员列表
	URL       string                                        `json:"url,omitempty"`        // 该清单分享的applink
	CreatedAt string                                        `json:"created_at,omitempty"` // 清单创建时间戳(ms)
	UpdatedAt string                                        `json:"updated_at,omitempty"` // 清单最后一次更新时间戳（ms)
}

// RemoveTaskTasklistMemberRespTasklistCreator ...
type RemoveTaskTasklistMemberRespTasklistCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员类型
	Role string `json:"role,omitempty"` // 成员角色
}

// RemoveTaskTasklistMemberRespTasklistMember ...
type RemoveTaskTasklistMemberRespTasklistMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员类型
	Role string `json:"role,omitempty"` // 成员角色
}

// RemoveTaskTasklistMemberRespTasklistOwner ...
type RemoveTaskTasklistMemberRespTasklistOwner struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员类型
	Role string `json:"role,omitempty"` // 成员角色
}

// removeTaskTasklistMemberResp ...
type removeTaskTasklistMemberResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *RemoveTaskTasklistMemberResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}