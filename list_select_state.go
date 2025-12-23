package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// ============================================
// State pattern for CourseListPane
// ============================================

type selectedPaneState struct {
	pane *CourseListPane
}

func (s *selectedPaneState) SelectStyleSet() {}
func (s *selectedPaneState) UnselectStyleSet() {
	p := s.pane

	if p.Width() <= 0 {
		return
	}

	p.SetWidth(p.Width() - 1)

	availableWidth := p.Width()

	p.activeDelegate.Styles.SelectedTitle = p.activeDelegate.Styles.SelectedTitle.
		Border(lipgloss.NormalBorder(), true, true, false, true).
		Width(availableWidth).
		UnsetMarginTop()
	p.activeDelegate.Styles.SelectedDesc = p.activeDelegate.Styles.SelectedDesc.
		Border(lipgloss.NormalBorder(), false, true, true, true).
		Width(availableWidth).
		UnsetMarginBottom()
	p.Model.SetDelegate(*p.activeDelegate)

	p.ToggleState()
}

type unselectedPaneState struct {
	pane *CourseListPane
}

func (u *unselectedPaneState) UnselectStyleSet() {}
func (u *unselectedPaneState) SelectStyleSet() {
	p := u.pane

	if p.Width() <= 0 {
		return
	}

	p.SetWidth(p.Width() + 1)

	setSelectedPaneStyle(p)

	p.ToggleState()
}

type initializedPaneState struct {
	pane *CourseListPane
}

func (i *initializedPaneState) SelectStyleSet() {
	p := i.pane

	if p.Width() <= 0 {
		return
	}
	setSelectedPaneStyle(p)
	p.changeToSelectState()
	fmt.Println(p.state)
}
func (i *initializedPaneState) UnselectStyleSet() {}

// Helper function to set selected state styles for CourseListPane
func setSelectedPaneStyle(p *CourseListPane) {
	availableWidth := p.Width()

	p.activeDelegate.Styles.SelectedTitle = p.activeDelegate.Styles.SelectedTitle.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Width(availableWidth).
		MarginTop(1)
	p.activeDelegate.Styles.SelectedDesc = p.activeDelegate.Styles.SelectedDesc.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		Width(availableWidth).
		MarginBottom(1)
	p.Model.SetDelegate(*p.activeDelegate)
}