package domain

import analysis "backgammon-teacher/internal/modules/analysis/domain"

type TeachingIntent struct {
	Themes []analysis.Theme
	Focus  string
}
