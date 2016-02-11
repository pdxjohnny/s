package variables

const (
	// APIPathLoginAuthServer is the path for a auth to login
	APIPathLoginAuthServer = "/auth/login"
	// APIPathLoginAuth is the for a auth to login
	APIPathLoginAuth = "/api" + APIPathLoginAuthServer
	// APIPathRefreshAuthServer is the path for a auth to refresh JWT
	APIPathRefreshAuthServer = "/auth/refresh"
	// APIPathRefreshAuth is the for a auth to refresh JWT
	APIPathRefreshAuth = "/api" + APIPathRefreshAuthServer
	// APIPathRegisterAuthServer is the path for a auth to register
	APIPathRegisterAuthServer = "/auth/register"
	// APIPathRegisterAuth is the for a auth to register
	APIPathRegisterAuth = "/api" + APIPathRegisterAuthServer
)
