# docker-health-check
A tiny health check tool for docker

**Usage**

`docker-health-check -U=http://www.google.com:80`
`docker-health-check -h` for help

`Dockerfile` can be used to output a tar holding the binary:

```
docker build .
docker run --rm <container_id> > hc.tar.gz
gunzip -c hc.tar.gz | tar xopf -
```

**Integration**

Two ways to incorporate the healthcheck into your `Dockerfile`.

Direct:

```
FROM ubuntu:latest as hc
ADD https://github.com/chrisaxiom/docker-health-check/blob/master/docker-health-check?raw=true /docker-health-check
RUN chmod a+x /docker-health-check
# your container
FROM scratch
COPY --from=hc /docker-health-check /docker-health-check
HEALTHCHECK --interval=8s --timeout=120s --retries=8 CMD ["/docker-health-check", "-url=http://127.0.0.1:8000/api/ping"]
...
```

Versioned (recommended):

```
FROM ubuntu:latest as hc
ADD https://github.com/chrisaxiom/docker-health-check/archive/v0.3.tar.gz /
RUN tar -xvzf /v0.3.tar.gz
# your container
FROM scratch
COPY --from=hc /docker-health-check-0.3/docker-health-check /docker-health-check
HEALTHCHECK --interval=8s --timeout=120s --retries=8 CMD ["/docker-health-check", "-url=http://127.0.0.1:8000/api/ping"]
...
```

**Not Supported (Yet)**

- Response body checking
- Request body


