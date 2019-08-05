package config

import (
	"os"
	"strconv"
)

var MembersToNotify int

func init() {
	var err error
	MembersToNotify, err = strconv.Atoi(os.Getenv("CLUSTER_MEMBERS"))
	if err != nil || MembersToNotify <= 0 {
		MembersToNotify = 2
	}
}

