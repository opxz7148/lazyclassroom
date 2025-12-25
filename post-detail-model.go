package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	contentTitleStyle    = TitleStyle.Padding(0, 1).Margin(1, 0, 0, 1)
	contentSubtitleStyle = SubtitleStyle.Margin(0, 0, 1, 1)
	contentInfo          = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return contentTitleStyle.BorderStyle(b)
	}()
	contentBodyStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)
)

type PostInfo interface {
	GetAuthor() string
	GetTitle() string
	GetContent() string
}

type PostDetailModel struct {
	postInfo        PostInfo
	height          int
	width           int
	contentViewPort viewport.Model
}

func NewPostModalModel() *PostDetailModel {
	return &PostDetailModel{
		postInfo: nil,
	}
}

func (pm *PostDetailModel) SetSize(width, height int) {
	pm.width, pm.height = width, height

	headerHeight := lipgloss.Height(pm.headerView())
	verticalMarginHeight := headerHeight + BorderOffset

	// Account for side borders (2 chars for left/right borders + padding)
	borderWidth := contentBodyStyle.GetHorizontalPadding() + contentBodyStyle.GetHorizontalBorderSize()

	pm.contentViewPort = viewport.New(width-borderWidth, height-verticalMarginHeight)
	pm.contentViewPort.YPosition = headerHeight
}

func (pm *PostDetailModel) SetPostInfo(postInfo PostInfo) {
	pm.postInfo = postInfo
}

// ============================================
// Implements tea.Model interface
// ============================================

func (pm *PostDetailModel) Init() tea.Cmd { return nil }

func (pm *PostDetailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return pm, nil
}

func (pm *PostDetailModel) headerView() string {

	if pm.postInfo == nil {
		return contentTitleStyle.Render("No post selected.")
	}

	title := contentTitleStyle.Render(pm.postInfo.GetTitle())
	author := contentSubtitleStyle.Render(fmt.Sprintf("%s", pm.postInfo.GetAuthor()))

	return title + " " + author
}

func (pm *PostDetailModel) View() string {
	if pm.postInfo == nil {
		return "No post selected."
	}

	wrappedContent := ContentWrappingStyle(pm.contentViewPort.Width - contentBodyStyle.GetHorizontalFrameSize()).Render(pm.postInfo.GetContent())

	pm.contentViewPort.SetContent(wrappedContent)

	header := pm.headerView()

	// Wrap viewport content with side borders
	content := contentBodyStyle.Render(pm.contentViewPort.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content)
}
