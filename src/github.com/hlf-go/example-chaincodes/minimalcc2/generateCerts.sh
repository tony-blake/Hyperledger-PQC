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

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

function print_and_verify_cert()
{
        CERT=$1
        CA_CERT=$2

        # Print new cert
        openssl x509 -engine qs_sig -in $CERT -noout -text

        # Verify the cert's classical signature
        openssl verify -CAfile $CA_CERT $CERT || exit 1
        echo "Classical verification success"

        # Verify the cert's alt signature
        openssl x509QSVerify -engine qs_sig -root $CA_CERT -untrusted $CERT  -cert $CERT || exit 1
        echo "Alt Signature verification success"
}

# Fetch CA cert
./estclient $VERBOSE_FLAG -g -u estuser -h estpwd -s $EST_HOST -p $EST_PORT -o $OUTPUT_DIR || exit 1
echo "Fetch CA cert success"

########################################
# estclient tests
########################################
date
# print CA cert
openssl base64 -d -in $OUTPUT_PKCS7_CACERT | openssl pkcs7 -engine qs_sig -inform DER -text -print_certs -noout

# Convert PKCS7 CA Cert to PEM
openssl base64 -d -in $OUTPUT_PKCS7_CACERT | openssl pkcs7 -inform DER  -print_certs -out $OUTPUT_PEM_CACERT || exit 1

# Enroll a new cert with new key
./estclient $VERBOSE_FLAG -e  --common-name "Newly Enrolled MPKAC" -u estuser -h estpwd -s $EST_HOST -p $EST_PORT -o $OUTPUT_DIR --pem-output || exit 1
date
sleep 13
print_and_verify_cert $OUTPUT_PEM_CERT $OUTPUT_PEM_CACERT
echo "New cert enrollment success"





########################################
# estclient-simple tests
########################################

#popd
#pushd example/client-simple

# Enroll new cert using estclient-simple
#./estclient_simple -u estuser -h estpwd -s $EST_HOST -p $EST_PORT || exit 1
# convert estclient-simple cert to PEM
#openssl base64 -d -in cert-b64.pkcs7 | openssl pkcs7 -inform DER  -print_certs -out $OUTPUT_DIR/estclient-simple.pem || exit 1
#sleep 13
#print_and_verify_cert $OUTPUT_DIR/estclient-simple.pem $OUTPUT_PEM_CACERT
#echo "New estclient_simple cert enrollment success"

echo
echo "--------------------"
echo "Success!!"
echo "--------------------"
echo


