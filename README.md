# slow-server

Golang web app that runs an inefficient recursive fibonacci sequence to simulate CPU load.

## Run

To run execute `go run main.go` and the web app will start serving at port 3000.

## Query

You can run a query by executing a GET request i.e. `http://localhost:3000/?n=45`

n=45 seems to be a good balance of computationally intensive, but responsive.

```bash
docker run -it --rm ddosify/ddosify ddosify -t "34.134.123.165:3000?n=40" -n 10 -d 20 -p HTTP -m GET -T 7
```

Options:

* -t is the target
* -n is number of total requests
* -d is the duration (seconds) over which to execute the requests
* -p is the protocol to use
* -m is the HTTP method
* -T is the request timeout
