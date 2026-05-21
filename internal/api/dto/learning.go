package dto

type LessonRecommendationResponse struct {
	LessonID string `json:"lessonId"`
	Reason   string `json:"reason"`
}
