# DRAFT STATUS!!!

# How to 

## Compile and build pull.go 

That's easy: 

    go build pull_osm.go


## Download an osm file

Download an .osm file (not a .osm.pbf file), for example download `central-america-latest.osm.bz2` from [download.geofabrik.de](http://download.geofabrik.de).  

Coffee time: in December 2014 this file was 354 MB in size, it may take a while to download it, depending on your internet connection. 

Download ready? Then unzip the file: 

    bunzip2 central-america-latest.osm.bz2

Coffee time again: this file unzips to 4.8G or more, which also takes awhile, so grab a coffee, but first check your disk-space ..


## Run your query


Get all the points in a 10km radius around Willemstad, Curacao:

    ./pull central-america-latest.osm 12.1166 -68.9333 10  > willemstad10k.csv

Coffee time again: scanning this file of nearly 5G takes a while, or what did you expect. So go head and grab a final coffee?

Before proceeding to the next step put on some instrumental music like the great Anoushka Shankar's debut album, which will assist you in concentrating! 

After this was finished, I ended up with 3544 lines of csv data: 

    wc -l willemstad10k.csv 
    3544 willemstad10k.csv


## Interactive querying using IPython & pandas 

Hopefully you are completely maxed out on coffee now, because it's time to filter your data further *interactively*! 












