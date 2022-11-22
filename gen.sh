
curl http://localhost:3000/openapi.json > openapi.json 
openapi-generator-cli generate -i openapi.json -g go -o ./golang --additional-properties packageName="cloudafrica"
