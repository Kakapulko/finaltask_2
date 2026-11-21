package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("неверный формат строки")
	}

	num, err := strconv.Atoi(data[0])
	if err != nil {
		return err
	}
	if num <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	dur, err := time.ParseDuration(data[2])
	if err != nil {
		return err
	}
	if dur <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}

	t.Steps = num
	t.TrainingType = data[1]
	t.Duration = dur

	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	result += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	result += fmt.Sprintf("Дистанция: %.2f км.\n", dist)
	result += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	result += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)

	return result, nil
}
