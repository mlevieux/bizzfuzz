package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// statsCalls wraps data and behaviours needed to process
// mostly used parameters in calls to fizzbuzz endpoint
type statsCalls struct {
	mux  *sync.Mutex
	info map[string]int
}

func (stats statsCalls) add(params string) {
	stats.mux.Lock()
	defer stats.mux.Unlock()

	stats.info[params]++
}

// most return the most used parameters it has been presented
func (stats statsCalls) most() (string, int) {

	var (
		mostParams string
		mostNbCalls int
	)

	stats.mux.Lock()
	defer stats.mux.Unlock()
	for params, nbCalls := range stats.info {
		if nbCalls > mostNbCalls {
			mostParams = params
			mostNbCalls = nbCalls
		}
	}

	return mostParams, mostNbCalls
}

// nMost returns the n most used strings it has been presented
// which is to say the n most used sets of parameters used in calls to
// fizzbuzz endpoint
func (stats statsCalls) nMost(n int) ([]string, []int) {

	type callInfo struct {
		params string
		nbCalls int
	}

	toSort := make([]callInfo, len(stats.info))
	i := 0

	stats.mux.Lock()
	for params, nbCalls := range stats.info {
		toSort[i] = callInfo{
			params:  params,
			nbCalls: nbCalls,
		}
		i++
	}
	stats.mux.Unlock()

	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i].nbCalls > toSort[j].nbCalls
	})

	resultStrings := make([]string, n)
	resultNumbers := make([]int, n)
	for i := 0 ; i < n && i < len(toSort); i++ {
		resultStrings[i] = toSort[i].params
		resultNumbers[i] = toSort[i].nbCalls
	}
	return resultStrings, resultNumbers
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