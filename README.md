# Encrypter
Test project to better understand how ransomware functions

# Structure

## Web server
Listens on HTTPS for incoming connections (just do http initially)

Logs:
    1. Private key
    2. Ip Address
    3. Hostname

Allows response to issue commands:
    1. Enumerate: hostname, OS version, current directory, ???
    2. Encrypt

## Encrypter
Opens a HTTPS connection back to web server (just do http initially)

Generates a private key

Has a enumerate function, triggered by receiving a particular string back

Has an encrypt function, trigger by receiving a particular string

## Goals
Learn Go

Do not read code for existing ransomware, it is okay to look for guides/code on individual functionality

