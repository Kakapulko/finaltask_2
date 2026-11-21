package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func check(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return errors.New("вес должен быть больше нуля")
	}
	if height <= 0 {
		return errors.New("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := check(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)
	mins := duration.Minutes()

	result := (weight * speed * mins) / minInH
	result = result * walkingCaloriesCoefficient

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := check(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)
	mins := duration.Minutes()

	result := (weight * speed * mins) / minInH

	return result, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}

	dist := Distance(steps, height)
	speed := dist / duration.Hours()

	return speed
}

func Distance(steps int, height float64) float64 {
	step := height * stepLengthCoefficient
	dist := float64(steps) * step / mInKm

	return dist
}
