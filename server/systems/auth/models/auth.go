package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// JwtClaims struct
type JwtClaims struct {
	ID    int32  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type AuthToken struct {
	AccessToken  string    `json:"access-token"`
	RefreshToken string    `json:"refresh-token"`
	AuthUser     string    `json:"auth-user"`
	ExpireTime   time.Time `json:"expire-time"`
	User         *User     `json:"user"`
}

type UserACL struct {
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	User        *User    `json:"user"`
}

type UserAuth struct {
	AuthToken AuthToken `json:"auth-token,omitempty"`
	UserACL   UserACL   `json:"user-acl"`
	User      *User     `json:"user"`
}

type Verify struct {
	Email     string `json:"email"`
	Subsystem string `json:"subsystem"`
	Message   []byte `json:"message"`
	Signature []byte `json:"signature"`
}

// HasPermission function check if the user has the given permission
// returns true if has the given permission
func (uc *UserACL) HasPermission(perm string) bool {

	for _, v := range uc.Permissions {
		if checkIndex(perm, v) { //TODO: (jwmdev) Generate unit testing to verify this function
			return true
		}
	}
	return false
}

// HasRole function check if user has the given role
// returns true if the user has the given role
func (uc *UserACL) HasRole(role string) bool {
	for _, v := range uc.Permissions {
		if v == role {
			return true
		}
	}
	return false
}

func checkIndex(str1, str2 string) bool {

	index := strings.LastIndex(str1, "/")

	lastString := str1[index:]
	_, err := strconv.Atoi(lastString)

	if str1 == str2 {
		return true
	} else if err != nil {
		str3 := str1[:index] + "/:id"
		if str3 == str2 {
			return true
		}
	}

	return false

}
