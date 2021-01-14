# analyserService [![Build Status](https://travis-ci.com/samtholiya/analyserService.svg?branch=main)](https://travis-ci.com/samtholiya/analyserService)

## Build

### Local
```
go build  -o /go/bin/server github.com/samtholiya/analyserService/cmd/server
```

### Docker 
```
docker pull samtholiya/analyse-service:latest
```

## Run

### Local
```
/go/bin/server
```
### Docker
```
docker run -it samtholiya/analyse-service:latest
```

