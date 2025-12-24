# Lazyclassroom
TUI application for inspecting Google Classroom from terminal

Frustrated with Google Classroom’s sluggish, complex UI? To download a material, you often need to open the browser and click through multiple pages to find what you need. lazyclassroom can help you

lazyclassroom is a friendly terminal UI for lazily browsing your classroom content inspired by [lazygit](https://github.com/jesseduffield/lazygit). It presents course lists and post tabs (Announcements, Materials, Coursework) with a clean, keyboard-first interface powered by Bubble Tea.

**What It Does**
- **Course List:** Browse available courses in a left pane.
- **Tabbed Posts:** View per-course tabs for **Announcements**, **Materials**, and **Coursework**.
- **Attachments:** Preview and download materials or attachments from posts to your machine. 
- **Keyboard Controls:** Navigate panes and tabs quickly.

**Libraries Used**
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea):** event loop, model/update/view architecture.
- **[Bubbles (list)](https://github.com/charmbracelet/bubbles):** high-level list component and delegate styling.
- **[Lipgloss](https://github.com/charmbracelet/lipgloss):** terminal styling, borders, alignment, colors.
- Go standard library.

The data layer is abstracted via a `ClassroomSource` interface, designed to be backed by the Google Classroom API (or mocks for local dev).

**How To Run**
Requirements: Go 1.25+

Download dependencies
```bash
go mod tidy
```

Run program directly
```bash
go run .
```

Optional:
- Build and run the binary

```bash
go build -o lazyclassroom
./lazyclassroom
```

**Limitations**
- **No file uploads or submissions:** Google does not allow attaching files or submitting coursework via the Classroom API, so this app focuses on viewing and downloading only.
- Features depend on the backing `ClassroomSource`; using a mock source will limit available data.
