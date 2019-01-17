package main

// START OMIT
import "expvar"

var hits = expvar.NewInt("hits")

func handler(w http.ResponseWriter, r *http.Request) {
	hits.Add(1)
	fmt.Fprintf(w, "Hello Gopher!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// END OMIT
