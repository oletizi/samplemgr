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

import (
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/rivo/tview"
)

type tInfoView struct {
	textView *tview.TextView
	logger   tui.Logger
	eh       tui.ErrorHandler
	display  view.Display
}

func (t *tInfoView) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *tInfoView) Close() error {
	//TODO implement me
	panic("implement me")
}

func (t *tInfoView) Clear() {
	//TODO implement me
	panic("implement me")
}

func (t *tInfoView) UpdateNode(node samplelib.Node) {
	//TODO implement me
	panic("implement me")
}

func (t *tInfoView) UpdateSample(sample samplelib.Sample) {
	//TODO implement me
	panic("implement me")
}