package domain

type Classification string

const (
	ClassificationBest    Classification = "best"
	ClassificationInacc   Classification = "inaccuracy"
	ClassificationMistake Classification = "mistake"
	ClassificationBlunder Classification = "blunder"
)
