#!/bin/bash

set -e

sudo -u cj git fetch --all
sudo -u cj git checkout --force "origin/master"
sudo -u cj docker-compose up -d --force-recreate --build

