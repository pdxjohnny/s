package auth

import (
	"github.com/pdxjohnny/easyreq"

	"github.com/pdxjohnny/s/variables"
)

// LoginAuth logs in a user
func LoginAuth(host, token string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathLoginAuth
	return easyreq.GenericRequest(host, path, token, doc)
}

// RefreshAuth gets a new toke for the user
func RefreshAuth(host, token string) (*map[string]interface{}, error) {
	path := variables.APIPathRefreshAuth
	return easyreq.GenericRequest(host, path, token, nil)
}

// RegisterAuth registers a user
func RegisterAuth(host, token string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathRegisterAuth
	return easyreq.GenericRequest(host, path, token, doc)
}
