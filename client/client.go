// Copyright © 2018 Bradley Weston <hello@bweston.me>.
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
package client // github.com/go-openwrks/openwrks/client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type (
	// Opt is a function that will manipulate the client.
	Opt func(*Client) error

	// OptError will provide the name of the option that failed and the
	// corrosponding message.
	OptError struct {
		Name    string
		Message string
	}

	// Client will communicate with Openwrks API.
	Client struct {
		http *http.Client
		env  Environment
		key  string
	}
)

// NewClient will take a client and apply the given options to it.
// If any of the options return an error, Client will be nil and the error
// from the option will be returned straight away which is of type OptError.
func NewClient(opts ...Opt) (*Client, error) {
	// We will initialise the client with sane defaults.
	c := &Client{
		env:  LiveEnvironment,
		http: http.DefaultClient,
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Do will attempt to make a HTTP request to Openwrks.
// URLs must start with a "/" and will be relative to the environment base URI.
func (c *Client) Do(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	if url[0] != '/' {
		return nil, errors.New("Openwrks URL must start with a /")
	}
	url = string(c.env) + url
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.key)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.http.Do(req)
}

// Error will convert the OptError into a human readable string. :D
func (err OptError) Error() string {
	return fmt.Sprintf("Option %s error: %s", err.Name, err.Message)
}

// WithHTTPClient will set the HTTPClient to use.
func WithHTTPClient(hc *http.Client) Opt {
	return func(c *Client) error {
		if hc == nil {
			return OptError{
				Name:    "WithHTTPClient",
				Message: "The client provided cannot be nil",
			}
		}
		c.http = hc
		return nil
	}
}

// WithEnv will set the environment to use.
func WithEnv(env Environment) Opt {
	return func(c *Client) error {
		c.env = env
		return nil
	}
}

// WithKey will set the API key to use.
func WithKey(key string) Opt {
	return func(c *Client) error {
		if key == "" {
			return OptError{
				Name:    "WithKey",
				Message: "API key cannot be empty.",
			}
		}

		c.key = key
		return nil
	}
}
