package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type CourseItem struct {
	Name    string `json:"name"`
	Section string `json:"section"`
}

func (ci CourseItem) FilterValue() string { return ci.Name }
func (ci CourseItem) Title() string       { return ci.Name }
func (ci CourseItem) Description() string { return ci.Section }

type ClassroomSource interface {
	GetCourseList() []list.Item
}

type ClassroomSessionModel struct {
	CourseList list.Model
	source     *ClassroomSource
	loading    bool
}

func NewClassroomSession(source ClassroomSource) *ClassroomSessionModel {
	// Initialize the list with empty items and default delegate
	items := []list.Item{}
	delegate := list.NewDefaultDelegate()
	courseList := list.New(items, delegate, 0, 0)
	courseList.Title = "Courses"

	return &ClassroomSessionModel{
		CourseList: courseList,
		source:     &source,
		loading:    false,
	}
}

func (cs *ClassroomSessionModel) RefreshCourseList() tea.Cmd {
	courseItems := (*cs.source).GetCourseList()
	return cs.CourseList.SetItems(courseItems)
}

func (cs *ClassroomSessionModel) IsLoading() bool         { return cs.loading }
func (cs *ClassroomSessionModel) SetLoading(loading bool) { cs.loading = loading }

func (cs *ClassroomSessionModel) Init() tea.Cmd {
	return cs.RefreshCourseList()
}

func (cs *ClassroomSessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cs.CourseList.SetWidth(msg.Width)
		cs.CourseList.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return cs, tea.Quit
		}
	}

	var cmd tea.Cmd
	cs.CourseList, cmd = cs.CourseList.Update(msg)
	return cs, cmd
}

func (cs *ClassroomSessionModel) View() string {
	return cs.CourseList.View()
}
