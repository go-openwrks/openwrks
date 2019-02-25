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

// Environment is a API URL prefix.
type Environment string

const (
	// SandboxEnvironment is to explore our API, make requests and receive responses with sample data
	SandboxEnvironment Environment = "https://api.openwrks.com/surface-sandbox/"
	// LabEnvironment once you’ve been on-boarded, you'll get access to a version of Flow that will
	// allow your users to connect their own bank accounts. This will allow you to perform full end to
	// end authentication and authorisation to individual banks.
	// You can test low volumes against real banks, real customers and ensure your integration is ready
	// ahead of a full live roll out
	LabEnvironment Environment = "https://api.openwrks.com/surface-lab/"
	// LiveEnvironment once you’ve got it all working in the Lab, this environment will use live data.
	LiveEnvironment Environment = "https://api.openwrks.com/surface/"
)
