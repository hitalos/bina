<?php
namespace Bina\Services;

class LdapSearcher {

    protected $host = '';
    protected $user = '';
    protected $pass = '';
    protected $basedn = '';
    public $filter = '(|(&(objectClass=user)(objectCategory=person)(&(objectCategory=person)(objectClass=contact)))(ipPhone=*))';
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
		'objectClass',
        'useraccountcontrol'
    ];
    public $list = [];

    protected $conn = false;

    public function __construct($host, $user, $pass, $basedn)
    {
        $this->host = $host;
        $this->user = $user;
        $this->pass = $pass;
        $this->basedn = $basedn;
    }

    public function connect()
    {
        if(!$this->conn){
            $this->conn = ldap_connect($this->host);
            ldap_bind($this->conn, $this->user, $this->pass);
        }
    }

    public function search(){
        $this->connect();
        $resource = ldap_search(
            $this->conn,
            $this->basedn,
            $this->filter,
            $this->attrs
        );
        $result = ldap_get_entries($this->conn, $resource);
        foreach($result as $key => &$person){
            if(is_numeric($key)){
                if(isset($person['useraccountcontrol']) and ($person['useraccountcontrol'][0] & 0x2)){
                    unset($result[$key]);
                }
                if(isset($person['displayname'])){
                    unset($person);
                }
            }
        }
        foreach($result as $key => $person){
            $contato = array();
            if(is_numeric($key)){
                foreach($this->attrs as $attr){
                    if(isset($person[strtolower($attr)])){
                        array_shift($person[strtolower($attr)]);
                        foreach($person[strtolower($attr)] as $info){
                            $contato[strtolower($attr)][] = utf8_encode(trim($info));
                        }
                    }
                }
                $this->list[] = $contato;
            }
        }
        usort($this->list, function($a, $b){
            if ($a['displayname'] == $b['displayname'])
                return 0;
            return $a['displayname'] < $b['displayname'] ? -1 : 1;
        });
        return $this->list;
    }

    public function cache($name = 'default')
    {
        $cachedir = __DIR__ . '/../../cache';
        if(!file_exists($cachedir)){
            mkdir($cachedir);
        }
        $cachefile = "$cachedir/$name.json";

        if (file_exists($cachefile) && time() - getenv('CACHETIME') < filemtime($cachefile)) {
            $this->list = json_decode(file_get_contents($cachefile), true);
        }
        else{
            $this->search();
            $cached = fopen($cachefile, 'w');
            fwrite($cached, json_encode($this->list));
            fclose($cached);
        }
        return $this->list;
    }
}
