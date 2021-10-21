#!/bin/sh
cp ipfs.service /etc/systemd/system/
cp artion.service /etc/systemd/system/
systemctl enable ipfs.service
systemctl enable artion.service
