FROM alpine
RUN apk add --no-cache gpsd
EXPOSE 2947
ENTRYPOINT ["gpsd", "-N", "-S", "2947", "-G", "/dev/gps"]
