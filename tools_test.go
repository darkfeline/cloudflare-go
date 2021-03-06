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

	"go.felesatra.moe/cloudflare"
)

func ExampleUpdateRecord() {
	ctx := context.Background()
	c := &cloudflare.Client{Token: "token"}
	cloudflare.UpdateRecord(ctx, c, "example.com", "A", "www", "1.2.3.4")
}
