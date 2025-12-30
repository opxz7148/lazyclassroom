package main

import (
	"github.com/charmbracelet/bubbles/key"
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
	Author() string
	PostTitle() string
	Content() string
	ExtraInfo() string
}

type PostDetailModel struct {
	postInfo        PostInfo
	height          int
	width           int
	contentViewPort viewport.Model
}

func NewPostDetailModel() *PostDetailModel {
	return &PostDetailModel{
		postInfo: nil,
	}
}

func (pm *PostDetailModel) SetSize(width, height int) {
	pm.width, pm.height = width, height

	// Get header height, safe even if postInfo is nil
	headerHeight := lipgloss.Height(pm.headerView())
	verticalMarginHeight := headerHeight +
		contentTitleStyle.GetVerticalMargins() +
		contentSubtitleStyle.GetVerticalMargins() +
		contentTitleStyle.GetVerticalFrameSize() +
		contentSubtitleStyle.GetVerticalFrameSize()

	// Account for side borders (2 chars for left/right borders + padding)
	borderWidth := contentBodyStyle.GetHorizontalPadding() + contentBodyStyle.GetHorizontalBorderSize()

	// Ensure dimensions are positive
	viewportWidth := width - borderWidth
	viewportHeight := height - verticalMarginHeight
	if viewportWidth < 1 {
		viewportWidth = 1
	}
	if viewportHeight < 1 {
		viewportHeight = 1
	}

	pm.contentViewPort = viewport.New(viewportWidth, viewportHeight)
	pm.contentViewPort.YPosition = headerHeight
}

func (pm *PostDetailModel) SetPostInfo(postInfo PostInfo) {
	pm.postInfo = postInfo

	// Update viewport content when postInfo changes
	if pm.contentViewPort.Width > 0 && postInfo != nil {
		wrappedContent := ContentWrappingStyle(pm.contentViewPort.Width - contentBodyStyle.GetHorizontalFrameSize()).Render(postInfo.Content())
		pm.contentViewPort.SetContent(wrappedContent)
	}
}

// ============================================
// Implements tea.Model interface
// ============================================

type CloseDetailMsg struct{}

func (pm *PostDetailModel) Init() tea.Cmd { return nil }

func (pm *PostDetailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {



	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		pm.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Back):
			return pm, func() tea.Msg { return CloseDetailMsg{} }
		}
	}

	return pm, nil
}

func (pm *PostDetailModel) headerView() string {

	if pm.postInfo == nil {
		return contentTitleStyle.Render("No post selected.")
	}

	title := contentTitleStyle.Render(pm.postInfo.PostTitle())
	extraInfo := contentSubtitleStyle.Render(pm.postInfo.Author() + " | " + pm.postInfo.ExtraInfo())

	return title + " " + extraInfo
}

func (pm *PostDetailModel) View() string {
	if pm.postInfo == nil {
		return "No post selected."
	}

	wrappedContent := ContentWrappingStyle(pm.contentViewPort.Width - contentBodyStyle.GetHorizontalFrameSize()).Render(pm.postInfo.Content())

	pm.contentViewPort.SetContent(wrappedContent)

	header := pm.headerView()

	// Wrap viewport content with side borders
	content := contentBodyStyle.Render(pm.contentViewPort.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content)
}
