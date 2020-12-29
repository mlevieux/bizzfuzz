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
		sh := newStatistics()
		for _, call := range set.calls {
			sh.newCall(call)
		}

		most := sh.most()
		if most != set.most {
			t.Logf("most should be '%s' but is '%s'\n", set.most, most)
			t.Fail()
		}
	}
}

type transformQueryTestSet struct {
	d1 int
	d2 int
	limit int
	s1 string
	s2 string

	result string
}

func TestTransformQuery(t *testing.T) {
	sets := transformationTestSets()

	for _, set := range sets {
		result := transformQuery(set.d1, set.d2, set.limit, set.s1, set.s2)
		if result != set.result {
			t.Logf("result should be '%s' with (%d, %d, %d, %s, %s) but is '%s'\n", set.result, set.d1, set.d2, set.limit, set.s1, set.s2, result)
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
			s1:     "fizz",
			s2:     "buzz",
			result: "4|6|19|fizz|buzz",
		},
	}
}

func TestGetQuery(t *testing.T) {
	sets := transformationTestSets()

	for _, set := range sets {
		d1, d2, limit, s1, s2 := getQuery(set.result)

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


		if s1 != set.s1 {
			t.Logf("s1 should be '%s' but is '%s' with '%s'\n", set.s1, s1, set.result)
			t.Fail()
		}

		if s2 != set.s2 {
			t.Logf("s2 should be '%s' but is '%s' with '%s'\n", set.s2, s2, set.result)
			t.Fail()
		}
	}
}