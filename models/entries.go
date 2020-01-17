package models

import (
	"sort"
	"sync"
	"unicode"

	"github.com/hitalos/bina/services/ldap"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var mutex sync.Mutex

// Entries slice of Entry
type Entries []Entry

func (e Entries) Len() int { return len([]Entry(e)) }

func (e Entries) Less(i, j int) bool {
	utf8ToASCII := func(s string) string {
		t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
			return unicode.Is(unicode.Mn, r)
		}), norm.NFC)
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
func GetContacts() (Entries, error) {
	result, err := ldap.GetContacts()
	if err != nil {
		return nil, err
	}
	var entry Entry
	mutex.Lock()
	contacts = make(Entries, len(result.Entries))
	for i, e := range result.Entries {
		entry.LoadFromLDAPEntry(e)
		contacts[i] = entry
	}
	sort.Sort(&contacts)
	mutex.Unlock()
	return contacts, nil
}
