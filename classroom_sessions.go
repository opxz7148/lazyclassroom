package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ClassroomSource interface {
	GetCourseList() []list.Item
	GetCourseAnnoucements(courseId string) []list.Item
	GetCourseMaterials(courseId string) []list.Item
	GetCourseWorks(courseId string) []list.Item
}

type ClassroomSessionModel struct {
	ClassRoom *ClassRoomModel
	source    *ClassroomSource
	loading   bool
}

func NewClassroomSession(source ClassroomSource) *ClassroomSessionModel {
	classRoom := NewClassRoomModel(source)

	return &ClassroomSessionModel{
		ClassRoom: classRoom,
		source:    &source,
		loading:   false,
	}
}

func (cs *ClassroomSessionModel) RefreshCourseList() tea.Cmd {
	courseItems := (*cs.source).GetCourseList()
	return cs.ClassRoom.SetItems(courseItems)
}

func (cs *ClassroomSessionModel) IsLoading() bool         { return cs.loading }
func (cs *ClassroomSessionModel) SetLoading(loading bool) { cs.loading = loading }

func (cs *ClassroomSessionModel) Init() tea.Cmd {
	return cs.RefreshCourseList()
}

func (cs *ClassroomSessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle app-level keys (quit, future: auth, modals)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return cs, tea.Quit
		}
	}

	// Delegate to ClassRoomModel
	updatedModel, cmd := cs.ClassRoom.Update(msg)
	cs.ClassRoom = updatedModel.(*ClassRoomModel)

	return cs, cmd
}

func (cs *ClassroomSessionModel) View() string {
	return cs.ClassRoom.View()
}
