package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, `Usage: 

%s osm-file latitude longitude max-distance

    eg. %s abc.osm 18.485493 -69.88211 25

The unit for maximum distance is km. 
`, os.Args[0], os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	if !fileExists(filename) {
		fmt.Fprintf(os.Stderr, "File does not exist: %s\n", os.Args[1])
		os.Exit(1)
	}

	lat, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Illegal value for latitude: %s\n", os.Args[2])
		os.Exit(1)
	}

	lon, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Illegal value for longitude: %s\n", os.Args[3])
		os.Exit(1)
	}

	dist, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Illegal value for max-distance: %s\n", os.Args[4])
		os.Exit(1)
	}

	// everything checks out, let's get started!

	// make a datachannel on which to receive the nodes
	dataChan := make(chan XNode)
	// launch a goroutine to process the file, and put the nodes on the chan
	go pull_nodes(filename, dataChan)

	for {
		n := <-dataChan
		if n.EndOfStream {
			break
		}
		rd := rough_distance(lat, lon, n.Lat, n.Lon)
		if rd < dist {
			fmt.Printf("%f,%f,%q,\"#\", %f, \"%s\"\n", n.Lat, n.Lon, stripComma(getName(n.XTags)), rd, stripComma(fmt.Sprintf("%v", n.XTags)))
		}
	}
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func stripComma(in string) string {
	return strings.Replace(in, ",", "|", -1)
}

func getName(xtsl []XTag) string {
	if len(xtsl) == 0 {
		return "NO_TAGS_FOUND"
	}
	for _, xt := range xtsl {
		if xt.K == "name" {
			return xt.V
		}
	}
	// it has tags, but no name, just return the first tag
	xt := xtsl[0]
	return fmt.Sprintf("%s %s", xt.K, xt.V)
	// it has tags more than 1 tags, but no name, let's just return the lot...
	//return fmt.Sprintf("%v",xtsl)
}

// define the XML structures used
type XTag struct {
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

type XNode struct {
	XMLName     xml.Name `xml:"node"`
	IdStr       string   `xml:"id,attr"`
	LatStr      string   `xml:"lat,attr"`
	LonStr      string   `xml:"lon,attr"`
	XTags       []XTag   `xml:"tag"`
	Id          int64
	Lat         float64
	Lon         float64
	EndOfStream bool
}

func pull_nodes(filename string, datachan chan XNode) {
	// read the file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	reNodeInline := regexp.MustCompile(`<node.[^>]*?/>`)
	reNode := regexp.MustCompile(`(?s)<node.*</node>`) // flag (?s): set the multiline

	var buf bytes.Buffer
	var line []byte
	for {
		line, err = reader.ReadBytes('\n')
		buf.Write(line)
		for buf.Len() > 0 {
			bufb := buf.Bytes()
			loc := reNodeInline.FindIndex(bufb)
			if loc != nil {
				// fmt.Printf("-----\nINLINE: %s\n", bufb[loc[0]:loc[1]])
				//IGNORE INLINE NODES
				buf.Reset()
				buf.Write(bufb[loc[1]:])
				continue
			}
			loc = reNode.FindIndex(bufb)
			if loc != nil {
				//fmt.Printf("-----\nREGULAR: %s\n", bufb[loc[0]:loc[1]])

				//handleNode( bufb[loc[0]:loc[1]] )
				// create the node, and put it on the chan
				n := XNode{EndOfStream: false}
				err := xml.Unmarshal(bufb[loc[0]:loc[1]], &n)
				if err == nil {
					n.Id, _ = strconv.ParseInt(n.IdStr, 10, 64) // check err?
					n.Lat, _ = strconv.ParseFloat(n.LatStr, 64)
					n.Lon, _ = strconv.ParseFloat(n.LonStr, 64)

					// okay, communicate it to the outside world
					datachan <- n
				} else {
					fmt.Printf("error: %v", err)
				}
				// carry on reading the file
				buf.Reset()
				buf.Write(bufb[loc[1]:])
				continue
			}
			if bytes.Index(bufb, []byte("<node")) == -1 {
				// there is nothing 'node' in the buffer so let's reset it
				buf.Reset()
			}
			break // nothing found, so get out
		}
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	datachan <- XNode{EndOfStream: true} // we are done!
}

/* from: http://www.movable-type.co.uk/scripts/latlong.html
   note: φ=lat λ=lon  in RADIANS!
   var x = (λ2-λ1) * Math.cos((φ1+φ2)/2);
   var y = (φ2-φ1);
   var d = Math.sqrt(x*x + y*y) * R;
*/

func rough_distance(lat1, lon1, lat2, lon2 float64) float64 {

	lat1 = lat1 * math.Pi / 180.0
	lon1 = lon1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0
	lon2 = lon2 * math.Pi / 180.0

	// convert to radians
	r := 6371. // km
	x := (lon2 - lon1) * math.Cos((lat1+lat2)/2)
	y := (lat2 - lat1)
	d := math.Sqrt(x*x+y*y) * r
	return d
}
