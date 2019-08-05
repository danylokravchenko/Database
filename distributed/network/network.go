package network

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"../data"
	"../cluster"
)

/**
 * Launch API
 * @db database
 */
func LaunchHTTPAPI(db *data.OneAndOnlyNumber) {

	go func() {
		m := mux.NewRouter()

		m.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
			val, _ := db.GetValue()
			fmt.Fprintf(w, "%v", val)
		})

		m.HandleFunc("/set/{newVal}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			newVal, err := strconv.Atoi(vars["newVal"])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "%v", err)
				return
			}

			db.SetValue(newVal)

			// Notify other members about changes
			ctx := context.Background()
			if name, err := os.Hostname(); err == nil {
				ctx = context.WithValue(ctx, "name", name)
			}
			members := cluster.GetOtherMembers(cluster.Cluster)
			ctx, _ = context.WithTimeout(ctx, time.Second * 2)
			go cluster.NotifyOthers(ctx, members, db)

			fmt.Fprintf(w, "%v", newVal)
		})

		m.HandleFunc("/notify/{curVal}/{curGeneration}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			curVal, err := strconv.Atoi(vars["curVal"])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "%v", err)
				return
			}
			curGeneration, err := strconv.Atoi(vars["curGeneration"])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "%v", err)
				return
			}

			if changed := db.NotifyValue(curVal, curGeneration); changed {
				log.Printf(
					"NewVal: %v Gen: %v Notifier: %v",
					curVal,
					curGeneration,
					r.URL.Query().Get("notifier"))
			}
			//log.Printf("Members: %v\n", cluster.Members())

			w.WriteHeader(http.StatusOK)
		})
		log.Fatal(http.ListenAndServe(":8080", m))
	}()

}