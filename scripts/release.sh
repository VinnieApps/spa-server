#!/bin/bash

set -e
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

./semantic-release -travis-com --token $GITHUB_TOKEN --slug VinnieApps/spa-server --ghr --vf
export VERSION=$(cat .version)

$SCRIPT_DIR/package.sh
ghr $(cat .ghr) build/packages
