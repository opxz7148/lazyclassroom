package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mockSource, err := NewMockClassroomSourceFromJSON("mockCourse.json")
	if err != nil {
		panic(err)
	}

	classroomSession := NewClassroomSession(mockSource)

	p := tea.NewProgram(classroomSession, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
