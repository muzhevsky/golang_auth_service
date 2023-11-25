package errorHandling

import (
	"fmt"
)

func LogError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
