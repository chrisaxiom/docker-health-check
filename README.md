# docker-health-check
A tiny health check tool for docker

**Usage**

`docker-health-check -U=http://www.google.com:80`
`docker-health-check -h` for help

**Docker**

`Dockerfile_build` can be used to output a tar holding the binary to be embedded into another container if you want to build it yourself.  Or you can just link to it directly in your own `Dcokerfile`:

```
FROM scratch
ADD https://github.com/chrisaxiom/docker-health-check /docker-health-check
...
```

**Not Supported (Yet)**

- Response body checking
- Request body


