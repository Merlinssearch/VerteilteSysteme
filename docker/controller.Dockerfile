FROM alpine:latest

COPY target/controller /usr/local/bin/controller
RUN chmod +x /usr/local/bin/controller

EXPOSE 3000

ENTRYPOINT ["controller"]
