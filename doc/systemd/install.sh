#!/bin/sh
cp opera.service /etc/systemd/system/
cp ipfs.service /etc/systemd/system/
cp artion.service /etc/systemd/system/

sudo systemctl daemon-reload

systemctl enable opera.service
systemctl enable ipfs.service
systemctl enable artion.service
