# Stage 1: Build voacapl from source in Alpine
FROM alpine AS voacap-builder
RUN apk add --no-cache git gcc gfortran make autoconf automake libtool musl-dev
WORKDIR /src
RUN git clone https://github.com/jawatson/voacapl.git .
RUN autoreconf -fi && ./configure && make && make install

# Stage 2: Build Go API
FROM golang:alpine AS go-builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o voacap-api

# Stage 3: Minimal runtime with Alpine
FROM alpine
RUN apk add --no-cache ca-certificates libgfortran
COPY --from=voacap-builder /usr/local/bin/voacapl /bin/
COPY --from=voacap-builder /usr/local/share/voacapl/itshfbc /itshfbc
COPY --from=go-builder /src/voacap-api /bin/voacap-api
RUN adduser -D -s /bin/sh voacap
USER voacap
WORKDIR /home/voacap
ENV VOACAP_DATA_DIR="/itshfbc"
ENV VOACAP_ADDR=":8080"
EXPOSE 8080
CMD ["/bin/voacap-api"]
