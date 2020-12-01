package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port = flag.Int("port",9090, "file server port")

func main()  {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: fs [flags] dir_path\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Printf("no path provide, see usage")
		os.Exit(2)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(paths[0])))
	if err != nil {
		fmt.Printf("file server exit: %s",err)
	}
}
