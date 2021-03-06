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

package main

import "strings"

// List of reserved words for files, includes old and new ones.
var posixReservedPrefix = []string{
	"$tmpfile",
	// Add new reserved words if any used in future.
}

// hasPosixReservedPrefix - has reserved prefix.
func hasPosixReservedPrefix(name string) (isReserved bool) {
	for _, reservedKey := range posixReservedPrefix {
		if strings.HasPrefix(name, reservedKey) {
			isReserved = true
			break
		}
		isReserved = false
	}
	return isReserved
}
