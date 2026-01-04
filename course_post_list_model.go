package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	AnnouncementTab = iota
	MaterialTab
	CourseWorkTab
)

type CoursePostListModel struct {
	PostTabList      []list.Model
	CourseID         string
	SelectedTab      int
	Tabs             []string
	width            int
	height           int
	postFetched      bool
	source           ClassroomSource
	style            lipgloss.Style
	activeTabStyle   lipgloss.Style
	inactiveTabStyle lipgloss.Style
	windowStyle      lipgloss.Style
}

func NewCoursePostListModel(CourseID string, source ClassroomSource) *CoursePostListModel {

	tabsList := []string{"Announcements", "Materials", "Course Works"}

	inactiveTabStyle := lipgloss.
		NewStyle().
		Border(inactiveTabBorder, true).
		BorderForeground(DetailUnSelectedColor).
		Padding(0, 1)

	return &CoursePostListModel{
		PostTabList: []list.Model{
			NewPostListModel(tabsList[AnnouncementTab]),
			NewPostListModel(tabsList[MaterialTab]),
			NewPostListModel(tabsList[CourseWorkTab]),
		},
		SelectedTab:      AnnouncementTab,
		Tabs:             tabsList,
		CourseID:         CourseID,
		source:           source,
		postFetched:      false,
		inactiveTabStyle: inactiveTabStyle,
		activeTabStyle: inactiveTabStyle.
			Border(activeTabBorder, true),
		windowStyle: lipgloss.NewStyle().
			BorderForeground(DetailUnSelectedColor).
			Padding(1).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop(),
	}
}

func (cplm *CoursePostListModel) NextTab() { cplm.SelectedTab = (cplm.SelectedTab + 1) % 3 }
func (cplm *CoursePostListModel) PrevTab() { cplm.SelectedTab = (cplm.SelectedTab + 2) % 3 }

func (cplm *CoursePostListModel) SetBorderColor(color lipgloss.TerminalColor) {
	cplm.inactiveTabStyle = cplm.inactiveTabStyle.BorderForeground(color)
	cplm.activeTabStyle = cplm.activeTabStyle.BorderForeground(color)
	cplm.windowStyle = cplm.windowStyle.BorderForeground(color)
}

type RouteAblePostListMsg interface{ GetCourseID() string }

type PostListModelMsgCommon struct {
	CourseID string
	TabIndex int
}

func (plm PostListModelMsgCommon) GetCourseID() string { return plm.CourseID }

type SetTabDataMsg struct {
	PostListModelMsgCommon
	UpdatedItem []list.Item
}

type AfterSetTabDataMsg struct {
	PostListModelMsgCommon
	OriginalMsg tea.Msg
}

func (cplm *CoursePostListModel) FetchPostData() tea.Cmd {
	if cplm.postFetched {
		return nil
	}

	announcementsCmd := func() tea.Msg {
		items := cplm.source.GetCourseAnnoucements(cplm.CourseID)
		return SetTabDataMsg{
			PostListModelMsgCommon: PostListModelMsgCommon{
				CourseID: cplm.CourseID,
				TabIndex: AnnouncementTab,
			},
			UpdatedItem: items,
		}
	}
	materialsCmd := func() tea.Msg {
		items := cplm.source.GetCourseMaterials(cplm.CourseID)
		return SetTabDataMsg{
			PostListModelMsgCommon: PostListModelMsgCommon{
				CourseID: cplm.CourseID,
				TabIndex: MaterialTab,
			},
			UpdatedItem: items,
		}
	}
	courseWorksCmd := func() tea.Msg {
		items := cplm.source.GetCourseWorks(cplm.CourseID)
		return SetTabDataMsg{
			PostListModelMsgCommon: PostListModelMsgCommon{
				CourseID: cplm.CourseID,
				TabIndex: CourseWorkTab,
			},
			UpdatedItem: items,
		}
	}

	cplm.postFetched = true
	return tea.Batch(announcementsCmd, materialsCmd, courseWorksCmd)
}

