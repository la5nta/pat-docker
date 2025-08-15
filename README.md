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
* [varahf](https://rosmodem.wordpress.com) (VARA HF modem, propriatery)
* [varafm](https://rosmodem.wordpress.com) (VARA FM modem, propriatery)
* [rigctld](https://github.com/Hamlib/Hamlib) (Rig control)
* [voacapl](https://github.com/jawatson/voacapl) (HF propagation prediction)

## Requirements

* Docker + Docker Compose
* Linux is recommended when using software modems

## Quick start guide

### Starting and stopping services

```
# Start Pat and ARDOP (in detached/background mode)
docker-compose up -d pat ardop

# Stream log output of all running containers
docker-compose logs -f

# Print status of all running/stopped containers
docker-compose ps -a

# Stop running containers
docker-compose stop
```

Run `docker-compose help` for more details.

### Updating images (app installations)

```
# Pull hosted images (i.e. Pat)
docker-compose pull

# Re-build unhosted images (from source)
docker-compose build
```

### Configuration

The most convenient method to configure Pat is though the web gui. For more
advanced configuration, edit docker-compose.yml to suite your needs. You'll
probably need to change `DEVNAME` to specify which soundcard to use with
ardopc/direwolf/vara.

The compose file also provides an example for rigctld, configured in Pat as a
rig named "dummy". You'll probably want to rename it, change the command
arguments to match your hardware, and add a devices section to mount the
correct tty device.
