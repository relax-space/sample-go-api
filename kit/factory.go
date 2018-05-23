package kit

import (
	"errors"
	"time"
)

func Scholarship(mathScore, historyScore float64) (money float64, err error) {
	math, err := CreateSubject("math")
	if err != nil {
		return
	}
	history, err := CreateSubject("history")
	if err != nil {
		return
	}
	newMathScore := math.Score(mathScore, func(t float64) float64 {
		if time.Now().Year() == 2017 { //This year's exam is very simple
			return t * 0.7
		}
		return t
	})
	money = (newMathScore + history.Score(historyScore)) * 100
	return
}

type Subject interface {
	Score(score float64, options ...func(float64) float64) float64
	Weights() float64
}

type Math struct{}

func (d Math) Score(score float64, options ...func(float64) float64) float64 {
	current := d.Weights() * score
	var newScore float64
	for _, v := range options {
		if v == nil {
			continue
		}
		newScore = v(current)
	}
	return newScore
}

func (Math) Weights() float64 {
	return 2
}

type History struct{}

func (d History) Score(score float64, options ...func(float64) float64) float64 {
	current := d.Weights() * score
	return current
}
func (History) Weights() float64 {
	return 3
}

type Geography struct{}

func (d Geography) Score(score float64, options ...func(float64) float64) float64 {
	current := d.Weights() * score
	return current
}
func (Geography) Weights() float64 {
	return 2
}

func CreateSubject(subjectName string) (subject Subject, err error) {
	switch subjectName {
	case "math":
		subject = new(Math)
	case "history":
		subject = new(History)
	case "geography":
		subject = new(Geography)
	default:
		err = errors.New("subjectName is not supported")
	}
	return
}
