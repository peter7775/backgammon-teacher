package domain

type LessonRecommended struct { UserID string }
func (LessonRecommended) Name() string { return "learning.lesson_recommended" }
