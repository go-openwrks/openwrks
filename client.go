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
package openwrks // github.com/go-openwrks/openwrks

import (
	"fmt"
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
		Env Environment
		Key string
	}
)

// NewClient will take a client and apply the given options to it.
// If any of the options return an error, Client will be nil and the error
// from the option will be returned straight away which is of type OptError.
func NewClient(opts ...Opt) (*Client, error) {
	// We will initialise the client with sane defaults.
	c := &Client{
		Env: LiveEnvironment,
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Error will convert the OptError into a human readable string. :D
func (err OptError) Error() string {
	return fmt.Sprintf("Option %s error: %s", err.Name, err.Message)
}

// WithEnv will set the environment to use.
func WithEnv(env Environment) Opt {
	return func(c *Client) error {
		c.Env = env
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

		c.Key = key
		return nil
	}
}
