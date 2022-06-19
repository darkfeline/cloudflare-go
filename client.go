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

// Package cloudflare provides a client for the Cloudflare API v4.
//
// This is not an official package.
//
// This package provides a simple client and functions that implement
// simple functionality like updating a single A record.
//
// This package handles authentication and basic marshaling of
// requests/unmarshaling of responses.  You will need to reference the
// API docs and do your own method calls for anything beyond the
// provided basics.
package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// A Client represents a Cloudflare API v4 HTTP RPC client.
type Client struct {
	C http.Client
	// See the API docs for the client token.
	Token string
}

// Call an API RPC.
//
// method is the HTTP method (e.g., GET).
// path is the HTTP path.
// data is something that can be encoded into JSON to pass as the body.
// If data is nil, no body is included (e.g., for GET calls).
func (c *Client) Call(ctx context.Context, method string, path string, data any) (*Response, error) {
	var b io.ReadWriter
	if data != nil {
		b = bytes.NewBuffer(nil)
		if err := json.NewEncoder(b).Encode(data); err != nil {
			return nil, fmt.Errorf("api call %s: %s", path, err)
		}
	}
	const base = "https://api.cloudflare.com/client/v4/"
	req, err := http.NewRequestWithContext(ctx, method, base+path, b)
	if err != nil {
		return nil, fmt.Errorf("api call %s: %s", path, err)
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.C.Do(req)
	if err != nil {
		return nil, fmt.Errorf("api call %s: %s", path, err)
	}
	defer resp.Body.Close()
	var r *Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("api call %s: %s", path, err)
	}
	return r, nil
}

// A Response represents an API call response.
// See the API docs for more info.
type Response struct {
	Success    bool        `json:"success"`
	Errors     []Error     `json:"errors"`
	Messages   []string    `json:"messages"`
	Result     any         `json:"result"`
	ResultInfo *ResultInfo `json:"result_info"`
}

func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// A ResultInfo represents some API response info.
// See the API docs for more info.
type ResultInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

func (r ResultInfo) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// An Error represents an API call error.
// See the API docs for more info.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) String() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b)
}
