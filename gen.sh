
curl http://localhost:3000/openapi.json > openapi.json 
npx openapi-generator-cli generate -i openapi.json -g go -o ./golang --additional-properties packageName="cloudafrica"
cp -r golang_custom/* golang/
