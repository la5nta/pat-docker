FROM debian:bookworm-slim
RUN apt-get update && apt-get install --yes git libasound2-dev build-essential

WORKDIR /src
RUN git clone https://github.com/pflarue/ardop
RUN cd ardop/ARDOPC && make && install ardopc /usr/local/bin/ardopc

WORKDIR /run
EXPOSE 8515 8516
ENTRYPOINT ardopc 8515 $DEVNAME $DEVNAME
