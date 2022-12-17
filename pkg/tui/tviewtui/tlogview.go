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

package tviewtui

import "github.com/rivo/tview"

type tLogView struct {
	textView *tview.TextView
}

func (t *tLogView) Close() error {
	//TODO implement me
	panic("implement me")
}

func (t *tLogView) Clear() {
	//TODO implement me
	panic("implement me")
}

func (t *tLogView) Write(p []byte) (n int, err error) {
	return t.textView.Write(p)
}
