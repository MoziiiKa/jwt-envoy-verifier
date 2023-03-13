#!/bin/bash

# generate private key in PEM format
ssh-keygen -t rsa -b 4096 -N "" -m PEM -f private.key

# generate public key in PEM format
ssh-keygen -f private.key -e -m PKCS8 > public.key

# copy files into jwkgenerator folder
cp private.key ./jwkgenerator/private.key
cp public.key ./jwkgenerator/public.key

# run jwkgenerator via docker
touch ./jwkgenerator/jwks.json
docker build -t jwkgenerator ./jwkgenerator
docker run -v $(pwd)/jwkgenerator/private.key:/app/private.key -v $(pwd)/jwkgenerator/jwks.json:/app/jwks.json -it --rm --name jwkgenerator jwkgenerator

# copy jwks.json into the envoy folder
cp ./jwkgenerator/jwks.json ./envoy/jwks.json

# copy private.key and public.key to application folder
cp private.key ./application/private.key
cp public.key ./application/public.key

# remove private.key and public.key from jwkgenerator folder
rm ./jwkgenerator/private.key
rm ./jwkgenerator/public.key

# remove private.key and public.key from root folder
rm private.key
rm private.key.pub
rm public.key

# run project via docker compose
docker compose up -d --build
