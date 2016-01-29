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
	if(isset($_SERVER['HTTP_IF_MODIFIED_SINCE']) && strtotime($_SERVER['HTTP_IF_MODIFIED_SINCE']) >= filemtime($file) && (isset($_SERVER['HTTP_CACHE_CONTROL']) && $_SERVER['HTTP_CACHE_CONTROL'] != 'no-cache')) {
    	header('HTTP/1.0 304 Not Modified');
	}
	else{
		header('Cache-Control: max-age=' . 60 * 60 * 24);
		header("Pragma: cache");
		header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', filemtime($file)));
		header('Expires: ' . gmdate('D, d M Y H:i:s \G\M\T', time() + 60 * 60 * 24));
		header('Content-Type: ' . $type);
		require_once($file);
	}
}
else{
	switch($path[1]){
		case '':
			$file = 'views/busca.html';
			$type = 'text/html';
			header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', filemtime($file)));
			break;
		case 'vcard':
			require_once('views/vcard.php');
			die();
		case 'index.json':
			$file = 'views/json.php';
			$type = 'application/json';
			header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', time()));
			break;
		case 'phonebook.xml':
		case 'gs_phonebook.xml':
			ini_set('zlib.output_compression', 'Off');
			header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', time()));
			$file = 'views/grandstream.php';
			$type = 'text/xml';
		case 'yealink.xml':
			$file = 'views/yealink.php';
			break;
		default:
			ini_set('zlib.output_compression', 'Off');
			return false;
	}
	header('Cache-Control: max-age=' . 60 * 60);
	header("Pragma: cache");
	header('Expires: ' . gmdate('D, d M Y H:i:s \G\M\T', time() + 60 * 60));
	header('Content-Type: ' . $type);
	require $file;
}
