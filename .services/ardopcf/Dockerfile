FROM debian:bookworm-slim
RUN apt-get update && apt-get install --yes git libasound2-dev build-essential

WORKDIR /src
RUN git clone https://github.com/pflarue/ardop
RUN cd ardop/ARDOPC && make && install ardopcf /usr/local/bin/ardopcf

WORKDIR /run
EXPOSE 8515 8516
ENTRYPOINT ardopcf 8515 $DEVNAME $DEVNAME
