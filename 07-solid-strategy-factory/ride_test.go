package solidstrategyfactory_test

import (
	"testing"
	"time"

	solidstrategyfactory "github.com/rntdoretto/ride-calculator/07-solid-strategy-factory"
	"github.com/stretchr/testify/assert"
)

type (
	input struct {
		distance float64
		datetime time.Time
	}
	testTable struct {
		name         string
		input        []input
		expectedFare float64
		expectedErr  string
	}
)

func TestExecute(t *testing.T) {
	testTable := []testTable{
		{
			name: "Should calculate the fare for a minimal ride",
			input: []input{
				{
					distance: 2.0,
					datetime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 10.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 21.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day with two segments",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 42.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day overnight",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 17, 23, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 39.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 29.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday overnight",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(2023, 6, 18, 23, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 50.0,
		},
		{
			name: "Should return err 'invalid distance' if distance is invalid",
			input: []input{
				{
					distance: -10.0,
					datetime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedErr: "invalid distance",
		},
		{
			name: "Should return err 'invalid date' if date is invalid",
			input: []input{
				{
					distance: 10.0,
					datetime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedErr: "invalid date",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ride := solidstrategyfactory.NewRide()
			for _, i := range tt.input {
				err := ride.AddSegment(i.distance, i.datetime)
				if err != nil {
					assert.EqualError(t, err, tt.expectedErr)
					t.SkipNow()
				}
			}
			fare, err := ride.CalculateFare()
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr)
			}
			assert.Equal(t, tt.expectedFare, fare)
		})
	}
}
