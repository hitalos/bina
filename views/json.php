<?php
$cachefile = "cache/" . $_SERVER['REQUEST_URI'];
$cachetime = 5 * 60; // 5 minutes
if (file_exists($cachefile) && (time() - $cachetime < filemtime($cachefile))) {
	header('Content-Type: application/json;charset=utf8');
    include($cachefile);
    exit;
}
ob_start();
$ldap = new ldapJFAL;
$result = $ldap->search();

$list = array();
foreach($result as $key => $person){
    $contato = array();
    if(is_numeric($key)){
        foreach($ldap->attrs as $attr){
            if(trim($person[strtolower($attr)][0]) != ''){
                $contato[strtolower($attr)] = utf8_encode(trim($person[strtolower($attr)][0]));
            }
        }
        $list[] = $contato;
    }
}

header('Content-Type: application/json;charset=utf8');
echo json_encode($list, JSON_UNESCAPED_UNICODE);


$fp = fopen($cachefile, 'w');
fwrite($fp, ob_get_contents());
fclose($fp);
ob_end_flush();
