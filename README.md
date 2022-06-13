# slow-server

Golang web app that runs an inefficient recursive fibonacci sequence to simulate CPU load.

## Run

To run execute `go run main.go` and the web app will start serving at port 3000.

## Query

You can run a query by executing a GET request i.e. `http://localhost:3000/?n=45`

n=45 seems to be a good balance of computationally intensive, but responsive.
