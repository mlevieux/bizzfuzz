package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

var (
	_ http.HandlerFunc = (&statsCalls{}).handleFizzBuzz
	_ http.HandlerFunc = (&statsCalls{}).handleStatistics
)

func (a statsCalls) handleFizzBuzz(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		values := r.URL.Query()
		d1String := values.Get("int1")
		d2String := values.Get("int2")
		limitString := values.Get("limit")
		str1 := values.Get("str1")
		str2 := values.Get("str2")

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
			http.Error(w, fmt.Sprintf("malformatted 'limit' value: %s", limitString), http.StatusBadRequest)
			return
		}

		p := transformQuery(int(d1), int(d2), int(limit), str1, str2)
		a[p]++

		_, err = w.Write(fizzBuzz(int(d1), int(d2), int(limit), str1, str2))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", http.MethodPost), http.StatusMethodNotAllowed)
		return
	}
}

func (a statsCalls) handleStatistics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		values := r.URL.Query()
		if nbString := values.Get("top") ; nbString != "" {
			nb, err := strconv.ParseInt(nbString, 10, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("malformatted 'top' value: '%s'", nbString), http.StatusBadRequest)
				return
			}

			mostN := a.nMost(int(nb))
			buf := new(bytes.Buffer)
			for _, most := range mostN {
				d1, d2, limit, str1, str2 := getQuery(most)
				_, err := buf.Write(formatRequestFromParams(d1, d2, limit, str1, str2))
				if err != nil {
					http.Error(w, "internal server error", http.StatusInternalServerError)
					return
				}

				buf.Write([]byte{'\n'})
			}

			_, err = w.Write(buf.Bytes())
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		d1, d2, limit, str1, str2 := getQuery(a.most())

		_, err := w.Write(formatRequestFromParams(d1, d2, limit, str1, str2))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", http.MethodPost), http.StatusMethodNotAllowed)
		return
	}
}

func formatRequestFromParams(d1 int, d2 int, limit int, str1 string, str2 string) []byte {
	return []byte(fmt.Sprintf("int1=%d ; int2=%d ; limit=%d ; str1 = %s ; str2 = %s", d1, d2, limit, str1, str2))
}
