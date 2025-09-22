package ldap

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"

	"github.com/hitalos/bina/internal/config"
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
		"(displayName=*)" +
		"(|" + phoneFields + ")" +
		"(objectCategory=person)" +
		"(!(UserAccountControl:1.2.840.113556.1.4.803:=2))" +
		"(|(objectClass=user)(objectClass=contact)))"

	if ldapTimeout, err := strconv.Atoi(p.Params["timeout"]); err == nil && ldapTimeout > 0 {
		ldap.DefaultTimeout = time.Duration(ldapTimeout) * time.Second
	}

	slog.Debug(fmt.Sprintf("timeout set to %d (in secs)", ldap.DefaultTimeout))

	_, err := strconv.ParseUint(p.Params["port"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("err: %w; invalid port: %s", err, p.Params["port"])
	}
	hostport := net.JoinHostPort(p.Params["host"], p.Params["port"])

	var ldapConn *ldap.Conn
	switch p.Params["schema"] {
	case "ldaps":
		tlsConf := &tls.Config{InsecureSkipVerify: p.IgnoreSSLVerification} //nolint:gosec // allow insecure skip verify
		ldapConn, err = ldap.DialURL("ldaps://"+hostport, ldap.DialWithTLSConfig(tlsConf))
	default:
		ldapConn, err = ldap.DialURL(p.Params["schema"] + "://" + hostport)
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = ldapConn.Close() }()

	if err = ldapConn.Bind(p.Params["user"], p.Params["pass"]); err != nil {
		return nil, err
	}

	request := ldap.NewSearchRequest(
		p.Params["base"],
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		1000,
		10,
		false,
		filter,
		getAttrFields(p),
		nil)

	result, err := ldapConn.Search(request)
	if err != nil {
		return nil, err
	}

	return result.Entries, nil
}
