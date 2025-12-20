package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Announcement = iota
	Material
	CourseWork
)

type CoursePostListModel struct {
	PostList     []list.Model
	SelectedList int
	Tabs         []string
	width        int
	height       int
	style	   	lipgloss.Style
	activeTabStyle   lipgloss.Style
	inactiveTabStyle lipgloss.Style
	windowStyle      lipgloss.Style
}

func NewCoursePostListModel() *CoursePostListModel {

	tabsList := []string{"Announcements", "Materials", "Course Works"}

	inactiveTabStyle := lipgloss.
			NewStyle().
			Border(inactiveTabBorder, true).
			BorderForeground(DetailUnSelectedColor).
			Padding(0, 1)

	return &CoursePostListModel{
		PostList: []list.Model{
			NewPlainTabListModel(tabsList[Announcement]),
			NewPlainTabListModel(tabsList[Material]),
			NewPlainTabListModel(tabsList[CourseWork]),
		},
		SelectedList: Announcement,
		Tabs:         tabsList,
		inactiveTabStyle:  inactiveTabStyle,
		activeTabStyle:    inactiveTabStyle.
			Border(activeTabBorder, true),
		windowStyle:       lipgloss.NewStyle().
			BorderForeground(DetailUnSelectedColor).
			Padding(2).
			Align(lipgloss.Center, lipgloss.Center).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop(),
	}
}

func (cplm *CoursePostListModel) NextTab() { cplm.SelectedList = (cplm.SelectedList + 1) % 3 }
func (cplm *CoursePostListModel) PrevTab() { cplm.SelectedList = (cplm.SelectedList - 1) % 3 }

func (cplm *CoursePostListModel) SetBorderColor(color lipgloss.TerminalColor) {
	cplm.inactiveTabStyle = cplm.inactiveTabStyle.BorderForeground(color)
	cplm.activeTabStyle = cplm.activeTabStyle.BorderForeground(color)
	cplm.windowStyle = cplm.windowStyle.BorderForeground(color)
}
func (cplm *CoursePostListModel) Select() { cplm.SetBorderColor(DetailSelectedColor) }
func (cplm *CoursePostListModel) Unselect() { cplm.SetBorderColor(DetailUnSelectedColor) }

func (cplm *CoursePostListModel) SetSize(width, height int) { cplm.width, cplm.height = width, height }

func (cplm *CoursePostListModel) SetTabData(tabIndex int, items []list.Item) {
	if tabIndex < 0 || tabIndex >= len(cplm.PostList) { return }
	cplm.PostList[tabIndex].SetItems(items)
}

func (cplm *CoursePostListModel) Init() tea.Cmd { return nil }

func (cplm *CoursePostListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.NextTab):
			cplm.NextTab()
			return cplm, nil
		case key.Matches(msg, keys.PrevTab):
			cplm.PrevTab()
			return cplm, nil
		}
	}

	cplm.PostList[cplm.SelectedList], _ = cplm.PostList[cplm.SelectedList].Update(msg)
	return cplm, nil
}

func (cplm *CoursePostListModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	// Calculate width per tab accounting for borders
	tabCount := len(cplm.Tabs)
	totalBorderWidth := tabCount * BorderOffset
	availableWidth := cplm.width
	tabContentWidth := (availableWidth - totalBorderWidth) / tabCount

	for i, t := range cplm.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(cplm.Tabs)-1, i == cplm.SelectedList
		if isActive {
			style = cplm.activeTabStyle
		} else {
			style = cplm.inactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border).Width(tabContentWidth).Align(lipgloss.Center)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row + "\n")

	contentWidth := lipgloss.Width(row) - BorderOffset

	doc.WriteString(
		cplm.windowStyle.
			Width(contentWidth).
			Height(cplm.height).
			Render(cplm.PostList[cplm.SelectedList].View()),
	)
	return doc.String()
}

func NewPlainTabListModel(title string) list.Model {
	delegate := list.NewDefaultDelegate()
	plainList := list.New([]list.Item{}, delegate, 0, 0)
	plainList.Title = title
	plainList.SetShowStatusBar(false)
	plainList.SetShowHelp(false)
	plainList.SetShowTitle(false)
	return plainList
}
