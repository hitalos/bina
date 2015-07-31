<?php
$ldap = new ldapJFAL;
$list = $ldap->search();

$doc = new DomDocument('1.0', 'UTF-8');
$root = $doc->appendChild($doc->createElement('AddressBook'));
$root->setAttribute('count', count($list));

foreach($list as $person){
	if(isset($person['ipphone'])){
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
}
header('Content-Type: text/xml;charset=utf8');
echo $doc->saveXML();
