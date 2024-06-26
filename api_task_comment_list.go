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

// GetTaskCommentList 给定一个资源, 返回该资源的评论列表。
//
// 支持分页。评论可以按照创建时间的正序（asc, 从最老到最新）, 或者逆序（desc, 从最老到最新）, 返回数据。
// 获取任务的评论列表需要任务的读取权限, 详见[任务是如何鉴权的？](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/faq)
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/list
func (r *TaskService) GetTaskCommentList(ctx context.Context, request *GetTaskCommentListReq, options ...MethodOptionFunc) (*GetTaskCommentListResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskCommentList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#GetTaskCommentList mock enable")
		return r.cli.mock.mockTaskGetTaskCommentList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskCommentList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskCommentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTaskCommentList mock TaskGetTaskCommentList method
func (r *Mock) MockTaskGetTaskCommentList(f func(ctx context.Context, request *GetTaskCommentListReq, options ...MethodOptionFunc) (*GetTaskCommentListResp, *Response, error)) {
	r.mockTaskGetTaskCommentList = f
}

// UnMockTaskGetTaskCommentList un-mock TaskGetTaskCommentList method
func (r *Mock) UnMockTaskGetTaskCommentList() {
	r.mockTaskGetTaskCommentList = nil
}

// GetTaskCommentListReq ...
type GetTaskCommentListReq struct {
	PageSize     *int64  `query:"page_size" json:"-"`     // 分页大小, 默认为50, 示例值: 50, 默认值: `50`, 取值范围: `1` ～ `100`
	PageToken    *string `query:"page_token" json:"-"`    // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: aWQ9NzEwMjMzMjMxMDE=, 最大长度: `100` 字符
	ResourceType *string `query:"resource_type" json:"-"` // 要获取评论列表的资源类型, 目前只支持"task", 默认为"task", 示例值: task, 默认值: `task`
	ResourceID   string  `query:"resource_id" json:"-"`   // 要获取评论的资源ID。例如要获取任务的评论列表, 此处应该填写任务全局唯一ID, 示例值: d300a75f-c56a-4be9-80d1-e47653028ceb
	Direction    *string `query:"direction" json:"-"`     // 返回数据的排序方式。"asc"表示从最老到最新顺序返回；"desc"表示从最新到最老顺序返回。默认为"asc", 示例值: asc, 可选值有: asc: 评论发表时间升序, desc: 评论发表时间降序, 默认值: `asc`
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
}

// GetTaskCommentListResp ...
type GetTaskCommentListResp struct {
	Items     []*GetTaskCommentListRespItem `json:"items,omitempty"`      // 评论列表数据
	PageToken string                        `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetTaskCommentListRespItem ...
type GetTaskCommentListRespItem struct {
	ID               string                             `json:"id,omitempty"`                  // 评论id
	Content          string                             `json:"content,omitempty"`             // 评论内容
	Creator          *GetTaskCommentListRespItemCreator `json:"creator,omitempty"`             // 评论创建人
	ReplyToCommentID string                             `json:"reply_to_comment_id,omitempty"` // 被回复的评论的ID, 如果不是回复评论, 则为空。
	CreatedAt        string                             `json:"created_at,omitempty"`          // 评论创建时间戳（ms)
	UpdatedAt        string                             `json:"updated_at,omitempty"`          // 评论更新时间戳（ms）
	ResourceType     string                             `json:"resource_type,omitempty"`       // 任务关联的资源类型
	ResourceID       string                             `json:"resource_id,omitempty"`         // 任务关联的资源ID
}

// GetTaskCommentListRespItemCreator ...
type GetTaskCommentListRespItemCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// getTaskCommentListResp ...
type getTaskCommentListResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *GetTaskCommentListResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
