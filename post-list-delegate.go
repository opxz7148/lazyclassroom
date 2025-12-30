package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ShowPostDetailMsg struct {
	postInfo PostInfo
}

func postListUpdateFunc(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Enter):
			selectedPost := m.SelectedItem()
			if postInfo, ok := selectedPost.(PostInfo); ok {
				return func() tea.Msg { return ShowPostDetailMsg{postInfo: postInfo} }
			}
		case key.Matches(msg, keys.Back):
			// Ignore Back/ESC here — do nothing and preserve original behavior
			return nil
		}
	}
	return nil
}

func newPostListDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.UpdateFunc = postListUpdateFunc
	return delegate
}
