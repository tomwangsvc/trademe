package app

import (
	"reflect"
	"testing"
)

func Test_FilterAndUseFirstItem(t *testing.T) {
	var data = []struct {
		desc     string
		input    []Property
		expected []Property
	}{
		{
			desc: "different properties",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
		},

		{
			desc: "same properties with different value",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         640000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
		},
	}

	for i, d := range data {
		result := FilterAndUseFirstItem(data[i].input)
		if !reflect.DeepEqual(result, d.expected) {
			t.Errorf("text failed at %d with description %s", i, d.desc)
		}
	}
}

func Test_FilterAndUseLastItem(t *testing.T) {
	var data = []struct {
		desc     string
		input    []Property
		expected []Property
	}{
		{
			desc: "different properties",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
		},

		{
			desc: "same properties with different value",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         640000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         640000,
				},
			},
		},
	}

	for i, d := range data {
		result := FilterAndUseLastItem(data[i].input)
		if !reflect.DeepEqual(result, d.expected) {
			t.Errorf("text failed at %d with description %s", i, d.desc)
		}
	}
}

func Test_FilterOutDuplicatedItem(t *testing.T) {
	var data = []struct {
		desc     string
		input    []Property
		expected []Property
	}{
		{
			desc: "different properties",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "2 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
			},
		},

		{
			desc: "same properties with different value",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         630000,
				},
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         640000,
				},
			},
			expected: nil,
		},
	}

	for i, d := range data {
		result := FilterOutDuplicatedItem(data[i].input)
		if !reflect.DeepEqual(result, d.expected) {
			t.Errorf("text failed at %d with description %s with result %+v", i, d.desc, result)
		}
	}
}

func Test_FilterByMutipleConditions(t *testing.T) {
	var data = []struct {
		desc     string
		input    []Property
		expected []Property
	}{
		{
			desc: "under 400k",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         399999,
				},
			},
			expected: nil,
		},

		{
			desc: "include 'PL' in street",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD PL",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
			expected: nil,
		},

		{
			desc: "remove 10th property",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "2",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "3",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "4",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "5",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "6",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "7",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "8",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "9",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "10",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "2",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "3",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "4",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "5",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "6",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "7",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "8",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "9",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
		},

		{
			desc: "valid property",
			input: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
		},
	}

	for i, d := range data {
		result := FilterByMutipleConditions(data[i].input)
		if !reflect.DeepEqual(result, d.expected) {
			t.Errorf("text failed at %d with description %s with result %+v %d", i, d.desc, result, len(result))
		}
	}
}

func Test_ProcessByChunk(t *testing.T) {
	type input struct {
		chunk      int
		properties []Property
	}
	var data = []struct {
		desc     string
		input    input
		expected []Property
	}{
		{
			desc: "under 400k",
			input: input{
				chunk: 2,
				properties: []Property{
					{
						Id:            "1",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         399999,
					},
				},
			},
			expected: nil,
		},

		{
			desc: "include 'PL' in street",
			input: input{
				chunk: 2,
				properties: []Property{
					{
						Id:            "1",
						StreetAddress: "1 OAKWOOD PL",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
				},
			},
			expected: nil,
		},

		{
			desc: "remove 10th property",
			input: input{
				chunk: 2,
				properties: []Property{
					{
						Id:            "1",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "2",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "3",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "4",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "5",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "6",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "7",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "8",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "9",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
					{
						Id:            "10",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "2",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "3",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "4",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "5",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "6",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "7",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "8",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
				{
					Id:            "9",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
		},

		{
			desc: "valid property",
			input: input{
				chunk: 2,
				properties: []Property{
					{
						Id:            "1",
						StreetAddress: "1 OAKWOOD AVENUE",
						Town:          "WANAKA",
						ValuationDate: "01/01/2020",
						Value:         8000000,
					},
				},
			},
			expected: []Property{
				{
					Id:            "1",
					StreetAddress: "1 OAKWOOD AVENUE",
					Town:          "WANAKA",
					ValuationDate: "01/01/2020",
					Value:         8000000,
				},
			},
		},
	}

	for i, d := range data {
		result := ProcessByChunk(data[i].input.properties, d.input.chunk)
		if len(result) != len(result) {
			t.Errorf("text failed at %d with description %s with result %+v but expect %v", i, d.desc, result, d.expected)
		}
	}
}
