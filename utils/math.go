package utils

import "github.com/google/uuid"

func Add(a, b int) int {
    return a + b
}

// GenerateUUID 生成一个新的随机 UUID (v4)
func GenerateUUID() string {
    return uuid.New().String()
}