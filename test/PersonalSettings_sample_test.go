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

func Test_PersonalSettings_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.PersonalSettings

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePersonalSettingsSystemStatus(ctx, &lark.CreatePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.PersonalSettings

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsCreatePersonalSettingsSystemStatus(func(ctx context.Context, request *lark.CreatePersonalSettingsSystemStatusReq, options ...lark.MethodOptionFunc) (*lark.CreatePersonalSettingsSystemStatusResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsCreatePersonalSettingsSystemStatus()

			_, _, err := moduleCli.CreatePersonalSettingsSystemStatus(ctx, &lark.CreatePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsDeletePersonalSettingsSystemStatus(func(ctx context.Context, request *lark.DeletePersonalSettingsSystemStatusReq, options ...lark.MethodOptionFunc) (*lark.DeletePersonalSettingsSystemStatusResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsDeletePersonalSettingsSystemStatus()

			_, _, err := moduleCli.DeletePersonalSettingsSystemStatus(ctx, &lark.DeletePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsUpdatePersonalSettingsSystemStatus(func(ctx context.Context, request *lark.UpdatePersonalSettingsSystemStatusReq, options ...lark.MethodOptionFunc) (*lark.UpdatePersonalSettingsSystemStatusResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsUpdatePersonalSettingsSystemStatus()

			_, _, err := moduleCli.UpdatePersonalSettingsSystemStatus(ctx, &lark.UpdatePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsGetPersonalSettingsSystemStatusList(func(ctx context.Context, request *lark.GetPersonalSettingsSystemStatusListReq, options ...lark.MethodOptionFunc) (*lark.GetPersonalSettingsSystemStatusListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsGetPersonalSettingsSystemStatusList()

			_, _, err := moduleCli.GetPersonalSettingsSystemStatusList(ctx, &lark.GetPersonalSettingsSystemStatusListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsBatchOpenPersonalSettingsSystemStatus(func(ctx context.Context, request *lark.BatchOpenPersonalSettingsSystemStatusReq, options ...lark.MethodOptionFunc) (*lark.BatchOpenPersonalSettingsSystemStatusResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsBatchOpenPersonalSettingsSystemStatus()

			_, _, err := moduleCli.BatchOpenPersonalSettingsSystemStatus(ctx, &lark.BatchOpenPersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockPersonalSettingsBatchClosePersonalSettingsSystemStatus(func(ctx context.Context, request *lark.BatchClosePersonalSettingsSystemStatusReq, options ...lark.MethodOptionFunc) (*lark.BatchClosePersonalSettingsSystemStatusResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockPersonalSettingsBatchClosePersonalSettingsSystemStatus()

			_, _, err := moduleCli.BatchClosePersonalSettingsSystemStatus(ctx, &lark.BatchClosePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.PersonalSettings

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePersonalSettingsSystemStatus(ctx, &lark.CreatePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePersonalSettingsSystemStatus(ctx, &lark.DeletePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePersonalSettingsSystemStatus(ctx, &lark.UpdatePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPersonalSettingsSystemStatusList(ctx, &lark.GetPersonalSettingsSystemStatusListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchOpenPersonalSettingsSystemStatus(ctx, &lark.BatchOpenPersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchClosePersonalSettingsSystemStatus(ctx, &lark.BatchClosePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.PersonalSettings
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePersonalSettingsSystemStatus(ctx, &lark.CreatePersonalSettingsSystemStatusReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePersonalSettingsSystemStatus(ctx, &lark.DeletePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePersonalSettingsSystemStatus(ctx, &lark.UpdatePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPersonalSettingsSystemStatusList(ctx, &lark.GetPersonalSettingsSystemStatusListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchOpenPersonalSettingsSystemStatus(ctx, &lark.BatchOpenPersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchClosePersonalSettingsSystemStatus(ctx, &lark.BatchClosePersonalSettingsSystemStatusReq{
				SystemStatusID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
