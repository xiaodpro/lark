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

func Test_TaskV1_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.TaskV1

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Follower(ctx, &lark.CreateTaskV1FollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.TaskV1

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CreateTaskV1Follower(func(ctx context.Context, request *lark.CreateTaskV1FollowerReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskV1FollowerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CreateTaskV1Follower()

			_, _, err := moduleCli.CreateTaskV1Follower(ctx, &lark.CreateTaskV1FollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1DeleteTaskV1Follower(func(ctx context.Context, request *lark.DeleteTaskV1FollowerReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskV1FollowerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1DeleteTaskV1Follower()

			_, _, err := moduleCli.DeleteTaskV1Follower(ctx, &lark.DeleteTaskV1FollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1BatchDeleteTaskV1Follower(func(ctx context.Context, request *lark.BatchDeleteTaskV1FollowerReq, options ...lark.MethodOptionFunc) (*lark.BatchDeleteTaskV1FollowerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1BatchDeleteTaskV1Follower()

			_, _, err := moduleCli.BatchDeleteTaskV1Follower(ctx, &lark.BatchDeleteTaskV1FollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskFollowerV1List(func(ctx context.Context, request *lark.GetTaskFollowerV1ListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskFollowerV1ListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskFollowerV1List()

			_, _, err := moduleCli.GetTaskFollowerV1List(ctx, &lark.GetTaskFollowerV1ListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CreateTaskV1Collaborator(func(ctx context.Context, request *lark.CreateTaskV1CollaboratorReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskV1CollaboratorResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CreateTaskV1Collaborator()

			_, _, err := moduleCli.CreateTaskV1Collaborator(ctx, &lark.CreateTaskV1CollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1DeleteTaskV1Collaborator(func(ctx context.Context, request *lark.DeleteTaskV1CollaboratorReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskV1CollaboratorResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1DeleteTaskV1Collaborator()

			_, _, err := moduleCli.DeleteTaskV1Collaborator(ctx, &lark.DeleteTaskV1CollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1BatchDeleteTaskV1Collaborator(func(ctx context.Context, request *lark.BatchDeleteTaskV1CollaboratorReq, options ...lark.MethodOptionFunc) (*lark.BatchDeleteTaskV1CollaboratorResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1BatchDeleteTaskV1Collaborator()

			_, _, err := moduleCli.BatchDeleteTaskV1Collaborator(ctx, &lark.BatchDeleteTaskV1CollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1CollaboratorList(func(ctx context.Context, request *lark.GetTaskV1CollaboratorListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskV1CollaboratorListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1CollaboratorList()

			_, _, err := moduleCli.GetTaskV1CollaboratorList(ctx, &lark.GetTaskV1CollaboratorListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CreateTaskV1Reminder(func(ctx context.Context, request *lark.CreateTaskV1ReminderReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskV1ReminderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CreateTaskV1Reminder()

			_, _, err := moduleCli.CreateTaskV1Reminder(ctx, &lark.CreateTaskV1ReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1ReminderList(func(ctx context.Context, request *lark.GetTaskV1ReminderListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskV1ReminderListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1ReminderList()

			_, _, err := moduleCli.GetTaskV1ReminderList(ctx, &lark.GetTaskV1ReminderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1DeleteTaskV1Reminder(func(ctx context.Context, request *lark.DeleteTaskV1ReminderReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskV1ReminderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1DeleteTaskV1Reminder()

			_, _, err := moduleCli.DeleteTaskV1Reminder(ctx, &lark.DeleteTaskV1ReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CreateTaskV1(func(ctx context.Context, request *lark.CreateTaskV1Req, options ...lark.MethodOptionFunc) (*lark.CreateTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CreateTaskV1()

			_, _, err := moduleCli.CreateTaskV1(ctx, &lark.CreateTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1(func(ctx context.Context, request *lark.GetTaskV1Req, options ...lark.MethodOptionFunc) (*lark.GetTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1()

			_, _, err := moduleCli.GetTaskV1(ctx, &lark.GetTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1List(func(ctx context.Context, request *lark.GetTaskV1ListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskV1ListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1List()

			_, _, err := moduleCli.GetTaskV1List(ctx, &lark.GetTaskV1ListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1DeleteTaskV1(func(ctx context.Context, request *lark.DeleteTaskV1Req, options ...lark.MethodOptionFunc) (*lark.DeleteTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1DeleteTaskV1()

			_, _, err := moduleCli.DeleteTaskV1(ctx, &lark.DeleteTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1UpdateTaskV1(func(ctx context.Context, request *lark.UpdateTaskV1Req, options ...lark.MethodOptionFunc) (*lark.UpdateTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1UpdateTaskV1()

			_, _, err := moduleCli.UpdateTaskV1(ctx, &lark.UpdateTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CompleteTaskV1(func(ctx context.Context, request *lark.CompleteTaskV1Req, options ...lark.MethodOptionFunc) (*lark.CompleteTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CompleteTaskV1()

			_, _, err := moduleCli.CompleteTaskV1(ctx, &lark.CompleteTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1UncompleteTaskV1(func(ctx context.Context, request *lark.UncompleteTaskV1Req, options ...lark.MethodOptionFunc) (*lark.UncompleteTaskV1Resp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1UncompleteTaskV1()

			_, _, err := moduleCli.UncompleteTaskV1(ctx, &lark.UncompleteTaskV1Req{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1CreateTaskV1Comment(func(ctx context.Context, request *lark.CreateTaskV1CommentReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskV1CommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1CreateTaskV1Comment()

			_, _, err := moduleCli.CreateTaskV1Comment(ctx, &lark.CreateTaskV1CommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1Comment(func(ctx context.Context, request *lark.GetTaskV1CommentReq, options ...lark.MethodOptionFunc) (*lark.GetTaskV1CommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1Comment()

			_, _, err := moduleCli.GetTaskV1Comment(ctx, &lark.GetTaskV1CommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1GetTaskV1CommentList(func(ctx context.Context, request *lark.GetTaskV1CommentListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskV1CommentListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1GetTaskV1CommentList()

			_, _, err := moduleCli.GetTaskV1CommentList(ctx, &lark.GetTaskV1CommentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1DeleteTaskV1Comment(func(ctx context.Context, request *lark.DeleteTaskV1CommentReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskV1CommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1DeleteTaskV1Comment()

			_, _, err := moduleCli.DeleteTaskV1Comment(ctx, &lark.DeleteTaskV1CommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockTaskV1UpdateTaskV1Comment(func(ctx context.Context, request *lark.UpdateTaskV1CommentReq, options ...lark.MethodOptionFunc) (*lark.UpdateTaskV1CommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskV1UpdateTaskV1Comment()

			_, _, err := moduleCli.UpdateTaskV1Comment(ctx, &lark.UpdateTaskV1CommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.TaskV1

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Follower(ctx, &lark.CreateTaskV1FollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Follower(ctx, &lark.DeleteTaskV1FollowerReq{
				TaskID:     "x",
				FollowerID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchDeleteTaskV1Follower(ctx, &lark.BatchDeleteTaskV1FollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskFollowerV1List(ctx, &lark.GetTaskFollowerV1ListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Collaborator(ctx, &lark.CreateTaskV1CollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Collaborator(ctx, &lark.DeleteTaskV1CollaboratorReq{
				TaskID:         "x",
				CollaboratorID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchDeleteTaskV1Collaborator(ctx, &lark.BatchDeleteTaskV1CollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1CollaboratorList(ctx, &lark.GetTaskV1CollaboratorListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Reminder(ctx, &lark.CreateTaskV1ReminderReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1ReminderList(ctx, &lark.GetTaskV1ReminderListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Reminder(ctx, &lark.DeleteTaskV1ReminderReq{
				TaskID:     "x",
				ReminderID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1(ctx, &lark.CreateTaskV1Req{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1(ctx, &lark.GetTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1List(ctx, &lark.GetTaskV1ListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1(ctx, &lark.DeleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateTaskV1(ctx, &lark.UpdateTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CompleteTaskV1(ctx, &lark.CompleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UncompleteTaskV1(ctx, &lark.UncompleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Comment(ctx, &lark.CreateTaskV1CommentReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1Comment(ctx, &lark.GetTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1CommentList(ctx, &lark.GetTaskV1CommentListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Comment(ctx, &lark.DeleteTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateTaskV1Comment(ctx, &lark.UpdateTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.TaskV1
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Follower(ctx, &lark.CreateTaskV1FollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Follower(ctx, &lark.DeleteTaskV1FollowerReq{
				TaskID:     "x",
				FollowerID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchDeleteTaskV1Follower(ctx, &lark.BatchDeleteTaskV1FollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskFollowerV1List(ctx, &lark.GetTaskFollowerV1ListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Collaborator(ctx, &lark.CreateTaskV1CollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Collaborator(ctx, &lark.DeleteTaskV1CollaboratorReq{
				TaskID:         "x",
				CollaboratorID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchDeleteTaskV1Collaborator(ctx, &lark.BatchDeleteTaskV1CollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1CollaboratorList(ctx, &lark.GetTaskV1CollaboratorListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Reminder(ctx, &lark.CreateTaskV1ReminderReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1ReminderList(ctx, &lark.GetTaskV1ReminderListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Reminder(ctx, &lark.DeleteTaskV1ReminderReq{
				TaskID:     "x",
				ReminderID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1(ctx, &lark.CreateTaskV1Req{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1(ctx, &lark.GetTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1List(ctx, &lark.GetTaskV1ListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1(ctx, &lark.DeleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateTaskV1(ctx, &lark.UpdateTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CompleteTaskV1(ctx, &lark.CompleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UncompleteTaskV1(ctx, &lark.UncompleteTaskV1Req{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateTaskV1Comment(ctx, &lark.CreateTaskV1CommentReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1Comment(ctx, &lark.GetTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetTaskV1CommentList(ctx, &lark.GetTaskV1CommentListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteTaskV1Comment(ctx, &lark.DeleteTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateTaskV1Comment(ctx, &lark.UpdateTaskV1CommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
