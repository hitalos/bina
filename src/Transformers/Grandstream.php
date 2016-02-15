<?php
namespace Bina\Transformers;

use Bina\Traits\Exporter;

class Grandstream {
    use Exporter;

    public function build($contatos)
    {
        $this->root = $this->doc->appendChild($this->doc->createElement('AddressBook'));
        foreach ($contatos as $key => $person) {
            if (is_numeric($key)) {
                $contato = $this->root->appendChild($this->doc->createElement('Contact'));
                $displayName = utf8_encode($person['displayname'][0]);
                $names = explode(' ', $displayName);
                $firstName = array_shift($names);
                $firstName .= ' ' . array_shift($names);
                $lastName = '';
                if (count($names)) {
                    $lastName = array_pop($names);
                }
                $contato->appendChild($this->doc->createElement('FirstName', $firstName));
                $contato->appendChild($this->doc->createElement('LastName', $lastName));
                $phone = $contato->appendChild($this->doc->createElement('Phone'));
                $phone->appendChild($this->doc->createElement('phonenumber', $person['ipphone'][0]));
            } else {
                unset($contato[$key]);
            }
        }
        $this->root->setAttribute('count', count($contatos));
    }
}