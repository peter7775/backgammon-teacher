package domain

type Step struct {
	From int8
	To   int8
	Pips int8
}

type Move struct {
	Steps []Step
}
