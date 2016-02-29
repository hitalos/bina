<?php
namespace Bina\Transformers;

use JeroenDesloovere\VCard\VCard;

/**
 * Transformer para gerar arquivos VCF
 */
class VcardCreator extends Vcard {

    /**
     * Constrói objeto Vcard e configura seus atributos
     *
     * @param     array $contato Array de atributos do contato
     * @return    void
     */
    public function __construct($contato)
    {
        $displayName = $contato['displayname'][0];

        /** @var array $names Quebra string para manipulação */
        $names = explode(' ', $displayName);
        $firstName = array_shift($names);
        $lastName = array_pop($names);
        $middleNames = '';

        if (count($names)) {
            $middleNames = implode(' ', $names);
        }

        $this->addName($lastName, $firstName, $middleNames, '', '');
        $this->addJobtitle($contato['title'][0]);
        $this->addCompany('Justiça Federal em Alagoas');
        $this->addNote('Lotação: ' . $contato['department'][0] . ' - ' . $contato['physicaldeliveryofficename'][0]);
        if(isset($contato['mail'])){
            foreach($contato['mail'] as $mail){
                $this->addEmail($mail, 'WORK');
            }
        }
        if(isset($contato['proxyaddresses'])){
            foreach($contato['proxyaddresses'] as $mail){
                $this->addEmail($mail, 'WORK');
            }
        }
        if(isset($contato['ipphone'])){
            foreach($contato['ipphone'] as $ipphone){
                $this->addPhoneNumber('082 2122-' . $ipphone, 'WORK');
            }
        }
        if(isset($contato['mobile'])){
            foreach($contato['mobile'] as $mobile){
                $this->addPhoneNumber('0' . $mobile, 'CELL');
            }
        }

        $this->addPhoto('http://www.jfal.jus.br/images/fotos3x4/' . $contato['samaccountname'][0] . '.jpg');
    }

    public function __toString(){
        return utf8_encode($this->getOutput());
    }
}
