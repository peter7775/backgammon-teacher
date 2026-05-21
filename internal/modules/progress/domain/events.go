package domain

type ProgressUpdated struct { UserID string }
func (ProgressUpdated) Name() string { return "progress.updated" }
