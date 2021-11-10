#!/bin/sh
mongoimport --db=artion --collection=accounts --file accounts.json
mongoimport --db=artion --collection=activities --file activities.json
mongoimport --db=artion --collection=auctions --file auctions.json
mongoimport --db=artion --collection=bids --file bids.json
mongoimport --db=artion --collection=collections --file collections.json
mongoimport --db=artion --collection=listings --file listings.json
mongoimport --db=artion --collection=observed --file observed.json
mongoimport --db=artion --collection=offers --file offers.json
mongoimport --db=artion --collection=ownerships --file ownerships.json
mongoimport --db=artion --collection=status --file status.json
mongoimport --db=artion --collection=tokens --file tokens.json

