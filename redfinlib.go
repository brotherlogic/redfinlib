package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/redfinlib/proto"
)

type info struct {
	Offers offers
	Name   string
}

type offers struct {
	Price int
}

func jsonExtract(stats *pb.Stats, jsonStr string) {
	i := &info{}
	json.Unmarshal([]byte(jsonStr), i)

	if i.Offers.Price > 0 {
		stats.CurrentPrice = int32(i.Offers.Price)
	}

}

func estimateExtract(stats *pb.Stats, str string) {
	re := regexp.MustCompile("Redfin Estimate:.*?value\">\\$(.*?)<")

	values := re.FindAllStringSubmatch(str, -1)
	for _, val := range values {
		es, _ := strconv.Atoi(strings.Replace(val[1], ",", "", -1))
		if es > 0 {
			stats.CurrentEstimate = int32(es)
		}
	}
}

func extract(data string) (*pb.Stats, error) {
	stats := &pb.Stats{}

	re := regexp.MustCompile("ld\\+json\">(.*?)<\\/script")
	strings := re.FindAllStringSubmatch(data, -1)
	for _, str := range strings {
		jsonExtract(stats, str[1])
	}

	estimateExtract(stats, data)

	return stats, nil
}

func main() {
	fmt.Printf("We do nothing\n")
}
