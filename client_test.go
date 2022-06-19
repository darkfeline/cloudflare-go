// Copyright (C) 2022  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudflare_test

import (
	"context"
	"fmt"

	"go.felesatra.moe/cloudflare"
)

func Example_GET() {
	ctx := context.Background()
	c := cloudflare.Client{
		Token: "secret token",
	}
	resp, err := c.Call(ctx, "GET", "zones?match=any&name=example.com", nil)
	if err != nil {
		panic(err)
	}
	if !resp.Success {
		panic("fail")
	}
	zoneID := resp.Result.([]any)[0].(map[string]any)["id"].(string)
	fmt.Printf(zoneID)
}

func Example_PATCH() {
	ctx := context.Background()
	c := cloudflare.Client{
		Token: "secret token",
	}
	resp, err := c.Call(ctx, "PATCH", "zones/zoneID/dns_records/recordID",
		map[string]any{"content": "1.2.3.4"})
	if err != nil {
		panic(err)
	}
	if !resp.Success {
		panic("fail")
	}
}
