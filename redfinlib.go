package redfinlib

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/redfinlib/proto"
)

type info struct {
	Offers offers
	Name   string
	Price  int
}

type offers struct {
	Price int
}

func jsonExtract(stats *pb.Stats, jsonStr string) {
	i := &info{}
	json.Unmarshal([]byte(jsonStr), i)

	if i.Offers.Price > 0 {
		stats.CurrentPrice = int32(i.Offers.Price)
		return
	}

	if i.Price > 0 {
		stats.CurrentPrice = int32(i.Price)
		return
	}
}

func stateExtract(stats *pb.Stats, str string) {
	re := regexp.MustCompile("Status:.*?clickable\">(.*?)<")

	values := re.FindAllStringSubmatch(str, -1)
	for _, val := range values {
		if val[1] == "Sold" {
			stats.State = pb.Stats_SOLD
			return
		}
		if val[1] == "Pending" {
			stats.State = pb.Stats_PENDING
			return
		}
	}
}

func estimateExtract(stats *pb.Stats, str string) {
	re := regexp.MustCompile("Redfin Estimate:.*?alue\">\\$(.*?)<")

	values := re.FindAllStringSubmatch(str, -1)
	for _, val := range values {
		es, _ := strconv.Atoi(strings.Replace(val[1], ",", "", -1))
		if es > 0 {
			stats.CurrentEstimate = int32(es)
			return
		}
	}

	// Option 2
	re = regexp.MustCompile("statsValue\">\\$(.*?)<.*?Redfin Estimate")
	values = re.FindAllStringSubmatch(str, -1)
	for _, val := range values {
		es, _ := strconv.Atoi(strings.Replace(val[1], ",", "", -1))
		if es > 0 {
			stats.CurrentEstimate = int32(es)
			return
		}
	}
}

// Extract performs the extraction
func Extract(data string) (*pb.Stats, error) {
	stats := &pb.Stats{}

	re := regexp.MustCompile("ld\\+json\">(.*?)<\\/script")
	strings := re.FindAllStringSubmatch(data, -1)
	for _, str := range strings {
		jsonExtract(stats, str[1])
	}

	estimateExtract(stats, data)
	stateExtract(stats, data)

	return stats, nil
}
