package main

import (
	"encoding/json"
	"fmt"
	"regexp"

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

func extract(data string) (*pb.Stats, error) {
	stats := &pb.Stats{}

	re := regexp.MustCompile("ld\\+json\">(.*?)<\\/script")
	strings := re.FindAllStringSubmatch(data, -1)
	for _, str := range strings {
		jsonExtract(stats, str[1])
	}

	return stats, nil
}

func main() {
	fmt.Printf("We do nothing\n")
}
