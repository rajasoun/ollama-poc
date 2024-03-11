#!/usr/bin/env bash

# Define the base directory for packages and source the library script
BASEDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)
source $BASEDIR/scripts/lib/bootstrap.sh

bootstrap_main
