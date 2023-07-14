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

// GetPersonalSettingsSystemStatusList 获取租户下所有系统状态。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/list
// new doc: https://open.feishu.cn/document/server-docs/personal_settings-v1/system_status/list
func (r *PersonalSettingsService) GetPersonalSettingsSystemStatusList(ctx context.Context, request *GetPersonalSettingsSystemStatusListReq, options ...MethodOptionFunc) (*GetPersonalSettingsSystemStatusListResp, *Response, error) {
	if r.cli.mock.mockPersonalSettingsGetPersonalSettingsSystemStatusList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] PersonalSettings#GetPersonalSettingsSystemStatusList mock enable")
		return r.cli.mock.mockPersonalSettingsGetPersonalSettingsSystemStatusList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "PersonalSettings",
		API:                   "GetPersonalSettingsSystemStatusList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/personal_settings/v1/system_statuses",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPersonalSettingsSystemStatusListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockPersonalSettingsGetPersonalSettingsSystemStatusList mock PersonalSettingsGetPersonalSettingsSystemStatusList method
func (r *Mock) MockPersonalSettingsGetPersonalSettingsSystemStatusList(f func(ctx context.Context, request *GetPersonalSettingsSystemStatusListReq, options ...MethodOptionFunc) (*GetPersonalSettingsSystemStatusListResp, *Response, error)) {
	r.mockPersonalSettingsGetPersonalSettingsSystemStatusList = f
}

// UnMockPersonalSettingsGetPersonalSettingsSystemStatusList un-mock PersonalSettingsGetPersonalSettingsSystemStatusList method
func (r *Mock) UnMockPersonalSettingsGetPersonalSettingsSystemStatusList() {
	r.mockPersonalSettingsGetPersonalSettingsSystemStatusList = nil
}

// GetPersonalSettingsSystemStatusListReq ...
type GetPersonalSettingsSystemStatusListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 50, 默认值: `50`, 取值范围: `1` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ[
}

// GetPersonalSettingsSystemStatusListResp ...
type GetPersonalSettingsSystemStatusListResp struct {
	Items     []*GetPersonalSettingsSystemStatusListRespItem `json:"items,omitempty"`      // 租户系统状态
	PageToken string                                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                           `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetPersonalSettingsSystemStatusListRespItem ...
type GetPersonalSettingsSystemStatusListRespItem struct {
	SystemStatusID string                                                  `json:"system_status_id,omitempty"` // 系统状态ID
	Title          string                                                  `json:"title,omitempty"`            // 系统状态名称, 名称字符数要在1到20范围内。不同系统状态的title不能重复, 注意: 1中文=2英文=2其他语言字符=2字符
	I18nTitle      *GetPersonalSettingsSystemStatusListRespItemI18nTitle   `json:"i18n_title,omitempty"`       // 系统状态国际化名称, 名称字符数要在1到20范围内。不同系统状态之间i18n_title中任何一种title都不能重复, 注意: 1中文=2英文=2其他语言字符=2字符
	IconKey        string                                                  `json:"icon_key,omitempty"`         // 图标, [了解icon_key可选值](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/overview), 可选值有: GeneralDoNotDisturb: GeneralDoNotDisturb, GeneralInMeetingBusy: GeneralInMeetingBusy, Coffee: Coffee, GeneralBusinessTrip: GeneralBusinessTrip, GeneralWorkFromHome: GeneralWorkFromHome, StatusEnjoyLife: StatusEnjoyLife, GeneralTravellingCar: GeneralTravellingCar, StatusBus: StatusBus, StatusInFlight: StatusInFlight, Typing: Typing, EatingFood: EatingFood, SICK: SICK, GeneralSun: GeneralSun, GeneralMoonRest: GeneralMoonRest, StatusReading: StatusReading, Status_PrivateMessage: Status_PrivateMessage, StatusFlashOfInspiration: StatusFlashOfInspiration, GeneralVacation: GeneralVacation
	Color          string                                                  `json:"color,omitempty"`            // 颜色, 可选值有: BLUE: 蓝色, GRAY: 灰色, INDIGO: 靛青色, WATHET: 浅蓝色, GREEN: 绿色, TURQUOISE: 绿松石色, YELLOW: 黄色, LIME: 酸橙色, RED: 红色, ORANGE: 橙色, PURPLE: 紫色, VIOLET: 紫罗兰色, CARMINE: 胭脂红色
	Priority       int64                                                   `json:"priority,omitempty"`         // 优先级, 数值越小, 客户端展示的优先级越高。不同系统状态的优先级不能一样。
	SyncSetting    *GetPersonalSettingsSystemStatusListRespItemSyncSetting `json:"sync_setting,omitempty"`     // 同步设置
}

// GetPersonalSettingsSystemStatusListRespItemI18nTitle ...
type GetPersonalSettingsSystemStatusListRespItemI18nTitle struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名
	EnUs string `json:"en_us,omitempty"` // 英文名
	JaJp string `json:"ja_jp,omitempty"` // 日文名
}

// GetPersonalSettingsSystemStatusListRespItemSyncSetting ...
type GetPersonalSettingsSystemStatusListRespItemSyncSetting struct {
	IsOpenByDefault bool                                                               `json:"is_open_by_default,omitempty"` // 是否默认开启
	Title           string                                                             `json:"title,omitempty"`              // 同步设置名称, 名称字符数要在1到30范围内, 注意: 1中文=2英文=2其他语言字符=2字符
	I18nTitle       *GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nTitle   `json:"i18n_title,omitempty"`         // 同步设置国际化名称, 名称字符数要在1到30范围内, 注意: 1中文=2英文=2其他语言字符=2字符
	Explain         string                                                             `json:"explain,omitempty"`            // 同步设置解释文案, 解释字符数要在1到60范围内, 注意: 1中文=2英文=2其他语言字符=2字符
	I18nExplain     *GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nExplain `json:"i18n_explain,omitempty"`       // 同步设置国际化解释文案, 解释字符数要在1到60范围内, 注意: 1中文=2英文=2其他语言字符=2字符
}

// GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nExplain ...
type GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nExplain struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名
	EnUs string `json:"en_us,omitempty"` // 英文名
	JaJp string `json:"ja_jp,omitempty"` // 日文名
}

// GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nTitle ...
type GetPersonalSettingsSystemStatusListRespItemSyncSettingI18nTitle struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名
	EnUs string `json:"en_us,omitempty"` // 英文名
	JaJp string `json:"ja_jp,omitempty"` // 日文名
}

// getPersonalSettingsSystemStatusListResp ...
type getPersonalSettingsSystemStatusListResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *GetPersonalSettingsSystemStatusListResp `json:"data,omitempty"`
}
