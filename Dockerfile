FROM golang:1.16.5 AS build

LABEL owner=digital
LABEL os=alpine
LABEL description='Run-time image for the GKE Secret Service'

# set the working directory
WORKDIR /build
RUN useradd -u 10001 scratchuser


COPY . .

#RUN mkdir etc && \
 #groupadd --gid 10001 myapp \
    #&& useradd --gid 10001 --uid 10001 myapp && \
    #echo 'nobody:x:65534:65534:nobody:/:' > etc/passwd && \
    #echo 'nobody:x:65534:' > etc/group && \
    #lsgo test -v ./... && \
    RUN CGO_ENABLED=0 go build -o logging-service ./cmd/logging-service/* && \
    #CGO_ENABLED=0 go build -ldflags='-extldflags=-static' -o logging-service ./cmd/logging-service/* && \
    chmod +777 logging-service && \
    chmod +777 build/packages/test2.yaml && \
    chmod +777 build/packages/testuserKeyfile.json

#FROM golang:1.16.5 AS final
#FROM gcr.io/distroless/base
#FROM gcr.io/distroless/static
FROM scratch

#COPY --from=build /tmp /tmp
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/build/packages/test2.yaml /etc/
#COPY --chown=65534:65534 --from=build /build/build/packages/testuserKeyfile.json /etc/
COPY --from=build /build/build/packages/testuserKeyfile.json /etc/
COPY --from=build /build/logging-service /usr/local/bin/logging-service
#COPY --from=build /build/etc/group /build/etc/passwd /etc/
COPY --from=0 /etc/passwd /etc/passwd
EXPOSE 9701

USER scratchuser
#USER nobody:nobody

ENTRYPOINT ["/usr/local/bin/logging-service"]
