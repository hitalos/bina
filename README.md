# Bina

A web App to search in a LDAP for users and contacts.

## Requirements

* NodeJS 6.x or newer

## Running

### Standalone

#### Install dependencies

    npm install && npm run build && npm run copy-env
or

    yarn && yarn run build && yarn run copy-env
    
#### Configuring

After install, edit file `.env` and set variables.
    
#### Start app
    npm start

If you need, use a third parameter to provide a config file (`.env` renamed for alternative or testing environment). 

## With Docker

  Copy `.env-example` to `.env` and edit variables.
  
Run:

    docker build -t bina .
    docker run -p 3000:3000 --env-file .env bina

or

    docker-compose up -d

## Debugging

Before running app, set environment variable `DEBUG`:

    DEBUG=Bina:* npm start
