package models

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-ldap/ldap/v3"

	"github.com/hitalos/bina/config"
)

// Entry struct of contact
type Entry struct {
	ID          string            `json:"id,omitempty"`
	FullName    string            `json:"fullName,omitempty"`
	ObjectClass string            `json:"objectClass,omitempty"`
	Phones      map[string]string `json:"phones,omitempty"`
	Emails      map[string]string `json:"emails,omitempty"`
	Others      map[string]string `json:"others,omitempty"`
	Photo       string            `json:"photo,omitempty"`
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
		body, err := io.ReadAll(resp.Body)
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

func makeMap(e *ldap.Entry, fields []string) map[string]string {
	list := make(map[string]string, len(fields))
	for _, field := range fields {
		if e.GetAttributeValue(field) != "" {
			list[field] = e.GetAttributeValue(field)
		}
	}
	return list
}

// LoadFromLDAPEntry convert LDAP type to this Entry struct
func (e *Entry) LoadFromLDAPEntry(entry *ldap.Entry, p config.Provider) {
	e.ID = entry.GetAttributeValue("sAMAccountName")
	e.FullName = entry.GetAttributeValue("displayName")
	e.ObjectClass = entry.GetAttributeValues("objectClass")[3]
	e.Emails = makeMap(entry, p.Fields.Emails)
	e.Phones = makeMap(entry, p.Fields.Phones)
	e.Others = makeMap(entry, p.Fields.Others)
}
