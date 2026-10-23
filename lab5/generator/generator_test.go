package generator

import (
	"slices"
	"testing"
)

func TestGenerateSliceInt(t *testing.T) {
	expected := []int{
		23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
		43, 44, 45, 46, 47,
	}

	result := GenerateSliceInt()

	if !slices.Equal(result, expected) {
		t.Errorf("результат %v, очікувано %v", result, expected)
	}
}

type testCase struct {
	name        string
	step        float64
	expected    []float64
	expectedErr string
}

func TestGenerateSliceFloat(t *testing.T) {
	cases := []testCase{
		{
			name:        "крок 0.25",
			step:        0.25,
			expected:    []float64{0, 0.25, 0.5, 0.75, 1},
			expectedErr: "",
		},
		{
			name:        "крок 0.3",
			step:        0.3,
			expected:    []float64{0, 0.3, 0.6, 0.9},
			expectedErr: "",
		},
		{
			name:        "крок 0 (помилка)",
			step:        0,
			expected:    nil,
			expectedErr: "крок повинен бути більшим за 0",
		},
		{
			name:        "від'ємний крок (помилка)",
			step:        -0.5,
			expected:    nil,
			expectedErr: "крок повинен бути більшим за 0",
		},
		{
			name:        "крок 1 (помилка)",
			step:        1.0,
			expected:    nil,
			expectedErr: "крок повинен бути меншим за 1",
		},
		{
			name:        "крок більший за 1 (помилка)",
			step:        1.5,
			expected:    nil,
			expectedErr: "крок повинен бути меншим за 1",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GenerateSliceFloat(tc.step)

			if tc.expectedErr == "" && err != nil {
				t.Fatalf("неочікувана помилка: %v", err)
			}

			if tc.expectedErr != "" {
				if err == nil {
					t.Fatalf("очікувано помилку %q, отримано nil", tc.expectedErr)
				}
				if err.Error() != tc.expectedErr {
					t.Errorf("помилка %q, очікувано %q", err.Error(), tc.expectedErr)
				}
			}

			if !slices.Equal(result, tc.expected) {
				t.Errorf("результат %v, очікувано %v", result, tc.expected)
			}
		})
	}
}
