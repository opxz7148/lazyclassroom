package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Setup debug logging
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Always restore terminal on exit
	defer func() {
		if r := recover(); r != nil {
			// Run reset command to fully restore terminal
			cmd := exec.Command("reset")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run() // Ignore errors

			// Print panic info after terminal is restored
			fmt.Fprintf(os.Stderr, "\n=== PANIC ===\n")
			fmt.Fprintf(os.Stderr, "Error: %v\n\n", r)
			fmt.Fprintf(os.Stderr, "Stack Trace:\n%s\n", debug.Stack())

			os.Exit(1)
		}
	}()

	mockSource, err := NewMockClassroomSourceFromJSON("mockCourse.json")
	if err != nil {
		panic(err)
	}

	classroomSession := NewClassroomSession(mockSource)

	log.Println("🚀 Starting application...")
	p := tea.NewProgram(classroomSession, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
