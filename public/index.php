<?php
if (PHP_SAPI == 'cli-server') {
    // Verifica se o arquivo solicitado existe no Filesystem
    if (is_file(__DIR__ . $_SERVER['REQUEST_URI'])) {
        return false;
    }
}

require __DIR__ . '/../vendor/autoload.php';

/** @var \Dotenv\Dotenv $dotenv Carrega configurações de ambiente */
$dotenv = new Dotenv\Dotenv(__DIR__ . '/..');
$dotenv->load();

// Cria aplicação
$app = new Bina\Services\App();

// Carrega rotas
require __DIR__ . '/../src/routes.php';

// Analisa requisição e envia resposta
$app->run();
