###################################################################################################################################
# builder image
#
FROM golang:latest as builder
WORKDIR /build
COPY . /build

ARG OWNER=dev-rodrigobaliza
ARG PROJECT=dimdim

# accept override of value from --build-args
ARG MY_VERSION=0.0.0
ARG MY_BUILTBY=unknown

# create module, fetch dependencies, then build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${MY_VERSION} -X main.BuiltBy=${MY_BUILTBY}" -o api cmd/api/main.go

###################################################################################################################################
# generate small final image for end users
#
# busybox-glibc (versus Alpine's musl) matches Debian, but that is not a techinical issue here. I simply chose to prefer glibc
# FROM busybox:1.36.1-glibc
FROM scratch

# copy golang binary into container
WORKDIR /app
COPY --from=builder /build/api ./

# executable
ENTRYPOINT [ "./api" ]
