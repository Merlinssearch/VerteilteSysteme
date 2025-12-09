FROM alpine:latest

COPY target/robot /usr/local/bin/robot
RUN chmod +x /usr/local/bin/robot

EXPOSE 3000

ENTRYPOINT ["robot"]
