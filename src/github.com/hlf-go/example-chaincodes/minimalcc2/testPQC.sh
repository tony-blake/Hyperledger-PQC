#!/bin/bash

PATH=/usr/local/ssl/bin:$PATH

# Get CA cert
rm 443-root.crt
wget http://test-pqpki.com/443-root.crt

export EST_OPENSSL_CACERT=`realpath 443-root.crt`

pushd example/client

export OUTPUT_DIR=$GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/go/minimalcc2/cryptogen
export OUTPUT_PKCS7_CACERT=$OUTPUT_DIR/cacert-0-0.pkcs7
export OUTPUT_PEM_CACERT=$OUTPUT_DIR/cacert-0-0.pem
export OUTPUT_PKCS7_CERT=$OUTPUT_DIR/cert-0-0.pkcs7
export OUTPUT_PEM_CERT=$OUTPUT_DIR/cert-0-0.pem

export EST_HOST=test-pqpki.com
export EST_PORT=443


export VERBOSE_FLAG=


# Print new cert
        openssl x509 -engine qs_sig -in $OUTPUT_PEM_CERT -noout -text

        # Verify the cert's classical signature
        openssl verify -CAfile $OUTPUT_PEM_CACERT $OUTPUT_PEM_CERT 
        echo "Classical verification success"

        # Verify the cert's alt signature
        openssl x509QSVerify -engine qs_sig -root $OUTPUT_PEM_CACERT -untrusted $OUTPUT_PEM_CERT  -cert $OUTPUT_PEM_CERT
