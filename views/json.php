<?php
$cachefile = "/tmp/cache/" . $_SERVER['REQUEST_URI'];
$cachetime = 5 * 60; // 5 minutes
$usecache = true;
if(isset($_SERVER['HTTP_CACHE_CONTROL']) and $_SERVER['HTTP_CACHE_CONTROL'] == 'no-cache'){
	$usecache = false;
}
if ($usecache && file_exists($cachefile) && (time() - $cachetime < filemtime($cachefile))) {
	if(isset($_SERVER['HTTP_IF_MODIFIED_SINCE']) && strtotime($_SERVER['HTTP_IF_MODIFIED_SINCE']) >= filemtime($cachefile)) {
		header('HTTP/1.0 304 Not Modified');
	}
	else{
		header('Content-Type: application/json;charset=utf8');
		header('Cache-Control: max-age=' . $cachetime);
		header("Pragma: cache");
		header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', filemtime($cachefile)));
		header('Expires: ' . gmdate('D, d M Y H:i:s \G\M\T', time() + $cachetime));
		require_once($cachefile);
	}
}
else{
	ob_start();
	$ldap = new ldapJFAL;
	$result = $ldap->search();

	$list = array();
	foreach($result as $key => $person){
	    $contato = array();
	    if(is_numeric($key)){
	        foreach($ldap->attrs as $attr){
	            if(isset($person[strtolower($attr)][0]) && trim($person[strtolower($attr)][0]) != ''){
	                $contato[strtolower($attr)] = utf8_encode(trim($person[strtolower($attr)][0]));
	            }
	        }
	        $list[] = $contato;
	    }
	}
	header('Cache-Control: max-age=' . $cachetime);
	header("Pragma: cache");
	header('Last-Modified: ' . gmdate('D, d M Y H:i:s \G\M\T', time()));
	header('Expires: ' . gmdate('D, d M Y H:i:s \G\M\T', time() + $cachetime));
	header('Content-Type: application/json;charset=utf8');
	echo json_encode($list, JSON_UNESCAPED_UNICODE);

	if(!file_exists('/tmp/cache')){
		mkdir("/tmp/cache");
	}
	if($fp = fopen($cachefile, 'w')){
		fwrite($fp, ob_get_contents());
		fclose($fp);
	}
	else{
		error_log("Erro ao tentar gravar arquivo no cache.");
	}
	ob_end_flush();
}
