cd ..

HOST=public.ecr.aws/l8z6j2l3
REPO=zexuz/crypto-idp
TAG=latest
NAME=$HOST/$REPO:$TAG

docker build -t $NAME .