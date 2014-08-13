#!/bin/bash

echo "Configuration vim"

TMP_SCRIPT="/tmp/vim_configuration.sh"

echo "#!/bin/bash

BUNDLE_PATH=/home/vagrant/.vim/bundle

mkdir -p \$BUNDLE_PATH

git clone https://github.com/gmarik/Vundle.vim.git \$BUNDLE_PATH/Vundle.vim

vim +BundleInstall +qall > /dev/null 2>&1

bash \$BUNDLE_PATH/YouCompleteMe/install.sh > /dev/null 2>&1

echo \"export GOPATH=$GOPATH:/home/vagrant/work\" >> /home/vagrant/.bashrc
" > $TMP_SCRIPT

chmod 0777 $TMP_SCRIPT

sudo -H -u vagrant bash -c $TMP_SCRIPT

rm $TMP_SCRIPT