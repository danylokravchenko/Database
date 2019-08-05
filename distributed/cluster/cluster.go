package cluster

import (
	"context"
	"fmt"
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"net/http"
	"../data"
	"../config"
)

var Cluster *serf.Serf

/**
 * Setup cluster with IP address of current node and IP address of cluster to join
 * @advertiseAddr IP address of current node
 * @clusterAddr IP address of cluster to join
 */
func SetupCluster(advertiseAddr string, clusterAddr string) (*serf.Serf, error) {

	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr

	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{clusterAddr}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	Cluster = cluster

	return cluster, nil

}

/**
 * Get members of the cluster except current node
 * @cluster cluster to search in
 */
func GetOtherMembers(cluster *serf.Serf) []serf.Member {
	members := cluster.Members()
	for i := 0; i < len(members); {
		if members[i].Name == cluster.LocalMember().Name || members[i].Status != serf.StatusAlive {
			if i < len(members)-1 {
				members = append(members[:i], members[i + 1:]...)
			} else {
				members = members[:i]
			}
		} else {
			i++
		}
	}
	return members
}

/**
 * Notify members of the cluster except current node
 * @ctx context
 * @otherMembers array of all members
 * @db database
 */
func NotifyOthers(ctx context.Context, otherMembers []serf.Member, db *data.OneAndOnlyNumber) {
	g, ctx := errgroup.WithContext(ctx)

	// It's possible to notify all members or just some of them
	if len(otherMembers) <= config.MembersToNotify {
		for _, member := range otherMembers {
			curMember := member
			g.Go(func() error {
				return NotifyMember(ctx, curMember.Addr.String(), db)
			})
		}
	} else {
		randIndex := rand.Int() % len(otherMembers)
		for i := 0; i < config.MembersToNotify; i++ {
			g.Go(func() error {
				return NotifyMember(
					ctx,
					otherMembers[(randIndex + i) % len(otherMembers)].Addr.String(),
					db)
			})
		}
	}

	err := g.Wait()
	if err != nil {
		log.Printf("Error when notifying other members: %v", err)
	}
}

/**
 * Notify a member of the cluster
 * @ctx context
 * @addr IP address of memeber to notify
 * @db database
 */
func NotifyMember(ctx context.Context, addr string, db *data.OneAndOnlyNumber) error {
	val, gen := db.GetValue()
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:8080/notify/%v/%v?notifier=%v", addr, val, gen, ctx.Value("name")), nil)
	if err != nil {
		return errors.Wrap(err, "Couldn't create request")
	}
	req = req.WithContext(ctx)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "Couldn't make request")
	}
	return nil
}



