#! /bin/bash

set -eu

MYSQL_CONF_DIR=/etc/mysql
NGINX_CONF_DIR=/etc/nginx

mkdir -p $HOME/$MYSQL_CONF_DIR
mkdir -p $HOME/$NGINX_CONF_DIR

TARGETS=(
        /etc/sysctl.conf
        /etc/mysql/my.cnf
        /etc/mysql/conf.d
        /etc/nginx/nginx.conf
        /etc/nginx/conf.d
        /etc/nginx/sites-enabled
        /etc/nginx/sites-available
)

for target in ${TARGETS[@]}
do
  if [[ ! -e $target.backup ]]; then
    sudo cp -r $target $target.backup
    sudo mv $target $HOME/$target
    # sudo chmod -R a+rw $HOME/$target
    sudo ln -s $HOME/$target $target
  fi
done
