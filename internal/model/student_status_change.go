package model

import "time"

type StudentStatusChange struct {
    ID                 int        `json:"id" db:"id"`
    StudentID          string     `json:"student_id" db:"student_id"` // varchar(20)なのでstring
    ChangeType         string     `json:"change_type" db:"change_type"`
    Reason             string     `json:"reason" db:"reason"`
    ApprovalDate       *time.Time `json:"approval_date" db:"approval_date"` // Nullableな日付はポインタにする
    ScheduledStartDate *time.Time `json:"scheduled_start_date" db:"scheduled_start_date"`
    ScheduledEndDate   *time.Time `json:"scheduled_end_date" db:"scheduled_end_date"`
    ActualStartDate    *time.Time `json:"actual_start_date" db:"actual_start_date"`
    ActualEndDate      *time.Time `json:"actual_end_date" db:"actual_end_date"`
}