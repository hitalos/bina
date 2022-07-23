# Bina

A web App to search in a LDAP for users and contacts.

## Requirements

* One or many LDAP servers like Active Directory

## Installing

* Download the released binary for your Operational System

## Configuring

After install, edit an `config.yml` file and set variables (like the [example](config_example.yml)).

## Docker

### To build a image

    make container_image

### To run

    docker run --rm -t -i -p 8000:8000 -v ./config.yml:/app/config.yml bina:latest

## Debugging

Before running app, set environment variable `DEBUG`:

    DEBUG=1
