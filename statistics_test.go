package main

import "testing"

type mostTestSet struct {
	calls []string
	most string
}

func TestStatisticsHandler_Most(t *testing.T) {

	sets := []mostTestSet{
		{
			calls: []string{
				"1",
				"1",
				"1",
				"1",
				"1",
				"2",
				"2",
				"2",
				"2",
				"2",
				"2",
				"3",
			},
			most: "2",
		},
	}

	for _, set := range sets {
		sh := make(statsCalls)
		for _, call := range set.calls {
			sh[call]++
		}

		most := sh.most()
		if most != set.most {
			t.Logf("most should be '%s' but is '%s'\n", set.most, most)
			t.Fail()
		}
	}
}

type nMostTestSet struct {
	calls []string
	most []string
	n int
}

func TestStatisticsHandler_NMost(t *testing.T) {

	sets := []nMostTestSet{
		{
			calls: []string{
				"1",
				"1",
				"1",
				"1",
				"1",
				"2",
				"2",
				"2",
				"2",
				"2",
				"2",
				"3",
			},
			most: []string{
				"2",
				"1",
			},
			n: 2,
		},
	}

	for _, set := range sets {
		sh := make(statsCalls)
		for _, call := range set.calls {
			sh[call]++
		}

		most := sh.nMost(set.n)
		for i, smost := range most {
			if smost != set.most[i] {
				t.Logf("most should be '%s' but is '%s'\n", set.most, most)
				t.Fail()
			}
		}
	}
}

type transformQueryTestSet struct {
	d1 int
	d2 int
	limit int
	str1 string
	str2 string

	result string
}

func TestTransformQuery(t *testing.T) {
	sets := transformationTestSets()

	for _, set := range sets {
		result := transformQuery(set.d1, set.d2, set.limit, set.str1, set.str2)
		if result != set.result {
			t.Logf("result should be '%s' with (%d, %d, %d, %s, %s) but is '%s'\n", set.result, set.d1, set.d2, set.limit, set.str1, set.str2, result)
			t.Fail()
		}
	}
}

func transformationTestSets() []transformQueryTestSet {
	return []transformQueryTestSet{
		{
			d1:     4,
			d2:     6,
			limit:  19,
			str1:     "fizz",
			str2:     "buzz",
			result: "4|6|19|fizz|buzz",
		},
	}
}

func TestGetQuery(t *testing.T) {
	sets := transformationTestSets()

	for _, set := range sets {
		d1, d2, limit, str1, str2 := getQuery(set.result)

		if d1 != set.d1 {
			t.Logf("d1 should be %d but is %d with '%s'\n", set.d1, d1, set.result)
			t.Fail()
		}

		if d2 != set.d2 {
			t.Logf("d2 should be %d but is %d with '%s'\n", set.d2, d2, set.result)
			t.Fail()
		}

		if limit != set.limit {
			t.Logf("limit should be %d but is %d with '%s'\n", set.limit, limit, set.result)
			t.Fail()
		}


		if str1 != set.str1 {
			t.Logf("str1 should be '%s' but is '%s' with '%s'\n", set.str1, str1, set.result)
			t.Fail()
		}

		if str2 != set.str2 {
			t.Logf("str2 should be '%s' but is '%s' with '%s'\n", set.str2, str2, set.result)
			t.Fail()
		}
	}
}