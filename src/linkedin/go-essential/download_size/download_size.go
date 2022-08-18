/* Calculate total download size for NYC taxi data for 2021

For each month, we have two files: green and yellow. For example:

    https://d37ci6vzurychx.cloudfront.net/trip-data/green_tripdata_2021-03.parquet
	https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2021-03.parquet

See also TLC Trip Record Data
https://www1.nyc.gov/site/tlc/about/tlc-trip-record-data.page
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	urlTemplate = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2021-%02d.parquet"
	colors      = []string{"green", "yellow"}
)

func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func main() {
	start := time.Now()
	size := 0
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplate, color, month)
			fmt.Println(url)
			n, err := downloadSize(url)
			if err != nil {
				log.Fatal(err)
			}
			size += n
		}
	}

	duration := time.Since(start)
	fmt.Println(size, duration)
}
