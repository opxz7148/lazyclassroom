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
	source         ClassroomSource
	selectPane     int
	paneList       map[int]SelectableModel
	width          int
	height         int
	coureseLoaded  bool
	postDetail     *PostDetailModel
	showPostDetail bool
}

// getCourseListPane returns typed access to the course list pane
func (m *ClassRoomModel) getCourseListPane() *CourseListPane {
	return m.paneList[CourseListPaneID].(*CourseListPane)
}

func NewClassRoomModel(source ClassroomSource) *ClassRoomModel {
	courseListPane := NewCourseListPane("Courses")

	m := &ClassRoomModel{
		source:        source,
		selectPane:    CourseListPaneID,
		paneList:      make(map[int]SelectableModel),
		coureseLoaded: false,
		postDetail:    NewPostDetailModel(),
	}

	m.paneList[CourseListPaneID] = courseListPane

	return m
}

// SetItems delegates to the course list pane
func (m *ClassRoomModel) SetItems(items []list.Item) tea.Cmd {
	m.coureseLoaded = true
	return m.getCourseListPane().SetItems(items)
}

// Pane management
func (m *ClassRoomModel) NextPane() { m.selectPane = (m.selectPane + 1) % 2 }

// GetSelectedCourse delegates to the course list pane
func (m *ClassRoomModel) GetSelectedCourse() *CourseItem {
	return m.getCourseListPane().GetSelectedCourse()
}

// ============================================
// Implements tea.Model interface
// ============================================
func (m *ClassRoomModel) Init() tea.Cmd { return nil }

func (m *ClassRoomModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	courseListPane := m.getCourseListPane()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		listWidth := int(float64(msg.Width)*CourseListPaneWidthRatio + 2)

		courseListPane.SetSize(listWidth, msg.Height-DetailPaneTopOffset)
		m.postDetail.Update(msg)

		// Resize all CoursePostLists
		for _, item := range courseListPane.Items() {
			if courseItem, ok := item.(*CourseItem); ok && courseItem.CoursePostListModel != nil {
				courseItem.SetSize(
					msg.Width-listWidth-SomeRandomConstantThatMakeItNotBreak-DetailPanePaddingOffset,
					msg.Height-SomeRandomConstantThatMakeItNotBreak-DetailPaneTopOffset,
				)
			}
		}

	case SetTabDataMsg:

		if courseItem := m.GetSelectedCourse(); courseItem != nil && courseItem.ClassIDChecked(msg.CourseID) {
			_, cmd := courseItem.CoursePostListModel.Update(msg)
			return m, cmd
		}

		for _, course := range courseListPane.Items() {
			if courseItem, ok := course.(*CourseItem); ok && courseItem.ClassIDChecked(msg.CourseID) {
				if courseItem.CoursePostListModel != nil {
					// Route to the specific tab's list
					_, cmd := courseItem.Update(msg)
					return m, cmd
				}
				break
			}
		}
		return m, nil

	case ShowPostDetailMsg:
		m.showPostDetail = true
		m.postDetail.SetPostInfo(msg.postInfo)
		return m, nil

	case CloseDetailMsg:
		m.showPostDetail = false
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.NextTab):
			m.NextPane()
			return m, nil
		}
	}
	cmds := []tea.Cmd{}

	if m.showPostDetail {
		updatedDetail, cmd := m.postDetail.Update(msg)
		m.postDetail = updatedDetail.(*PostDetailModel)
		return m, cmd
	}

	// Update course list pane and reassign to paneList
	if pane, exists := m.paneList[m.selectPane]; exists {
		updatedPane, cmd := pane.Update(msg)
		cmds = append(cmds, cmd)
		m.paneList[m.selectPane] = updatedPane.(SelectableModel)
	}

	// Update pane list with current detail pane
	if selected := m.GetSelectedCourse(); selected != nil && selected.CoursePostListModel != nil {
		if !selected.IsFetched() {
			announcements := m.source.GetCourseAnnoucements(selected.ClassRoomId)
			materials := m.source.GetCourseMaterials(selected.ClassRoomId)
			courseWorks := m.source.GetCourseWorks(selected.ClassRoomId)
			cmds = append(cmds, selected.InsertCoursePosts(announcements, materials, courseWorks))
		}
		m.paneList[CoursePostPaneID] = selected
	}
	return m, tea.Batch(cmds...)
}

func (m *ClassRoomModel) View() string {

	if m.showPostDetail && m.postDetail != nil {
		return m.postDetail.View()
	}

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
		detailView = selectedCourse.View()
	} else {	
		detailView = "No Course Selected"
	}

	detailView = detailStyle.Render(detailView)
	courseListView := listStyle.Render(courseListPane.View())

	content := lipgloss.JoinHorizontal(lipgloss.Top, courseListView, detailView)

	return content
}
