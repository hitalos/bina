
# BINA

Para baixar o código:  
`git clone [usuário@]dev:/repo/bina.git`

## Dependências
Para instalação das dependências precisa-se dos seguintes executáveis no SO:

* `php` (com as devidas extensões)
 * `dom`
 * `json`
 * `ldap`
 * `xml`
 * `zlib`
* `bower` (gerenciador de pacotes JavaScript)  
* `composer` (gerenciador de pacotes do PHP)

## Instalação das dependências
Execute os comandos abaixo:  
`bower install`  
`composer install`

## Ambiente de execução
Pode-se usar qualquer servidor web ou um container `docker`.  
#### Para execução em ambiente docker
Para gerar a imagem:  
`docker build -t bina .`  

Para executar o container (trocar a porta se necessário):  
`docker run -d --name jacunbina -v $PWD:/var/www -p 80:80 bina`

#### Para execução em ambiente Apache
* Habilite os módulos:
 * `mod_rewrite`
 * `mod_expires`
* Configure o apache para trabalhar com o PHP
