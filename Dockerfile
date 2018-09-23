FROM golang:1.10.3

# hajimehoshi/ebiten dependencies: https://github.com/hajimehoshi/ebiten/wiki/Linux
RUN apt update && \
    apt install -y \
    libglu1-mesa-dev \
    libgles2-mesa-dev \
    libxrandr-dev \
    libxcursor-dev \
    libxinerama-dev \
    libxi-dev \
    libasound2-dev

