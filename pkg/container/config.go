// Copyright 2022 Chainguard, Inc.
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

package container

import (
	"log"
)

type BindMount struct {
	Source      string
	Destination string
}

type Capabilities struct {
	Networking bool
}

type Config struct {
	Mounts       []BindMount
	Capabilities Capabilities
	Logger       *log.Logger
	Environment  map[string]string
	ImgDigest    string
	PodID        string
}
