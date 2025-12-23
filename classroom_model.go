package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Pane identifiers - use PaneID suffix to avoid conflict with CourseListPane type
const (
	CourseListPaneID = iota
	CoursePostPaneID
)

// Layout constants
const (
	CourseListPaneWidthRatio             = 0.25
	DetailPanePaddingOffset              = 2
	DetailPaneTopOffset                  = 2
	SomeRandomConstantThatMakeItNotBreak = 4
)

type ListSelectedState interface {
	SelectStyleSet()
	UnselectStyleSet()
}

type Selectable interface {
	Select()
	Unselect()
}

type SelectableModel interface {
	Selectable
	tea.Model
}

// ClassRoomModel orchestrates the course list pane and detail pane
type ClassRoomModel struct {
	// Pane management
	source    ClassroomSource
	selectPane int
	paneList   map[int]SelectableModel
	width      int
	height     int
}

// getCourseListPane returns typed access to the course list pane
func (m *ClassRoomModel) getCourseListPane() *CourseListPane {
	return m.paneList[CourseListPaneID].(*CourseListPane)
}

func NewClassRoomModel(source ClassroomSource) *ClassRoomModel {
	courseListPane := NewCourseListPane("Courses")

	m := &ClassRoomModel{
		source:    source,
		selectPane: CourseListPaneID,
		paneList:   make(map[int]SelectableModel),
	}

	m.paneList[CourseListPaneID] = courseListPane

	return m
}

// SetItems delegates to the course list pane
func (m *ClassRoomModel) SetItems(items []list.Item) tea.Cmd {
	return m.getCourseListPane().SetItems(items)
}

// Pane management
func (m *ClassRoomModel) NextPane() { m.selectPane = (m.selectPane + 1) % 2 }

// GetSelectedCourse delegates to the course list pane
func (m *ClassRoomModel) GetSelectedCourse() *CourseItem {
	return m.getCourseListPane().GetSelectedCourse()
}

func (m *ClassRoomModel) Init() tea.Cmd { return nil }

func (m *ClassRoomModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	courseListPane := m.getCourseListPane()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		listWidth := int(float64(msg.Width)*CourseListPaneWidthRatio + 2)

		courseListPane.SetSize(listWidth, msg.Height-DetailPaneTopOffset)

		// Resize all CoursePostLists
		for _, item := range courseListPane.Items() {
			if courseItem, ok := item.(*CourseItem); ok && courseItem.CoursePostListModel != nil {
				courseItem.SetSize(
					msg.Width-listWidth-SomeRandomConstantThatMakeItNotBreak-DetailPanePaddingOffset,
					msg.Height-SomeRandomConstantThatMakeItNotBreak-DetailPaneTopOffset,
				)
			}
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.NextTab):
			m.NextPane()
			return m, nil
		}
	}

	// Update course list pane and reassign to paneList
	updatedPane, cmd := m.paneList[m.selectPane].Update(msg)
	m.paneList[m.selectPane] = updatedPane.(SelectableModel)
	
	// Update pane list with current detail pane
	if selected := m.GetSelectedCourse(); selected != nil && selected.CoursePostListModel != nil {
		m.paneList[CoursePostPaneID] = selected
	}
	return m, cmd
}

func (m *ClassRoomModel) View() string {
	// Apply select/unselect styles to panes
	for i, pane := range m.paneList {
		if i == m.selectPane {
			pane.Select()
		} else {
			pane.Unselect()
		}
	}

	courseListPane := m.getCourseListPane()
	listWidth := courseListPane.Width()

	listStyle := lipgloss.NewStyle().
		Margin(0, 2, 0, 0).
		Padding(1)

	detailWidth := m.width - listWidth - listStyle.GetHorizontalMargins()

	detailStyle := lipgloss.NewStyle().
		Padding(1).
		Width(detailWidth).
		Height(m.height-2).
		Align(lipgloss.Left, lipgloss.Top)

	selectedCourse := m.GetSelectedCourse()

	var detailView string
	if selectedCourse != nil && selectedCourse.CoursePostListModel != nil {
		detailView = selectedCourse.CoursePostListModel.View()
	} else {
		detailView = "No Course Selected"
	}

	detailView = detailStyle.Render(detailView)
	courseListView := listStyle.Render(courseListPane.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, courseListView, detailView)
}
