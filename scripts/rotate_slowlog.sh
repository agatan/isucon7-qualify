#! /bin/bash

set -eu

MYSQL_USER='isucon'
MYSQL_PASSWD='isucon'
MYSQL_HOST='db'

SLOW_LOG_FILE='/tmp/slowlog.log'

ROTATE_SLOW_LOG_DIR="$HOME/slowlog"
ROTATE_SLOW_LOG_FILE="$ROTATE_SLOW_LOG_DIR/slowlog-$(date '+%Y%m%d%H%M%S').log"

LONG_QUERY_TIME=0

if [ ! -z "${MYSQL_PASSWD-}" ]; then
  MYSQL="mysql -u $MYSQL_USER -p $MYSQL_PASSWD -h $MYSQL_HOST"
else
  MYSQL="mysql -u $MYSQL_USER -h $MYSQL_HOST"
fi

rotate() {
  echo 'rotate'

  mkdir -p "$ROTATE_SLOW_LOG_DIR"
  if [ -f $SLOW_LOG_FILE ]; then
    sudo mv "$SLOW_LOG_FILE" "$ROTATE_SLOW_LOG_FILE"
    sudo chown "$(whoami)" "$ROTATE_SLOW_LOG_FILE"
  fi
}

describe_current_conf(){
  echo ''
  echo 'Describe current conf'
  echo '====================='

  eval "$MYSQL --skip-column-names" << EOS
show variables like 'slow_query_log%';
show variables like 'long_query_time';
EOS

  echo '====================='
  echo ''
}

disable_slow_log() {
  echo 'disable slow log'

  eval "$MYSQL" << EOS
set global slow_query_log=0;
EOS
}

enable_slow_log() {
  echo 'enable slow log'

  eval "$MYSQL" << EOS
set global slow_query_log=1;
set global slow_query_log_file='$SLOW_LOG_FILE';
set global long_query_time=$LONG_QUERY_TIME;
EOS
}

describe_current_conf
disable_slow_log
rotate
enable_slow_log
describe_current_conf
