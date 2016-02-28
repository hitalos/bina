<?php
namespace Bina\Services;

class App extends \Slim\App {

    public function __construct()
    {
        parent::__construct([
            'settings' => ['displayErrorDetails' => getenv('DEBUG')]
        ]);

        $container = $this->getContainer();

        // ldap searcher
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
