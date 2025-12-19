package main

type CourseItem struct {
	Name           string `json:"name"`
	Section        string `json:"section"`
	ClassRoomId    string `json:"id"`
	CoursePostList *CoursePostListModel
}

func (ci *CourseItem) FilterValue() string { return ci.Name }
func (ci *CourseItem) Title() string       { return ci.Name }
func (ci *CourseItem) Description() string { return ci.Section }
func (ci *CourseItem) InitializeCoursePosts() {
	ci.CoursePostList = NewCoursePostListModel()
}