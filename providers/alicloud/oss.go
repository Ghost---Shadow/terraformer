// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alicloud

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// OSSGenerator Generator for OSS service
type OSSGenerator struct {
	AliCloudService
}

// InitResources Need bucket name as ID for terraform resource
func (g *OSSGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	raw, err := client.WithOssClient(func(ossClient *oss.Client) (interface{}, error) {
		return ossClient.ListBuckets()
	})

	bucketResult := (raw).(oss.ListBucketsResult)
	if err != nil {
		return err
	}
	s, _ := json.MarshalIndent(bucketResult, "", "\t")
	fmt.Println(string(s))

	return nil
}
