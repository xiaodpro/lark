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
package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

func (r *Mindnote) meta(ctx context.Context) (*lark.GetDriveFileMetaRespMeta, error) {
	resp, _, err := r.larkClient.Drive.GetDriveFileMeta(ctx, &lark.GetDriveFileMetaReq{
		RequestDocs: []*lark.GetDriveFileMetaReqRequestDocs{{
			DocToken: r.token, DocType: r.typ,
		}},
	})
	if err != nil {
		return nil, err
	} else if len(resp.Metas) == 0 {
		return nil, nil
	}
	return resp.Metas[0], nil
}
