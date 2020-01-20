package ldap

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"

	"github.com/hitalos/bina/config"
)

func getAttrFields(p config.Provider) []string {
	fields := []string{"objectClass", p.Fields.Identifier, p.Fields.FullName}
	fields = append(fields, p.Fields.Phones...)
	fields = append(fields, p.Fields.Emails...)
	return append(fields, p.Fields.Others...)
}

// GetContacts searches contacts on LDAP server
func GetContacts(p config.Provider) ([]*ldap.Entry, error) {
	phoneFields := ""
	for _, field := range p.Fields.Phones {
		phoneFields += fmt.Sprintf("(%s=*)", field)
	}

	filter := "(&" +
		"(|" + phoneFields + ")" +
		"(objectCategory=person)" +
		"(!(UserAccountControl:1.2.840.113556.1.4.803:=2))" +
		"(|(objectClass=user)(objectClass=contact)))"

	if ldapTimeout, err := strconv.Atoi(p.Params["timeout"]); err == nil && ldapTimeout > 0 {
		ldap.DefaultTimeout = time.Duration(ldapTimeout) * time.Second
	}
	if os.Getenv("DEBUG") == "1" {
		log.Println("timeout set to:", ldap.DefaultTimeout)
	}

	ldapConn, err := ldap.DialURL("ldap://" + p.Params["host"] + ":389")
	if err != nil {
		return nil, err
	}
	defer ldapConn.Close()
	if err = ldapConn.Bind(p.Params["user"], p.Params["pass"]); err != nil {
		return nil, err
	}
	request := ldap.NewSearchRequest(p.Params["base"], ldap.ScopeWholeSubtree, ldap.DerefAlways, 1000, 10, false, filter, getAttrFields(p), nil)
	result, err := ldapConn.Search(request)
	if err != nil {
		return nil, err
	}
	return result.Entries, nil
}
