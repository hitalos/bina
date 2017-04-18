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

## With Docker

  Copy `.env-example` to `.env` and edit variables.
  
Run:

    docker build -t bina .
    docker run -p 3000:3000 --env-file .env bina

or

    docker-compose up -d

## Trying without LDAP Server

Comment this line in `src/routes/contacts.js`:

    const ldapService = require('../ldapService')

Create contacts folder:

    mkdir public/contacts

In this new folder, create a file `all.json` with "fake" contacts in this format:

```
[
  {
    "id": "user_id_1",
    "fullName": "Jonh Doe",
    "phones": {
      "ipPhone": "5555",
      "mobile": "9999-9999"
    },
    "emails": {
      "mail": "user_id_1@domain.com"
    },
    "department": "IT",
    "title": "Employee",
    "objectClass": "user"
  },
  {
    "id": "user_id_2",
    "fullName": "Jane Doe",
    "phones": {
      "ipPhone": "5556",
      "telephoneNumber": "555-5555"
    },
    "emails": {
      "mail": "user_id_2@domain.com"
    },
    "department": "marketing",
    "title": "Outsourced",
    "objectClass": "contact"
  }
]
```
Follow **install** instructions.

## Debugging

Before running app, set environment variable `DEBUG`:

    DEBUG=Bina:* npm start
