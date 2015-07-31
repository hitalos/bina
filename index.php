<?php
require_once('classes/ldapJFAL.class.php');

switch($_SERVER['REQUEST_URI']){
	case '/':
		require_once('views/busca.html');
		break;
	case '/index.json':
		require_once('views/json.php');
		break;
	case '/phonebook.xml':
	case '/gs_phonebook.xml':
		require_once('views/grandstream.php');
		break;
	default:
		return false;
}
//Saída para log
$stdout = fopen('php://stdout', 'a');
fwrite($stdout, $_SERVER['REMOTE_ADDR'] . ' ' . $_SERVER['REQUEST_URI'] . "\n");
fclose($stdout);
