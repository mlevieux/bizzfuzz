package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type statsCalls map[string]int

func (s statsCalls) most() string {

	var (
		mostParams string
		mostNbCalls int
	)

	for params, nbCalls := range s {
		if nbCalls > mostNbCalls {
			mostParams = params
			mostNbCalls = nbCalls
		}
	}

	return mostParams
}

func (s statsCalls) nMost(n int) []string {

	type callInfo struct {
		params string
		nbCalls int
	}

	toSort := make([]callInfo, len(s))
	i := 0
	for params, nbCalls := range s {
		toSort[i] = callInfo{
			params:  params,
			nbCalls: nbCalls,
		}
		i++
	}

	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i].nbCalls > toSort[j].nbCalls
	})

	result := make([]string, n)
	for i := 0 ; i < n && i < len(toSort); i++ {
		result[i] = toSort[i].params
	}
	return result
}

func transformQuery(d1, d2, limit int, str1, str2 string) string {
	return fmt.Sprintf("%d|%d|%d|%s|%s", d1, d2, limit, str1, str2)
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
	str1, str2 := params[:i], params[i+1:]

	return d1, d2, limit, str1, str2
}