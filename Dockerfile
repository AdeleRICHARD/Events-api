FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY  events-api /app/events-api
RUN apk --no-cache add curl \
    libc6-compat
CMD ["./events-api"]