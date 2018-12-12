FROM golang:1.10.3

# veandco/go-sdl2: https://github.com/veandco/go-sdl2/#requirements
RUN apt update && \
    apt install -y \
    libsdl2-image-dev \
    libsdl2-mixer-dev \
    libsdl2-ttf-dev \
    libsdl2-gfx-dev
