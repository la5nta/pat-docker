# Run Pat with Docker Compose

The purpose of this repository is to serve as a building block for a
[containerized](https://en.wikipedia.org/wiki/Containerization_(computing))
Winlink station, deployed using [Docker
Compose](https://docs.docker.com/compose/).

The primary benefit of such a deployment is the packaging and isolation of
required software components, so they can run independent of the host's
operating system and system libraries. Docker Compose also handles
[orchestration](https://en.wikipedia.org/wiki/Orchestration_(computing)),
making it easier to control the life-cycle of the applications.

## Available services

* [pat](https://getpat.io) (Winlink client)
* [ardopcf](https://github.com/pflarue/ardop) (ARDOP modem)
* [direwolf](https://github.com/wb2osz/direwolf) (Packet modem)
* [rigctld](https://github.com/Hamlib/Hamlib) (Rig control)

## Requirements

* Docker + Docker Compose
* Linux is recommended when using software modems

## Quick start guide 

### Starting and stopping services

See `docker-compose up --help`.

### Updating the docker images

```
# Pull hosted images (i.e. Pat)
docker-compose pull

# Re-build unhosted images (from source)
docker-compose build --pull --no-cache
```

### Configuration

Edit docker-compose.yml to suite your needs. As a bare minimum you should set
the environment variables `PAT_MYCALL` and `PAT_LOCATOR`, and possibly
`DEVNAME` to specify which soundcard to use with ardopc/direwolf.

The compose file also provides an example for rigctld, configured in Pat as a
rig named "dummy". You'll probably want to rename it, change the command
arguments to match your hardware, and add a devices section to mount the
correct tty device.
