package utils

import guuid "github.com/google/uuid"

func GetUUID() string {
	id := guuid.New()
	return id.String()
}
