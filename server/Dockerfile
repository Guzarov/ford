FROM golang:1.15-alpine AS BUILD
COPY . /src
RUN cd /src && go build -o app

FROM alpine
WORKDIR /app
COPY --from=BUILD /src/app /app/
ENV DB_URL mongodb://base:27017
ENV DB_NAME ford
ENTRYPOINT ["/app/app"]
