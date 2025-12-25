package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// CourseListPane wraps list.Model and implements SelectableModel
// It handles the visual state (selected/unselected) via state pattern
type CourseListPane struct {
	list.Model
	activeDelegate *list.DefaultDelegate
	stateMap       map[int]ListSelectedState
	state          int
}

func NewCourseListPane(title string) *CourseListPane {
	items := []list.Item{}
	delegate := list.NewDefaultDelegate()
	courseList := list.New(items, delegate, 0, 0)
	courseList.Title = title

	pane := &CourseListPane{
		Model:          courseList,
		activeDelegate: &delegate,
		state:          0,
	}

	pane.stateMap = map[int]ListSelectedState{
		1:  &selectedPaneState{pane: pane},
		0:  &initializedPaneState{pane: pane},
		-1: &unselectedPaneState{pane: pane},
	}

	return pane
}

// State management
func (p *CourseListPane) ToggleState()         { p.state *= -1 }
func (p *CourseListPane) changeToSelectState() { p.state = 1 }

// ============================================
// Implements Selectable interface
// ============================================
func (p *CourseListPane) Select()   { p.stateMap[p.state].SelectStyleSet() }
func (p *CourseListPane) Unselect() { p.stateMap[p.state].UnselectStyleSet() }

// GetSelectedCourse returns the currently selected CourseItem
func (p *CourseListPane) GetSelectedCourse() *CourseItem {
	selectedItem := p.SelectedItem()
	if selectedItem == nil {
		return nil
	}
	if courseItem, ok := selectedItem.(*CourseItem); ok {
		return courseItem
	}
	return nil
}

// ============================================
// Implements tea.Model interface
// ============================================
func (p *CourseListPane) Init() tea.Cmd { return nil }

func (p *CourseListPane) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.Model, cmd = p.Model.Update(msg)
	return p, cmd
}

func (p *CourseListPane) View() string {
	return p.Model.View()
}
