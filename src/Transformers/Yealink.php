<?php
namespace Bina\Transformers;

use Bina\Traits\XMLExporter;

class Yealink {
    use XMLExporter;

    public function build($contatos){
        $this->root = $this->doc->appendChild($this->doc->createElement('JFALIPPhoneDirectory'));
        foreach($contatos as $key => $person){
            if(is_numeric($key)){
                $contato = $this->root->appendChild($this->doc->createElement('DirectoryEntry'));
                $contato->appendChild($this->doc->createElement('Name', ($person['displayname'][0])));

                $phones = [];
                if(isset($person['ipphone']) and is_array($person['ipphone'])){
                    $phones = $person['ipphone'];
                }
                if(isset($person['mobile']) and is_array($person['mobile'])){
                    $phones += $person['mobile'];
                }
                if(isset($person['telephonenumber']) and is_array($person['telephonenumber'])){
                    $phones += $person['telephonenumber'];
                }
                foreach($phones as $phone){
                    $contato->appendChild($this->doc->createElement('Telephone', $phone));
                }
            }
            else {
                unset($contatos[$key]);
            }
        }
    }
}
