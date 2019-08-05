package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"./data"
	"./cluster"
	"./network"
)

func main() {

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	Cluster, err := cluster.SetupCluster(
		os.Getenv("ADVERTISE_ADDR"),  // The address of our instance in the network
		os.Getenv("CLUSTER_ADDR")) // The -optional- address of the cluster to join
	if err != nil {
		log.Fatal(err)
	}
	defer Cluster.Leave()

	theOneAndOnlyNumber := data.InitTheNumber(42)
	network.LaunchHTTPAPI(theOneAndOnlyNumber)

	ctx := context.Background()
	if name, err := os.Hostname(); err == nil {
		ctx = context.WithValue(ctx, "name", name)
	}

	//debugDataPrinterTicker := time.Tick(time.Second * 5)
	numberBroadcastTicker := time.Tick(time.Second * 5)

	for {
		select {
		case <-numberBroadcastTicker:
			// Notifications go here...
			members := cluster.GetOtherMembers(Cluster)
			ctx, _ := context.WithTimeout(ctx, time.Second * 2)
			go cluster.NotifyOthers(ctx, members, theOneAndOnlyNumber)
		//case <-debugDataPrinterTicker:
		//	log.Printf("Members: %v\n", cluster.Cluster.Members())
		//
		//	curVal, curGen := theOneAndOnlyNumber.GetValue()
		//	log.Printf("State: Val: %v Gen: %v\n", curVal, curGen)
		}
	}

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is recieved
	<-ch
	fmt.Println("Stopping the server")

}



