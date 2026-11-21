package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("неверный формат строки")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	time1, err := time.ParseDuration(slice[1])
	if err != nil {
		return err
	}
	if time1 <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = time1

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	km := spentenergy.Distance(ds.Steps, ds.Height)

	calor, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, km, calor)

	return res, nil
}
