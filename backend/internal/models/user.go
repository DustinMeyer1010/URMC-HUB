package models

import (
	"github.com/go-ldap/ldap/v3"
)

// Used for the serach page to show a breif description of the user
type UserSimpleInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	NetID    string `json:"net_id"`
	URID     string `json:"urid"`
	Email    string `json:"email"`
	OU       string `json:"ou"`
}

// Fills in UserFullInfo struct will attributes for ldap search
func (u *UserSimpleInfo) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("name")
	u.Username = user.GetAttributeValue("sAMAccountName")
	u.Email = user.GetAttributeValue("mail")
	u.NetID = user.GetAttributeValue("uid")
	u.URID = user.GetAttributeValue("URID")
	u.OU = user.GetAttributeValue("distinguishedName")
}

// Used for the user page on frontend to display all infromation about the user
type UserFullInfo struct {
	Username           string            `json:"username"`
	Email              string            `json:"email"`
	LastPasswordSet    string            `json:"last_password_set"`
	NetID              string            `json:"net_id"`
	Department         string            `json:"department"`
	Title              string            `json:"title"`
	URID               string            `json:"urid"`
	Phone              string            `json:"phone"`
	Location           string            `json:"location"`
	RelationshipStatus string            `json:"relationship_status"`
	Name               string            `json:"name"`
	OU                 string            `json:"ou"`
	Description        string            `json:"description"`
	MemberOf           []GroupSimpleInfo `json:"member_of"`
}

// Fills in UserFullInfo struct will attributes for ldap search
func (u *UserFullInfo) FillAttributes(user *ldap.Entry) []string {
	u.Name = user.GetAttributeValue("cn")
	u.Username = user.GetAttributeValue("sAMAccountName")
	u.NetID = user.GetAttributeValue("uid")
	u.URID = user.GetAttributeValue("URID")
	u.Email = user.GetAttributeValue("mail")
	u.Phone = user.GetAttributeValue("telephoneNumber")
	u.Department = user.GetAttributeValue("department")
	u.Title = user.GetAttributeValue("title")
	u.OU = user.GetAttributeValue("distinguishedName")
	u.LastPasswordSet = TimeConvert(user.GetAttributeValue("pwdLastSet"))
	u.RelationshipStatus = user.GetAttributeValue("URRoleStatus")
	u.Location = user.GetAttributeValue("physicalDeliveryOfficeName")
	u.Description = user.GetAttributeValue("description")

	return user.GetAttributeValues("memberOf")
}

type GroupModify struct {
	Users  []string `json:"users"`
	Groups []string `json:"groups"`
}
