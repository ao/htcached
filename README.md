# htcached

A in-memory http caching server to sit in front of a traditional http server

## how to use

syntax:

`go run . <port> <end_resource>`

example:

`go run . 80 https://ao.gl`


This will run the server locally on port 80 and will pipe resources from end resource of https://ao.gl
