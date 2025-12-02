package model

type TimetableItem struct {
	CourseID       int64  `db:"course_id" json:"courseId"`
	CourseName     string `db:"course_name" json:"courseName"`
	InstructorName string `db:"instructor_name" json:"instructorName"`
	CourseType     string `db:"course_type" json:"courseType"`
	DayOfWeek      string `db:"day_of_week" json:"dayOfWeek"`
	Period         int    `db:"period" json:"period"`
	Status         string `db:"status" json:"status"`
}
