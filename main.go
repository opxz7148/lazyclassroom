package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Example usage of ClassroomSessionModel with MockClassroomSource
	mockSource, err := NewMockClassroomSourceFromJSON("mockCourse.json")
	if err != nil {
		panic(err)
	}

	classroomSession := NewClassroomSession(mockSource)

	p := tea.NewProgram(classroomSession, tea.WithAltScreen())
	if err, _ := p.Run(); err != nil {
		panic(err)
	}
}