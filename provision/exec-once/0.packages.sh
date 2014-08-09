#!/bin/sh

echo "Install basic packages"

apt-get update > /dev/null 2>&1
apt-get install -y vim mc htop \
    curl git mercurial \
    cmake autoconf automake build-essential \
    libpython-dev \
    binutils bison > /dev/null 2>&1

