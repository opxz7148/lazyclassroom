package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type MockClassroomSource struct {
	MockCourses []CourseItem
}

func (mcs *MockClassroomSource) GetCourseList() []list.Item {

	courseList := make([]list.Item, len(mcs.MockCourses))
	for i := range mcs.MockCourses {
		mcs.MockCourses[i].InitializeCoursePosts()
		courseList[i] = &mcs.MockCourses[i]
	}
	
	return courseList
}

func NewMockClassroomSourceFromJSON(filename string) (*MockClassroomSource, error) {

	type courseList struct {
		Courses []CourseItem `json:"courses"`
	}

	var courses courseList

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonData), &courses)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return &MockClassroomSource{MockCourses: courses.Courses}, nil
}
