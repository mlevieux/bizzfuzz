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

// handleFizzBuzz handles calls on fizzbuzz endpoint.
// It takes 5 parameters, namely int1, int2, limit, str1, and str2
// int1, int2 and limit should be integers that fit in an int32
// It responds with the fizzbuzz sequence using int1 and int2 as divisors,
// limit as upper bound, and str1 and str2 as replacers.
func (stats statsCalls) handleFizzBuzz(w http.ResponseWriter, r *http.Request) {

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

		if int32Overflow(w, d1, "int1") || int32Overflow(w, d2, "int2") || int32Overflow(w, limit, "limit") {
			return
		}

		p := transformQuery(int(d1), int(d2), int(limit), str1, str2)
		stats.add(p)

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

func int32Overflow(w http.ResponseWriter, i int64, intName string) bool {
	if int64(int32(i)) != i {
		http.Error(w, fmt.Sprintf("Overflow value ( > math.MaxInt32) for '%s': %d", intName, i), http.StatusBadRequest)
		return true
	}
	return false
}

// handleStatistics handles calls on statistics endpoint
// It responds with a pretty formatted representation of the most used parameters
// for calls to fizzbuzz endpoints.
// Additionally, it is possible to add a "top" parameter, representing an integer,
// so that the response will contain the representation of the 'top' most used sets
// of parameters in calls to fizzbuzz endpoints.
func (stats statsCalls) handleStatistics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		if len(stats.info) == 0 {
			_, err := w.Write([]byte("No call performed on the server yet"))
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		values := r.URL.Query()
		if nbString := values.Get("top"); nbString != "" {
			nb, err := strconv.ParseInt(nbString, 10, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("malformatted 'top' value: '%s'", nbString), http.StatusBadRequest)
				return
			}

			mostNStrings, mostNNumbers := stats.nMost(int(nb))
			buf := new(bytes.Buffer)
			for i, most := range mostNStrings {
				d1, d2, limit, str1, str2 := getQuery(most)
				_, err := buf.Write(formatRequestFromParams(d1, d2, limit, str1, str2, mostNNumbers[i]))
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

		most, calls := stats.most()
		d1, d2, limit, str1, str2 := getQuery(most)

		_, err := w.Write(formatRequestFromParams(d1, d2, limit, str1, str2, calls))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("unsupported method: %s", http.MethodPost), http.StatusMethodNotAllowed)
		return
	}
}

func formatRequestFromParams(d1 int, d2 int, limit int, str1 string, str2 string, nbCalls int) []byte {
	return []byte(fmt.Sprintf("int1=%d ; int2=%d ; limit=%d ; str1 = %s ; str2 = %s -- called %d times", d1, d2, limit, str1, str2, nbCalls))
}
