// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetUserByID
//
// 根据用户邮箱或手机号查询用户 open_id 和 user_id，支持批量查询。<br>
// 调用该接口需要申请 `通过手机号或者邮箱获取用户ID` 权限。<br>只能查询到应用可用性范围内的用户 ID，不在范围内的用户会表现为不存在。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN
func (r *ContactService) BatchGetUserByID(ctx context.Context, request *BatchGetUserByIDReq, options ...MethodOptionFunc) (*BatchGetUserByIDResp, *Response, error) {
	if r.cli.mock.mockContactBatchGetUserByID != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#BatchGetUserByID mock enable")
		return r.cli.mock.mockContactBatchGetUserByID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "BatchGetUserByID",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/user/v1/batch_get_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetUserByIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactBatchGetUserByID(f func(ctx context.Context, request *BatchGetUserByIDReq, options ...MethodOptionFunc) (*BatchGetUserByIDResp, *Response, error)) {
	r.mockContactBatchGetUserByID = f
}

func (r *Mock) UnMockContactBatchGetUserByID() {
	r.mockContactBatchGetUserByID = nil
}

type BatchGetUserByIDReq struct {
	Emails  []string `query:"emails" json:"-"`  // 要查询的用户邮箱，最多 50 条。
	Mobiles []string `query:"mobiles" json:"-"` // 要查询的用户手机号，最多 50 条。<br>非中国大陆地区的手机号需要添加以 “+” 开头的国家 / 地区代码，并且需要进行 URL 转义。<br>
}

type batchGetUserByIDResp struct {
	Code int64                 `json:"code,omitempty"` // 返回码，非 0 表示失败。
	Msg  string                `json:"msg,omitempty"`  // 对返回码的文本描述。
	Data *BatchGetUserByIDResp `json:"data,omitempty"` // -
}

type BatchGetUserByIDResp struct {
	EmailUsers      map[string][]*BatchGetUserByIDRespEmailUser `json:"email_users,omitempty"`       // 根据邮箱查询到的用户，key 为邮箱，value 为查询到用户的 array。<br>目前同一个邮箱最多只能查询到一个用户。
	EmailsNotExist  []string                                    `json:"emails_not_exist,omitempty"`  // 没有匹配记录的邮箱。
	MobileUsers     map[string][]*BatchGetUserByIDRespEmailUser `json:"mobile_users,omitempty"`      // 根据手机号查询到的用户，key 为手机号，value 为查询到用户的 array。<br>目前同一个手机号最多只能查询到一个用户。
	MobilesNotExist []string                                    `json:"mobiles_not_exist,omitempty"` // 没有匹配记录的手机号。
}

type BatchGetUserByIDRespEmailUser struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id。
	UserID string `json:"user_id,omitempty"` // 用户的 user_id。<br>只有已申请 `获取用户UserID` 权限的企业自建应用返回此字段。
}
