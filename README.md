# lazyclassroom
TUI application for inspecting Google Classroom from terminal
# LazyClassroom

A friendly terminal UI for lazily browsing your classroom content. It presents course lists and post tabs (Announcements, Materials, Coursework) with a clean, keyboard-first interface powered by Bubble Tea. ✨

**What It Does**
- **Course List:** Browse available courses in a left pane.
- **Tabbed Posts:** View per-course tabs for **Announcements**, **Materials**, and **Coursework**.
- **Attachments:** Preview and download materials or attachments from posts to your machine. 📥
- **Keyboard Controls:** Navigate panes and tabs quickly; lightweight, snappy TUI.
- **Adaptive Styling:** Uses Lipgloss for borders, colors, and layout for a cohesive look.

**Libraries Used**
- **Bubble Tea:** event loop, model/update/view architecture.
- **Bubbles (list):** high-level list component and delegate styling.
- **Lipgloss:** terminal styling, borders, alignment, colors.
- Go standard library.

The data layer is abstracted via a `ClassroomSource` interface, designed to be backed by the Google Classroom API (or mocks for local dev).

**How To Run**
Requirements: Go 1.25+

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

Enjoy a calmer way to skim class updates right from your terminal. 🧘‍♂️
