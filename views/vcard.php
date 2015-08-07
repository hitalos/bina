<?php
require_once('vendor/autoload.php');

$ldap = new ldapJFAL;
$ldap->filter = '(&(objectClass=user)(objectCategory=person)(sAMAccountName=' . $path[2] . '))';
$result = $ldap->search();

// Preparando dados
$displayName = utf8_encode($result[0]['displayname'][0]);
$names = explode(' ', $displayName);
$firstName = array_shift($names);
$lastName = array_pop($names);
if(count($names)){
    $middleNames = implode(' ', $names);
}

//Gerando o vcard
use JeroenDesloovere\VCard\VCard;
$vcard = new VCard();
$vcard->addName($lastName, $firstName, $middleNames, '', '');
$vcard->addJobtitle($result[0]['title'][0]);
$vcard->addCompany('Justiça Federal em Alagoas');
$vcard->addNote('Lotação: ' . utf8_encode($result[0]['department'][0] . ' - ' . $result[0]['physicaldeliveryofficename'][0]));
$vcard->addEmail($result[0]['mail'][0], 'WORK');
$vcard->addPhoneNumber('082 2122-' . $result[0]['ipphone'][0], 'PREF;WORK');
if(isset($result[0]['mobile'][0])){
    $vcard->addPhoneNumber('0' . $result[0]['mobile'][0], 'CELL');
}
$vcard->addPhoto('http://www.jfal.jus.br/images/fotos3x4/' . $path[2] . '.jpg');

return $vcard->download();
