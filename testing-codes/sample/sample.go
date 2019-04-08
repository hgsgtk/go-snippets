package sample

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func SayHello() string {
	return "hellox"
}

func GetNum(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.Wrapf(err, "GetNum failed converting %#v", str)
	}
	return num, nil
}

func GetTomorrow(tm time.Time) time.Time {
	return tm.AddDate(0, 0, 1)
}
