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

// GetApprovalExternalApproval 根据 Approval Code 获取之前同步的某个三方审批定义的详情数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_approval/get
func (r *ApprovalService) GetApprovalExternalApproval(ctx context.Context, request *GetApprovalExternalApprovalReq, options ...MethodOptionFunc) (*GetApprovalExternalApprovalResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalExternalApproval != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalExternalApproval mock enable")
		return r.cli.mock.mockApprovalGetApprovalExternalApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApprovalExternalApproval",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/external_approvals/:approval_code",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalExternalApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalGetApprovalExternalApproval mock ApprovalGetApprovalExternalApproval method
func (r *Mock) MockApprovalGetApprovalExternalApproval(f func(ctx context.Context, request *GetApprovalExternalApprovalReq, options ...MethodOptionFunc) (*GetApprovalExternalApprovalResp, *Response, error)) {
	r.mockApprovalGetApprovalExternalApproval = f
}

// UnMockApprovalGetApprovalExternalApproval un-mock ApprovalGetApprovalExternalApproval method
func (r *Mock) UnMockApprovalGetApprovalExternalApproval() {
	r.mockApprovalGetApprovalExternalApproval = nil
}

// GetApprovalExternalApprovalReq ...
type GetApprovalExternalApprovalReq struct {
	ApprovalCode string  `path:"approval_code" json:"-"` // 调用[创建三方审批定义](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_approval/create)时返回的审批定义code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	UserIDType   *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApprovalExternalApprovalResp ...
type GetApprovalExternalApprovalResp struct {
	ApprovalName  string                                         `json:"approval_name,omitempty"`  // 审批定义名称
	ApprovalCode  string                                         `json:"approval_code,omitempty"`  // 创建三方审批定义时传入的定义code
	GroupCode     string                                         `json:"group_code,omitempty"`     // 审批定义所属分组
	GroupName     string                                         `json:"group_name,omitempty"`     // 分组名称
	Description   string                                         `json:"description,omitempty"`    // 审批定义的说明
	External      *GetApprovalExternalApprovalRespExternal       `json:"external,omitempty"`       // 三方审批定义相关
	Viewers       []*GetApprovalExternalApprovalRespViewer       `json:"viewers,omitempty"`        // 可见人列表
	I18nResources []*GetApprovalExternalApprovalRespI18nResource `json:"i18n_resources,omitempty"` // 国际化文案
	Managers      []string                                       `json:"managers,omitempty"`       // 流程管理员
}

// GetApprovalExternalApprovalRespExternal ...
type GetApprovalExternalApprovalRespExternal struct {
	BizName                     string `json:"biz_name,omitempty"`                      // 列表中用于提示审批来自哪里, i18n key, 注意不需要“来自”前缀, 审批中心会拼上前缀
	BizType                     string `json:"biz_type,omitempty"`                      // 审批定义业务类别
	CreateLinkMobile            string `json:"create_link_mobile,omitempty"`            // 移动端发起链接, 如果设置了该链接, 则会在移动端审批发起页展示该审批, 用户点击后会跳转到该链接进行发起； 如果不填, 则在mobile端不显示该审批
	CreateLinkPc                string `json:"create_link_pc,omitempty"`                // PC端发起链接, 如果设置了该链接, 则会在PC端审批发起页展示该审批, 用户点击后会跳转到该链接进行发起； 如果不填, 则在PC端不显示该审批
	SupportPc                   bool   `json:"support_pc,omitempty"`                    // 审批实例、审批任务、审批抄送是否要在PC端展示, 如果为 true, 则PC端列表会展示该定义下的实例信息；否则, 不展示
	SupportMobile               bool   `json:"support_mobile,omitempty"`                // 审批实例、审批任务、审批抄送是否要在移动端展示, 如果为 true, 则移动端列表会展示该定义下的实例信息；否则, 不展示
	SupportBatchRead            bool   `json:"support_batch_read,omitempty"`            // 是否支持批量已读
	EnableMarkReaded            bool   `json:"enable_mark_readed,omitempty"`            // 是否支持标注可读
	EnableQuickOperate          bool   `json:"enable_quick_operate,omitempty"`          // 是否支持快速操作
	ActionCallbackURL           string `json:"action_callback_url,omitempty"`           // 三方系统的操作回调 url, [待审批]列表的任务审批人点同意或拒绝操作后, 审批中心调用该地址通知三方系统, 回调地址相关信息可参考: [三方快捷审批回调, ](https://open.feishu.cn/document/ukTMukTMukTM/ukjNyYjL5YjM24SO2IjN/quick-approval-callback)
	ActionCallbackToken         string `json:"action_callback_token,omitempty"`         // 回调时带的 token, 用于业务系统验证请求来自审批, 具体参考: [回调token部分](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
	ActionCallbackKey           string `json:"action_callback_key,omitempty"`           // 请求参数加密密钥, 如果配置了该参数, 则会对请求参数进行加密, 业务需要对请求进行解密, 加解密算法参考:[加密部分](https://open.feishu.cn/document/ukTMukTMukTM/uADM4QjLwADO04CMwgDN)
	AllowBatchOperate           bool   `json:"allow_batch_operate,omitempty"`           // 是否支持批量审批
	ExcludeEfficiencyStatistics bool   `json:"exclude_efficiency_statistics,omitempty"` // 审批流程数据是否不纳入效率统计
}

// GetApprovalExternalApprovalRespI18nResource ...
type GetApprovalExternalApprovalRespI18nResource struct {
	Locale    string                                             `json:"locale,omitempty"`     // 语言, 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	Texts     []*GetApprovalExternalApprovalRespI18nResourceText `json:"texts,omitempty"`      // 文案 key, value, i18n key 以 @i18n@ 开头； 该字段主要用于做国际化, 允许用户同时传多个语言的文案, 审批中心会根据用户当前的语音环境使用对应的文案, 如果没有传用户当前的语音环境文案, 则会使用默认的语言文案
	IsDefault bool                                               `json:"is_default,omitempty"` // 是否默认语言, 默认语言需要包含所有key, 非默认语言如果key不存在会使用默认语言代替
}

// GetApprovalExternalApprovalRespI18nResourceText ...
type GetApprovalExternalApprovalRespI18nResourceText struct {
	Key   string `json:"key,omitempty"`   // 文案key
	Value string `json:"value,omitempty"` // 文案value
}

// GetApprovalExternalApprovalRespViewer ...
type GetApprovalExternalApprovalRespViewer struct {
	ViewerType         string `json:"viewer_type,omitempty"`          // 可见人类型, 可选值有: TENANT: 租户内可见, DEPARTMENT: 指定部门, USER: 指定用户, NONE: 任何人都不可见
	ViewerUserID       string `json:"viewer_user_id,omitempty"`       // 当 viewer_type 是 USER, 根据user_id_type返回用户id
	ViewerDepartmentID string `json:"viewer_department_id,omitempty"` // 当 view_type 为DEPARTMENT, 根据department_id_type返回部门id
}

// getApprovalExternalApprovalResp ...
type getApprovalExternalApprovalResp struct {
	Code  int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                           `json:"msg,omitempty"`  // 错误描述
	Data  *GetApprovalExternalApprovalResp `json:"data,omitempty"`
	Error *ErrorDetail                     `json:"error,omitempty"`
}
