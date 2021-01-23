# analyserService [![Build Status](https://travis-ci.com/samtholiya/analyserService.svg?branch=main)](https://travis-ci.com/samtholiya/analyserService)  [![codecov](https://codecov.io/gh/samtholiya/analyserService/branch/main/graph/badge.svg)](https://codecov.io/gh/samtholiya/analyserService)

## Build

### Local
```
go build  -o /go/bin/server github.com/samtholiya/analyserService/cmd/server
```

### Docker 
```
docker pull samtholiya/analyser-service:latest
```

## Run

### Local
```
/go/bin/server
```
### Docker
```
docker run -it samtholiya/analyser-service:latest
```

