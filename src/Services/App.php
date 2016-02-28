<?php
namespace Bina\Services;

/**
 * Extensão da aplicação padrão do Slim
 */
class App extends \Slim\App {

    /**
     * Constrói o objeto App e configura os serviços necessários
     *
     * @return    void
     */
    public function __construct()
    {
        parent::__construct([
            'settings' => ['displayErrorDetails' => getenv('DEBUG')]
        ]);

        $container = $this->getContainer();

        // Serviço ldap searcher
        $container['ldap'] = function($c){
            return new LdapSearcher(
                getenv('LDAP_HOST'),
                getenv('LDAP_USER'),
                getenv('LDAP_PASS'),
                getenv('LDAP_BASE')
            );
        };
    }
}
