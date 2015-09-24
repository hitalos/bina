<?php
class ldapJFAL {
    const host = '172.16.0.11';
    const base = 'ou=JFAL - Usuarios,dc=jfal,dc=jus,dc=br';
    const user = 'bina@jfal.jus.br';
    const pass = 'bina@p0rt4r14';
    public $filter = '(&(objectClass=user)(objectCategory=person)(ipPhone=*))';
    public $attrs = [
        'DisplayName',
        'sAMAccountName',
        'ipPhone',
        'mobile',
        'homePhone',
        'facsimileTelephoneNumber',
        'mail',
        'Department',
        'PhysicalDeliveryOfficeName',
        'title',
        'employeeID',
		'proxyAddresses',
        'useraccountcontrol'
    ];
    private $conn = false;

    public function __construct(){
        if(!$this->conn){
            $this->conn = ldap_connect(self::host);
            ldap_bind($this->conn, self::user, self::pass);
        }
    }

    public function search(){
        $resource = ldap_search($this->conn, self::base, $this->filter, $this->attrs);
        ldap_sort($this->conn, $resource, 'DisplayName');
        $list = ldap_get_entries($this->conn, $resource);
        foreach($list as $key => &$person){
            if(is_numeric($key)){
                if(isset($person['useraccountcontrol']) and ($person['useraccountcontrol'][0] & 0x2)){
                    unset($list[$key]);
                }
            }
        }
        return $list;
    }
}
