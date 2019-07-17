#!/usr/bin/bash -x

LOGLEVEL=${LOGLEVEL:-debug}

bash bindata.sh

go run . serve --addr :1234 --verbosity ${LOGLEVEL} --metrics ./example/{default_site,127.0.0.1} 