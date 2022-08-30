@echo off

cd ./client/my-community/

yarn build

cd ..
cd ..

cd ./server/

./build-server.bat