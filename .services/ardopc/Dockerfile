FROM debian:bookworm
RUN apt-get update && apt-get install --yes wget libasound2-dev build-essential unzip

WORKDIR /src
RUN wget -q https://www.cantab.net/users/john.wiseman/Downloads/Beta/ARDOPProjects.zip && unzip ARDOPProjects.zip
RUN cd ARDOPProjects/ARDOPC && make

WORKDIR /run
ENTRYPOINT /src/ARDOPProjects/ARDOPC/ardopc 8515 $DEVNAME $DEVNAME --leaderlength=$LEADERLENGTH -L
