package models

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

// Entry struct of contact
type Entry struct {
	ID                         string            `json:"id,omitempty"`
	FullName                   string            `json:"fullName,omitempty"`
	Phones                     map[string]string `json:"phones,omitempty"`
	Emails                     map[string]string `json:"emails"`
	Department                 string            `json:"department,omitempty"`
	Title                      string            `json:"title,omitempty"`
	EmployeeID                 string            `json:"employeeId,omitempty"`
	PhysicalDeliveryOfficeName string            `json:"physicalDeliveryOfficeName,omitempty"`
	ObjectClass                string            `json:"objectClass,omitempty"`
	Photo                      string            `json:"photo,omitempty"`
}

// FirstName returns the first name of contact
func (e Entry) FirstName() string {
	return strings.Split(e.FullName, " ")[0]
}

// LastName returns the last name of contact
func (e Entry) LastName() string {
	names := strings.Split(e.FullName, " ")
	return names[len(names)-1]
}

// AttachPhoto loads a phto from URL and attach to Entry struct
func (e *Entry) AttachPhoto(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		e.Photo = base64.StdEncoding.EncodeToString(body)
	}
	return nil
}

// GetByAccount loads data using account name
func (e *Entry) GetByAccount(account string) error {
	for _, entry := range []Entry(contacts) {
		if entry.ID == account {
			*e = entry
			return nil
		}
	}
	return errors.New("not found")
}

// LoadFromLDAPEntry convert LDAP type to this Entry struct
func (e *Entry) LoadFromLDAPEntry(entry *ldap.Entry) {
	e.ID = entry.GetAttributeValue("sAMAccountName")
	e.FullName = entry.GetAttributeValue("displayName")
	e.Department = entry.GetAttributeValue("department")
	e.Title = entry.GetAttributeValue("title")
	e.EmployeeID = entry.GetAttributeValue("employeeID")
	e.PhysicalDeliveryOfficeName = entry.GetAttributeValue("physicalDeliveryOfficeName")
	e.ObjectClass = entry.GetAttributeValues("objectClass")[3]
	e.Emails = make(map[string]string)
	for _, emailField := range strings.Split(os.Getenv("EMAIL_FIELDS"), ",") {
		if entry.GetAttributeValue(emailField) != "" {
			e.Emails[emailField] = entry.GetAttributeValue(emailField)
		}
	}
	e.Phones = make(map[string]string)
	for _, f := range strings.Split(os.Getenv("PHONE_FIELDS"), ",") {
		if entry.GetAttributeValue(f) != "" {
			e.Phones[f] = entry.GetAttributeValue(f)
		}
	}
}
