package main


/* A commandline tool that converts times to different timezones

Steps to take:
[] Parse commandline args
    [] pass in a time and list of locations and get that time in those locations
        [] pass in the timestamp
        # TODO Current stuck here parsing a datetime string on commandline is a pain here!
        [] take a default as now
    [] will need to map the locations to timezone abbreviations
[] 
*/

import (
        "flag"
        "fmt"
        "time"
)

func main() {

    isoStr := "2006-01-02T03:04"
    nowStr := time.Now().Format(isoStr)
        
    var inTime string
    flag.StringVar(&inTime, "time", nowStr, "a time")
    flag.Parse()

    fmt.Println(inTime)
   
    myTime, err := time.Parse(isoStr, inTime)
	if err != nil {
		panic(err)
	}

    timezonesToCheck := flag.Args()

    for _, element := range timezonesToCheck {
            location, err := time.LoadLocation(element)
            if err != nil {
                panic(err)
            }
	        fmt.Println(element, myTime.In(location).Format(isoStr))
        }

}
