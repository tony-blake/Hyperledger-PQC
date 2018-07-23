#!/bin/bash

PATH=/usr/local/ssl/bin:$PATH

# Get CA cert
#rm 443-root.crt
#wget http://test-pqpki.com/443-root.crt

export EST_OPENSSL_CACERT=`realpath 443-root.crt`

pushd example/client

export OUTPUT_DIR=/Users/tonyblake/Desktop/marbles/libest/cryptogen
export OUTPUT_PKCS7_CACERT=$OUTPUT_DIR/cacert-0-0.pkcs7
export OUTPUT_PEM_CACERT=$OUTPUT_DIR/cacert-0-0.pem
export OUTPUT_PKCS7_CERT=$OUTPUT_DIR/cert-0-0.pkcs7
export OUTPUT_PEM_CERT=$OUTPUT_DIR/cert-0-0.pem

export EST_HOST=test-pqpki.com
export EST_PORT=443


openssl x509QSVerify -engine qs_sig -root $CA_CERT -untrusted $CERT  -cert $CERT
