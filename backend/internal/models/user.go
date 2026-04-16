package models

import (
	"github.com/go-ldap/ldap/v3"
)

type UserDetails struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

func (u *UserDetails) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("name")
	u.Username = user.GetAttributeValue("sAMAccountName")
	u.Email = user.GetAttributeValue("mail")
	u.Department = user.GetAttributeValue("department")
}

type UserModifyRequest struct {
	GroupDN string `json:"group"`
}

type UserModifyResults struct {
	GroupDN      string `json:"group"`
	Successful   bool   `json:"successful"`
	ErrorMessage string `json:"errorMessage"`
}
