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

// GetOKRPeriodList 获取OKR周期列表
//
// 使用tenant_access_token需要额外申请权限以应用身份访问OKR信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/list
func (r *OKRService) GetOKRPeriodList(ctx context.Context, request *GetOKRPeriodListReq, options ...MethodOptionFunc) (*GetOKRPeriodListResp, *Response, error) {
	if r.cli.mock.mockOKRGetOKRPeriodList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetOKRPeriodList mock enable")
		return r.cli.mock.mockOKRGetOKRPeriodList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetOKRPeriodList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/periods",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getOKRPeriodListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRGetOKRPeriodList mock OKRGetOKRPeriodList method
func (r *Mock) MockOKRGetOKRPeriodList(f func(ctx context.Context, request *GetOKRPeriodListReq, options ...MethodOptionFunc) (*GetOKRPeriodListResp, *Response, error)) {
	r.mockOKRGetOKRPeriodList = f
}

// UnMockOKRGetOKRPeriodList un-mock OKRGetOKRPeriodList method
func (r *Mock) UnMockOKRGetOKRPeriodList() {
	r.mockOKRGetOKRPeriodList = nil
}

// GetOKRPeriodListReq ...
type GetOKRPeriodListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标志page_token, 示例值："xaasdasdax"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小，默认10, 示例值：10, 默认值: `10`
}

// getOKRPeriodListResp ...
type getOKRPeriodListResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetOKRPeriodListResp `json:"data,omitempty"`
}

// GetOKRPeriodListResp ...
type GetOKRPeriodListResp struct {
	PageToken string                      `json:"page_token,omitempty"` // 分页标志
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否有更多
	Items     []*GetOKRPeriodListRespItem `json:"items,omitempty"`      // 数据项
}

