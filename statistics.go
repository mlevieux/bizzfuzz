package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// statsCalls wraps data and behaviours needed to process
// mostly used parameters in calls to fizzbuzz endpoint
type statsCalls map[string]int

// most return the most used parameters it has been presented
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

// nMost returns the n most used strings it has been presented
// which is to say the n most used sets of parameters used in calls to
// fizzbuzz endpoint
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

// transformQuery takes a set of parameters representing a call to
// fizzbuzz endpoint and transforms it to an easily parsable string
func transformQuery(d1, d2, limit int, str1, str2 string) string {
	return fmt.Sprintf("%d|%d|%d|%s|%s", d1, d2, limit, str1, str2)
}

// getQuery takes a 'transformQuery'-formatted string and returns
// the original set of parameters it represent.
// Both functions work as a dependant couple to easily store call
// parameters in a map[string]int.
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