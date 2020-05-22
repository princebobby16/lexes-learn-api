package Godep

import (
	"strings"
)

func GetSchemaName(email string) string {
	index := strings.Index(email, ".")
	indexOfAt := strings.Index(email, "@")

	schema_namespace := email[indexOfAt + 1:index]

	//log.Println(email[indexOfAt + 1:index])
	return schema_namespace
}
