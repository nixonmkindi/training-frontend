package helpers

import (
	"errors"
	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/auth/models"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

const aclSuffix = "-acl"
const ResponseMsgKey = "respmsg"

type ResponseMessage struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

func Init() {
	Cache = cache.New(5*time.Minute, 15*time.Minute)
}

// Map defines a generic map of type `map[string]interface{}`.
type Map map[string]interface{}

func Serve(c echo.Context, data map[string]interface{}) Map {
	_, campusID, email := auth.GetUserFromContext(c)
	aclKey := GetACLKey(email)
	acl, isCached := GetCache(aclKey)

	if isCached && acl != nil {
		user := acl.(models.UserAuth)
		data["roles"] = user.UserACL.Roles
		data["permissions"] = user.UserACL.Permissions
		data["name"] = user.User.Name
		data["email"] = email
		data["campus_id"] = campusID
		data[errorMessageKey] = GetErrorMessage(c)
		data[infoMessageKey] = GetInfoMessage(c)
		//data["csrf"] = c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)

	} else {
		data[errorMessageKey] = GetErrorMessage(c)
		data[infoMessageKey] = GetInfoMessage(c)
		//data["csrf"] = c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	}

	//Get ResponseMessage
	e, r := GetCache(ResponseMsgKey)
	if r && e != nil {
		errMsg := e.(*ResponseMessage)
		data[ResponseMsgKey] = errMsg
	}
	data["path"] = c.Request().URL.String()
	return data
}

func StoreCache(key string, value interface{}) error {
	if Cache == nil {
		Cache = cache.New(30*time.Minute, 90*time.Minute)
	}
	_, ok := Cache.Get(key)
	if ok {
		return nil
	}
	return Cache.Add(key, value, 30*time.Minute)
}

func GetCache(key string) (interface{}, bool) {
	return Cache.Get(key)
}

func ClearCache(key string) {
	Cache.Set(key, nil, 1*time.Second)

}

func GetACLKey(email string) string {
	return email + aclSuffix
}

func SetResponseMessage(isError bool, message ...string) error {
	err := &ResponseMessage{
		Error:   isError,
		Message: message,
	}
	return Cache.Add(ResponseMsgKey, err, 500*time.Millisecond)
}

func HasPermission(email, res string) (bool, error) {

	key := GetACLKey(email)
	acl, ok := GetCache(key)

	if !ok {
		return false, errors.New("error getting cached permission")
	}
	userAuth := acl.(models.UserAuth)
	p := userAuth.UserACL.HasPermission(res)
	return p, nil
}
