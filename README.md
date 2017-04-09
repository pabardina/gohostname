Build
====

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  .
docker build -t pabardina/gohostname .
```

Run
===

`docker run --rm -it -p 8000:8000 pabardina/gohostname`
