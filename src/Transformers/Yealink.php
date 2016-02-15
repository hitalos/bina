<?php
namespace Bina\Transformers;

use Bina\Traits\Exporter;

class Yealink {
    use Exporter;

    public function build($contatos){
        $this->root = $this->doc->appendChild($this->doc->createElement('JFALIPPhoneDirectory'));
        foreach($contatos as $key => $person){
            if(is_numeric($key)){
                $contato = $this->root->appendChild($this->doc->createElement('DirectoryEntry'));
                $contato->appendChild($this->doc->createElement('Name', ($person['displayname'][0])));

                if(isset($person['ipphone'])){
                    if(is_array($person['ipphone'])){
                        $contato->appendChild($this->doc->createElement('Telephone', $person['ipphone'][0]));
                    }
                    else{
                        $contato->appendChild($this->doc->createElement('Telephone', $person['ipphone']));
                    }
                }
                if(isset($person['mobile'])){
                    if(is_array($person['mobile'])){
                        $contato->appendChild($this->doc->createElement('Telephone', $person['mobile'][0]));
                    }
                    else{
                        $contato->appendChild($this->doc->createElement('Telephone', $person['mobile']));
                    }
                }
                if(isset($person['telephonenumber'])){
                    if(is_array($person['telephonenumber'])){
                        $contato->appendChild($this->doc->createElement('Telephone', $person['telephonenumber'][0]));
                    }
                    else{
                        $contato->appendChild($this->doc->createElement('Telephone', $person['telephonenumber']));
                    }
                }
            }
            else {
                unset($contatos[$key]);
            }
        }
    }
}