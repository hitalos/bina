<?php
if (PHP_SAPI == 'cli-server') {
    // Verifica se o arquivo solicitado existe
    if (is_file(__DIR__ . $_SERVER['REQUEST_URI'])) {
        return false;
    }
}

require __DIR__ . '/../vendor/autoload.php';

/** @var \Dotenv\Dotenv $dotenv Carrega configurações de ambiente */
$dotenv = new Dotenv\Dotenv(__DIR__ . '/..');
$dotenv->load();

$app = new \Slim\App([
    'settings' => ['displayErrorDetails' => getenv('DEBUG')]
]);

$container = $app->getContainer();

// ldap searcher
use \Bina\Services\LdapSearcher as Ldap;
$container['ldap'] = function($c){
    return new Ldap(getenv('LDAP_HOST'), getenv('LDAP_USER'), getenv('LDAP_PASS'), getenv('LDAP_BASE'));
};

// Carrega rotas
require __DIR__ . '/../src/routes.php';

$app->run();
