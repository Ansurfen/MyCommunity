@echo off

cd ./client/my-community/

yarn build

cd ..
cd ..

cd ./server/

go build main.go