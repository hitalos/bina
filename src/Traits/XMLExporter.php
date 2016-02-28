<?php
namespace Bina\Traits;


trait XMLExporter
{
    protected $doc;
    protected $root;

    public function __construct(){
        $this->doc = new \DomDocument('1.0', 'UTF-8');
    }

    public function __toString(){
        return $this->doc->saveXML();
    }
}
