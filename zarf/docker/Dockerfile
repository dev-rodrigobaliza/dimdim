#
# builder image
#
FROM golang:latest as builder
RUN mkdir /build
ADD * /build/
WORKDIR /build

ARG OWNER=dev-rodrigobaliza
ARG PROJECT=dimdim

# accept override of value from --build-args
ARG MY_VERSION=0.0.0
ARG MY_BUILTBY=unknown

# create module, fetch dependencies, then build
RUN go mod init ${OWNER}/${PROJECT} \
   && go mod tidy \
   && CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${MY_VERSION} -X main.BuiltBy=${MY_BUILTBY}" main.go

#
# generate small final image for end users
#
# busybox-glibc (versus Alpine's musl) matches Debian, but that is not a techinical issue here. I simply chose to prefer glibc
FROM busybox:1.36.1-glibc

# copy golang binary into container
WORKDIR /root
COPY --from=builder /build/main .

# executable
ENTRYPOINT [ "./main" ]