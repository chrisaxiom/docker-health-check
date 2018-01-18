# docker-health-check
A tiny health check tool for docker

**Usage**

`docker-health-check -U=http://www.google.com:80`
`docker-health-check -h` for help

**Docker**

`Dockerfile_build` can be used to output a tar holding the binary to be embedded into another container if you want to build it yourself.  Or you can just link to it directly in your own `Dcokerfile`:

```
FROM scratch
ADD https://github.com/chrisaxiom/docker-health-check/blob/master/docker-health-check?raw=true /docker-health-check
HEALTHCHECK --interval=8s --timeout=120s --retries=8 CMD ["/docker-health-check", "-url=http://127.0.0.1:8000/api/ping"]
...
```

**Not Supported (Yet)**

- Response body checking
- Request body


