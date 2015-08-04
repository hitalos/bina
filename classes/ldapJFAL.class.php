<?php
class ldapJFAL {
    const host = '172.16.0.11';
    const base = 'ou=JFAL - Usuarios,dc=jfal,dc=jus,dc=br';
    const user = 'bina@jfal.jus.br';
    const pass = 'bina@p0rt4r14';
    const filter = '(&(objectClass=user)(objectCategory=person)(ipPhone=*))';
    public $attrs = array('DisplayName', 'ipPhone', 'mail', 'Department', 'PhysicalDeliveryOfficeName', 'title', 'employeeID');
    private $conn = false;

    public function __construct(){
        if(!$this->conn){
            $this->conn = ldap_connect(self::host);
            ldap_bind($this->conn, self::user, self::pass);
        }
    }

    public function search(){
        $resource = ldap_search($this->conn, self::base, self::filter, $this->attrs);
        return ldap_get_entries($this->conn, $resource);
    }
}
