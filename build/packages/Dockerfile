FROM golang:1.16.5 AS build

LABEL owner=digital
LABEL os=alpine
LABEL description='Run-time image for the GKE Secret Service'

# set the working directory
WORKDIR /build

# credentials for package management via eterprise Artifactory.
#ARG ARTIFACTORY_USERNAME_READONLY
#ARG ARTIFACTORY_PASSWORD_READONLY

# ensure package management goes through Artifactory
#ENV GOPROXY="https://${ARTIFACTORY_USERNAME_READONLY}:${ARTIFACTORY_PASSWORD_READONLY}@af.cds.bns/artifactory/api/go/virtual-go-gwrt"

#COPY certs /etc/ssl/certs
#COPY apk/repositories /etc/apk/repositories
#COPY ./build/packages/test2.yaml /
COPY . .

RUN mkdir etc && \
    echo 'nobody:x:65534:65534:nobody:/:' > etc/passwd && \
    echo 'nobody:x:65534:' > etc/group && \
    #lsgo test -v ./... && \
    CGO_ENABLED=0 go build -o logging-service ./cmd/logging-service/* && \
    chmod +x logging-service

FROM scratch

COPY --from=build /build/build/packages/test2.yaml /etc/
COPY --chown=65534:65534 --from=build /build/build/packages/testuserKeyfile.json /etc/
COPY --from=build /build/logging-service /usr/local/bin/logging-service
COPY --from=build /build/etc/group /build/etc/passwd /etc/

EXPOSE 9701

USER nobody:nobody

ENTRYPOINT ["/usr/local/bin/logging-service"]
