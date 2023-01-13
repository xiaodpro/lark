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

// UpdateTask 该接口用于修改任务的标题、描述、时间、来源等相关信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/patch
func (r *TaskService) UpdateTask(ctx context.Context, request *UpdateTaskReq, options ...MethodOptionFunc) (*UpdateTaskResp, *Response, error) {
	if r.cli.mock.mockTaskUpdateTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#UpdateTask mock enable")
		return r.cli.mock.mockTaskUpdateTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "UpdateTask",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskUpdateTask mock TaskUpdateTask method
func (r *Mock) MockTaskUpdateTask(f func(ctx context.Context, request *UpdateTaskReq, options ...MethodOptionFunc) (*UpdateTaskResp, *Response, error)) {
	r.mockTaskUpdateTask = f
}

// UnMockTaskUpdateTask un-mock TaskUpdateTask method
func (r *Mock) UnMockTaskUpdateTask() {
	r.mockTaskUpdateTask = nil
}

// UpdateTaskReq ...
type UpdateTaskReq struct {
	TaskID       string             `path:"task_id" json:"-"`        // 任务 ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	UserIDType   *IDType            `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Task         *UpdateTaskReqTask `json:"task,omitempty"`          // 被更新的任务实体基础信息
	UpdateFields []string           `json:"update_fields,omitempty"` // 指定需要更新的任务字段, 否则服务端将不知道更新哪些字段, 示例值: ["summary"]
}

// UpdateTaskReqTask ...
type UpdateTaskReqTask struct {
	Summary         *string                          `json:"summary,omitempty"`          // 任务的标题, 类型为文本字符串, 如果要在任务标题中插入 URL 或者 @某个用户, 请使用rich_summary字段, 创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段, 任务标题和任务富文本标题同时存在时只使用富文本标题, 示例值: "完成本季度OKR编写", 长度范围: `0` ～ `1000` 字符
	Description     *string                          `json:"description,omitempty"`      // 任务的描述, 类型为文本字符串, 如果要在任务描述中插入 URL 或者 @某个用户, 请使用rich_description字段, 任务备注和任务富文本备注同时存在时只使用富文本备注, 示例值: "对本次会议内容复盘总结, 编写更新本季度OKR", 长度范围: `0` ～ `65536` 字符
	Extra           *string                          `json:"extra,omitempty"`            // 附属信息, 接入方可以传入base64 编码后的自定义的数据。用户如果需要对当前任务备注信息, 但对外不显示, 可使用该字段进行存储, 该数据会在获取任务详情时, 原样返回给用户, 示例值: "dGVzdA[", 长度范围: `0` ～ `65536` 字符
	Due             *UpdateTaskReqTaskDue            `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *UpdateTaskReqTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	CanEdit         *bool                            `json:"can_edit,omitempty"`         // 此字段用于控制该任务在飞书任务中心是否可编辑, 默认为false, 已经废弃, 向前兼容故仍然保留, 但不推荐使用, 示例值: true, 默认值: `false`
	Custom          *string                          `json:"custom,omitempty"`           // 自定义完成配置, 此字段用于设置完成任务时的页面跳转, 或展示提示语。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields), 示例值: "{\"custom_complete\":{\"android\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}, \"ios\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}, \"pc\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}}}", 长度范围: `0` ～ `65536` 字符
	Followers       []*UpdateTaskReqTaskFollower     `json:"followers,omitempty"`        // 任务的关注者
	Collaborators   []*UpdateTaskReqTaskCollaborator `json:"collaborators,omitempty"`    // 任务的执行者
	CollaboratorIDs []string                         `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值: ["ou_1400208f15333e20e11339d39067844b", "ou_84ed6312949945c8ae6168f10829e9e6"], 最大长度: `100`
	FollowerIDs     []string                         `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值: ["ou_1400208f15333e20e11339d39067844b", "ou_84ed6312949945c8ae6168f10829e9e6"], 最大长度: `100`
	RepeatRule      *string                          `json:"repeat_rule,omitempty"`      // 重复任务的规则表达式, 语法格式参见[RRule语法规范](https://www.ietf.org/rfc/rfc2445.txt) 4.3.10小节, 示例值: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO, TU, WE, TH, FR"
	RichSummary     *string                          `json:"rich_summary,omitempty"`     // 富文本任务标题。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 。创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段, 示例值: "完成本季度OKR编写\[飞书开放平台](https://open.feishu.cn/)", 长度范围: `0` ～ `1000` 字符
	RichDescription *string                          `json:"rich_description,omitempty"` // 富文本任务备注。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 示例值: "对本次会议内容复盘总结, 编写更新本季度OKR\[飞书开放平台](https://open.feishu.cn/)", 长度范围: `0` ～ `65536` 字符
}

// UpdateTaskReqTaskCollaborator ...
type UpdateTaskReqTaskCollaborator struct {
	ID     *string  `json:"id,omitempty"`      // 任务执行者的 ID, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 已经废弃, 为了向前兼容早期只支持单次添加一个人的情况而保留, 但不再推荐使用, 建议使用id_list字段, 示例值: "ou_99e1a581b36ecc4862cbfbce473f1234"
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值: ["ou_550cc75233d8b7b9fcbdad65f34433f4", "ou_d1e9d27cf3235b40ca9a67c67ef088b0"]
}

// UpdateTaskReqTaskDue ...
type UpdateTaskReqTaskDue struct {
	Time     *string `json:"time,omitempty"`       // 表示截止时间的Unix时间戳（单位为秒）, 示例值: "1623124318"
	Timezone *string `json:"timezone,omitempty"`   // 截止时间对应的时区, 传入值需要符合IANA Time Zone Database标准, 规范见[Time Zone Database](https://www.iana.org/time-zones), 示例值: "Asia/Shanghai", 默认值: `Asia/Shanghai`
	IsAllDay *bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务, 包括如下取值: true: 表示是全天任务, 全天任务的截止时间为当天 UTC 时间的 0 点, false: 表示不是全天任务, 示例值: false, 默认值: `false`
}

// UpdateTaskReqTaskFollower ...
type UpdateTaskReqTaskFollower struct {
	ID     *string  `json:"id,omitempty"`      // 任务关注人 ID, 示例值: "ou_99e1a581b36ecc4862cbfbce473f3123"
	IDList []string `json:"id_list,omitempty"` // 要删除的关注人ID列表, 示例值: [, "ou_550cc75233d8b7b9fcbdad65f34433f4", "ou_d1e9d27cf3235b40ca9a67c67ef088b0", ]
}

// UpdateTaskReqTaskOrigin ...
type UpdateTaskReqTaskOrigin struct {
	PlatformI18nName string                       `json:"platform_i18n_name,omitempty"` // 任务来源的名称, 用于在任务中心详情页展示。需要提供一个字典, 支持多种语言名称映射。应用在使用不同语言时, 导入来源也将展示对应的内容。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields), 示例值: "{\"zh_cn\": \"IT 工作台\", \"en_us\": \"IT Workspace\"}", 长度范围: `0` ～ `1024` 字符
	Href             *UpdateTaskReqTaskOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// UpdateTaskReqTaskOriginHref ...
type UpdateTaskReqTaskOriginHref struct {
	URL   *string `json:"url,omitempty"`   // 具体链接地址, URL仅支持解析http、https。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields), 示例值: "https://support.feishu.com/internal/foo-bar", 长度范围: `0` ～ `1024` 字符
	Title *string `json:"title,omitempty"` // 链接对应的标题, 示例值: "反馈一个问题, 需要协助排查", 长度范围: `0` ～ `512` 字符
}

// UpdateTaskResp ...
type UpdateTaskResp struct {
	Task *UpdateTaskRespTask `json:"task,omitempty"` // 返回修改后的任务详情
}

// UpdateTaskRespTask ...
type UpdateTaskRespTask struct {
	ID              string                            `json:"id,omitempty"`               // 任务的唯一ID, 例如"83912691-2e43-47fc-94a4-d512e03984fa"
	Summary         string                            `json:"summary,omitempty"`          // 任务的标题, 类型为文本字符串, 如果要在任务标题中插入 URL 或者 @某个用户, 请使用rich_summary字段, 创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段, 任务标题和任务富文本标题同时存在时只使用富文本标题。
	Description     string                            `json:"description,omitempty"`      // 任务的描述, 类型为文本字符串, 如果要在任务描述中插入 URL 或者 @某个用户, 请使用rich_description字段, 任务备注和任务富文本备注同时存在时只使用富文本备注。
	CompleteTime    string                            `json:"complete_time,omitempty"`    // 任务的完成时间戳（单位为秒）, 完成时间为0则表示任务尚未完成, 不支持部分完成, 只有整个任务完成, 该字段才会有非0值。
	CreatorID       string                            `json:"creator_id,omitempty"`       // 任务的创建者 ID, 其中查询字段 user_id_type 是用于控制返回体中 creator_id 的类型, 不传时默认返回 open_id, 特别的, 使用tenant_access_token 调用接口时, 如果是user_id_type是openid, 则返回代表该应用身份的openid；当user_id_type为user_id时, 不返回creator_id。原因是user_id代表一个真实飞书用户的id, 应用身份没有user_id。使用user_access_token调用接口正常返回创建者。
	Extra           string                            `json:"extra,omitempty"`            // 附属信息, 接入方可以传入base64 编码后的自定义的数据。用户如果需要对当前任务备注信息, 但对外不显示, 可使用该字段进行存储, 该数据会在获取任务详情时, 原样返回给用户。
	CreateTime      string                            `json:"create_time,omitempty"`      // 任务的创建时间的Unix时间戳（单位为秒）
	UpdateTime      string                            `json:"update_time,omitempty"`      // 任务的更新时间的Unix时间戳（单位为秒）, 创建任务时update_time与create_time相同
	Due             *UpdateTaskRespTaskDue            `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *UpdateTaskRespTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	Custom          string                            `json:"custom,omitempty"`           // 自定义完成配置, 此字段用于设置完成任务时的页面跳转, 或展示提示语。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Source          int64                             `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知类型, 1: 来源任务中心创建, 2: 来源消息转任务, 3: 来源云文档, 4: 来源文档单品, 5: 来源PANO, 6: 来源tenant_access_token创建的任务, 7: 来源user_access_token创建的任务, 8: 来源新版云文档
	Followers       []*UpdateTaskRespTaskFollower     `json:"followers,omitempty"`        // 任务的关注者
	Collaborators   []*UpdateTaskRespTaskCollaborator `json:"collaborators,omitempty"`    // 任务的执行者
	CollaboratorIDs []string                          `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
	FollowerIDs     []string                          `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
	RepeatRule      string                            `json:"repeat_rule,omitempty"`      // 重复任务的规则表达式, 语法格式参见[RRule语法规范](https://www.ietf.org/rfc/rfc2445.txt) 4.3.10小节
	RichSummary     string                            `json:"rich_summary,omitempty"`     // 富文本任务标题。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 。创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段。
	RichDescription string                            `json:"rich_description,omitempty"` // 富文本任务备注。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
}

// UpdateTaskRespTaskCollaborator ...
type UpdateTaskRespTaskCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 已经废弃, 为了向前兼容早期只支持单次添加一个人的情况而保留, 但不再推荐使用, 建议使用id_list字段
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
}

