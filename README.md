
# BINA

## Dependências
Para instalação das dependências precisa-se dos seguintes executáveis no SO:

* `php` (com as devidas extensões)
 * `dom`
 * `json`
 * `ldap`
 * `xml`
 * `zlib`
* `npm` (gerenciador de pacotes do **nodeJS** usado para instalar o bower)
* `composer` (gerenciador de pacotes do PHP)

## Instalação das dependências
Execute os comandos abaixo:  
```
sudo npm install -g bower
bower install
composer install
```

## Ambiente de execução
Pode-se usar qualquer servidor web com suporte a PHP ou um container `docker`.  
Segue instruções para as 3 opções mais práticas.
#### Docker
Gerar a imagem:  
```
docker build -t bina .
```
Instalar as dependências:
```
#Javascript
docker run --rm -it -v $PWD:/var/www node npm --prefix /var/www install bower
docker run --rm -it -v $PWD:/var/www node /var/www/node_modules/bower/bin/bower --allow-root

#PHP
docker run --rm -it -v $PWD:/var/www bina curl -o composer.phar https://getcomposer.org/composer.phar
docker run --rm -it -v $PWD:/var/www bina php composer.phar install
```
Executar o container (trocar a porta se necessário ou usar um proxy reverso como o **NGINX**):  
```
docker run -d --name jacunbina -v $PWD:/var/www -p 80:80 bina
```

#### Apache
* Habilite os módulos:
 * `mod_rewrite`
 * `mod_expires`
* Configure o apache para trabalhar com o PHP

#### PHP (Built-in Server)
Execute o comando:  
```
php -S 0.0.0.0:80 index.php
```
