FROM golang:alpine AS build-env
ENV CGO_ENABLED 1
ADD ./ /go/src/github.com/essial/AthenaEngine
RUN apk add --no-cache git
RUN go get github.com/essial/AthenaEngine
RUN go build -o /athenaengine_server github.com/essial/AthenaEngine/cmd/server

FROM alpine
EXPOSE 3000
WORKDIR /
COPY --from=build-env /athenaengine_server /
COPY --from=build-env /go/src/github.com/essial/AthenaEngine/aehttp /aehttp
CMD ["/athenaengine_server"]
