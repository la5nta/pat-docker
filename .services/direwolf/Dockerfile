FROM debian:bookworm-slim
RUN apt-get update && apt-get install --yes git libasound2-dev build-essential
RUN apt-get install --yes libudev-dev cmake gettext-base

WORKDIR /src
# Bust cache if the master branch is updated
ADD https://api.github.com/repos/wb2osz/direwolf/commits?per_page=1 latest_commit
RUN git clone https://www.github.com/wb2osz/direwolf
RUN cd direwolf && mkdir build && cd build && cmake .. && make -j4

WORKDIR /run
ADD direwolf.conf.subst direwolf.conf.subst
ENTRYPOINT /bin/sh -c "envsubst < direwolf.conf.subst > direwolf.conf; /src/direwolf/build/src/direwolf"
