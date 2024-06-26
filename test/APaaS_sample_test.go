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

func Test_APaaS_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.APaaS

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AgreeAPaaSApprovalTask(ctx, &lark.AgreeAPaaSApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.APaaS

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAPaaSAgreeAPaaSApprovalTask(func(ctx context.Context, request *lark.AgreeAPaaSApprovalTaskReq, options ...lark.MethodOptionFunc) (*lark.AgreeAPaaSApprovalTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAPaaSAgreeAPaaSApprovalTask()

			_, _, err := moduleCli.AgreeAPaaSApprovalTask(ctx, &lark.AgreeAPaaSApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAPaaSRejectAPaaSApprovalTask(func(ctx context.Context, request *lark.RejectAPaaSApprovalTaskReq, options ...lark.MethodOptionFunc) (*lark.RejectAPaaSApprovalTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAPaaSRejectAPaaSApprovalTask()

			_, _, err := moduleCli.RejectAPaaSApprovalTask(ctx, &lark.RejectAPaaSApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAPaaSTransferAPaaSApprovalTask(func(ctx context.Context, request *lark.TransferAPaaSApprovalTaskReq, options ...lark.MethodOptionFunc) (*lark.TransferAPaaSApprovalTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAPaaSTransferAPaaSApprovalTask()

			_, _, err := moduleCli.TransferAPaaSApprovalTask(ctx, &lark.TransferAPaaSApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAPaaSAddAssigneeAPaaSApprovalTask(func(ctx context.Context, request *lark.AddAssigneeAPaaSApprovalTaskReq, options ...lark.MethodOptionFunc) (*lark.AddAssigneeAPaaSApprovalTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAPaaSAddAssigneeAPaaSApprovalTask()

			_, _, err := moduleCli.AddAssigneeAPaaSApprovalTask(ctx, &lark.AddAssigneeAPaaSApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.APaaS

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AgreeAPaaSApprovalTask(ctx, &lark.AgreeAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RejectAPaaSApprovalTask(ctx, &lark.RejectAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TransferAPaaSApprovalTask(ctx, &lark.TransferAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddAssigneeAPaaSApprovalTask(ctx, &lark.AddAssigneeAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.APaaS
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AgreeAPaaSApprovalTask(ctx, &lark.AgreeAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RejectAPaaSApprovalTask(ctx, &lark.RejectAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TransferAPaaSApprovalTask(ctx, &lark.TransferAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddAssigneeAPaaSApprovalTask(ctx, &lark.AddAssigneeAPaaSApprovalTaskReq{
				ApprovalTaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
