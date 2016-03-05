<?php
namespace Bina\Traits;

/**
 * Inclui propriedades e métodos comuns para classes que exportam XML
 */
trait XMLExporter
{
    /** @var \DomDocument $doc Objeto para manipulação do DOM */
    protected $doc;

    /** @var \DomElement $root Elemento raiz do DOM */
    protected $root;

    /**
     * Inicializa objeto DOM e armazena na propriedade $doc
     *
     * @return    void
     */
    public function __construct()
    {
        $this->doc = new \DomDocument('1.0', 'UTF-8');
    }

    /**
     * Método 'mágico' do PHP, retorna string com o XML atual
     *
     * @return    string
     */
    public function __toString()
    {
        return $this->doc->saveXML();
    }
}
