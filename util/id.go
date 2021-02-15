package id

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUuid() string {
    return strings.Replace(uuid.New().String(), "-", "", -1)
}