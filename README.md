# nighscout-openapi-client

Go client for NightScout API. Built using https://github.com/OpenAPITools/openapi-generator.

# How to build

MacOS:
```
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/nightscout/cgm-remote-monitor/master/lib/server/swagger.yaml \
    -g go \
    -o . && \
    cd generated && go mod tidy
```

or for a local copy

```
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/swagger.yaml \
    -g go \
    -o /local/generated \
    --git-user-id smyachenkov \
    --git-repo-id nightscout-go-client/generated && \
    cd generated && \
    go get . && \
    go mod tidy
```