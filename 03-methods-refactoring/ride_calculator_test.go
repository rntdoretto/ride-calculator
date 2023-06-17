package methodsrefactoring_test

import (
	"testing"
	"time"

	methodsrefactoring "github.com/rntdoretto/ride-calculator/03-methods-refactoring"
	"github.com/stretchr/testify/assert"
)

type (
	input struct {
		distance float64
		dateTime time.Time
	}
	testTable struct {
		name     string
		input    []input
		expected float64
	}
)

func TestExecute(t *testing.T) {
	testTable := []testTable{
		{
			name: "Should calculate the fare for a minimal ride",
			input: []input{
				{
					distance: 2.0,
					dateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 10.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 21.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day with two segments",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 42.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day overnight",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 17, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 39.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 29.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday overnight",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(2023, 6, 18, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 50.0,
		},
		{
			name: "Should return -1 if distance is invalid",
			input: []input{
				{
					distance: -10.0,
					dateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: -1,
		},
		{
			name: "Should return -2 if date is invalid",
			input: []input{
				{
					distance: 10.0,
					dateTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: -2,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			segments := []methodsrefactoring.Segment{}
			for _, i := range tt.input {
				segment := methodsrefactoring.Segment{i.distance, i.dateTime}
				segments = append(segments, segment)
			}
			fare := methodsrefactoring.CalculateRide(segments)
			assert.Equal(t, tt.expected, fare)
		})
	}
}
