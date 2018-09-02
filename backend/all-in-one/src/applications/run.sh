#! /bin/bash

echo "...pull..."
git pull

echo "...stop container tourjpn..."
docker stop tourjpn

echo "...remove container tourjpn..."
docker rm tourjpn

echo "...remove image tourjprun..."
docker rmi tourjprun

echo "...copy config file..."
cp ../utils/config.yaml ./

echo "...set linux temp env..."
CGO_ENABLED=0 GOOS=linux 

echo "...building golang..."
go build -a -installsuffix cgo  -tags netgo -o tourjprun

echo "...building docker image..."
docker build -t tourjprun:latest .

echo "...start docker..."
docker run -itd -p 805:805 -v /self/prod/imagesviews:/self/prod/aaa  --name tourjpn --network=yuyang_nw  tourjprun:latest

echo "...done..."

