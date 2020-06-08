FROM golang:alpine AS build-env
ENV CGO_ENABLED 0
ADD . /go/src/athena_engine
RUN go build -o /server athena_engine

FROM alpine
EXPOSE 5432
WORKDIR /
COPY --from=build-env /server /
CMD ["/server"]
