package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type ListSelectedState interface {
	SelectStyleSet()
	UnselectStyleSet()
}

type CourseListModel struct {
	list.Model
	activeDelegate *list.DefaultDelegate
	stateMap	map[int]ListSelectedState
	state int
}

func NewCourseListModel() *CourseListModel {
	items := []list.Item{}
	delegate := list.NewDefaultDelegate()
	courseList := list.New(items, delegate, 0, 0)	
	courseList.Title = "Courses"

	m := CourseListModel{
		Model:          courseList,
		activeDelegate: &delegate,
	}
	m.stateMap = map[int]ListSelectedState{
		1: &selectedState{model: &m},
		0: &intializedState{model: &m},
		-1: &unselectedState{model: &m},
	}
	m.state = 0

	return &m
}

func (clm *CourseListModel) ToggleState() { clm.state *= -1 }
func (clm *CourseListModel) changeToSelectState() { clm.state = 1 }
func (clm *CourseListModel) Select() {
	// // Get available width (list width minus some padding for borders)

	// clm.SetWidth(clm.Width() - 1)

	// availableWidth := clm.Width()

	// clm.activeDelegate.Styles.SelectedTitle = clm.activeDelegate.Styles.SelectedTitle.
	// 	Border(lipgloss.NormalBorder(), true, true, false, true).
	// 	Width(availableWidth)
	// clm.activeDelegate.Styles.SelectedDesc = clm.activeDelegate.Styles.SelectedDesc.
	// 	Border(lipgloss.NormalBorder(), false, true, true, true).
	// 	Width(availableWidth)
	// clm.Model.SetDelegate(*clm.activeDelegate)
	clm.stateMap[clm.state].SelectStyleSet()
}

func (clm *CourseListModel) Unselect() {
	// clm.SetWidth(clm.Width() + 1)

// availableWidth := clm.Width()		

	// clm.activeDelegate.Styles.SelectedTitle = clm.activeDelegate.Styles.SelectedTitle.
	// 	Border(lipgloss.NormalBorder(), false, false, false, true).
	// 	Width(availableWidth)
	// clm.activeDelegate.Styles.SelectedDesc = clm.activeDelegate.Styles.SelectedDesc.
	// 	Border(lipgloss.NormalBorder(), false, false, false, true).
	// 	Width(availableWidth)
	// clm.Model.SetDelegate(*clm.activeDelegate)
	clm.stateMap[clm.state].UnselectStyleSet()
}
