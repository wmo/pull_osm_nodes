# DRAFT STATUS!!!

# How to 

## Compile and build pull.go 


## Download an osm file

Download an .osm file (not a .osm.pbf file), for example download `central-america-latest.osm.bz2` from [download.geofabrik.de](http://download.geofabrik.de).  

Coffee time: in December 2014 this file was 354 MB in size, it may take a while to download it, depending on your internet connection. 

Download ready? Then unzip the file: 

    bunzip2 central-america-latest.osm.bz2

Coffee time again: this file unzips to 4.8G or more, which also takes awhile, so grab a coffee, but first check your disk-space ..


## Run your query


Get all the points in a 10km radius around Willemstad, Curacao:

    ./pull central-america-latest.osm 12.1166 -68.9333 10  > willemstad10k.csv

Coffee time again: this file unzips to 4.8G or more, so first check if you have enough disk-space. Also, this process takes a couple of minutes, so why not grab a final coffee?


## Interactive queruing using IPython & pandas 

Hopefully you've have topped up on coffee, because now it's time to to filter your data further interactively! 









