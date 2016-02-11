package api

import (
	"strings"

	"github.com/pdxjohnny/s/variables"
)

// GetAccount retrives an account
func GetAccount(host, token, id string) (*map[string]interface{}, error) {
	path := variables.APIPathAccount
	path = strings.Replace(path, ":id", id, 1)
	return GenericRequest(host, path, token, nil)
}

// SaveAccount saves an account
func SaveAccount(host, token, id string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathAccount
	path = strings.Replace(path, ":id", id, 1)
	doc["_id"] = id
	return GenericRequest(host, path, token, doc)
}
