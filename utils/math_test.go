package utils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		// 正数相加
		{
			name:     "positive numbers",
			a:        3,
			b:        5,
			expected: 8,
		},
		{
			name:     "large positive numbers",
			a:        1000000,
			b:        2000000,
			expected: 3000000,
		},
		{
			name:     "addition with one being zero",
			a:        42,
			b:        0,
			expected: 42,
		},

		// 负数相加
		{
			name:     "negative numbers",
			a:        -3,
			b:        -5,
			expected: -8,
		},
		{
			name:     "one positive one negative (negative result)",
			a:        3,
			b:        -10,
			expected: -7,
		},
		{
			name:     "one positive one negative (positive result)",
			a:        15,
			b:        -5,
			expected: 10,
		},

		// 零值情况
		{
			name:     "both zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "negative zero and positive zero",
			a:        -0,
			b:        0,
			expected: 0,
		},
		{
			name:     "adding to negative zero",
			a:        -0,
			b:        10,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d, expected %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
