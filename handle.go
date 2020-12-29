package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var (
	_ http.HandlerFunc = (&statisticsHandler{}).handleFizzBuzz
	_ http.HandlerFunc = (&statisticsHandler{}).handleStatistics
)

func (sh *statisticsHandler) handleFizzBuzz(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		values := r.URL.Query()
		d1String, d2String, limitString, s1, s2 := values["int1"][0], values["int2"][0], values["limit"][0], values["s1"][0], values["s2"][0]

		d1, err := strconv.ParseInt(d1String, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("malformatted 'int1' value: %s", d1String), http.StatusBadRequest)
			return
		}

		d2, err := strconv.ParseInt(d2String, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("malformatted 'int2' value: %s", d2String), http.StatusBadRequest)
			return
		}

		limit, err := strconv.ParseInt(limitString, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("malformatted 'int1' value: %s", limitString), http.StatusBadRequest)
			return
		}

		sh.newCall(transformQuery(int(d1), int(d2), int(limit), s1, s2))
		_, err = w.Write([]byte(fizzBuzz(int(d1), int(d2), int(limit), s1, s2)))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", http.MethodPost), http.StatusMethodNotAllowed)
		return
	}
}

func (sh *statisticsHandler) handleStatistics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		d1, d2, limit, s1, s2 := getQuery(sh.most())

		_, err := w.Write([]byte(fmt.Sprintf("int1=%d ; int2=%d ; limit=%d ; s1 = %s ; s2 = %s", d1, d2, limit, s1, s2)))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", http.MethodPost), http.StatusMethodNotAllowed)
		return
	}
}