// UpdateTaskRespTaskDue ...
type UpdateTaskRespTaskDue struct {
	Time     string `json:"time,omitempty"`       // 表示截止时间的Unix时间戳（单位为秒）。
	Timezone string `json:"timezone,omitempty"`   // 截止时间对应的时区, 传入值需要符合IANA Time Zone Database标准, 规范见[Time Zone Database](https://www.iana.org/time-zones)。
	IsAllDay bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务, 包括如下取值: true: 表示是全天任务, 全天任务的截止时间为当天 UTC 时间的 0 点, false: 表示不是全天任务。
}

// UpdateTaskRespTaskFollower ...
type UpdateTaskRespTaskFollower struct {
	ID     string   `json:"id,omitempty"`      // 任务关注人 ID
	IDList []string `json:"id_list,omitempty"` // 要删除的关注人ID列表
}

// UpdateTaskRespTaskOrigin ...
type UpdateTaskRespTaskOrigin struct {
	PlatformI18nName string                        `json:"platform_i18n_name,omitempty"` // 任务来源的名称, 用于在任务中心详情页展示。需要提供一个字典, 支持多种语言名称映射。应用在使用不同语言时, 导入来源也将展示对应的内容。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Href             *UpdateTaskRespTaskOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// UpdateTaskRespTaskOriginHref ...
type UpdateTaskRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 具体链接地址, URL仅支持解析http、https。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// updateTaskResp ...
type updateTaskResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateTaskResp `json:"data,omitempty"`
}
