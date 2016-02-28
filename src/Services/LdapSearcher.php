<?php
namespace Bina\Services;

/**
 * Serviço para buscar informações num servidor LDAP
 */
class LdapSearcher {

    /** @var string $host IP do servidor LDAP */
    protected $host = '';

    /** @var string $user Nome de usuário para realizar a conexão LDAP */
    protected $user = '';

    /** @var string $pass Senha para conexão LDAP */
    protected $pass = '';

    /** @var string $basedn Caminho do LDAP onde serão executadas as buscas */
    protected $basedn = '';

    /** @var string $filter Query usada na busca LDAP */
    public $filter = '(|(&(objectClass=user)(objectCategory=person)(&(objectCategory=person)(objectClass=contact)))(ipPhone=*))';

    /** @var string[] $attrs Atributos desejados no retorno das buscas */
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

    /** @var array $list Resultado da busca */
    public $list = [];

    /** @var resource $conn objeto de conexão (nativo do PHP) */
    protected $conn = false;

    /**
     * Recebe parâmetros de conexão e configura no objeto
     *
     * @param     string $host
     * @param     string $user
     * @param     string $pass
     * @param     string $basedn
     * @return    void
     */
    public function __construct($host, $user, $pass, $basedn)
    {
        $this->host = $host;
        $this->user = $user;
        $this->pass = $pass;
        $this->basedn = $basedn;
    }

    /**
     * Cria conexão se ainda não foi solicitado
     *
     * @return    void
     */
    public function connect()
    {
        if(!$this->conn){
            $this->conn = ldap_connect($this->host);
            ldap_bind($this->conn, $this->user, $this->pass);
        }
    }

    /**
     * Executa busca no LDAP e retorna um array com o resultado
     *
     * @return    array
     */
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

    /**
     * Executa buscas, lê e guarda resultado em cache
     *
     * @param     string $name Nome dado ao objeto em cache
     * @return    array
     */
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
