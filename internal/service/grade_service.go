package service

import (
	"math"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type GradeService struct {
	repo *repository.GradeRepository
}

func NewGradeService(r *repository.GradeRepository) *GradeService {
	return &GradeService{repo: r}
}

// GradeResponse は成績APIのレスポンス用構造体
type GradeResponse struct {
	Year          int                  `json:"year"`
	Semester      string               `json:"semester"`
	TotalCredits  int                  `json:"totalCredits"`
	GPA           float64              `json:"gpa"`
	Items         []model.GradeItem    `json:"items"`
}

// GetGrades はユーザーの成績一覧とGPA等を返す
func (s *GradeService) GetGrades(userID string, year int, semester string) (*GradeResponse, error) {
	year, semester = resolveYearSemester(year, semester)

	items, err := s.repo.GetGradesByUserID(userID, year, semester)
	if err != nil {
		return nil, err
	}

	totalCredits := 0
	var totalPoints float64

	for _, item := range items {
		if item.Credits <= 0 {
			continue
		}
		if item.Grade == nil {
			continue
		}

		point := gradeToPoint(*item.Grade)
		if point < 0 {
			// 評価に対応するポイントがない場合は計算から除外
			continue
		}

		totalCredits += item.Credits
		totalPoints += float64(item.Credits) * point
	}

	var gpa float64
	if totalCredits > 0 {
		gpa = totalPoints / float64(totalCredits)
		// 小数第3位で四捨五入（例: 3.175 -> 3.18）
		gpa = math.Round(gpa*100) / 100
	}

	return &GradeResponse{
		Year:         year,
		Semester:     semester,
		TotalCredits: totalCredits,
		GPA:          gpa,
		Items:        items,
	}, nil
}

// gradeToPoint は成績記号をGPAポイントに変換する
// 例: S=4.0, A=3.0, B=2.0, C=1.0, D/F=0.0
func gradeToPoint(grade string) float64 {
	switch grade {
	case "S":
		return 4.0
	case "A":
		return 3.0
	case "B":
		return 2.0
	case "C":
		return 1.0
	case "D", "F":
		return 0.0
	default:
		return -1.0
	}
}


