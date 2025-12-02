package model

type Profile struct {
	Name          string `json:"name"`
	StudentID     string `json:"studentId"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Department    string `json:"department"`
	Grade         string `json:"grade"`
	EnrollmentYear string `json:"enrollmentYear"`
	Birthday      string `json:"birthday"`
}

