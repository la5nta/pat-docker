# Using Debian Bullseye because the VARA installer segfaults with newer WINE versions
FROM debian:bullseye-slim as installer
RUN apt-get -qq update && apt-get install --yes curl wget libasound2-dev xvfb xterm gettext-base p7zip-full cabextract unzip
RUN dpkg --add-architecture i386 && apt-get -qq update && apt-get install --yes wine32

ENV WINEARCH=win32
ENV WINEPREFIX /wine
RUN xvfb-run wine wineboot && wineserver -k

WORKDIR /src/winetricks
ARG WINETRICKS_TAG=20240105
RUN wget -q https://raw.githubusercontent.com/Winetricks/winetricks/${WINETRICKS_TAG}/src/winetricks
RUN xvfb-run sh winetricks -q vb6run pdh_nt4 sound=alsa && wineserver -k

WORKDIR /src/ahk
ARG AHK_VERSION=1.1.36.02
RUN wget -q https://github.com/AutoHotkey/AutoHotkey/releases/download/v${AHK_VERSION}/AutoHotkey_${AHK_VERSION}_setup.exe
RUN 7z e AutoHotkey_${AHK_VERSION}_setup.exe AutoHotkeyU32.exe -y -bsp0 -bso0
ADD install.ahk install.ahk

WORKDIR /src/vara
ARG VARA_PRODUCTS_URL=https://downloads.winlink.org/VARA%20Products/
# Bust cache if the source URL is updated
ADD ${VARA_PRODUCTS_URL} index.html
RUN wget --no-parent -r -nd -l1 -A 'VARA HF*.zip,VARA FM*.zip' ${VARA_PRODUCTS_URL}
RUN 7z e VARA\ HF*.zip && 7z e VARA\ FM*.zip
RUN stat "VARA setup (Run as Administrator).exe" > /dev/null
RUN stat "VARA FM setup (Run as Administrator).exe" > /dev/null
RUN xvfb-run wine /src/ahk/AutoHotkeyU32.exe /src/ahk/install.ahk "VARA setup (Run as Administrator).exe" && wineserver -k
RUN stat /wine/drive_c/VARA/VARA.exe > /dev/null
RUN xvfb-run wine /src/ahk/AutoHotkeyU32.exe /src/ahk/install.ahk "VARA FM setup (Run as Administrator).exe" && wineserver -k
RUN stat /wine/drive_c/VARA\ FM/VARAFM.exe > /dev/null

# Using Debian Bookworm for the runtime because we can
FROM debian:bookworm-slim
RUN apt-get update && apt-get install --yes libasound2-dev gettext-base xvfb xterm
RUN dpkg --add-architecture i386 && apt-get update && apt-get install --yes wine32
COPY --from=installer /wine /wine
ENV WINEDEBUG=-all
ENV WINEARCH=win32
ENV WINEPREFIX /wine
ADD ["./VARAFM.ini", "/wine/drive_c/VARA FM/VARAFM.ini"]
ADD ["./VARAHF.ini", "/wine/drive_c/VARA/VARA.ini"]
ADD varahf.sh /bin/varahf
ADD varafm.sh /bin/varafm
RUN chown root:root -R /wine
ADD asoundrc.subst .
