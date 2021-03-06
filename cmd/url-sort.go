/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "net/url"

type byHostPath []*url.URL

func (s byHostPath) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byHostPath) Len() int {
	return len(s)
}

// Note: Host in url.URL includes the port too.
func (s byHostPath) Less(i, j int) bool {
	return (s[i].Host + s[i].Path) < (s[j].Host + s[j].Path)
}
