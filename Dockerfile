FROM golang:latest AS build
# install file
RUN apt-get update -y && apt-get install -y file && apt-get install -y upx
# add source files
COPY . /go/src/github.com/chrisaxiom/docker-health-check
# set workdir
WORKDIR /go/src/github.com/chrisaxiom/docker-health-check
# build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags="-s -w"
# do static compilation check
RUN ls -lah ./docker-health-check && file ./docker-health-check
# upx
RUN upx ./docker-health-check
# do static compilation check after
RUN ls -lah ./docker-health-check && file ./docker-health-check
# tar up result
RUN tar cfvzp /tmp/app.tar.gz ./docker-health-check
# entrypoint
ENTRYPOINT ["cat", "/tmp/app.tar.gz"]