package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, str := range dataset {
		err := dp.Parse(str)
		if err != nil {
			log.Println("ошибка парсинга:", err)
			continue
		}

		text, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка:", err)
			continue
		}

		fmt.Println(text)
	}
}
