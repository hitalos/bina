<?php
ini_set('zlib.output_compression', 'On');
require_once('classes/ldapJFAL.class.php');

$path = explode('/', $_SERVER['REQUEST_URI']);
switch($path[1]){
	case '':
		require_once('views/busca.html');
		break;
	case 'vcard':
		require_once('views/vcard.php');
		break;
	case 'index.json':
		require_once('views/json.php');
		break;
	case 'phonebook.xml':
	case 'gs_phonebook.xml':
		require_once('views/grandstream.php');
		break;
	default:
		ini_set('zlib.output_compression', 'Off');
		return false;
}
//Saída para log
$stdout = fopen('php://stdout', 'a');
fwrite($stdout, $_SERVER['REMOTE_ADDR'] . ' ' . $_SERVER['REQUEST_URI'] . "\n");
fclose($stdout);
