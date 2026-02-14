package utils

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
)

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

func TestGenerateUUID(t *testing.T) {
	t.Run("valid format", func(t *testing.T) {
		id := GenerateUUID()

		// UUID 格式：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		uuidRegex := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
		if !uuidRegex.MatchString(id) {
			t.Errorf("GenerateUUID() returned invalid format: %s", id)
		}
	})

	t.Run("can be parsed", func(t *testing.T) {
		id := GenerateUUID()

		_, err := uuid.Parse(id)
		if err != nil {
			t.Errorf("GenerateUUID() returned unparseable UUID: %s, error: %v", id, err)
		}
	})

	t.Run("uniqueness", func(t *testing.T) {
		// 生成多个 UUID 并验证它们各不相同
		uuids := make(map[string]bool)
		count := 100

		for i := 0; i < count; i++ {
			id := GenerateUUID()
			if uuids[id] {
				t.Errorf("GenerateUUID() generated duplicate UUID: %s", id)
				return
			}
			uuids[id] = true
		}

		if len(uuids) != count {
			t.Errorf("Expected %d unique UUIDs, got %d", count, len(uuids))
		}
	})

	t.Run("length", func(t *testing.T) {
		id := GenerateUUID()
		expectedLength := 36 // 32 hex chars + 4 hyphens

		if len(id) != expectedLength {
			t.Errorf("GenerateUUID() returned length %d, expected %d", len(id), expectedLength)
		}
	})

	t.Run("version 4", func(t *testing.T) {
		id := GenerateUUID()

		// v4 UUID 的第 9 个字符（索引 14）应该是 '4'
		if id[14] != '4' {
			t.Errorf("GenerateUUID() should return v4 UUID, got: %s", id)
		}
	})
}
