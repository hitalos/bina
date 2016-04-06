<?php
// Rota padrão para HTML de busca
$app->get('/', function (\Slim\Http\Request $req, \Slim\Http\Response $res) {
    return $res->write(file_get_contents(__DIR__ . '/../Templates/index.html'));
});

// Listagem completa de contatos codificada em JSON
$app->get('/contatos/json', function (\Slim\Http\Request $req, \Slim\Http\Response $res) {
    return $res->withJson($this->ldap->cache())
        ->withHeader('Access-Control-Allow-Origin', '*');
});

// Contato no formato VCF, para importação no Android, iOS, Outlook...
$app->get('/vcard/{contato}', function (\Slim\Http\Request $req, \Slim\Http\Response $res, $args) {
    $this->ldap->filter = '(&(objectClass=user)(objectCategory=person)(sAMAccountName=' . $args['contato'] . '))';
    $contato = $this->ldap->cache($args['contato']);

    if (count($contato)) {
        $card = new Bina\Transformers\VcardCreator($contato[0]);
        return $res->withHeader('Content-Type', 'text/x-vcard; charset=utf-8')
            ->withHeader('Content-Disposition', 'attachment; filename="' . $card->getFilename() . '.vcf"')
            ->write($card->buildVCard());
    } else {
        return $res->withStatus(404, 'Not Found');
    }
});

// Listagem completa em formato XML, para importação nos aparelhos VOIP
$app->get('/xml/{format}', function (\Slim\Http\Request  $req, \Slim\Http\Response $res, $args) {
    switch ($args['format']) {
        case 'yealink':
            $doc = new Bina\Transformers\Yealink();
            break;
        case 'grandstream':
            $doc = new Bina\Transformers\Grandstream();
    }
    $doc->build($this->ldap->cache());

    return $res->write($doc)->withHeader('Content-Type', 'text/xml; charset=UTF-8');
});
