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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Chat_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Chat

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Chat

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatCreateChat(func(ctx context.Context, request *lark.CreateChatReq, options ...lark.MethodOptionFunc) (*lark.CreateChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatCreateChat()

			_, _, err := moduleCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChat(func(ctx context.Context, request *lark.GetChatReq, options ...lark.MethodOptionFunc) (*lark.GetChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChat()

			_, _, err := moduleCli.GetChat(ctx, &lark.GetChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatOld(func(ctx context.Context, request *lark.GetChatOldReq, options ...lark.MethodOptionFunc) (*lark.GetChatOldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatOld()

			_, _, err := moduleCli.GetChatOld(ctx, &lark.GetChatOldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatUpdateChat(func(ctx context.Context, request *lark.UpdateChatReq, options ...lark.MethodOptionFunc) (*lark.UpdateChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatUpdateChat()

			_, _, err := moduleCli.UpdateChat(ctx, &lark.UpdateChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatDeleteChat(func(ctx context.Context, request *lark.DeleteChatReq, options ...lark.MethodOptionFunc) (*lark.DeleteChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatDeleteChat()

			_, _, err := moduleCli.DeleteChat(ctx, &lark.DeleteChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatListOfSelf(func(ctx context.Context, request *lark.GetChatListOfSelfReq, options ...lark.MethodOptionFunc) (*lark.GetChatListOfSelfResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatListOfSelf()

			_, _, err := moduleCli.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatSearchChat(func(ctx context.Context, request *lark.SearchChatReq, options ...lark.MethodOptionFunc) (*lark.SearchChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatSearchChat()

			_, _, err := moduleCli.SearchChat(ctx, &lark.SearchChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatMemberList(func(ctx context.Context, request *lark.GetChatMemberListReq, options ...lark.MethodOptionFunc) (*lark.GetChatMemberListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatMemberList()

			_, _, err := moduleCli.GetChatMemberList(ctx, &lark.GetChatMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatIsInChat(func(ctx context.Context, request *lark.IsInChatReq, options ...lark.MethodOptionFunc) (*lark.IsInChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatIsInChat()

			_, _, err := moduleCli.IsInChat(ctx, &lark.IsInChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatCreateChatManager(func(ctx context.Context, request *lark.CreateChatManagerReq, options ...lark.MethodOptionFunc) (*lark.CreateChatManagerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatCreateChatManager()

			_, _, err := moduleCli.CreateChatManager(ctx, &lark.CreateChatManagerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatDeleteChatManager(func(ctx context.Context, request *lark.DeleteChatManagerReq, options ...lark.MethodOptionFunc) (*lark.DeleteChatManagerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatDeleteChatManager()

			_, _, err := moduleCli.DeleteChatManager(ctx, &lark.DeleteChatManagerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatAddChatMember(func(ctx context.Context, request *lark.AddChatMemberReq, options ...lark.MethodOptionFunc) (*lark.AddChatMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatAddChatMember()

			_, _, err := moduleCli.AddChatMember(ctx, &lark.AddChatMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatDeleteChatMember(func(ctx context.Context, request *lark.DeleteChatMemberReq, options ...lark.MethodOptionFunc) (*lark.DeleteChatMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatDeleteChatMember()

			_, _, err := moduleCli.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatJoinChat(func(ctx context.Context, request *lark.JoinChatReq, options ...lark.MethodOptionFunc) (*lark.JoinChatResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatJoinChat()

			_, _, err := moduleCli.JoinChat(ctx, &lark.JoinChatReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatModeration(func(ctx context.Context, request *lark.GetChatModerationReq, options ...lark.MethodOptionFunc) (*lark.GetChatModerationResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatModeration()

			_, _, err := moduleCli.GetChatModeration(ctx, &lark.GetChatModerationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatUpdateChatModeration(func(ctx context.Context, request *lark.UpdateChatModerationReq, options ...lark.MethodOptionFunc) (*lark.UpdateChatModerationResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatUpdateChatModeration()

			_, _, err := moduleCli.UpdateChatModeration(ctx, &lark.UpdateChatModerationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatUpdateChatTopNotice(func(ctx context.Context, request *lark.UpdateChatTopNoticeReq, options ...lark.MethodOptionFunc) (*lark.UpdateChatTopNoticeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatUpdateChatTopNotice()

			_, _, err := moduleCli.UpdateChatTopNotice(ctx, &lark.UpdateChatTopNoticeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatDeleteChatTopNotice(func(ctx context.Context, request *lark.DeleteChatTopNoticeReq, options ...lark.MethodOptionFunc) (*lark.DeleteChatTopNoticeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatDeleteChatTopNotice()

			_, _, err := moduleCli.DeleteChatTopNotice(ctx, &lark.DeleteChatTopNoticeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGenChatShareLink(func(ctx context.Context, request *lark.GenChatShareLinkReq, options ...lark.MethodOptionFunc) (*lark.GenChatShareLinkResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGenChatShareLink()

			_, _, err := moduleCli.GenChatShareLink(ctx, &lark.GenChatShareLinkReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatAnnouncement(func(ctx context.Context, request *lark.GetChatAnnouncementReq, options ...lark.MethodOptionFunc) (*lark.GetChatAnnouncementResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatAnnouncement()

			_, _, err := moduleCli.GetChatAnnouncement(ctx, &lark.GetChatAnnouncementReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatUpdateChatAnnouncement(func(ctx context.Context, request *lark.UpdateChatAnnouncementReq, options ...lark.MethodOptionFunc) (*lark.UpdateChatAnnouncementResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatUpdateChatAnnouncement()

			_, _, err := moduleCli.UpdateChatAnnouncement(ctx, &lark.UpdateChatAnnouncementReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatCreateChatTab(func(ctx context.Context, request *lark.CreateChatTabReq, options ...lark.MethodOptionFunc) (*lark.CreateChatTabResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatCreateChatTab()

			_, _, err := moduleCli.CreateChatTab(ctx, &lark.CreateChatTabReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatDeleteChatTab(func(ctx context.Context, request *lark.DeleteChatTabReq, options ...lark.MethodOptionFunc) (*lark.DeleteChatTabResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatDeleteChatTab()

			_, _, err := moduleCli.DeleteChatTab(ctx, &lark.DeleteChatTabReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatGetChatTabList(func(ctx context.Context, request *lark.GetChatTabListReq, options ...lark.MethodOptionFunc) (*lark.GetChatTabListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatGetChatTabList()

			_, _, err := moduleCli.GetChatTabList(ctx, &lark.GetChatTabListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatUpdateChatTab(func(ctx context.Context, request *lark.UpdateChatTabReq, options ...lark.MethodOptionFunc) (*lark.UpdateChatTabResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatUpdateChatTab()

			_, _, err := moduleCli.UpdateChatTab(ctx, &lark.UpdateChatTabReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockChatSortChatTab(func(ctx context.Context, request *lark.SortChatTabReq, options ...lark.MethodOptionFunc) (*lark.SortChatTabResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockChatSortChatTab()

			_, _, err := moduleCli.SortChatTab(ctx, &lark.SortChatTabReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Chat

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChat(ctx, &lark.GetChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatOld(ctx, &lark.GetChatOldReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChat(ctx, &lark.UpdateChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChat(ctx, &lark.DeleteChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchChat(ctx, &lark.SearchChatReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatMemberList(ctx, &lark.GetChatMemberListReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.IsInChat(ctx, &lark.IsInChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChatManager(ctx, &lark.CreateChatManagerReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatManager(ctx, &lark.DeleteChatManagerReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddChatMember(ctx, &lark.AddChatMemberReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.JoinChat(ctx, &lark.JoinChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatModeration(ctx, &lark.GetChatModerationReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatModeration(ctx, &lark.UpdateChatModerationReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatTopNotice(ctx, &lark.UpdateChatTopNoticeReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatTopNotice(ctx, &lark.DeleteChatTopNoticeReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GenChatShareLink(ctx, &lark.GenChatShareLinkReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatAnnouncement(ctx, &lark.GetChatAnnouncementReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatAnnouncement(ctx, &lark.UpdateChatAnnouncementReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChatTab(ctx, &lark.CreateChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatTab(ctx, &lark.DeleteChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatTabList(ctx, &lark.GetChatTabListReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatTab(ctx, &lark.UpdateChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SortChatTab(ctx, &lark.SortChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Chat
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChat(ctx, &lark.CreateChatReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChat(ctx, &lark.GetChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatOld(ctx, &lark.GetChatOldReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChat(ctx, &lark.UpdateChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChat(ctx, &lark.DeleteChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchChat(ctx, &lark.SearchChatReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatMemberList(ctx, &lark.GetChatMemberListReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.IsInChat(ctx, &lark.IsInChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChatManager(ctx, &lark.CreateChatManagerReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatManager(ctx, &lark.DeleteChatManagerReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddChatMember(ctx, &lark.AddChatMemberReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.JoinChat(ctx, &lark.JoinChatReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatModeration(ctx, &lark.GetChatModerationReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatModeration(ctx, &lark.UpdateChatModerationReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatTopNotice(ctx, &lark.UpdateChatTopNoticeReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatTopNotice(ctx, &lark.DeleteChatTopNoticeReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GenChatShareLink(ctx, &lark.GenChatShareLinkReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatAnnouncement(ctx, &lark.GetChatAnnouncementReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatAnnouncement(ctx, &lark.UpdateChatAnnouncementReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateChatTab(ctx, &lark.CreateChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteChatTab(ctx, &lark.DeleteChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetChatTabList(ctx, &lark.GetChatTabListReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateChatTab(ctx, &lark.UpdateChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SortChatTab(ctx, &lark.SortChatTabReq{
				ChatID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
