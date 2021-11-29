#!/bin/sh
if [ $# -eq 0 ]; then
    echo "This script start re-scanning of Artion database"
    echo "Use as: sudo $0 (block-number) (reason-for-log)"
    exit 1
fi
if [ $USER != "root" ]; then
    echo "Script must be run as root"
    exit 1
fi

version=`/home/opera/go/src/Artion-API-GraphQL/build/artionapi -v 2>/dev/null | grep "Commit Hash:"`
time=`date --iso-8601=seconds`
log="$time Scanning from block $1 ($version) reason: $@"
echo "Log: $log"

systemctl stop artion.service
echo "$log" >> /var/log/artion-scan.log
mongo --eval "db.status.update({ _id: \"lastSeenBlock\" }, { \$set: { block: NumberLong($1) } });" artion
systemctl start artion.service
systemctl status artion.service
exit 0

