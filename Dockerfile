FROM golang:alpine AS build

ARG RELEASE_VERSION="dev"
ARG RELEASE_GIT_COMMIT="build"

WORKDIR /build
COPY . /build

RUN go build \
  -ldflags "-X github.com/taskmedia/nuCal/pkg/http/rest.version=${RELEASE_VERSION}-${RELEASE_GIT_COMMIT}" \
  cmd/nuCal/nuCal.go

FROM alpine

RUN addgroup -S tm && \
  adduser -S tm -G tm

COPY --from=build /build/nuCal /bin/nuCal

USER tm
EXPOSE 8080

ENTRYPOINT [ "/bin/nuCal" ]
