openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private_key.pem -out public_key.pem

mkdir -p ../assets/keys

mv private_key.pem ../assets/keys/private_key.pem
mv public_key.pem ../assets/keys/public_key.pem