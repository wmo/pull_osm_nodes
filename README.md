# What? 

Before I go on holiday, I like to gather the OpenStreetMap waypoints for the area I'll be staying at. I may not have online access while there, so I prefer to store this information as a file on my laptop.

Usually the .osm files are big blobs of data, and it takes a while to pull the requested data out. So in a first step I pull out every node in a radius of say 25km in one go, and after that post-process (filter, subset,... ) the resulting information with Pandas.

These are the steps to take:

* 1. clone this repo and compile `pull.go` 
* 2. download an .osm file
* 3. pull out all the nodes of the area of interest
* 4. post-process the resulting .csv with Python Pandas  

Prerequisites: 

- have a `go` compiler installed on your computer
- have a working installation of python and the pandas library



## Compile and build pull.go 

Clone this repo and compile the go file. It's this easy: 

    git clone https://github.com/wmo/pull_osm_nodes
    cd pull_osm_nodes
    go build pull.go


## Download an osm file

Download an .osm file (*not* an .osm.pbf file), for example download `central-america-latest.osm.bz2` from [download.geofabrik.de](http://download.geofabrik.de).  

Coffee time: this file may take a while to download: in December 2014 this file was 354 MB in size. Depends on your internet connection, of course.

Download ready? Then unzip the file: 

    bunzip2 central-america-latest.osm.bz2

Coffee time again: this file unzips to 4.8G or more, which also takes a while, so grab another coffee, but first check your disk-space ..



## Pull out the nodes

Get all the points in a 10km radius around Willemstad, Curacao:

    ./pull central-america-latest.osm 12.1166 -68.9333 10  > willemstad10k.csv

Coffee time again: scanning this file of nearly 5G takes a while, or what did you expect? Yes, go ahead and grab a final coffee!

Before proceeding to the next step, put on some instrumental music like the great Anoushka Shankar's debut album, which will be very good for your concentration. 

After running the above command, I ended up with 3544 lines of csv data: 

    wc -l willemstad10k.csv 
    3544 willemstad10k.csv


## Interactive querying using IPython & pandas 

Hopefully you are completely maxed out on coffee now, because it's time to further filter your data *interactively*! 

(coming soon)









