
## Using Docker for MongoDB
```
$ docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4
```
## Test it
```
$ curl -v localhost:27017
```
