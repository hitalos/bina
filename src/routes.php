<?php

$app->get('/', function(\Slim\Http\Request $req, \Slim\Http\Response $res){
    return $res->write(file_get_contents(__DIR__ . '/../Templates/index.html'));
});


$app->get('/contatos/json', function(\Slim\Http\Request $req, \Slim\Http\Response $res){
    return $res->withJson($this->ldap->cache());
});


$app->get('/vcard/{contato}', function(\Slim\Http\Request $req, \Slim\Http\Response $res, $args){
    $this->ldap->filter = '(&(objectClass=user)(objectCategory=person)(sAMAccountName=' . $args['contato'] . '))';
    $contato = $this->ldap->cache($args['contato']);

    $card = new Bina\Transformers\VcardCreator($contato[0]);
    return $res->withHeader('Content-Type', 'text/x-vcard; charset=utf-8')
        ->withHeader('Content-Disposition', 'attachment; filename="' . $card->getFilename() . '.vcf"')
        ->write($card->buildVCard());
});


$app->get('/xml/{format}', function(\Slim\Http\Request  $req, \Slim\Http\Response $res, $args){
    switch($args['format']){
        case 'yealink':
            $this->ldap->attrs = ['DisplayName', 'ipPhone', 'mobile', 'telephoneNumber'];
            $doc = new Bina\Transformers\Yealink();
            break;
        case 'grandstream':
            $this->ldap->attrs = ['DisplayName', 'ipPhone'];
            $doc = new Bina\Transformers\Grandstream();
    }
    $doc->build($this->ldap->cache($args['format']));

    return $res->write($doc)->withHeader('Content-Type', 'text/xml; charset=UTF-8');
});
