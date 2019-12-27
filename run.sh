#! /bin/bash

# Change
LOGIN="YOUR_LOGIN"
export COOKIE="YOUR_COOKIE"

# Local
ROOT="${PWD}/"
WWW="${ROOT}html/index.html"

# Change values for server
if [ "$HOSTNAME" == "lamp18" ]; then
	echo "Server!"
	ROOT="/SERVER/PATH"
	WWW="/var/www/YOUR_FOLDER/public_html/index.html"
fi

# Reading these in app
export DATA_FOLDER="${ROOT}data/"
export TEMPLATE="${ROOT}html/template.html"
export INDEX="${ROOT}html/index.html"
export USERS="${ROOT}users.json"
# Delay (ms) between requests
export DELAY="1200"
export SHOW_USERNAMES="TRUE"

LOG="${ROOT}nolife.log"
CRAWLER="${ROOT}bin/crawler"
PARSER="${ROOT}bin/parser"

#$CRAWLER | tee /dev/tty > $LOG
date | tee /dev/tty >> $LOG
$PARSER

# Copy index.html only on server
if [ "${INDEX}" != "${WWW}" ]; then
	echo "Copy index..."
	cp $INDEX $WWW
fi
exit 0
