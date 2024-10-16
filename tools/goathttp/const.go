package goathttp

const (
	_contentTypeHeader   = "Content-Type"
	_contentTypeJSON     = "application/json"
	_authorizationHeader = "Authorization"

	_accessControlAllowOriginHeader       = "Access-Control-Allow-Origin"
	_accessControlAllowMethodsHeader      = "Access-Control-Allow-Methods"
	_accessControlAllowHeaders            = "Access-Control-Allow-Headers"
	_accessControlAllowsCredentialsHeader = "Access-Control-Allow-Credentials"

	_allowedOrigins = "*" //TODO: как перейдем на сервер поменять разрешенные хосты
	_allowedMethods = "GET, POST, PUT, DELETE, OPTIONS"
	_allowedHeaders = "Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,Keep-Alive,X-Requested-With,If-Modified-Since,x-referer"
)

func AuthorizationHeader() string {
	return _authorizationHeader
}
