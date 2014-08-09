#!/bin/bash

shopt -s dotglob

for FILE in "/vagrant/provision/dot/*.*"
do
    FILE_NAME=$(basename "$FILE")

    cp -r $FILE /home/vagrant/
    cp -r $FILE /root/

    chown -R vagrant:vagrant /home/vagrant/$FILE_NAME
    chown -R root:root /root/$FILE_NAME
done
