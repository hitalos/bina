<?php
$ldap = new ldapJFAL;
$result = $ldap->search();

$list = array();
foreach($result as $key => $person){
    $contato = array();
    if(is_numeric($key)){
        foreach($ldap->attrs as $attr){
            $contato[strtolower($attr)] = utf8_encode($person[strtolower($attr)][0]);
        }
    $list[] = $contato;
    }
}

header('Content-Type: application/json;charset=utf8');
echo json_encode($list, JSON_UNESCAPED_UNICODE | JSON_PRETTY_PRINT);
