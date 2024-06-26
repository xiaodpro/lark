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

// RemoveTaskReminder 将一个提醒从任务中移除。
//
// 如果要移除的提醒本来就不存在, 本接口将直接返回成功。
// 需要任务的编辑权限。详情见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“任务是如何鉴权的？”章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_reminders
func (r *TaskService) RemoveTaskReminder(ctx context.Context, request *RemoveTaskReminderReq, options ...MethodOptionFunc) (*RemoveTaskReminderResp, *Response, error) {
	if r.cli.mock.mockTaskRemoveTaskReminder != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#RemoveTaskReminder mock enable")
		return r.cli.mock.mockTaskRemoveTaskReminder(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "RemoveTaskReminder",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid/remove_reminders",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(removeTaskReminderResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskRemoveTaskReminder mock TaskRemoveTaskReminder method
func (r *Mock) MockTaskRemoveTaskReminder(f func(ctx context.Context, request *RemoveTaskReminderReq, options ...MethodOptionFunc) (*RemoveTaskReminderResp, *Response, error)) {
	r.mockTaskRemoveTaskReminder = f
}

// UnMockTaskRemoveTaskReminder un-mock TaskRemoveTaskReminder method
func (r *Mock) UnMockTaskRemoveTaskReminder() {
	r.mockTaskRemoveTaskReminder = nil
}

// RemoveTaskReminderReq ...
type RemoveTaskReminderReq struct {
	TaskGuid    string   `path:"task_guid" json:"-"`     // 要移除提醒的任务全局唯一ID, 示例值: "d300a75f-c56a-4be9-80d1-e47653028ceb", 最大长度: `100` 字符
	UserIDType  *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	ReminderIDs []string `json:"reminder_ids,omitempty"` // 要移除的reminder的id列表, 示例值: ["7202449098622894100"]
}

// RemoveTaskReminderResp ...
type RemoveTaskReminderResp struct {
	Task *RemoveTaskReminderRespTask `json:"task,omitempty"` // 移除后任务的提醒列表
}

// RemoveTaskReminderRespTask ...
type RemoveTaskReminderRespTask struct {
	Guid           string                                    `json:"guid,omitempty"`             // 任务的全局唯一ID
	Summary        string                                    `json:"summary,omitempty"`          // 任务标题
	Description    string                                    `json:"description,omitempty"`      // 任务备注
	Due            *RemoveTaskReminderRespTaskDue            `json:"due,omitempty"`              // 任务截止时间
	Reminders      []*RemoveTaskReminderRespTaskReminder     `json:"reminders,omitempty"`        // 任务的提醒配置列表。目前每个任务最多有1个。
	Creator        *RemoveTaskReminderRespTaskCreator        `json:"creator,omitempty"`          // 任务创建者
	Members        []*RemoveTaskReminderRespTaskMember       `json:"members,omitempty"`          // 任务成员列表
	CompletedAt    string                                    `json:"completed_at,omitempty"`     // 任务完成的时间戳(ms)
	Attachments    []*RemoveTaskReminderRespTaskAttachment   `json:"attachments,omitempty"`      // 任务的附件列表
	Origin         *RemoveTaskReminderRespTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息。创建是设置后就不可更改。
	Extra          string                                    `json:"extra,omitempty"`            // 任务附带的自定义数据。
	Tasklists      []*RemoveTaskReminderRespTaskTasklist     `json:"tasklists,omitempty"`        // 任务所属清单的名字。调用者只能看到有权限访问的清单的列表。
	RepeatRule     string                                    `json:"repeat_rule,omitempty"`      // 任务重复规则
	ParentTaskGuid string                                    `json:"parent_task_guid,omitempty"` // 如果当前任务为某个任务的子任务, 返回父任务的guid
	Mode           int64                                     `json:"mode,omitempty"`             // 任务的模式。1 - 会签任务；2 - 或签任务
	Source         int64                                     `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知来源, 1: 任务中心, 2: 群组任务/消息转任务, 6: 通过开放平台以tenant_access_token授权创建的任务, 7: 通过开放平台以user_access_token授权创建的任务, 8: 文档任务
	CustomComplete *RemoveTaskReminderRespTaskCustomComplete `json:"custom_complete,omitempty"`  // 任务的自定义完成配置
	TaskID         string                                    `json:"task_id,omitempty"`          // 任务界面上的代码
	CreatedAt      string                                    `json:"created_at,omitempty"`       // 任务创建时间戳(ms)
	UpdatedAt      string                                    `json:"updated_at,omitempty"`       // 任务最后一次更新的时间戳(ms)
	Status         string                                    `json:"status,omitempty"`           // 任务的状态, 支持"todo"和"done"两种状态
	URL            string                                    `json:"url,omitempty"`              // 任务的分享链接
	Start          *RemoveTaskReminderRespTaskStart          `json:"start,omitempty"`            // 任务的开始时间
	SubtaskCount   int64                                     `json:"subtask_count,omitempty"`    // 该任务的子任务的个数。
}

// RemoveTaskReminderRespTaskAttachment ...
type RemoveTaskReminderRespTaskAttachment struct {
	Guid       string                                        `json:"guid,omitempty"`        // 附件guid
	FileToken  string                                        `json:"file_token,omitempty"`  // 附件在云文档系统中的token
	Name       string                                        `json:"name,omitempty"`        // 附件名
	Size       int64                                         `json:"size,omitempty"`        // 附件的字节大小
	Resource   *RemoveTaskReminderRespTaskAttachmentResource `json:"resource,omitempty"`    // 附件归属的资源
	Uploader   *RemoveTaskReminderRespTaskAttachmentUploader `json:"uploader,omitempty"`    // 附件上传者
	IsCover    bool                                          `json:"is_cover,omitempty"`    // 是否是封面图
	UploadedAt string                                        `json:"uploaded_at,omitempty"` // 上传时间戳(ms)
}

// RemoveTaskReminderRespTaskAttachmentResource ...
type RemoveTaskReminderRespTaskAttachmentResource struct {
	Type string `json:"type,omitempty"` // 资源类型
	ID   string `json:"id,omitempty"`   // 资源ID
}

// RemoveTaskReminderRespTaskAttachmentUploader ...
type RemoveTaskReminderRespTaskAttachmentUploader struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// RemoveTaskReminderRespTaskCreator ...
type RemoveTaskReminderRespTaskCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// RemoveTaskReminderRespTaskCustomComplete ...
type RemoveTaskReminderRespTaskCustomComplete struct {
	Pc      *RemoveTaskReminderRespTaskCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *RemoveTaskReminderRespTaskCustomCompleteIos     `json:"ios,omitempty"`     // ios端的自定义完成配置
	Android *RemoveTaskReminderRespTaskCustomCompleteAndroid `json:"android,omitempty"` // android端的自定义完成配置
}

// RemoveTaskReminderRespTaskCustomCompleteAndroid ...
type RemoveTaskReminderRespTaskCustomCompleteAndroid struct {
	Href string                                              `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *RemoveTaskReminderRespTaskCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// RemoveTaskReminderRespTaskCustomCompleteAndroidTip ...
type RemoveTaskReminderRespTaskCustomCompleteAndroidTip struct {
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

// RemoveTaskReminderRespTaskCustomCompleteIos ...
type RemoveTaskReminderRespTaskCustomCompleteIos struct {
	Href string                                          `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *RemoveTaskReminderRespTaskCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// RemoveTaskReminderRespTaskCustomCompleteIosTip ...
type RemoveTaskReminderRespTaskCustomCompleteIosTip struct {
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

// RemoveTaskReminderRespTaskCustomCompletePc ...
type RemoveTaskReminderRespTaskCustomCompletePc struct {
	Href string                                         `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *RemoveTaskReminderRespTaskCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// RemoveTaskReminderRespTaskCustomCompletePcTip ...
type RemoveTaskReminderRespTaskCustomCompletePcTip struct {
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

// RemoveTaskReminderRespTaskDue ...
type RemoveTaskReminderRespTaskDue struct {
	Timestamp string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// RemoveTaskReminderRespTaskMember ...
type RemoveTaskReminderRespTaskMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// RemoveTaskReminderRespTaskOrigin ...
type RemoveTaskReminderRespTaskOrigin struct {
	PlatformI18nName *RemoveTaskReminderRespTaskOriginPlatformI18nName `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。需提供多语言版本。
	Href             *RemoveTaskReminderRespTaskOriginHref             `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// RemoveTaskReminderRespTaskOriginHref ...
type RemoveTaskReminderRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 链接对应的地址
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// RemoveTaskReminderRespTaskOriginPlatformI18nName ...
type RemoveTaskReminderRespTaskOriginPlatformI18nName struct {
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

// RemoveTaskReminderRespTaskReminder ...
type RemoveTaskReminderRespTaskReminder struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间分钟数。例如30表示截止时间前30分钟提醒；0表示截止时提醒。
}

// RemoveTaskReminderRespTaskStart ...
type RemoveTaskReminderRespTaskStart struct {
	Timestamp string `json:"timestamp,omitempty"`  // 开始时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// RemoveTaskReminderRespTaskTasklist ...
type RemoveTaskReminderRespTaskTasklist struct {
	TasklistGuid string `json:"tasklist_guid,omitempty"` // 任务所在清单的guid
	SectionGuid  string `json:"section_guid,omitempty"`  // 任务所在清单的自定义分组guid
}

// removeTaskReminderResp ...
type removeTaskReminderResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *RemoveTaskReminderResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
