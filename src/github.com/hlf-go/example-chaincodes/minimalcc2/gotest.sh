#!/bin/bash


openssl x509QSVerify -engine qs_sig -root 443-root.crt -untrusted cert-0-0.pem  -cert cert-0-0.pem 

echo "Alt Signature verification success"
