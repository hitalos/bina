<?php
$cachefile = "/tmp/cache/" . $_SERVER['REQUEST_URI'];
$cachetime = 5 * 60; // 5 minutes
if (file_exists($cachefile) && (time() - $cachetime < filemtime($cachefile))) {
	header('Content-Type: text/xml;charset=utf8');
    include($cachefile);
    exit;
}
ob_start();

$ldap = new ldapJFAL;
$ldap->attrs = ['DisplayName', 'ipPhone'];
$list = $ldap->search();

$doc = new DomDocument('1.0', 'UTF-8');
$root = $doc->appendChild($doc->createElement('AddressBook'));

foreach($list as $key => $person){
	if(is_numeric($key)){
		$contato = $root->appendChild($doc->createElement('Contact'));
		$displayName = utf8_encode($person['displayname'][0]);
		$names = explode(' ', $displayName);
		$firstName = array_shift($names);
		$firstName .= ' ' . array_shift($names);
		$lastName = '';
		if(count($names)){
			$lastName = array_pop($names);
		}
		$contato->appendChild($doc->createElement('FirstName', $firstName));
		$contato->appendChild($doc->createElement('LastName', $lastName));
		$phone = $contato->appendChild($doc->createElement('Phone'));
		$phone->appendChild($doc->createElement('phonenumber', $person['ipphone'][0]));
	}
	else {
		unset($list[$key]);
	}
}
$root->setAttribute('count', count($list));

header('Content-Type: text/xml;charset=utf8');
echo $doc->saveXML();

if(!file_exists('/tmp/cache')){
	mkdir("/tmp/cache");
}
if(file_exists($cachefile)){
	$fp = fopen($cachefile, 'w');
	fwrite($fp, ob_get_contents());
	fclose($fp);
}
ob_end_flush();
