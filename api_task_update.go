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

// UpdateTask 该接口用于修改任务的标题、描述、截止时间等信息。
//
// 更新时, 将`update_fields`字段中填写所有要修改的任务字段名, 同时在`task`字段中填写要修改的字段的新值即可。如果`update_fields`中设置了要变更一个字段的名字, 但是task里没设置新的值, 则表示将该字段清空。调用约定详情见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 关于资源的更新”章节。
// 目前支持更新的字段包括:
// * `summary` - 任务标题
// * `description` - 任务描述
// * `start` - 任务开始时间
// * `due` - 任务截止时间
// * `completed_at` - 用于标记任务完成/未完成
// * `extra` - 任务附带自定义数据
// * `custom_complete` - 任务自定义完成配置。
// * `repeat_rule` -  重复任务规则。
// * `mode` - 任务完成模式。
// * `is_milestone` - 是否是里程碑任务。
// * `custom_fields` - 自定义字段值。
// 该接口可以用于完成任务和将任务恢复至未完成, 只需要修改`completed_at`字段即可。但留意, 目前不管任务本身是会签任务还是或签任务, oapi对任务进行完成只能实现“整体完成”, 不支持个人单独完成。此外, 不能对已经完成的任务再次完成, 但可以将其恢复到未完成的状态(设置`completed_at`为"0")。
// 如更新自定义字段的值, 需要调用身份同时拥有任务的编辑权限和自定义字段的编辑权限。详情见[自定义字段功能概览](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/custom-field-overview)。更新时, 只有填写在`task.custom_fields`的自定义字段值会被更新, 不填写的不会被改变。
// 任务成员/提醒/清单数据不能使用本接口进行更新。
// * 如要修改任务成员, 需要使用[添加任务成员](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/add_members)
// 和[移除任务成员](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_members)接口。
// * 如要修改任务提醒, 需要使用[添加任务提醒](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/add_reminders)和[移除任务提醒](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_reminders)接口。
// * 如要变更任务所在的清单, 需要使用[任务加入清单](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/add_tasklist)和[任务移出清单]( /ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_tasklist)接口。
// 修改时需要调用身份对任务有编辑权限。详情见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“任务是如何鉴权的？”章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/patch
func (r *TaskService) UpdateTask(ctx context.Context, request *UpdateTaskReq, options ...MethodOptionFunc) (*UpdateTaskResp, *Response, error) {
	if r.cli.mock.mockTaskUpdateTask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#UpdateTask mock enable")
		return r.cli.mock.mockTaskUpdateTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "UpdateTask",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid",
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
	TaskGuid     string             `path:"task_guid" json:"-"`      // 要更新的任务全局唯一ID, 示例值: "e297ddff-06ca-4166-b917-4ce57cd3a7a0", 最大长度: `100` 字符
	UserIDType   *IDType            `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	Task         *UpdateTaskReqTask `json:"task,omitempty"`          // 要更新的任务数据, 只需要设置出现在`update_fields`中的字段即可。如果`update_fields`设置了要变更一个字段名, 但是`task`里没设置新的值, 则表示将该字段清空。
	UpdateFields []string           `json:"update_fields,omitempty"` // 设置需要修改的字段, summary: 任务标题, description: 任务描, start: 任务开始时间, due: 任务截止时间, completed_at: 任务完成时间, extra: 任务附属自定义数据, custom_complete: 任务自定义完成规则, repeat_rule: 任务重复规则, mode: 任务完成模式, is_milestone: 是否是里程碑任务, , 示例值: ["summary"]
}

// UpdateTaskReqTask ...
type UpdateTaskReqTask struct {
	Summary        *string                          `json:"summary,omitempty"`         // 任务标题。如更新标题, 不可将任务标题设为空。标题最大支持3000个utf8 字符, 示例值: "针对全年销售进行一次复盘"
	Description    *string                          `json:"description,omitempty"`     // 任务描述。描述最大支持3000个utf8字符, 示例值: "需要事先阅读复盘总结文档"
	Due            *UpdateTaskReqTaskDue            `json:"due,omitempty"`             // 任务截止时间。详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何使用开始时间和截止时间？”章节, 示例值: 1675742789470
	Extra          *string                          `json:"extra,omitempty"`           // 调用者可以传入的任意附带到任务上的数据。在获取任务详情时会原样返回, 示例值: "dGVzdA[", 最大长度: `65536` 字符
	CompletedAt    *string                          `json:"completed_at,omitempty"`    // 任务的完成时刻时间戳(ms), 示例值: "1675742789470", 默认值: `0`, 最大长度: `20` 字符
	RepeatRule     *string                          `json:"repeat_rule,omitempty"`     // 如果设置, 则该任务为“重复任务”。详见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“如何使用重复任务？”章节, 示例值: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO, TU, WE, TH, FR", 最大长度: `1000` 字符
	CustomComplete *UpdateTaskReqTaskCustomComplete `json:"custom_complete,omitempty"` // 任务自定义完成配置。详见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“ 如何使用任务自定义完成？”章节。
	Start          *UpdateTaskReqTaskStart          `json:"start,omitempty"`           // 任务的开始时间(ms)。详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何使用开始时间和截止时间？”章节。
	Mode           *int64                           `json:"mode,omitempty"`            // 任务的完成模式。1 - 会签任务；2 - 或签任务, 示例值: 2, 取值范围: `1` ～ `2`
	IsMilestone    *bool                            `json:"is_milestone,omitempty"`    // 是否是里程碑任务, 示例值: false
	CustomFields   []*UpdateTaskReqTaskCustomField  `json:"custom_fields,omitempty"`   // 自定义字段值, 如要更新, 每个字段的值根据字段类型填写相应的字段, * 当`type`为"number"时, 应使用`number_value`字段, 表示数字类型自定义字段的值；, * 当`type`为"member"时, 应使用`member_value`字段, 表示人员类型自定义字段的值；, * 当`type`为"datetime"时, 应使用`datetime_value`字段, 表示日期类型自定义字段的值；, * 当`type`为"single_select"时, 应使用`single_select_value`字段, 表示单选类型自定义字段的值；, * 当`type`为"multi_select"时, 应使用`multi_select_value`字段, 表示多选类型自定义字段的值；, * 当`type`为"text"时, 应使用`text_value`字段, 表示文本类型自定义字段的值。
}

// UpdateTaskReqTaskCustomComplete ...
type UpdateTaskReqTaskCustomComplete struct {
	Pc      *UpdateTaskReqTaskCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *UpdateTaskReqTaskCustomCompleteIos     `json:"ios,omitempty"`     // ios端的自定义完成配置
	Android *UpdateTaskReqTaskCustomCompleteAndroid `json:"android,omitempty"` // android端的自定义完成配置
}

// UpdateTaskReqTaskCustomCompleteAndroid ...
type UpdateTaskReqTaskCustomCompleteAndroid struct {
	Href *string                                    `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *UpdateTaskReqTaskCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskReqTaskCustomCompleteAndroidTip ...
type UpdateTaskReqTaskCustomCompleteAndroidTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// UpdateTaskReqTaskCustomCompleteIos ...
type UpdateTaskReqTaskCustomCompleteIos struct {
	Href *string                                `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *UpdateTaskReqTaskCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskReqTaskCustomCompleteIosTip ...
type UpdateTaskReqTaskCustomCompleteIosTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// UpdateTaskReqTaskCustomCompletePc ...
type UpdateTaskReqTaskCustomCompletePc struct {
	Href *string                               `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *UpdateTaskReqTaskCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskReqTaskCustomCompletePcTip ...
type UpdateTaskReqTaskCustomCompletePcTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// UpdateTaskReqTaskCustomField ...
type UpdateTaskReqTaskCustomField struct {
	Guid              string                                     `json:"guid,omitempty"`                // 自定义字段guid, 示例值: "73b21903-0041-4796-a11e-f8be919a7063"
	NumberValue       *string                                    `json:"number_value,omitempty"`        // 数字类型的自定义字段值, 填写一个合法数字的字符串表示, 空字符串表示设为空, 示例值: "10.23", 最大长度: `20` 字符
	MemberValue       []*UpdateTaskReqTaskCustomFieldMemberValue `json:"member_value,omitempty"`        // 人员类型的自定义字段值。可以设置1个或多个用户的id（遵循member格式, 只支持user类型）。当字段设为只不能多选时只能输入一个值。设为空数组表示设为空, 最大长度: `50`
	DatetimeValue     *string                                    `json:"datetime_value,omitempty"`      // 日期类型自定义字段值, 可以输入一个表示日期的以毫秒为单位的字符串。设为空字符串表示设为空, 示例值: "1698192000000"
	SingleSelectValue *string                                    `json:"single_select_value,omitempty"` // 单选类型字段值, 填写一个字段选项的option_guid。设置为空字符串表示设为空, 示例值: "73b21903-0041-4796-a11e-f8be919a7063", 最大长度: `100` 字符
	MultiSelectValue  []string                                   `json:"multi_select_value,omitempty"`  // 多选类型字段值, 可以填写一个或多个本字段的option_guid。设为空数组表示设为空, 示例值: ["73b21903-0041-4796-a11e-f8be919a7063"], 最大长度: `50`
	TextValue         *string                                    `json:"text_value,omitempty"`          // 文本类型字段值。可以填写最多3000字符。使用空字符串表示设为空, 示例值: "文本类型字段值。可以输入一段文本。空字符串表示清空。"
}

// UpdateTaskReqTaskCustomFieldMemberValue ...
type UpdateTaskReqTaskCustomFieldMemberValue struct {
	ID   *string `json:"id,omitempty"`   // 表示member的id, 示例值: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f", 最大长度: `100` 字符
	Type *string `json:"type,omitempty"` // 成员的类型, 示例值: "user", 默认值: `user`
}

// UpdateTaskReqTaskDue ...
type UpdateTaskReqTaskDue struct {
	Timestamp *string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true, 示例值: "1675454764000"
	IsAllDay  *bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储, 示例值: true
}

// UpdateTaskReqTaskStart ...
type UpdateTaskReqTaskStart struct {
	Timestamp *string `json:"timestamp,omitempty"`  // 开始时间的时间戳, 距1970-01-01 00:00:00的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true, 示例值: "1675454764000"
	IsAllDay  *bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储, 示例值: true
}

// UpdateTaskResp ...
type UpdateTaskResp struct {
	Task *UpdateTaskRespTask `json:"task,omitempty"` // 更新后的任务
}

// UpdateTaskRespTask ...
type UpdateTaskRespTask struct {
	Guid           string                            `json:"guid,omitempty"`             // 任务guid, 任务的唯一ID
	Summary        string                            `json:"summary,omitempty"`          // 任务标题
	Description    string                            `json:"description,omitempty"`      // 任务描述
	Due            *UpdateTaskRespTaskDue            `json:"due,omitempty"`              // 任务截止时间
	Reminders      []*UpdateTaskRespTaskReminder     `json:"reminders,omitempty"`        // 任务的提醒配置列表。目前每个任务最多有1个。
	Creator        *UpdateTaskRespTaskCreator        `json:"creator,omitempty"`          // 任务创建者
	Members        []*UpdateTaskRespTaskMember       `json:"members,omitempty"`          // 任务成员列表
	CompletedAt    string                            `json:"completed_at,omitempty"`     // 任务完成的时间戳(ms)
	Attachments    []*UpdateTaskRespTaskAttachment   `json:"attachments,omitempty"`      // 任务的附件列表
	Origin         *UpdateTaskRespTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息。创建是设置后就不可更改。
	Extra          string                            `json:"extra,omitempty"`            // 任务附带的自定义数据。
	Tasklists      []*UpdateTaskRespTaskTasklist     `json:"tasklists,omitempty"`        // 任务所属清单的名字。调用者只能看到有权限访问的清单的列表。
	RepeatRule     string                            `json:"repeat_rule,omitempty"`      // 如果任务为重复任务, 返回重复任务的配置
	ParentTaskGuid string                            `json:"parent_task_guid,omitempty"` // 如果当前任务为某个任务的子任务, 返回父任务的guid
	Mode           int64                             `json:"mode,omitempty"`             // 任务的模式。1 - 会签任务；2 - 或签任务
	Source         int64                             `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知来源, 1: 任务中心, 2: 群组任务/消息转任务, 6: 通过开放平台以tenant_access_token授权创建的任务, 7: 通过开放平台以user_access_token授权创建的任务, 8: 文档任务
	CustomComplete *UpdateTaskRespTaskCustomComplete `json:"custom_complete,omitempty"`  // 任务的自定义完成配置
	TaskID         string                            `json:"task_id,omitempty"`          // 任务界面上的代码
	CreatedAt      string                            `json:"created_at,omitempty"`       // 任务创建时间戳(ms)
	UpdatedAt      string                            `json:"updated_at,omitempty"`       // 任务最后一次更新的时间戳(ms)
	Status         string                            `json:"status,omitempty"`           // 任务的状态, 支持"todo"和"done"两种状态
	URL            string                            `json:"url,omitempty"`              // 任务的分享链接
	Start          *UpdateTaskRespTaskStart          `json:"start,omitempty"`            // 任务的开始时间
	SubtaskCount   int64                             `json:"subtask_count,omitempty"`    // 该任务的子任务的个数。
	IsMilestone    bool                              `json:"is_milestone,omitempty"`     // 是否是里程碑任务
	CustomFields   []*UpdateTaskRespTaskCustomField  `json:"custom_fields,omitempty"`    // 任务的自定义字段值
	Dependencies   []*UpdateTaskRespTaskDependencie  `json:"dependencies,omitempty"`     // 任务依赖
}

// UpdateTaskRespTaskAttachment ...
type UpdateTaskRespTaskAttachment struct {
	Guid       string                                `json:"guid,omitempty"`        // 附件guid
	FileToken  string                                `json:"file_token,omitempty"`  // 附件在云文档系统中的token
	Name       string                                `json:"name,omitempty"`        // 附件名
	Size       int64                                 `json:"size,omitempty"`        // 附件的字节大小
	Resource   *UpdateTaskRespTaskAttachmentResource `json:"resource,omitempty"`    // 附件归属的资源
	Uploader   *UpdateTaskRespTaskAttachmentUploader `json:"uploader,omitempty"`    // 附件上传者
	IsCover    bool                                  `json:"is_cover,omitempty"`    // 是否是封面图
	UploadedAt string                                `json:"uploaded_at,omitempty"` // 上传时间戳(ms)
}

// UpdateTaskRespTaskAttachmentResource ...
type UpdateTaskRespTaskAttachmentResource struct {
	Type string `json:"type,omitempty"` // 资源类型
	ID   string `json:"id,omitempty"`   // 资源ID
}

// UpdateTaskRespTaskAttachmentUploader ...
type UpdateTaskRespTaskAttachmentUploader struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 任务角色
}

// UpdateTaskRespTaskCreator ...
type UpdateTaskRespTaskCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 任务角色
}

// UpdateTaskRespTaskCustomComplete ...
type UpdateTaskRespTaskCustomComplete struct {
	Pc      *UpdateTaskRespTaskCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *UpdateTaskRespTaskCustomCompleteIos     `json:"ios,omitempty"`     // ios端的自定义完成配置
	Android *UpdateTaskRespTaskCustomCompleteAndroid `json:"android,omitempty"` // android端的自定义完成配置
}

// UpdateTaskRespTaskCustomCompleteAndroid ...
type UpdateTaskRespTaskCustomCompleteAndroid struct {
	Href string                                      `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *UpdateTaskRespTaskCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskRespTaskCustomCompleteAndroidTip ...
type UpdateTaskRespTaskCustomCompleteAndroidTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// UpdateTaskRespTaskCustomCompleteIos ...
type UpdateTaskRespTaskCustomCompleteIos struct {
	Href string                                  `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *UpdateTaskRespTaskCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskRespTaskCustomCompleteIosTip ...
type UpdateTaskRespTaskCustomCompleteIosTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// UpdateTaskRespTaskCustomCompletePc ...
type UpdateTaskRespTaskCustomCompletePc struct {
	Href string                                 `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *UpdateTaskRespTaskCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// UpdateTaskRespTaskCustomCompletePcTip ...
type UpdateTaskRespTaskCustomCompletePcTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// UpdateTaskRespTaskCustomField ...
type UpdateTaskRespTaskCustomField struct {
	Guid              string                                      `json:"guid,omitempty"`                // 字段GUID
	Type              string                                      `json:"type,omitempty"`                // 自定义字段类型, 支持"member", "datetime", "number", "single_select", "multi_select"五种类型
	NumberValue       string                                      `json:"number_value,omitempty"`        // 数字类型的自定义字段值, 填写一个合法数字的字符串表示, 空字符串表示设为空。
	DatetimeValue     string                                      `json:"datetime_value,omitempty"`      // 日期类型自定义字段值。可以输入一个表示日期的以毫秒为单位的字符串。设为空字符串表示设为空。
	MemberValue       []*UpdateTaskRespTaskCustomFieldMemberValue `json:"member_value,omitempty"`        // 人员类型的自定义字段值, 可以设置1个或多个用户的id（遵循member格式, 只支持user类型）。当该字段的设置为“不能多选”时只能输入一个值。设为空数组表示设为空。
	SingleSelectValue string                                      `json:"single_select_value,omitempty"` // 单选类型字段值, 填写一个字段选项的option_guid。设置为空字符串表示设为空。
	MultiSelectValue  []string                                    `json:"multi_select_value,omitempty"`  // 多选类型字段值, 可以填写一个或多个本字段的option_guid。设为空数组表示设为空。
	Name              string                                      `json:"name,omitempty"`                // 自定义字段名
	TextValue         string                                      `json:"text_value,omitempty"`          // 文本类型字段值。可以输入一段文本。空字符串表示清空。
}

// UpdateTaskRespTaskCustomFieldMemberValue ...
type UpdateTaskRespTaskCustomFieldMemberValue struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// UpdateTaskRespTaskDependencie ...
type UpdateTaskRespTaskDependencie struct {
	Type     string `json:"type,omitempty"`      // 依赖类型, 可选值有: prev: 前置依赖, next: 后置依赖
	TaskGuid string `json:"task_guid,omitempty"` // 依赖任务的GUID
}

// UpdateTaskRespTaskDue ...
type UpdateTaskRespTaskDue struct {
	Timestamp string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// UpdateTaskRespTaskMember ...
type UpdateTaskRespTaskMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 任务角色
}

// UpdateTaskRespTaskOrigin ...
type UpdateTaskRespTaskOrigin struct {
	PlatformI18nName *UpdateTaskRespTaskOriginPlatformI18nName `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。需提供多语言版本。
	Href             *UpdateTaskRespTaskOriginHref             `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// UpdateTaskRespTaskOriginHref ...
type UpdateTaskRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 链接对应的地址, 如填写必须以https://或http://开头
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// UpdateTaskRespTaskOriginPlatformI18nName ...
type UpdateTaskRespTaskOriginPlatformI18nName struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// UpdateTaskRespTaskReminder ...
type UpdateTaskRespTaskReminder struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间分钟数。例如30表示截止时间前30分钟提醒；0表示截止时提醒。
}

// UpdateTaskRespTaskStart ...
type UpdateTaskRespTaskStart struct {
	Timestamp string `json:"timestamp,omitempty"`  // 开始时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// UpdateTaskRespTaskTasklist ...
type UpdateTaskRespTaskTasklist struct {
	TasklistGuid string `json:"tasklist_guid,omitempty"` // 任务所在清单的guid
	SectionGuid  string `json:"section_guid,omitempty"`  // 任务所在清单的自定义分组guid
}

// updateTaskResp ...
type updateTaskResp struct {
	Code  int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string          `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateTaskResp `json:"data,omitempty"`
	Error *ErrorDetail    `json:"error,omitempty"`
}
