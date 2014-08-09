#!/bin/bash

echo "Install Go SDK"

TMP_SCRIPT="/tmp/install_go.sh"

echo "#!/bin/bash
GVM_INSTALLER=/home/vagrant/gvm-instaler.sh

wget -O \$GVM_INSTALLER https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer > /dev/null 2>&1
chmod 0777 \$GVM_INSTALLER

bash \$GVM_INSTALLER

source /home/vagrant/.gvm/scripts/gvm
gvm install go1.3

gvm use go1.3 --default

for GOOS in linux darwin windows
do
    for GOARCH in amd64 386
    do
        gvm cross \$GOOS \$GOARCH > /dev/null 2>&1
    done
done

gvm cross linux arm > /dev/null 2>&1

rm \$GVM_INSTALLER

export GOPATH=$GOPATH:/home/vagrant/work
" > $TMP_SCRIPT

chmod 0777 $TMP_SCRIPT

sudo -H -u vagrant bash -c $TMP_SCRIPT

rm $TMP_SCRIPT