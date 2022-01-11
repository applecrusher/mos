//
// Copyright (c) 2014-2019 Cesanta Software Limited
// All rights reserved
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

//go:generate go-bindata -pkg esp32c3 -nocompress -modtime 1 -mode 420 stub/stub.json

package esp32c3

import (
	"fmt"

	"github.com/mongoose-os/mos/cli/flash/esp"
)

var (
	FlashSizeToId = map[string]int{
		// +1, to distinguish from null-value
		"8m":   1,
		"16m":  2,
		"32m":  3,
		"64m":  4,
		"128m": 5,
	}

	FlashSizes = map[int]int{
		0: 1048576,
		1: 2097152,
		2: 4194304,
		3: 8388608,
		4: 16777216,
	}
)

func GetChipDescr(rrw esp.RegReaderWriter) (string, error) {
	chip_rev := 3
	chip_pkg := "ESP32C3"

	return fmt.Sprintf("%s R%d", chip_pkg, chip_rev), nil
}