// ============================================
// Implements Selectable interface
// ============================================
func (cplm *CoursePostListModel) Select()   { cplm.SetBorderColor(DetailSelectedColor) }
func (cplm *CoursePostListModel) Unselect() { cplm.SetBorderColor(DetailUnSelectedColor) }

func (cplm *CoursePostListModel) SetSize(width, height int) {
	cplm.width, cplm.height = width, height

	// Account for: tab bar (~3 lines) + window style padding (2*2) + borders (2)
	// tabBarHeight := 3
	windowPadding := cplm.windowStyle.GetVerticalPadding() + cplm.windowStyle.GetVerticalBorderSize()
	listHeight := height - windowPadding
	listWidth := width - cplm.windowStyle.GetHorizontalPadding() - cplm.windowStyle.GetHorizontalBorderSize()

	// Update size for all tab lists and trigger their update to re-render
	for i := range cplm.PostTabList {
		cplm.PostTabList[i].SetSize(listWidth-4, listHeight)
		// Force an update cycle by passing a window size message
		cplm.PostTabList[i], _ = cplm.PostTabList[i].Update(tea.WindowSizeMsg{Width: listWidth - 4, Height: listHeight})
	}
}

func (cplm *CoursePostListModel) SetTabData(tabIndex int, items []list.Item) tea.Cmd {
	if tabIndex < 0 || tabIndex >= len(cplm.PostTabList) {
		return nil
	}
	cmd := cplm.PostTabList[tabIndex].SetItems(items)

	// If no cmd returned (filtering not active), no need to wrap
	if cmd == nil {
		return nil
	}

	return func() tea.Msg {
		msg := cmd()
		return AfterSetTabDataMsg{
			OriginalMsg: msg,
			PostListModelMsgCommon: PostListModelMsgCommon{
				CourseID: cplm.CourseID,
				TabIndex: tabIndex,
			},
		}
	}
}

// ============================================
// Implements tea.Model interface
// ============================================
func (cplm *CoursePostListModel) Init() tea.Cmd { return cplm.FetchPostData() }

func (cplm *CoursePostListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SetTabDataMsg:
		var cmd tea.Cmd
		cmd = cplm.SetTabData(msg.TabIndex, msg.UpdatedItem)
		return cplm, cmd
	case AfterSetTabDataMsg:
		var cmd tea.Cmd
		cplm.PostTabList[msg.TabIndex], cmd = cplm.PostTabList[msg.TabIndex].Update(msg.OriginalMsg)
		return cplm, cmd
	case tea.KeyMsg:

		if cplm.PostTabList[cplm.SelectedTab].FilterState() == list.Filtering {
			// Delegate to the active tab's list filtering
			updatedList, cmd := cplm.PostTabList[cplm.SelectedTab].Update(msg)
			cplm.PostTabList[cplm.SelectedTab] = updatedList
			return cplm, cmd
		}

		switch {
		case key.Matches(msg, keys.Right):
			cplm.NextTab()
			return cplm, nil
		case key.Matches(msg, keys.Left):
			cplm.PrevTab()
			return cplm, nil
		}
	}

	var cmd tea.Cmd
	cplm.PostTabList[cplm.SelectedTab], cmd = cplm.PostTabList[cplm.SelectedTab].Update(msg)
	return cplm, cmd
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
		isFirst, isLast, isActive := i == 0, i == len(cplm.Tabs)-1, i == cplm.SelectedTab
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
			MaxHeight(cplm.height).
			Render(cplm.PostTabList[cplm.SelectedTab].View()),
	)
	return doc.String()
}

func NewPostListModel(title string) list.Model {
	delegate := newPostListDelegate()
	plainList := list.New([]list.Item{}, delegate, 0, 0)
	plainList.Title = title
	plainList.SetShowTitle(false)
	plainList.SetShowPagination(false)
	return plainList
}
