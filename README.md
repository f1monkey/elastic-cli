# Elasticsearch cli tools

Set of CLI tools for ElasticSearch

## Usage

Example with custom elastic server URL:
```
$ ELASTICSEARCH_URL=http://localhost:9200 elastic-cli index list
```
### Index commands

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