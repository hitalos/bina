package ldap

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func getFields(envVar string) []string {
	return strings.Split(os.Getenv(envVar), ",")
}

func getPhoneFilters() string {
	output := ""
	for _, field := range getFields("PHONE_FIELDS") {
		output += fmt.Sprintf("(%s=*)", field)
	}
	return output
}

func getAttrFields() []string {
	fields := []string{"objectClass", os.Getenv("IDENT_FIELD"), os.Getenv("FULL_NAME_FIELD")}
	fields = append(fields, getFields("PHONE_FIELDS")...)
	fields = append(fields, getFields("EMAIL_FIELDS")...)
	fields = append(fields, getFields("OTHER_FIELDS")...)
	return fields
}

// GetContacts searches contacts on LDAP server
func GetContacts() (*ldap.SearchResult, error) {
	host := os.Getenv("LDAP_HOST")
	user := os.Getenv("LDAP_USER")
	pass := os.Getenv("LDAP_PASS")
	base := os.Getenv("LDAP_BASE")

	filter := "(&" +
		"(|" + getPhoneFilters() + ")" +
		"(objectCategory=person)" +
		"(!(UserAccountControl:1.2.840.113556.1.4.803:=2))" +
		"(|(objectClass=user)(objectClass=contact)))"

	ldapConn, err := ldap.DialURL("ldap://" + host + ":389")
	if err != nil {
		return nil, err
	}
	if err = ldapConn.Bind(user, pass); err != nil {
		return nil, err
	}
	request := ldap.NewSearchRequest(base, ldap.ScopeWholeSubtree, ldap.DerefAlways, 1000, 10, false, filter, getAttrFields(), nil)
	return ldapConn.Search(request)
}
