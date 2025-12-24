package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type CourseItem struct {
	Name           string `json:"name"`
	Section        string `json:"section"`
	ClassRoomId    string `json:"id"`
	fetched 	 bool		   `json:"-"`
	*CoursePostListModel
}

func (ci *CourseItem) FilterValue() string { return ci.Name }
func (ci *CourseItem) Title() string       { return ci.Name }
func (ci *CourseItem) Description() string { return ci.Section }
func (ci *CourseItem) InitializeCoursePosts() {
	ci.CoursePostListModel = NewCoursePostListModel(ci.ClassRoomId)
	ci.fetched = false
}

func (ci *CourseItem) IsFetched() bool { return ci.fetched }
func (ci *CourseItem) ClassIDChecked(id string) bool { return ci.ClassRoomId == id }

func (ci *CourseItem) InsertCoursePosts (
	announcements []list.Item,
	materials []list.Item,
	courseWorks []list.Item,
) tea.Cmd {
	cmds := []tea.Cmd{}
	cmds = append(cmds, ci.SetTabData(AnnouncementTab, announcements))
	cmds = append(cmds, ci.SetTabData(MaterialTab, materials))
	cmds = append(cmds, ci.SetTabData(CourseWorkTab, courseWorks))
	ci.fetched = true
	return tea.Batch(cmds...)
}

func (ci *CourseItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := ci.CoursePostListModel.Update(msg)
	ci.CoursePostListModel = updatedModel.(*CoursePostListModel)
	return ci, cmd
}
