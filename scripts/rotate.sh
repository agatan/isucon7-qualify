#! /bin/bash

set -eu

SCRIPTS_DIR=$(dirname "$0")

echo '>>> Rotating alp log... <<<'
$SCRIPTS_DIR/rotate_alplog.sh
echo '>>> Rotating alp log... Success! <<<'
echo ''
echo '>>> Rotating slow log... <<<'
$SCRIPTS_DIR/rotate_slowlog.sh
echo '>>> Rotating slow log... Success!<<<'
