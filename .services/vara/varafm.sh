#!/bin/sh
envsubst < asoundrc.subst > ~/.asoundrc
tail -F -n0 /wine/drive_c/VARA\ FM/VARAFM.log &
rm -f /tmp/.X*-lock # Work-around for dirty termination of Xvfb
wineserver -k; xvfb-run wine /wine/drive_c/VARA\ FM/VARAFM.exe
