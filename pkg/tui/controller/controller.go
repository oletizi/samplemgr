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
	"github.com/oletizi/samplemgr/pkg/audio"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"log"
	"sync"
)

//go:generate mockgen -destination=../../../mocks/tui/controller/controller.go . Controller
type Controller interface {
	UpdateNode(node samplelib.Node)
}

type controller struct {
	mu            sync.Mutex
	ac            audio.Context
	ds            samplelib.DataSource
	eh            tui.ErrorHandler
	nv            view.NodeView
	iv            view.InfoView
	lv            view.LogView
	logger        tui.Logger
	currentPlayer audio.Player
	currentSample samplelib.Sample
}

// UpdateNode tells the controller to update the UI for a new node
func (c *controller) UpdateNode(node samplelib.Node) {
	c.logger.Print("Calling UpdateNode on node: " + node.Name())
	c.nv.UpdateNode(c.ds, node, c.nodeSelected, c.sampleSelected, c.nodeChosen, c.sampleChosen)
	c.iv.UpdateNode(c.ds, node)
}

// nodeSelected callback function for when a node is selected in the node view
func (c *controller) nodeSelected(node samplelib.Node) {
	c.iv.UpdateNode(c.ds, node)
}

// sampleSelected callback function for when a sample is selected in the node view
func (c *controller) sampleSelected(sample samplelib.Sample) {
	c.iv.UpdateSample(c.ds, sample)
}

// nodeChosen callback function for when a node is chosen in the node view
func (c *controller) nodeChosen(node samplelib.Node) {
	c.logger.Print("In controller nodeChosen: " + node.Name())
	c.UpdateNode(node)
}

// sampleChosen callback function for when a sample is chosen in the node view
func (c *controller) sampleChosen(newSample samplelib.Sample) {
	var err error
	currentSample, currentPlayer := c.getCurrentPlayer()

	// stop current playback, if any
	if currentPlayer != nil && currentPlayer.Playing() {
		err = currentPlayer.Stop()
		c.eh.Handle(err)
		// If the current sample is the same as the newSample, don't play the sample again.
		// This is the play/pause toggle condition.
		if newSample.Equal(currentSample) {
			return
		}
	}

	// if the chosen newSample is different than the current newSample, create a new newPlayer
	// and start playback
	newPlayer, err := c.ac.PlayerFor(newSample.Path())
	if err != nil {
		// notest
		c.eh.Handle(err)
		return
	}
	// Play the chosen newSample
	err = newPlayer.Play(func() {
		c.logger.Println("Done playing newSample! Closing the newPlayer...")
		err := newPlayer.Close()
		c.eh.Handle(err)
	})
	if err != nil {
		// notest
		c.eh.Handle(err)
		return
	}
	c.setCurrentPlayer(newSample, newPlayer)
}

func (c *controller) setCurrentPlayer(sample samplelib.Sample, player audio.Player) {
	c.mu.Lock()
	c.currentPlayer = player
	c.currentSample = sample
	c.mu.Unlock()
}

func (c *controller) getCurrentPlayer() (samplelib.Sample, audio.Player) {
	c.mu.Lock()
	p := c.currentPlayer
	s := c.currentSample
	c.mu.Unlock()
	return s, p
}

func New(
	ac audio.Context,
	ds samplelib.DataSource,
	eh tui.ErrorHandler,
	nodeView view.NodeView,
	infoView view.InfoView,
	logView view.LogView,
) Controller {
	return &controller{
		ac:     ac,
		ds:     ds,
		eh:     eh,
		nv:     nodeView,
		iv:     infoView,
		lv:     logView,
		logger: log.New(logView, "", 0)}
}
