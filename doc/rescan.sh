#!/bin/sh
echo "Starting re-scanning from block $1, reason: $2"
version=`/home/opera/go/src/Artion-API-GraphQL/build/artionapi -v 2>/dev/null | grep "Commit Hash:"`
time=`date --iso-8601=seconds`
log="$time Scanning from block $1 ($version) reason: $2"
echo "Log: $log"

systemctl stop artion.service
echo "$log" >> /var/log/artion-scan.log
mongo --eval "db.status.update({ _id: \"lastSeenBlock\" }, { \$set: { block: NumberLong($1) } });" artion
systemctl start artion.service
systemctl status artion.service

