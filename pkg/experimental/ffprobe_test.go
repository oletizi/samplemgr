/*
 * Copyright (c) 2022 Orion Letizi
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
 *
 */

package experimental

import (
	"context"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"testing"
	"time"
)

func Test_ffProbe(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, "../../test/data/library/one-level/hh.wav")
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	} else {
		log.Println("format: " + data.Format.FormatName)
	}

}