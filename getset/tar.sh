go build

rm -r getset
rm -rf STO
rm -rf STO.tar
mkdir -p STO

cp -R $(ls . | grep -v .git) STO

tar -cvf ./STO.tar ./STO

rm -rf STO
