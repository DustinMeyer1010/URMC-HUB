package models

import "github.com/go-ldap/ldap/v3"

// Used for the serach page to show a breif description of the user
type UserSimpleInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	NetID    string `json:"net_id"`
	OU       string `json:"ou"`
}

// Fills in UserFullInfo struct will attributes for ldap search
func (u *UserSimpleInfo) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("name")
	u.Username = user.GetAttributeValue("sAMAccountName")
	u.Email = user.GetAttributeValue("mail")
	u.NetID = user.GetAttributeValue("uid")
	u.OU = user.GetAttributeValue("distinguishedName")
}

// Used for the user page on frontend to display all infromation about the user
type UserFullInfo struct {
	Username           string `json:"username"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	RelationshipStatus string `json:"relationship_status"`
	Department         string `json:"department"`
	Title              string `json:"title"`
	Phone              string `json:"phone"`
	Location           string `json:"location"`
	LastPasswordSet    string `json:"last_password_set"`
	URID               string `json:"urid"`
	OU                 string `json:"ou"`
	NetID              string `json:"net_id"`
	Description        string `json:"description"`
}

// Fills in UserFullInfo struct will attributes for ldap search
func (u *UserFullInfo) FillAttributes(user *ldap.Entry) {
	u.Name = user.GetAttributeValue("cn")
	u.Username = user.GetAttributeValue("Username")
	u.NetID = user.GetAttributeValue("uid")
	u.URID = user.GetAttributeValue("URID")
	u.Email = user.GetAttributeValue("mail")
	u.Phone = user.GetAttributeValue("telephoneNumber")
	u.Department = user.GetAttributeValue("department")
	u.Title = user.GetAttributeValue("title")
	u.OU = user.GetAttributeValue("distinguishedName")
	u.LastPasswordSet = user.GetAttributeValue("pwdLastSet")
	u.RelationshipStatus = user.GetAttributeValue("URRoleStatus")
	u.Location = user.GetAttributeValue("physicalDeliveryOfficeName")
	u.Description = user.GetAttributeValue("description")
}
