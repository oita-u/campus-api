package model

import "time"

type Course struct {
	ID             int64     `db:"id" json:"id"`
	CourseName     string    `db:"course_name" json:"courseName"`
	InstructorName string    `db:"instructor_name" json:"instructorName"`
	CourseType     string    `db:"course_type" json:"courseType"`
	DayOfWeek      string    `db:"day_of_week" json:"dayOfWeek"`
	Period         int       `db:"period" json:"period"`
	AcademicYear   int       `db:"academic_year" json:"academicYear"`
	Semester       string    `db:"semester" json:"semester"`
	CourseCode     string    `db:"course_code" json:"courseCode"`
	Credits        int       `db:"credits" json:"credits"`
	DepartmentID   *int64    `db:"department_id" json:"departmentId,omitempty"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
}
