package sample

import (
	"encoding/json"
	"net/http"
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

func InStatusList(x string) bool {
	ls := []string{"drafted", "published"}
	for _, s := range ls {
		if x == s {
			return true
		}
	}
	return false
}

func GetTomorrow(tm time.Time) time.Time {
	return tm.AddDate(0, 0, 1)
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Status string `json:"status"`
	}
	body := Body{Status: "OK"}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
