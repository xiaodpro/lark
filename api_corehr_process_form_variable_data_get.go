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
 
 // GetCoreHRProcessFormVariableData 根据流程实例 id（process_id）获取流程表单字段数据, 包括表单里的业务字段和自定义字段。仅支持飞书人事、假勤相关业务流程。
 //
 // 注: 旧版 API 文档已移动到[历史版本]目录。
 //
 // doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/process-form_variable_data/get
 func (r *CoreHRService) GetCoreHRProcessFormVariableData(ctx context.Context, request *GetCoreHRProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHRProcessFormVariableDataResp, *Response, error) {
	 if r.cli.mock.mockCoreHRGetCoreHRProcessFormVariableData != nil {
		 r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRProcessFormVariableData mock enable")
		 return r.cli.mock.mockCoreHRGetCoreHRProcessFormVariableData(ctx, request, options...)
	 }
 
	 req := &RawRequestReq{
		 Scope:                 "CoreHR",
		 API:                   "GetCoreHRProcessFormVariableData",
		 Method:                "GET",
		 URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/processes/:process_id/form_variable_data",
		 Body:                  request,
		 MethodOption:          newMethodOption(options),
		 NeedTenantAccessToken: true,
	 }
	 resp := new(getCoreHRProcessFormVariableDataResp)
 
	 response, err := r.cli.RawRequest(ctx, req, resp)
	 return resp.Data, response, err
 }
 
 // MockCoreHRGetCoreHRProcessFormVariableData mock CoreHRGetCoreHRProcessFormVariableData method
 func (r *Mock) MockCoreHRGetCoreHRProcessFormVariableData(f func(ctx context.Context, request *GetCoreHRProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHRProcessFormVariableDataResp, *Response, error)) {
	 r.mockCoreHRGetCoreHRProcessFormVariableData = f
 }
 
 // UnMockCoreHRGetCoreHRProcessFormVariableData un-mock CoreHRGetCoreHRProcessFormVariableData method
 func (r *Mock) UnMockCoreHRGetCoreHRProcessFormVariableData() {
	 r.mockCoreHRGetCoreHRProcessFormVariableData = nil
 }
 
 // GetCoreHRProcessFormVariableDataReq ...
 type GetCoreHRProcessFormVariableDataReq struct {
	 ProcessID        string            `path:"process_id" json:"-"`          // 流程实例 id, 示例值: "7341373094948242956"
	 UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	 DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 department_id 来标识部门, 默认值: `people_corehr_department_id`
 }
 
 // GetCoreHRProcessFormVariableDataResp ...
 type GetCoreHRProcessFormVariableDataResp struct {
	 FieldVariableValues []*GetCoreHRProcessFormVariableDataRespFieldVariableValue `json:"field_variable_values,omitempty"` // 表单数据
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValue struct {
	 VariableApiName string                                                               `json:"variable_api_name,omitempty"` // 变量唯一标识
	 VariableName    *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName  `json:"variable_name,omitempty"`     // 变量名称
	 VariableValue   *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue `json:"variable_value,omitempty"`    // 变量值
	 SubValues       []*GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValue    `json:"sub_values,omitempty"`        // 在 list_values 和 record_values 中引用的变量
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValue struct {
	 Key   string                                                               `json:"key,omitempty"`   // 用于关联 list 和 record 类型变量值中的 key
	 Value *GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValue `json:"value,omitempty"` // 变量值
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValue struct {
	 TextValue       *string                                                                          `json:"text_value,omitempty"`       // 文本值
	 BoolValue       *bool                                                                            `json:"bool_value,omitempty"`       // 布尔值
	 NumberValue     *string                                                                          `json:"number_value,omitempty"`     // 数字值
	 EnumValue       *string                                                                          `json:"enum_value,omitempty"`       // 枚举值, 这里是枚举的 id
	 DateValue       *string                                                                          `json:"date_value,omitempty"`       // 从 1970 开始的天数
	 DateTimeValue   *string                                                                          `json:"date_time_value,omitempty"`  // 时间戳, 毫秒
	 I18nValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueI18nValue   `json:"i18n_value,omitempty"`       // 多语字段值
	 ObjectValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueObjectValue `json:"object_value,omitempty"`     // 对象值, 包括对象 id 和对象类型
	 UserValue       *string                                                                          `json:"user_value,omitempty"`       // 用户 id, 根据 user_type 选择对应的用户 id
	 DepartmentValue *string                                                                          `json:"department_value,omitempty"` // 部门 id, 根据入参选择对应的部门 id
	 RecordValues     []*GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueRecordValue `json:"record_values,omitempty"`     // 记录类型字段值
	 EmploymentValue *string                                                                          `json:"employment_value,omitempty"` // 员工类型字段值, 为用户 id, 根据入参选择返回的用户 id
	 ListValues      []string                                                                        `json:"list_values,omitempty"`      // 数组类型值, 里面包含多个值, 每个元素都对应 subValues 中的数组下标
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueI18nValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueI18nValue struct {
	 ZhCn string `json:"zh_cn,omitempty"` // 中文值
	 EnUs string `json:"en_us,omitempty"` // 英文值
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueObjectValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueObjectValue struct {
	 WkID      string `json:"wk_id,omitempty"`       // 飞书人事 Wukong 元数据的对象唯一标识
	 WkApiName string `json:"wk_api_name,omitempty"` // 飞书人事 Wukong 元数据唯一标识
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueRecordValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueSubValueValueRecordValue struct {
	 VariableApiName string `json:"variable_api_name,omitempty"` // 变量唯一标识
	 SubValueKey   string `json:"sub_value_key,omitempty"`    // 变量值, 对应 subValues 中的 key
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName struct {
	 ZhCn string `json:"zh_cn,omitempty"` // 中文值
	 EnUs string `json:"en_us,omitempty"` // 英文值
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue struct {
	 TextValue       *string                                                                          `json:"text_value,omitempty"`       // 文本值
	 BoolValue       *bool                                                                            `json:"bool_value,omitempty"`       // 布尔值
	 NumberValue     *string                                                                          `json:"number_value,omitempty"`     // 数字值
	 EnumValue       *string                                                                          `json:"enum_value,omitempty"`       // 枚举值, 这里是枚举的 id
	 DateValue       *string                                                                          `json:"date_value,omitempty"`       // 从 1970 开始的天数
	 DateTimeValue   *string                                                                          `json:"date_time_value,omitempty"`  // 时间戳, 毫秒
	 I18nValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue   `json:"i18n_value,omitempty"`       // 多语字段值
	 ObjectValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue `json:"object_value,omitempty"`     // 对象值, 包括对象 id 和对象类型
	 UserValue       *string                                                                          `json:"user_value,omitempty"`       // 用户 id, 根据 user_type 选择对应的用户 id
	 DepartmentValue *string                                                                          `json:"department_value,omitempty"` // 部门 id, 根据入参选择对应的部门 id
	 RecordValues     []*GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueRecordValue `json:"record_values,omitempty"`     // 记录类型字段值
	 EmploymentValue *string                                                                          `json:"employment_value,omitempty"` // 员工类型字段值, 为用户 id, 根据入参选择返回的用户 id
	 ListValues      []string                                                                        `json:"list_values,omitempty"`      // 数组类型值, 里面包含多个值, 每个元素都对应 subValues 中的数组下标
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue struct {
	 ZhCn string `json:"zh_cn,omitempty"` // 中文值
	 EnUs string `json:"en_us,omitempty"` // 英文值
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue struct {
	 WkID      string `json:"wk_id,omitempty"`       // 飞书人事 Wukong 元数据的对象唯一标识
	 WkApiName string `json:"wk_api_name,omitempty"` // 飞书人事 Wukong 元数据唯一标识
 }
 
 // GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueRecordValue ...
 type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueRecordValue struct {
	 VariableApiName string `json:"variable_api_name,omitempty"` // 变量唯一标识
	 SubValueKey     string `json:"sub_value_key,omitempty"`     // 变量值, 对应 subValues 中的 key
 }
 
 // getCoreHRProcessFormVariableDataResp ...
 type getCoreHRProcessFormVariableDataResp struct {
	 Code  int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	 Msg   string                                `json:"msg,omitempty"`  // 错误描述
	 Data  *GetCoreHRProcessFormVariableDataResp `json:"data,omitempty"`
	 Error *ErrorDetail                          `json:"error,omitempty"`
 }
 