# Elasticsearch cli tools

Set of CLI tools for ElasticSearch

## Usage

Example with custom elastic server URL:
```
$ ELASTICSEARCH_URL=http://localhost:9200 elastic-cli index list
```
#### Index commands

* List existing indexes
```
$ elastic-cli index list
```
* Create index
```
$ elastic-cli index create %name%
```
* Check if index exists
```
$ elastic-cli index exists %name%
```
* Delete index
```
$ elastic-cli index delete %name%
```

#### Alias commands

* List existing aliases
```
$ elastic-cli alias list
```
* Create index alias
```
$ elastic-cli alias create %index_name% %alias_name%
```
* Delete index alias
```
$ elastic-cli alias delete %index_name% %alias_name%
```

## Development

* Install required VS Code extensions:
```
$ cat .vscode/.vscode-extensions | xargs -L 1 code --install-extension
```
* Copy `docker-compose.override.yml.dist` to `docker-compose.override.yml`
```
$ cp docker-compose.override.yml.dist docker-compose.override.yml
```
* Create docker network (if not exists)
```
$ docker network create f1monkey
```
* Start container
```
$ docker-compose up -d
```
* Connect to the docker container using [Remote Containers](https://code.visualstudio.com/docs/remote/containers) extension