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
$ldap->attrs = ['DisplayName', 'ipPhone', 'mobile', 'telephoneNumber'];
$list = $ldap->search();

$doc = new DomDocument('1.0', 'UTF-8');

$root = $doc->appendChild($doc->createElement('JFALIPPhoneDirectory'));

foreach($list as $key => $person){
	if(is_numeric($key)){
		$contato = $root->appendChild($doc->createElement('DirectoryEntry'));
		$contato->appendChild($doc->createElement('Name', utf8_encode($person['displayname'][0])));

        if(isset($person['ipphone'])){
            if(is_array($person['ipphone'])){
                $contato->appendChild($doc->createElement('Telephone', $person['ipphone'][0]));
            }
            else{
                $contato->appendChild($doc->createElement('Telephone', $person['ipphone']));
            }
        }
        if(isset($person['mobile'])){
            if(is_array($person['mobile'])){
                $contato->appendChild($doc->createElement('Telephone', $person['mobile'][0]));
            }
            else{
                $contato->appendChild($doc->createElement('Telephone', $person['mobile']));
            }
        }
        if(isset($person['telephonenumber'])){
            if(is_array($person['telephonenumber'])){
                $contato->appendChild($doc->createElement('Telephone', $person['telephonenumber'][0]));
            }
            else{
                $contato->appendChild($doc->createElement('Telephone', $person['telephonenumber']));
            }
        }
	}
	else {
		unset($list[$key]);
	}
}

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
