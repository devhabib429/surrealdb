// Copyright © 2016 Abcum Ltd
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

package math

func Bottom(vals []float64, take int) []float64 {

	var out []float64

	if len(vals) == 0 {
		return nil
	}

	sort := Sort(vals)

	for i := 0; i < take && i < len(vals); i++ {
		out = append(out, sort[i])
	}

	return out
}