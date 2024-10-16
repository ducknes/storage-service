package storagecontext

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"storage-service/settings"
	"storage-service/tools/goathttp"

	"github.com/GOAT-prod/goatlogger"
)

const (
	_authTest = "bruhmagedon"
)

type StorageContext struct {
	ctx    context.Context
	logger *goatlogger.Logger
	auth   goathttp.Auth
}

func New(r *http.Request) StorageContext {
	logger := goatlogger.New(settings.GetAppName())

	return StorageContext{
		ctx:    r.Context(),
		logger: &logger,
		auth:   parseAuth(r.Header.Get(goathttp.AuthorizationHeader())),
	}
}

func (sc *StorageContext) SetCtx(ctx context.Context) {
	sc.ctx = ctx
}

func (sc *StorageContext) Ctx() context.Context {
	return sc.ctx
}

func (sc *StorageContext) Log() *goatlogger.Logger {
	return sc.logger
}

func (sc *StorageContext) SetLogTag(tag string) {
	sc.logger.SetTag(tag)
}

func parseAuth(token string) (auth goathttp.Auth) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}

	if err = json.Unmarshal(decodedToken, &auth); err != nil {
		return
	}

	return
}

func (sc *StorageContext) IsAuthorized() bool {
	return sc.auth.UserName == _authTest
}
