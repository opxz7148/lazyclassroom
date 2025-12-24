package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type MockClassroomSource struct {
	MockCourses []CourseItem
	MockAnnouncements []AnnouncementItem
	MockMaterials []CourseWorkMaterialItem
	MockCourseWorks []CourseWorkItem
}

func (mcs *MockClassroomSource) GetCourseList() []list.Item {

	courseList := make([]list.Item, len(mcs.MockCourses))
	for i := range mcs.MockCourses {
		mcs.MockCourses[i].InitializeCoursePosts()
		courseList[i] = &mcs.MockCourses[i]
	}
	
	return courseList
}

func (mcs *MockClassroomSource) GetCourseAnnoucements(courseId string) []list.Item {
	announcementList := make([]list.Item, len(mcs.MockAnnouncements))
	for i := range mcs.MockAnnouncements {
		announcementList[i] = &mcs.MockAnnouncements[i]
	}
	return announcementList
}

func (mcs *MockClassroomSource) GetCourseMaterials(courseId string) []list.Item {
	materialList := make([]list.Item, len(mcs.MockMaterials))
	for i := range mcs.MockMaterials {
		materialList[i] = &mcs.MockMaterials[i]
	}
	return materialList
}

func (mcs *MockClassroomSource) GetCourseWorks(courseId string) []list.Item {
	courseWorkList := make([]list.Item, len(mcs.MockCourseWorks))
	for i := range mcs.MockCourseWorks {
		courseWorkList[i] = &mcs.MockCourseWorks[i]
	}
	return courseWorkList
}

func NewMockClassroomSourceFromJSON(filename string) (*MockClassroomSource, error) {

	type courseList struct {
		Courses []CourseItem `json:"courses"`
		Annoucement []AnnouncementItem `json:"announcements"`
		Materials [] CourseWorkMaterialItem `json:"courseWorkMaterial"`
		CourseWorks []CourseWorkItem `json:"courseWork"`
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
	return &MockClassroomSource{
		MockCourses: courses.Courses,
		MockAnnouncements: courses.Annoucement,
		MockMaterials: courses.Materials,
		MockCourseWorks: courses.CourseWorks,
	}, nil
}
