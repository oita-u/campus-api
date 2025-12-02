package model

// GradeItem は成績一覧の1行分を表す
type GradeItem struct {
	CourseID       int64   `db:"course_id" json:"courseId"`
	CourseCode     string  `db:"course_code" json:"courseCode"`
	CourseName     string  `db:"course_name" json:"courseName"`
	InstructorName string  `db:"instructor_name" json:"instructorName"`
	Credits        int     `db:"credits" json:"credits"`
	Semester       string  `db:"semester" json:"semester"`
	AcademicYear   int     `db:"academic_year" json:"academicYear"`
	Status         string  `db:"status" json:"status"`
	Grade          *string `db:"grade" json:"grade"` // NULL 許容
}


