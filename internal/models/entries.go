package models

import (
	"sort"
	"sync"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/hitalos/bina/internal/config"
	"github.com/hitalos/bina/internal/services/ldap"
)

var mutex sync.Mutex

// Entries slice of Entry
type Entries []Entry

func (e Entries) Len() int { return len([]Entry(e)) }

func (e Entries) Less(i, j int) bool {
	utf8ToASCII := func(s string) string {
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		result, _, _ := transform.String(t, s)
		return result
	}
	return utf8ToASCII([]Entry(e)[i].FullName) < utf8ToASCII([]Entry(e)[j].FullName)
}

func (e *Entries) Swap(i, j int) {
	list := []Entry(*e)
	list[i], list[j] = list[j], list[i]
	*e = Entries(list)
}

var contacts Entries

// GetContacts returns a list of all Entry
func GetContacts(providers []config.Provider) (Entries, error) {
	mutex.Lock()
	contacts = Entries{}
	e := Entry{}
	for _, provider := range providers {
		entries, err := ldap.GetContacts(provider)
		if err != nil {
			return nil, err
		}
		for _, ldapEntry := range entries {
			e.LoadFromLDAPEntry(ldapEntry, provider)
			contacts = append(contacts, e)
		}
	}
	sort.Sort(&contacts)
	mutex.Unlock()

	return contacts, nil
}
