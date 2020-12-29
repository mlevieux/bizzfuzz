package main

import (
	"fmt"
	"strconv"
	"strings"
)

type statisticsHandler struct {
	calls map[string]int
}

func newStatistics() *statisticsHandler {
	return &statisticsHandler{calls: make(map[string]int)}
}

func (s *statisticsHandler) newCall(params string) {
	s.calls[params]++
}

func (s *statisticsHandler) most() string {

	var (
		bestParams string
		bestNbCalls int
	)

	for params, nbCalls := range s.calls {
		if nbCalls > bestNbCalls {
			bestParams = params
			bestNbCalls = nbCalls
		}
	}

	return bestParams
}


func transformQuery(d1, d2, limit int, s1, s2 string) string {
	return fmt.Sprintf("%d|%d|%d|%s|%s", d1, d2, limit, s1, s2)
}

func getQuery(params string) (int, int, int, string, string) {
	var i int

	for i = 0 ; params[i] != '|' ; i++ {}
	d1, _ := strconv.Atoi(params[:i])
	params = params[i+1:]

	for i = 0 ; params[i] != '|' ; i++ {}
	d2, _ := strconv.Atoi(params[:i])
	params = params[i+1:]

	for i = 0 ; params[i] != '|' ; i++ {}
	limit, _ := strconv.Atoi(params[:i])
	params = params[i+1:]

	i = strings.IndexByte(params, '|')
	s1, s2 := params[:i], params[i+1:]

	return d1, d2, limit, s1, s2
}