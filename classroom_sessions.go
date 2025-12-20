package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	CourseList = iota
	CoursePostList
)

const (
	CourseListPaneWidthRatio = 0.25
	DetailPanePaddingOffset = 2
	DetailPaneTopOffset = 2
	SomeRandomConstantThatMakeItNotBreak = 4
)
type ClassroomSource interface {
	GetCourseList() []list.Item
}

type Selectable interface {
	Select()
	Unselect()
}

type ClassroomSessionModel struct {
	CourseList     *CourseListModel
	source         *ClassroomSource
	SelectPane     int
	loading        bool
	width          int
	height         int
	PaneList 	 map[int]SelectableModel
	initialized    bool
}

type SelectableModel interface {
	Selectable
	tea.Model
}

func NewClassroomSession(source ClassroomSource) *ClassroomSessionModel {
	// Initialize the list with empty items and default delegate
	courseList := NewCourseListModel("Courses")

	return &ClassroomSessionModel{
		CourseList: courseList,
		source:     &source,
		loading:    false,
		SelectPane: CourseList,
		PaneList: map[int]SelectableModel {
			CourseList: courseList,
		},
	}
}

func (cs *ClassroomSessionModel) RefreshCourseList() tea.Cmd {
	courseItems := (*cs.source).GetCourseList()
	return cs.CourseList.SetItems(courseItems)
}

func (cs *ClassroomSessionModel) GetSelectedCourse() *CourseItem {
	
	selectedItem := cs.CourseList.SelectedItem()
	
	if selectedItem == nil { return nil }
	if courseItem, ok := selectedItem.(*CourseItem); ok { return courseItem }
	
	return nil
}

func (cs *ClassroomSessionModel) IsLoading() bool         { return cs.loading }
func (cs *ClassroomSessionModel) SetLoading(loading bool) { cs.loading = loading }

func (cs *ClassroomSessionModel) NextPane() { cs.SelectPane = (cs.SelectPane + 1) % 2 }

func (cs *ClassroomSessionModel) Init() tea.Cmd { 
	return cs.RefreshCourseList() 
}

func (cs *ClassroomSessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cs.width, cs.height = msg.Width, msg.Height
		listWidth := int(float64(msg.Width) * CourseListPaneWidthRatio + 2)
		
		cs.CourseList.SetSize(listWidth, msg.Height-DetailPaneTopOffset)

		// Apply initial selected state styles on first window size
		cs.initialized = true

		courseItemList := cs.CourseList.Items()
		for _, item := range courseItemList {
			if courseItem, ok := item.(*CourseItem); ok && courseItem.CoursePostList != nil {
				courseItem.CoursePostList.SetSize(
					msg.Width - listWidth - SomeRandomConstantThatMakeItNotBreak - DetailPanePaddingOffset, 
					msg.Height - SomeRandomConstantThatMakeItNotBreak - DetailPaneTopOffset,
				)
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return cs, tea.Quit
		case key.Matches(msg, keys.NextTab):
			cs.NextPane()
			return cs, nil
		}
	}

	updatedModel, cmd := cs.CourseList.Update(msg)
	cs.CourseList = updatedModel.(*CourseListModel)

	selectedCourse := cs.GetSelectedCourse().CoursePostList
	cs.PaneList[CoursePostList] = selectedCourse

	return cs, cmd
}

func (cs *ClassroomSessionModel) View() string {

	for i, pane := range cs.PaneList {
		if i == cs.SelectPane {
			pane.Select()
		} else {
			pane.Unselect()
		}
	}

	listWidth := cs.CourseList.Width()

	listStyle := lipgloss.NewStyle().
		Margin(0, 2, 0, 0).
		Padding(1)

	// Calculate detail width accounting for list width, margins, padding, and border
	detailWidth := cs.width - listWidth - listStyle.GetHorizontalMargins() // 2 for border

	detailStyle := lipgloss.NewStyle().
		Padding(1).
		Width(detailWidth).
		Height(cs.height-2).
		Align(lipgloss.Left, lipgloss.Top)

	selectedCourse := cs.GetSelectedCourse()

	var detailView string
	if selectedCourse != nil && selectedCourse.CoursePostList != nil {
		detailView = selectedCourse.CoursePostList.View()
	} else {
		detailView = "No Course Selected"
	}

	// selectedCourse.CoursePostList.SetSize(detailWidth, cs.height-4) // Adjust for border and padding

	detailView = detailStyle.Render(detailView)
	courseListView := listStyle.Render(cs.CourseList.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, courseListView, detailView)
}
