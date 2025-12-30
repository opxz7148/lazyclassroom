package main

import (
	"strings"
	"time"
)

type CoursePostItem struct {
	CreatorId    string     `json:"creatorUserId"`
	CreatorName  string     `json:"-"`
	CourseId     string     `json:"courseId"`
	Id           string     `json:"id"`
	CreationTime time.Time  `json:"creationTime"`
	UpdateTime   time.Time  `json:"updateTime"`
	Materials    []Material `json:"materials"`
}

func (cpi *CoursePostItem) Author() string    { return cpi.CreatorId }
func (cpi *CoursePostItem) ExtraInfo() string { return "" }

// AnnouncementItem represents a class announcement
// Unique field: Text (announcement content)
type AnnouncementItem struct {
	CoursePostItem
	Text string `json:"text"`
}

// ============================================
// Implements list.Item interface
// ============================================
func (ai *AnnouncementItem) FilterValue() string { return ai.Text }
func (ai *AnnouncementItem) Title() string { return strings.Split(ai.Text, "\n")[0] }
// func (ai *AnnouncementItem) Description() string { return ai.CreationTime.Format("2006-01-02") }
func (ai *AnnouncementItem) Description() string { return ai.CreatorName }
// ============================================
// Implements PostInfo interface
// ============================================
func (ai *AnnouncementItem) PostTitle() string { return "Announcement" }
func (ai *AnnouncementItem) Content() string   { return ai.Text }

// CourseWorkMaterialItem represents course materials (lectures, resources)
// Has Title, Desc, TopicId but NO grading or due dates
type CourseWorkMaterialItem struct {
	CoursePostItem
	Desc            string `json:"description"`
	CourseWorkTitle string `json:"title"`
	TopicId         string `json:"topicId"`
}

// ============================================
// Implements list.Item interface
// ============================================
func (cwmi *CourseWorkMaterialItem) FilterValue() string {
	return cwmi.CourseWorkTitle + " " + cwmi.Desc
}
func (cwmi *CourseWorkMaterialItem) Title() string       { return cwmi.CourseWorkTitle }
func (cwmi *CourseWorkMaterialItem) Description() string { return cwmi.Desc }

// ============================================
// Implements PostInfo interface
// ============================================
func (cwmi *CourseWorkMaterialItem) PostTitle() string { return cwmi.CourseWorkTitle }
func (cwmi *CourseWorkMaterialItem) Content() string   { return cwmi.Desc }

// CourseWorkItem represents assignments with grades and due dates
// Extends CourseWorkMaterialItem with dueDate, dueTime, maxPoints, workType
type CourseWorkItem struct {
	CourseWorkMaterialItem
	DueDateStruct struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"dueDate"`
	DueTimeStruct struct {
		Hours   int `json:"hours"`
		Minutes int `json:"minutes"`
	} `json:"dueTime"`
	MaxPoints                  float64 `json:"maxPoints"`
	WorkType                   string  `json:"workType"`
	SubmissionModificationMode string  `json:"submissionModificationMode"`
}

// DueDateTime parses the DueDateStruct and DueTimeStruct and converts to Thai timezone
// Returns formatted string in "2006-01-02 15:04" format in Asia/Bangkok timezone
func (cwi *CourseWorkItem) DueDateTime() string {
	// Check if due date is set
	if cwi.DueDateStruct.Year == 0 {
		return "No due date"
	}

	// Create time in UTC (Google Classroom API returns UTC)
	dueTime := time.Date(
		cwi.DueDateStruct.Year,
		time.Month(cwi.DueDateStruct.Month),
		cwi.DueDateStruct.Day,
		cwi.DueTimeStruct.Hours,
		cwi.DueTimeStruct.Minutes,
		0, // seconds
		0, // nanoseconds
		time.UTC,
	)

	// Load Thai timezone (Asia/Bangkok)
	thaiLocation, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		// Fallback to UTC+7 if location loading fails
		thaiLocation = time.FixedZone("ICT", 7*60*60)
	}

	// Convert to Thai timezone
	thaiTime := dueTime.In(thaiLocation)

	// Format as string
	return thaiTime.Format("2006-01-02 15:04")
}

// ============================================
// Implements PostInfo interface
// ============================================
func (cwi *CourseWorkItem) ExtraInfo() string { return "Due: " + cwi.DueDateTime() }
