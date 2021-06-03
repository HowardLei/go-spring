/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package toml

import (
	"github.com/pelletier/go-toml"
)

// Parse 将字节数组解析成 map 结构。
func Parse(b []byte) (map[string]interface{}, error) {
	tree, err := toml.LoadBytes(b)
	if err != nil {
		return nil, err
	}
	return tree.ToMap(), nil
}
