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

// GetWikiTask 该方法用于获取wiki异步任务的结果
//
// 知识库权限要求:
// - 为任务创建者（用户或应用/机器人）
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get
func (r *DriveService) GetWikiTask(ctx context.Context, request *GetWikiTaskReq, options ...MethodOptionFunc) (*GetWikiTaskResp, *Response, error) {
	if r.cli.mock.mockDriveGetWikiTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetWikiTask mock enable")
		return r.cli.mock.mockDriveGetWikiTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetWikiTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/tasks/:task_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getWikiTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetWikiTask mock DriveGetWikiTask method
func (r *Mock) MockDriveGetWikiTask(f func(ctx context.Context, request *GetWikiTaskReq, options ...MethodOptionFunc) (*GetWikiTaskResp, *Response, error)) {
	r.mockDriveGetWikiTask = f
}

// UnMockDriveGetWikiTask un-mock DriveGetWikiTask method
func (r *Mock) UnMockDriveGetWikiTask() {
	r.mockDriveGetWikiTask = nil
}

// GetWikiTaskReq ...
type GetWikiTaskReq struct {
	TaskID   string `path:"task_id" json:"-"`    // 任务id, 示例值: "7037044037068177428-075c9481e6a0007c1df689dfbe5b55a08b6b06f7"
	TaskType string `query:"task_type" json:"-"` // 任务类型, 示例值: "move", 可选值有: `move`: MoveDocsToWiki任务
}

// GetWikiTaskResp ...
type GetWikiTaskResp struct {
	Task *GetWikiTaskRespTask `json:"task,omitempty"` // 任务结果
}

// GetWikiTaskRespTask ...
type GetWikiTaskRespTask struct {
	TaskID     string                           `json:"task_id,omitempty"`     // 任务id
	MoveResult []*GetWikiTaskRespTaskMoveResult `json:"move_result,omitempty"` // MoveDocsToWiki任务结果
}

// GetWikiTaskRespTaskMoveResult ...
type GetWikiTaskRespTaskMoveResult struct {
	Node      *GetWikiTaskRespTaskMoveResultNode `json:"node,omitempty"`       // 移动完成的节点信息
	Status    int64                              `json:"status,omitempty"`     // 节点移动状态码
	StatusMsg string                             `json:"status_msg,omitempty"` // 节点移动状态信息
}

// GetWikiTaskRespTaskMoveResultNode ...
type GetWikiTaskRespTaskMoveResultNode struct {
	SpaceID         string `json:"space_id,omitempty"`          // 知识库id
	NodeToken       string `json:"node_token,omitempty"`        // 节点token
	ObjToken        string `json:"obj_token,omitempty"`         // 文档token, 可以根据obj_type判断是属于doc、sheet还是mindnote的token(对于快捷方式, 该字段是对应的实体的obj_token)
	ObjType         string `json:"obj_type,omitempty"`          // 文档类型, 对于快捷方式, 该字段是对应的实体的obj_type, 可选值有: `doc`: doc, `sheet`: sheet, `mindnote`: mindnote, `bitable`: bitable, `file`: file, `docx`: docx
	ParentNodeToken string `json:"parent_node_token,omitempty"` // 节点的父亲token。当节点为一级节点时, 父亲token为空。
	NodeType        string `json:"node_type,omitempty"`         // 节点类型, 可选值有: `origin`: 实体, `shortcut`: 快捷方式
	OriginNodeToken string `json:"origin_node_token,omitempty"` // 快捷方式对应的实体node_token, 当创建节点为快捷方式时, 需要传该值
	OriginSpaceID   string `json:"origin_space_id,omitempty"`   // 快捷方式对应的实体所在的spaceid
	HasChild        bool   `json:"has_child,omitempty"`         // 是否有子节点
	Title           string `json:"title,omitempty"`             // 文档标题
	ObjCreateTime   string `json:"obj_create_time,omitempty"`   // 文档创建时间
	ObjEditTime     string `json:"obj_edit_time,omitempty"`     // 文档最近编辑时间
	NodeCreateTime  string `json:"node_create_time,omitempty"`  // 节点创建时间
}

// getWikiTaskResp ...
type getWikiTaskResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetWikiTaskResp `json:"data,omitempty"`
}
