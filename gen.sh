
curl http://localhost:3000/swagger.json > openapi.json 
/home/andre/.asdf/installs/golang/1.19.2/packages/bin/oapi-codegen --package cloudafrica openapi.json  > cloudafrica.gen.go
