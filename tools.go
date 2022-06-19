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

package cloudflare

import (
	"context"
	"fmt"
)

// UpdateRecord updates a DNS zone record.
func UpdateRecord(ctx context.Context, c *Client, zone, recordType, recordName, value string) error {
	zoneID, err := getZoneID(ctx, c, zone)
	if err != nil {
		return fmt.Errorf("cloudflare update record: %s", err)
	}
	recordID, err := getRecordID(ctx, c, zoneID, recordType, recordName)
	if err != nil {
		return fmt.Errorf("cloudflare update record: %s", err)
	}
	if err := patchRecord(ctx, c, zoneID, recordID, value); err != nil {
		return fmt.Errorf("cloudflare update record: %s", err)
	}
	return nil
}

func getZoneID(ctx context.Context, c *Client, zone string) (string, error) {
	resp, err := c.Call(ctx, "GET", "zones?match=any&name="+zone, nil)
	if err != nil {
		return "", err
	}
	res, err := getFirstResult(resp)
	if err != nil {
		return "", fmt.Errorf("get zone id: %s", err)
	}
	zoneID, ok := res["id"].(string)
	if !ok {
		return "", fmt.Errorf("get zone id: unexpected id type: %q", resp)
	}
	if zoneID == "" {
		return "", fmt.Errorf("get zone id: missing id: %q", resp)
	}
	return zoneID, nil
}

func getRecordID(ctx context.Context, c *Client, zoneID, recordType, recordName string) (string, error) {
	path := "zones/" + zoneID + "/dns_records?match=any&type=" + recordType + "&name=" + recordName
	resp, err := c.Call(ctx, "GET", path, nil)
	if err != nil {
		return "", err
	}
	res, err := getFirstResult(resp)
	if err != nil {
		return "", fmt.Errorf("get record id: %s", err)
	}
	recordID, ok := res["id"].(string)
	if !ok {
		return "", fmt.Errorf("get record id: unexpected id type: %q", resp)
	}
	if recordID == "" {
		return "", fmt.Errorf("get record id: missing id: %q", resp)
	}
	return recordID, nil
}

func patchRecord(ctx context.Context, c *Client, zoneID, recordID, value string) error {
	data := map[string]any{"content": value}
	resp, err := c.Call(ctx, "PATCH", "zones/"+zoneID+"/dns_records/"+recordID, data)
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("patch record: no success: %q", resp)
	}
	return nil
}

func getFirstResult(resp *Response) (map[string]any, error) {
	if !resp.Success {
		return nil, fmt.Errorf("get first result: no success: %q", resp)
	}
	things, ok := resp.Result.([]any)
	if !ok {
		return nil, fmt.Errorf("get first result: unexpected result: %q", resp)
	}
	if len(things) < 1 {
		return nil, fmt.Errorf("get first result: no results: %q", resp)
	}
	thing, ok := things[0].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("get first result: unexpected result item: %q", resp)
	}
	return thing, nil
}
