// Copyright 2023 The Bazel Authors. All rights reserved.
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

package python

import (
	"bufio"
	_ "embed"
	"strings"
)

var (
	//go:embed stdlib_list.txt
	stdlibList string
	stdModules map[string]struct{}
)

func init() {
	stdModules = make(map[string]struct{})
	scanner := bufio.NewScanner(strings.NewReader(stdlibList))
	for scanner.Scan() {
		stdModules[scanner.Text()] = struct{}{}
	}
}

func isStdModule(m Module) bool {
	_, ok := stdModules[m.Name]
	return ok
}
