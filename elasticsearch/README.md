## Elasticsearch with Golang to Create , Delete indexs and Search for data inside it:

- This is an entry point for dealing with Elasticsearch using its golang client and it is in progress.
- This example is based on Elasticsearch 7.* and later
- Install Elasticsearch in Docker as the below :
```
$ docker run -- name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.2.0
```

- To test it :
```
$ curl -v http://localhost:9200
```
- run the code :
```
$ go get -u -v github.com/olivere/elastic
& go run main.go
```
- To look for Created Index :
```
$ curl 'localhost:9200/_cat/indices'
```
- To Look for Data of a Document in "Test" Index :
```
$ curl -XGET http://localhost:9200/Test/_search?pretty=true
```
