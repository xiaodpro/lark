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

// GetTask 该接口用于获取任务详情, 包括任务标题、描述、时间、成员等信息。
//
// 获取任务详情需要调用身份拥有对任务的可读取权限。详情见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“任务是如何鉴权的？”章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/get
func (r *TaskService) GetTask(ctx context.Context, request *GetTaskReq, options ...MethodOptionFunc) (*GetTaskResp, *Response, error) {
	if r.cli.mock.mockTaskGetTask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#GetTask mock enable")
		return r.cli.mock.mockTaskGetTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTask mock TaskGetTask method
func (r *Mock) MockTaskGetTask(f func(ctx context.Context, request *GetTaskReq, options ...MethodOptionFunc) (*GetTaskResp, *Response, error)) {
	r.mockTaskGetTask = f
}

// UnMockTaskGetTask un-mock TaskGetTask method
func (r *Mock) UnMockTaskGetTask() {
	r.mockTaskGetTask = nil
}

// GetTaskReq ...
type GetTaskReq struct {
	TaskGuid   string  `path:"task_guid" json:"-"`     // 要获取的任务guid, 示例值: "e297ddff-06ca-4166-b917-4ce57cd3a7a0", 最大长度: `100` 字符
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
}

// GetTaskResp ...
type GetTaskResp struct {
	Task *GetTaskRespTask `json:"task,omitempty"` // 获得的任务实体
}

// GetTaskRespTask ...
type GetTaskRespTask struct {
	Guid           string                         `json:"guid,omitempty"`             // 任务guid, 任务的唯一ID
	Summary        string                         `json:"summary,omitempty"`          // 任务标题
	Description    string                         `json:"description,omitempty"`      // 任务描述
	Due            *GetTaskRespTaskDue            `json:"due,omitempty"`              // 任务截止时间
	Reminders      []*GetTaskRespTaskReminder     `json:"reminders,omitempty"`        // 任务的提醒配置列表。目前每个任务最多有1个。
	Creator        *GetTaskRespTaskCreator        `json:"creator,omitempty"`          // 任务创建者
	Members        []*GetTaskRespTaskMember       `json:"members,omitempty"`          // 任务成员列表
	CompletedAt    string                         `json:"completed_at,omitempty"`     // 任务完成的时间戳(ms)
	Attachments    []*GetTaskRespTaskAttachment   `json:"attachments,omitempty"`      // 任务的附件列表
	Origin         *GetTaskRespTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息。创建是设置后就不可更改。
	Extra          string                         `json:"extra,omitempty"`            // 任务附带的自定义数据。
	Tasklists      []*GetTaskRespTaskTasklist     `json:"tasklists,omitempty"`        // 任务所属清单的名字。调用者只能看到有权限访问的清单的列表。
	RepeatRule     string                         `json:"repeat_rule,omitempty"`      // 如果任务为重复任务, 返回重复任务的配置
	ParentTaskGuid string                         `json:"parent_task_guid,omitempty"` // 如果当前任务为某个任务的子任务, 返回父任务的guid
	Mode           int64                          `json:"mode,omitempty"`             // 任务的模式。1 - 会签任务；2 - 或签任务
	Source         int64                          `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知来源, 1: 任务中心, 2: 群组任务/消息转任务, 6: 通过开放平台以tenant_access_token授权创建的任务, 7: 通过开放平台以user_access_token授权创建的任务, 8: 文档任务
	CustomComplete *GetTaskRespTaskCustomComplete `json:"custom_complete,omitempty"`  // 任务的自定义完成配置
	TaskID         string                         `json:"task_id,omitempty"`          // 任务界面上的代码
	CreatedAt      string                         `json:"created_at,omitempty"`       // 任务创建时间戳(ms)
	UpdatedAt      string                         `json:"updated_at,omitempty"`       // 任务最后一次更新的时间戳(ms)
	Status         string                         `json:"status,omitempty"`           // 任务的状态, 支持"todo"和"done"两种状态
	URL            string                         `json:"url,omitempty"`              // 任务的分享链接
	Start          *GetTaskRespTaskStart          `json:"start,omitempty"`            // 任务的开始时间
	SubtaskCount   int64                          `json:"subtask_count,omitempty"`    // 该任务的子任务的个数。
	IsMilestone    bool                           `json:"is_milestone,omitempty"`     // 是否是里程碑任务
	CustomFields   []*GetTaskRespTaskCustomField  `json:"custom_fields,omitempty"`    // 任务的自定义字段值
	Dependencies   []*GetTaskRespTaskDependencie  `json:"dependencies,omitempty"`     // 任务依赖
}

// GetTaskRespTaskAttachment ...
type GetTaskRespTaskAttachment struct {
	Guid       string                             `json:"guid,omitempty"`        // 附件guid
	FileToken  string                             `json:"file_token,omitempty"`  // 附件在云文档系统中的token
	Name       string                             `json:"name,omitempty"`        // 附件名
	Size       int64                              `json:"size,omitempty"`        // 附件的字节大小
	Resource   *GetTaskRespTaskAttachmentResource `json:"resource,omitempty"`    // 附件归属的资源
	Uploader   *GetTaskRespTaskAttachmentUploader `json:"uploader,omitempty"`    // 附件上传者
	IsCover    bool                               `json:"is_cover,omitempty"`    // 是否是封面图
	UploadedAt string                             `json:"uploaded_at,omitempty"` // 上传时间戳(ms)
}

// GetTaskRespTaskAttachmentResource ...
type GetTaskRespTaskAttachmentResource struct {
	Type string `json:"type,omitempty"` // 资源类型
	ID   string `json:"id,omitempty"`   // 资源ID
}

// GetTaskRespTaskAttachmentUploader ...
type GetTaskRespTaskAttachmentUploader struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 角色
}

// GetTaskRespTaskCreator ...
type GetTaskRespTaskCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// GetTaskRespTaskCustomComplete ...
type GetTaskRespTaskCustomComplete struct {
	Pc      *GetTaskRespTaskCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *GetTaskRespTaskCustomCompleteIos     `json:"ios,omitempty"`     // ios端的自定义完成配置
	Android *GetTaskRespTaskCustomCompleteAndroid `json:"android,omitempty"` // android端的自定义完成配置
}

// GetTaskRespTaskCustomCompleteAndroid ...
type GetTaskRespTaskCustomCompleteAndroid struct {
	Href string                                   `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *GetTaskRespTaskCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// GetTaskRespTaskCustomCompleteAndroidTip ...
type GetTaskRespTaskCustomCompleteAndroidTip struct {
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

// GetTaskRespTaskCustomCompleteIos ...
type GetTaskRespTaskCustomCompleteIos struct {
	Href string                               `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *GetTaskRespTaskCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// GetTaskRespTaskCustomCompleteIosTip ...
type GetTaskRespTaskCustomCompleteIosTip struct {
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

// GetTaskRespTaskCustomCompletePc ...
type GetTaskRespTaskCustomCompletePc struct {
	Href string                              `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *GetTaskRespTaskCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// GetTaskRespTaskCustomCompletePcTip ...
type GetTaskRespTaskCustomCompletePcTip struct {
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

// GetTaskRespTaskCustomField ...
type GetTaskRespTaskCustomField struct {
	Guid              string                                   `json:"guid,omitempty"`                // 字段GUID
	Type              string                                   `json:"type,omitempty"`                // 自定义字段类型, 支持"member", "datetime", "number", "single_select", "multi_select"五种类型
	NumberValue       string                                   `json:"number_value,omitempty"`        // 数字类型的自定义字段值, 填写一个合法数字的字符串表示, 空字符串表示设为空。
	DatetimeValue     string                                   `json:"datetime_value,omitempty"`      // 日期类型自定义字段值。可以输入一个表示日期的以毫秒为单位的字符串。设为空字符串表示设为空。
	MemberValue       []*GetTaskRespTaskCustomFieldMemberValue `json:"member_value,omitempty"`        // 人员类型的自定义字段值, 可以设置1个或多个用户的id（遵循member格式, 只支持user类型）。当该字段的设置为“不能多选”时只能输入一个值。设为空数组表示设为空。
	SingleSelectValue string                                   `json:"single_select_value,omitempty"` // 单选类型字段值, 填写一个字段选项的option_guid。设置为空字符串表示设为空。
	MultiSelectValue  []string                                 `json:"multi_select_value,omitempty"`  // 多选类型字段值, 可以填写一个或多个本字段的option_guid。设为空数组表示设为空。
	Name              string                                   `json:"name,omitempty"`                // 自定义字段名
	TextValue         string                                   `json:"text_value,omitempty"`          // 文本类型字段值。可以输入一段文本。空字符串表示清空。
}

// GetTaskRespTaskCustomFieldMemberValue ...
type GetTaskRespTaskCustomFieldMemberValue struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// GetTaskRespTaskDependencie ...
type GetTaskRespTaskDependencie struct {
	Type     string `json:"type,omitempty"`      // 依赖类型, 可选值有: prev: 前置依赖, next: 后置依赖
	TaskGuid string `json:"task_guid,omitempty"` // 依赖任务的GUID
}

// GetTaskRespTaskDue ...
type GetTaskRespTaskDue struct {
	Timestamp string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// GetTaskRespTaskMember ...
type GetTaskRespTaskMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// GetTaskRespTaskOrigin ...
type GetTaskRespTaskOrigin struct {
	PlatformI18nName *GetTaskRespTaskOriginPlatformI18nName `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。需提供多语言版本。
	Href             *GetTaskRespTaskOriginHref             `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// GetTaskRespTaskOriginHref ...
type GetTaskRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 链接对应的地址
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// GetTaskRespTaskOriginPlatformI18nName ...
type GetTaskRespTaskOriginPlatformI18nName struct {
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

// GetTaskRespTaskReminder ...
type GetTaskRespTaskReminder struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间分钟数。例如30表示截止时间前30分钟提醒；0表示截止时提醒。
}

// GetTaskRespTaskStart ...
type GetTaskRespTaskStart struct {
	Timestamp string `json:"timestamp,omitempty"`  // 开始时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// GetTaskRespTaskTasklist ...
type GetTaskRespTaskTasklist struct {
	TasklistGuid string `json:"tasklist_guid,omitempty"` // 任务所在清单的guid
	SectionGuid  string `json:"section_guid,omitempty"`  // 任务所在清单的自定义分组guid
}

// getTaskResp ...
type getTaskResp struct {
	Code  int64        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string       `json:"msg,omitempty"`  // 错误描述
	Data  *GetTaskResp `json:"data,omitempty"`
	Error *ErrorDetail `json:"error,omitempty"`
}
