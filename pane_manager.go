package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Selectable interface {
	Select()
	Unselect()
}

type SelectableModel interface {
	Selectable
	tea.Model
}

// PaneManager manages pane selection and switching
type PaneManager struct {
	paneList   map[int]SelectableModel
	selectPane int
	numPanes   int
}

// NewPaneManager creates a new PaneManager
func NewPaneManager(numPanes int) *PaneManager {
	return &PaneManager{
		paneList:   make(map[int]SelectableModel),
		selectPane: CourseListPaneID,
		numPanes:   numPanes,
	}
}

// NextPane switches to the next pane
func (pm *PaneManager) NextPane() {
	pm.selectPane = (pm.selectPane + 1) % pm.numPanes
}

// GetPane returns the pane at the given index
func (pm *PaneManager) GetPane(paneID int) (SelectableModel, bool) {
	pane, exists := pm.paneList[paneID]
	return pane, exists
}

// SetPane sets the pane at the given index
func (pm *PaneManager) SetPane(paneID int, pane SelectableModel) {
	if paneID < 0 || paneID >= pm.numPanes {
		return
	}
	pm.paneList[paneID] = pane
}

// GetSelectedPane returns the currently selected pane
func (pm *PaneManager) GetSelectedPane() (SelectableModel, bool) {
	return pm.GetPane(pm.selectPane)
}

// GetSelectedPaneID returns the ID of the currently selected pane
func (pm *PaneManager) GetSelectedPaneID() int {
	return pm.selectPane
}

// GetAllPanes returns all panes
func (pm *PaneManager) GetAllPanes() map[int]SelectableModel {
	return pm.paneList
}

// Update updates the currently selected pane with the given message
// and reassigns it internally. Returns the command from the update.
func (pm *PaneManager) Update(msg tea.Msg) tea.Cmd {
	pane, exists := pm.GetSelectedPane()
	if !exists {
		return nil
	}

	updatedPane, cmd := pane.Update(msg)
	pm.SetPane(pm.GetSelectedPaneID(), updatedPane.(SelectableModel))
	return cmd
}

// UpdateSelected applies select/unselect styles to all panes
func (pm *PaneManager) UpdateSelected() {
	for i, pane := range pm.paneList {
		if i == pm.selectPane {
			pane.Select()
		} else {
			pane.Unselect()
		}
	}
}
