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

package controller

import (
	"github.com/golang/mock/gomock"
	mocksamplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mocktui "github.com/oletizi/samplemgr/mocks/tui"
	mockview "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	ctl := gomock.NewController(t)
	ds := mocksamplelib.NewMockDataSource(ctl)
	errorHandler := mocktui.NewMockErrorHandler(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	logView := mockview.NewMockLogView(ctl)

	c := New(ds, errorHandler, nodeView, infoView, logView)
	assert.NotNil(t, c)
}

func TestController_UpdateNode(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ds := mocksamplelib.NewMockDataSource(ctl)
	eh := mocktui.NewMockErrorHandler(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	logView := mockview.NewMockLogView(ctl)
	node := mocksamplelib.NewMockNode(ctl)

	// make a new controller
	c := &controller{
		ds:     ds,
		nv:     nodeView,
		iv:     infoView,
		lv:     logView,
		eh:     eh,
		logger: log.Default(),
	}
	// XXX: can't figure out how to get function arguments to match
	nodeView.EXPECT().UpdateNode(ds, node, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
	infoView.EXPECT().UpdateNode(node)
	node.EXPECT().Name()
	assert.NotNil(t, c)
	c.UpdateNode(node)
}

func TestController_selectNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	infoView := mockview.NewMockInfoView(ctl)
	node := mocksamplelib.NewMockNode(ctl)
	sample := mocksamplelib.NewMockSample(ctl)

	c := &controller{
		iv: infoView,
	}

	infoView.EXPECT().UpdateNode(node)
	c.nodeSelected(node)

	infoView.EXPECT().UpdateSample(sample)
	c.sampleSelected(sample)
}

func TestController_chooseNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ds := mocksamplelib.NewMockDataSource(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	node := mocksamplelib.NewMockNode(ctl)
	sample := mocksamplelib.NewMockSample(ctl)

	c := &controller{
		ds:     ds,
		nv:     nodeView,
		iv:     infoView,
		logger: tui.NewLogger(log.Default()),
	}

	nodeView.EXPECT().UpdateNode(ds, node, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	infoView.EXPECT().UpdateNode(node)
	node.EXPECT().Name().Return("node name").AnyTimes()
	c.nodeChosen(node)

	infoView.EXPECT().UpdateSample(sample)
	c.sampleChosen(sample)
}
