// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// AddContactGroupMember 向用户组中添加成员(目前成员仅支持用户，未来会支持部门)，如果应用的通讯录权限范围是“全部员工”，则可将任何成员添加到任何用户组。如果应用的通讯录权限范围不是“全部员工”，则仅可将通讯录权限范围中的成员添加到通讯录权限范围的用户组中，[点击了解通讯录权限范围](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/overview)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/add
func (r *ContactService) AddContactGroupMember(ctx context.Context, request *AddContactGroupMemberReq, options ...MethodOptionFunc) (*AddContactGroupMemberResp, *Response, error) {
	if r.cli.mock.mockContactAddContactGroupMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#AddContactGroupMember mock enable")
		return r.cli.mock.mockContactAddContactGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "AddContactGroupMember",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/group/:group_id/member/add",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(addContactGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactAddContactGroupMember(f func(ctx context.Context, request *AddContactGroupMemberReq, options ...MethodOptionFunc) (*AddContactGroupMemberResp, *Response, error)) {
	r.mockContactAddContactGroupMember = f
}

func (r *Mock) UnMockContactAddContactGroupMember() {
	r.mockContactAddContactGroupMember = nil
}

type AddContactGroupMemberReq struct {
	GroupID      string `path:"group_id" json:"-"`        // 用户组ID, 示例值："g281721"
	MemberType   string `json:"member_type,omitempty"`    // 用户组成员的类型，取值为 user, 示例值："user", 可选值有: `user`：user, 默认值: `user`
	MemberIDType IDType `json:"member_id_type,omitempty"` // 当member_type =user时候，member_id_type表示user_id_type，枚举值为open_id, union_id, user_id, 示例值："open_id", 可选值有: `open_id`：member_type =user时候，表示用户的open_id, `union_id`：member_type =user时候，表示用户的union_id, `user_id`：member_type =user时候，表示用户的user_id
	MemberID     string `json:"member_id,omitempty"`      // 添加的成员ID, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

type addContactGroupMemberResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *AddContactGroupMemberResp `json:"data,omitempty"`
}

type AddContactGroupMemberResp struct{}