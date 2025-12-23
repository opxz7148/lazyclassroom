package main

import "time"

type CoursePostMaterial struct {
	CoursePostId string `json:"coursePostId"`
	Id           string `json:"id"`
	Link         string `json:"link"`
	Title        string `json:"title"`
}

type CoursePostItem struct {
	CreatorId    string               `json:"creatorUserId"`
	CreatorName  string               `json:"-"`
	CourseId     string               `json:"courseId"`
	Id           string               `json:"id"`
	CreationTime time.Time            `json:"creationTime"`
	UpdateTime   time.Time            `json:"updateTime"`
	Material     []CoursePostMaterial `json:"materials"`
}

// AnnouncementItem represents a class announcement
// Unique field: Text (announcement content)
type AnnouncementItem struct {
	CoursePostItem
	Text string `json:"text"`
}

func (ai *AnnouncementItem) FilterValue() string {return ai.Text}
func (ai *AnnouncementItem) Title() string       {return ai.Text}


// CourseWorkMaterialItem represents course materials (lectures, resources)
// Has Title, Desc, TopicId but NO grading or due dates
type CourseWorkMaterialItem struct {
	CoursePostItem
	Desc    string `json:"description"`
	CourseWorkTitle   string `json:"title"`
	TopicId string `json:"topicId"`
}

func (cwmi *CourseWorkMaterialItem) FilterValue() string {
	return cwmi.CourseWorkTitle + " " + cwmi.Desc
}
func (cwmi *CourseWorkMaterialItem) Title() string { return cwmi.CourseWorkTitle }
func (cwmi *CourseWorkMaterialItem) Description() string { return cwmi.Desc }

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
