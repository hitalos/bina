<?php
ini_set('zlib.output_compression', 'On');
require_once('classes/ldapJFAL.class.php');

$extension = array_pop(explode('.', $_SERVER['REQUEST_URI']));
$file = substr($_SERVER['REQUEST_URI'], 1);
$path = explode('/', $_SERVER['REQUEST_URI']);
$type = '';
switch($extension){
	case 'css':
		$type = 'text/css';
		break;
	case 'js':
		$type = 'text/javascript';
		break;
}
if(file_exists($file)){
	if(isset($_SERVER['HTTP_IF_MODIFIED_SINCE']) && strtotime($_SERVER['HTTP_IF_MODIFIED_SINCE']) >= filemtime($file)) {
    	header('HTTP/1.0 304 Not Modified');
	}
	else{
		header('Cache-Control: max-age=' . 60 * 60);
		header("Pragma: cache");
		header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', filemtime($file)));
		header('Expires: ' . gmdate('D, d M Y H:i:s \G\M\T', time() + 60 * 60));
		header('Content-Type: ' . $type);
		require_once($file);
	}
}
else{
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
			ini_set('zlib.output_compression', 'Off');
			require_once('views/grandstream.php');
			break;
		default:
			ini_set('zlib.output_compression', 'Off');
			return false;
	}
}
//Saída para log
$stdout = fopen('php://stdout', 'a');
fwrite($stdout, $_SERVER['REMOTE_ADDR'] . ' ' . $_SERVER['REQUEST_URI'] . "\n");
fclose($stdout);
