#!/bin/sh
# export info about last processed block first
mongoexport --db=artion --collection=status --out status.json

# export scanned data next
mongoexport --db=artion --collection=activities --out activities.json
mongoexport --db=artion --collection=auctions --out auctions.json
mongoexport --db=artion --collection=bids --out bids.json
mongoexport --db=artion --collection=collections --out collections.json
mongoexport --db=artion --collection=listings --out listings.json
mongoexport --db=artion --collection=observed --out observed.json
mongoexport --db=artion --collection=offers --out offers.json
mongoexport --db=artion --collection=ownerships --out ownerships.json
mongoexport --db=artion --collection=tokens --out tokens.json
mongoexport --db=artion --collection=burns --out burns.json

# export notification data as the last
mongoexport --db=artion --collection=notification_tpl --out notification_tpl.json
mongoexport --db=artion --collection=notifications --out notifications.json
