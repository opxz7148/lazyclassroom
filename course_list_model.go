package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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

func NewCourseListModel(title string) *CourseListModel {
	items := []list.Item{}
	delegate := list.NewDefaultDelegate()
	courseList := list.New(items, delegate, 0, 0)	
	courseList.Title = title
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
func (clm *CourseListModel) Select() { clm.stateMap[clm.state].SelectStyleSet() }
func (clm *CourseListModel) Unselect() { clm.stateMap[clm.state].UnselectStyleSet() }

func (clm *CourseListModel) Init() tea.Cmd { return nil}
func (clm *CourseListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	clm.Model.Title =  fmt.Sprintf("%d", clm.state)
	clm.Model, _ = clm.Model.Update(msg)
	return clm, nil
}