// GetOKRPeriodListRespItem ...
type GetOKRPeriodListRespItem struct {
	ID     string `json:"id,omitempty"`      // id
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
	Status int64  `json:"status,omitempty"`  // 启用状态, 可选值有: `0`：正常状态, `1`：暂不处理, `2`：标记失效, `3`：隐藏周期
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// GetUserOKRList 根据用户的id获取OKR列表
//
// 使用tenant_access_token需要额外申请权限以应用身份访问OKR信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/user-okr/list
func (r *OKRService) GetUserOKRList(ctx context.Context, request *GetUserOKRListReq, options ...MethodOptionFunc) (*GetUserOKRListResp, *Response, error) {
	if r.cli.mock.mockOKRGetUserOKRList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetUserOKRList mock enable")
		return r.cli.mock.mockOKRGetUserOKRList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetUserOKRList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/users/:user_id/okrs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getUserOKRListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRGetUserOKRList mock OKRGetUserOKRList method
func (r *Mock) MockOKRGetUserOKRList(f func(ctx context.Context, request *GetUserOKRListReq, options ...MethodOptionFunc) (*GetUserOKRListResp, *Response, error)) {
	r.mockOKRGetUserOKRList = f
}

// UnMockOKRGetUserOKRList un-mock OKRGetUserOKRList method
func (r *Mock) UnMockOKRGetUserOKRList() {
	r.mockOKRGetUserOKRList = nil
}

// GetUserOKRListReq ...
type GetUserOKRListReq struct {
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, `people_admin_id`：以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求:  获取用户 user ID
	Offset     string   `query:"offset" json:"-"`       // 请求列表的偏移，offset>=0，请求Query中, 示例值："0"
	Limit      string   `query:"limit" json:"-"`        // 请求列表的长度，0<limit<=10，请求Query中, 示例值："0"
	Lang       *string  `query:"lang" json:"-"`         // 请求OKR的语言版本（比如@的人名），lang=en_us/zh_cn，请求 Query中, 示例值："zh_cn", 默认值: `zh_cn`
	PeriodIDs  []string `query:"period_ids" json:"-"`   // period_id列表，最多10个, 示例值：["6951461264858777132"], 最大长度：`10`
	UserID     string   `path:"user_id" json:"-"`       // 目标用户id, 示例值："ou-asdasdasdasdasd"
}

// getUserOKRListResp ...
type getUserOKRListResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetUserOKRListResp `json:"data,omitempty"`
}

// GetUserOKRListResp ...
type GetUserOKRListResp struct {
	Total   int64                    `json:"total,omitempty"`    // OKR周期总数
	OKRList []*GetUserOKRListRespOKR `json:"okr_list,omitempty"` // OKR 列表
}

// GetUserOKRListRespOKR ...
type GetUserOKRListRespOKR struct {
	ID            string                            `json:"id,omitempty"`             // id
	Permission    int64                             `json:"permission,omitempty"`     // OKR的访问权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	PeriodID      string                            `json:"period_id,omitempty"`      // period_id
	Name          string                            `json:"name,omitempty"`           // 名称
	ObjectiveList []*GetUserOKRListRespOKRObjective `json:"objective_list,omitempty"` // Objective列表
}

// GetUserOKRListRespOKRObjective ...
type GetUserOKRListRespOKRObjective struct {
	ID                    string                                             `json:"id,omitempty"`                      // Objective ID
	Permission            int64                                              `json:"permission,omitempty"`              // 权限, 可选值有: `0`：此时OKR只返回id, `1`：返回OKR的其他具体字段
	Content               string                                             `json:"content,omitempty"`                 // Objective 内容
	ProgressReport        string                                             `json:"progress_report,omitempty"`         // Objective 进度记录内容
	Score                 int64                                              `json:"score,omitempty"`                   // Objective 分数（0 - 100）
	Weight                float64                                            `json:"weight,omitempty"`                  // Objective的权重（0 - 100）
	ProgressRate          *GetUserOKRListRespOKRObjectiveProgressRate        `json:"progress_rate,omitempty"`           // Objective进度
	KrList                []*GetUserOKRListRespOKRObjectiveKr                `json:"kr_list,omitempty"`                 // Objective KeyResult 列表
	AlignedObjectiveList  []*GetUserOKRListRespOKRObjectiveAlignedObjective  `json:"aligned_objective_list,omitempty"`  // 对齐到该Objective的Objective列表
	AligningObjectiveList []*GetUserOKRListRespOKRObjectiveAligningObjective `json:"aligning_objective_list,omitempty"` // 该Objective对齐到的Objective列表
}

// GetUserOKRListRespOKRObjectiveProgressRate ...
type GetUserOKRListRespOKRObjectiveProgressRate struct {
	Percent int64  `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

// GetUserOKRListRespOKRObjectiveKr ...
type GetUserOKRListRespOKRObjectiveKr struct {
	ID           string                                        `json:"id,omitempty"`            // Key Result ID
	Content      string                                        `json:"content,omitempty"`       // KeyResult 内容
	Score        int64                                         `json:"score,omitempty"`         // KeyResult打分（0 - 100）
	Weight       int64                                         `json:"weight,omitempty"`        // KeyResult权重（0 - 100）（废弃）
	KrWeight     float64                                       `json:"kr_weight,omitempty"`     // KeyResult的权重（0 - 100）
	ProgressRate *GetUserOKRListRespOKRObjectiveKrProgressRate `json:"progress_rate,omitempty"` // KR进度
}

// GetUserOKRListRespOKRObjectiveKrProgressRate ...
type GetUserOKRListRespOKRObjectiveKrProgressRate struct {
	Percent int64  `json:"percent,omitempty"` // Objective 进度百分比 >= 0
	Status  string `json:"status,omitempty"`  // Objective 进度状态, 可选值有: `-1`：未更新, `0`：正常, `1`：有风险, `2`：已延期
}

// GetUserOKRListRespOKRObjectiveAlignedObjective ...
type GetUserOKRListRespOKRObjectiveAlignedObjective struct {
	ID    string                                               `json:"id,omitempty"`     // Objective的ID
	OKRID string                                               `json:"okr_id,omitempty"` // OKR的ID
	Owner *GetUserOKRListRespOKRObjectiveAlignedObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

// GetUserOKRListRespOKRObjectiveAlignedObjectiveOwner ...
type GetUserOKRListRespOKRObjectiveAlignedObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
	UserID string `json:"user_id,omitempty"` // 用户的 user_id
}

// GetUserOKRListRespOKRObjectiveAligningObjective ...
type GetUserOKRListRespOKRObjectiveAligningObjective struct {
	ID    string                                                `json:"id,omitempty"`     // Objective的ID
	OKRID string                                                `json:"okr_id,omitempty"` // OKR的ID
	Owner *GetUserOKRListRespOKRObjectiveAligningObjectiveOwner `json:"owner,omitempty"`  // 该Objective的Owner
}

// GetUserOKRListRespOKRObjectiveAligningObjectiveOwner ...
type GetUserOKRListRespOKRObjectiveAligningObjectiveOwner struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
	UserID string `json:"user_id,omitempty"` // 用户的 user_id
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// CreateSearchDataSource 创建一个数据源
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create
func (r *SearchService) CreateSearchDataSource(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchCreateSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#CreateSearchDataSource mock enable")
		return r.cli.mock.mockSearchCreateSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "CreateSearchDataSource",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchCreateSearchDataSource mock SearchCreateSearchDataSource method
func (r *Mock) MockSearchCreateSearchDataSource(f func(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error)) {
	r.mockSearchCreateSearchDataSource = f
}

// UnMockSearchCreateSearchDataSource un-mock SearchCreateSearchDataSource method
func (r *Mock) UnMockSearchCreateSearchDataSource() {
	r.mockSearchCreateSearchDataSource = nil
}

// CreateSearchDataSourceReq ...
type CreateSearchDataSourceReq struct {
	Name        string  `json:"name,omitempty"`        // data_source的展示名称, 示例值："客服工单"
	State       *int64  `json:"state,omitempty"`       // 数据源状态，0-未上线，1-已上线, 示例值：0, 可选值有: `0`：未上线, `1`：已上线
	Description *string `json:"description,omitempty"` // 对于数据源的描述, 示例值："搜索客服工单数据"
}

// createSearchDataSourceResp ...
type createSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateSearchDataSourceResp `json:"data,omitempty"`
}

// CreateSearchDataSourceResp ...
type CreateSearchDataSourceResp struct {
	DataSource *CreateSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源实例
}

// CreateSearchDataSourceRespDataSource ...
type CreateSearchDataSourceRespDataSource struct {
	ID            string `json:"id,omitempty"`              // 数据源的唯一标识
	Name          string `json:"name,omitempty"`            // data_source的展示名称
	State         int64  `json:"state,omitempty"`           // 数据源状态，0-未上线，1-已上线, 可选值有: `0`：未上线, `1`：已上线
	Description   string `json:"description,omitempty"`     // 对于数据源的描述
	CreateTime    string `json:"create_time,omitempty"`     // 创建时间，使用Unix时间戳，单位为“秒”
	UpdateTime    string `json:"update_time,omitempty"`     // 更新时间，使用Unix时间戳，单位为“秒”
	IsExceedQuota bool   `json:"is_exceed_quota,omitempty"` // 是否超限
}
