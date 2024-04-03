FROM debian:bookworm-slim
RUN apt-get update && apt-get install --yes git libasound2-dev build-essential

WORKDIR /src
# Bust cache if the master branch is updated
ADD https://api.github.com/repos/pflarue/ardop/commits?per_page=1 latest_commit
RUN git clone https://github.com/pflarue/ardop
RUN cd ardop/ARDOPC && make && install ardopcf /usr/local/bin/ardopcf

WORKDIR /run
EXPOSE 8515 8516
ENTRYPOINT ardopcf 8515 $DEVNAME $DEVNAME
