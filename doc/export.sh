#!/bin/sh
mongoexport --db=artion --collection=activities --out activities.json
mongoexport --db=artion --collection=auctions --out auctions.json
mongoexport --db=artion --collection=bids --out bids.json
mongoexport --db=artion --collection=collections --out collections.json
mongoexport --db=artion --collection=listings --out listings.json
mongoexport --db=artion --collection=observed --out observed.json
mongoexport --db=artion --collection=offers --out offers.json
mongoexport --db=artion --collection=ownerships --out ownerships.json
mongoexport --db=artion --collection=status --out status.json
mongoexport --db=artion --collection=tokens --out tokens.json

