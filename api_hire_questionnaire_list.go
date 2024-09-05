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

// GetHireQuestionnaireList 获取面试满意度问卷列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/questionnaire/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/delivery-process-management/interview/list-2
func (r *HireService) GetHireQuestionnaireList(ctx context.Context, request *GetHireQuestionnaireListReq, options ...MethodOptionFunc) (*GetHireQuestionnaireListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireQuestionnaireList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireQuestionnaireList mock enable")
		return r.cli.mock.mockHireGetHireQuestionnaireList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireQuestionnaireList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/questionnaires",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireQuestionnaireListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireQuestionnaireList mock HireGetHireQuestionnaireList method
func (r *Mock) MockHireGetHireQuestionnaireList(f func(ctx context.Context, request *GetHireQuestionnaireListReq, options ...MethodOptionFunc) (*GetHireQuestionnaireListResp, *Response, error)) {
	r.mockHireGetHireQuestionnaireList = f
}

// UnMockHireGetHireQuestionnaireList un-mock HireGetHireQuestionnaireList method
func (r *Mock) UnMockHireGetHireQuestionnaireList() {
	r.mockHireGetHireQuestionnaireList = nil
}

// GetHireQuestionnaireListReq ...
type GetHireQuestionnaireListReq struct {
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize        *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 100, 默认值: `1`
	ApplicationID   *string `query:"application_id" json:"-"`    // 投递 ID, 示例值: 6985833807195212076
	InterviewID     *string `query:"interview_id" json:"-"`      // 面试 ID, 示例值: 7038435261598763308
	UpdateStartTime *string `query:"update_start_time" json:"-"` // 最早更新时间, 示例值: 1638848468868
	UpdateEndTime   *string `query:"update_end_time" json:"-"`   // 最晚更新时间, 示例值: 1638848468869
}

// GetHireQuestionnaireListResp ...
type GetHireQuestionnaireListResp struct {
	HasMore   bool                                `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                              `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetHireQuestionnaireListRespItem `json:"items,omitempty"`      // 满意度评价列表
}

// GetHireQuestionnaireListRespItem ...
type GetHireQuestionnaireListRespItem struct {
	QuestionnaireID string                                      `json:"questionnaire_id,omitempty"` // 问卷 ID
	ApplicationID   string                                      `json:"application_id,omitempty"`   // 投递 ID；当「面试满意度问卷发送时间」选项选择「面试流程结束后」, 将返回 投递 ID
	InterviewID     string                                      `json:"interview_id,omitempty"`     // 面试 ID；当「面试满意度问卷发送时间」选项选择「第一次面试后」、「每次面试后」将返回 面试 ID
	Version         int64                                       `json:"version,omitempty"`          // 问卷版本
	Questions       []*GetHireQuestionnaireListRespItemQuestion `json:"questions,omitempty"`        // 题目列表
	HasAnswers      bool                                        `json:"has_answers,omitempty"`      // 是否完成作答
	UpdateTime      string                                      `json:"update_time,omitempty"`      // 更新时间
}

// GetHireQuestionnaireListRespItemQuestion ...
type GetHireQuestionnaireListRespItemQuestion struct {
	QuestionID             string                                                          `json:"question_id,omitempty"`               // 题目 ID
	QuestionName           string                                                          `json:"question_name,omitempty"`             // 题目中文名称
	QuestionEnName         string                                                          `json:"question_en_name,omitempty"`          // 题目英文名称
	QuestionDesc           string                                                          `json:"question_desc,omitempty"`             // 题目中文描述
	QuestionEnDesc         string                                                          `json:"question_en_desc,omitempty"`          // 题目英文描述
	QuestionType           int64                                                           `json:"question_type,omitempty"`             // 题目类型, 可选值有: 1: 单选题, 2: 多选题, 3: 描述题, 4: 评分题
	IsRequired             bool                                                            `json:"is_required,omitempty"`               // 是否必填
	SelectOptionResultList []*GetHireQuestionnaireListRespItemQuestionSelectOptionResult   `json:"select_option_result_list,omitempty"` // 选项题回答列表（单选题及多选题）
	FiveStartScoringResult *GetHireQuestionnaireListRespItemQuestionFiveStartScoringResult `json:"five_start_scoring_result,omitempty"` // 评分题回答
	DescriptionResult      string                                                          `json:"description_result,omitempty"`        // 描述题回答
}

// GetHireQuestionnaireListRespItemQuestionFiveStartScoringResult ...
type GetHireQuestionnaireListRespItemQuestionFiveStartScoringResult struct {
	HighestScoreDesc   string  `json:"highest_score_desc,omitempty"`    // 最高分中文描述
	HighestScoreEnDesc string  `json:"highest_score_en_desc,omitempty"` // 最高分英文描述
	LowestScoreDesc    string  `json:"lowest_score_desc,omitempty"`     // 最低分中文描述
	LowestScoreEnDesc  string  `json:"lowest_score_en_desc,omitempty"`  // 最低分英文描述
	ScoreResult        float64 `json:"score_result,omitempty"`          // 评分分数
}

// GetHireQuestionnaireListRespItemQuestionSelectOptionResult ...
type GetHireQuestionnaireListRespItemQuestionSelectOptionResult struct {
	OptionID     string `json:"option_id,omitempty"`      // 选项 ID
	OptionName   string `json:"option_name,omitempty"`    // 选项中文名称
	OptionEnName string `json:"option_en_name,omitempty"` // 选项英文名称
	OptionDesc   string `json:"option_desc,omitempty"`    // 选项中文描述
	OptionEnDesc string `json:"option_en_desc,omitempty"` // 选项英文描述
	IsSelected   bool   `json:"is_selected,omitempty"`    // 是否选择
}

// getHireQuestionnaireListResp ...
type getHireQuestionnaireListResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireQuestionnaireListResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}