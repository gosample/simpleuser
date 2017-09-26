package main

import (
	"github.com/yaronsumel/simpleuser/server/semaphore"
	"github.com/yaronsumel/simpleuser/server/storage"
	"github.com/yaronsumel/simpleuser/server/user"
	"log"
	"net/http"
)

// sort users by counter
// db.getCollection('simpleuserns').find({}).sort({"timesreceived":-1})

const maxConnectionToHandle = 50

func main() {

	sm := semaphore.NewSemaphore(maxConnectionToHandle)
	storage := storage.NewHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		sm.Wait()
		defer sm.Release()
		// serve just one path
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		defer req.Body.Close()
		// parse user request
		var u *user.Object = &user.Object{}
		if err := u.Decode(req.Body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		storage.Insert(u)
	})

	// serve on :8080
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}
