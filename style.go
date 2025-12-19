package main

import "github.com/charmbracelet/lipgloss"

const (
	BorderOffset = 2 // Each tab has left and right border (1 char each)
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	DetailSelectedColor   = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	DetailUnSelectedColor = lipgloss.AdaptiveColor{Light: "#000000ff", Dark: "#ffffffff"}

	SelectedListItemColor  = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	SelectedListItemSubcolor = lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}
)
