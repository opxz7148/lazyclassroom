package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type selectedState struct {
	model *CourseListModel
}

func (s *selectedState) SelectStyleSet() {}
func (s *selectedState) UnselectStyleSet() {

	clm := s.model

	// Only set styles if width is valid (after WindowSizeMsg)
	if clm.Width() <= 0 {
		return
	}

	clm.SetWidth(clm.Width() - 1)

	availableWidth := clm.Width()

	clm.activeDelegate.Styles.SelectedTitle = clm.activeDelegate.Styles.SelectedTitle.
		Border(lipgloss.NormalBorder(), true, true, false, true).
		Width(availableWidth).
		UnsetMarginTop()
	clm.activeDelegate.Styles.SelectedDesc = clm.activeDelegate.Styles.SelectedDesc.
		Border(lipgloss.NormalBorder(), false, true, true, true).
		Width(availableWidth).
		UnsetMarginBottom()
	clm.Model.SetDelegate(*clm.activeDelegate)

	clm.ToggleState()
}

type unselectedState struct {
	model *CourseListModel
}

func (u *unselectedState) UnselectStyleSet() {}
func (u *unselectedState) SelectStyleSet() {
	clm := u.model

	// Only set styles if width is valid (after WindowSizeMsg)
	if clm.Width() <= 0 { return }

	clm.SetWidth(clm.Width() + 1)

	setSelectedState(clm)

	clm.ToggleState()
}


type intializedState struct {
	model *CourseListModel
}

func (i *intializedState) SelectStyleSet()   {
	clm := i.model

	// Only set styles if width is valid (after WindowSizeMsg)
	if clm.Width() <= 0 { return }
	setSelectedState(clm)
	clm.changeToSelectState()
	fmt.Println(clm.state)
}
func (u *intializedState) UnselectStyleSet() {}


// Helper function to set selected state
func setSelectedState(clm *CourseListModel) {
	availableWidth := clm.Width()

	clm.activeDelegate.Styles.SelectedTitle = clm.activeDelegate.Styles.SelectedTitle.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Width(availableWidth).
		MarginTop(1)
	clm.activeDelegate.Styles.SelectedDesc = clm.activeDelegate.Styles.SelectedDesc.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Width(availableWidth).
		MarginBottom(1)
	clm.Model.SetDelegate(*clm.activeDelegate)
}