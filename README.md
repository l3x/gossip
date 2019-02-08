# About this app

* Demonstrates the gossip protocol using Go.
* Has ReactJS frontend code get a Ticket [here](https://goo.gl/q7rKCi) and I'll send it to you.
* Is part of a 1 Day [course](https://www.eventbrite.com/e/cryptocurrencies-developers-class-tickets-54789982312)


## Every time you change code and want to run app
The following will start 4 peers.  The first (running on port 7000) is the bootstrap node.
```
go install
gossip -p 7000 -b 7000
gossip -p 7001 -b 7000
gossip -p 7002 -b 7000
gossip -p 7003 -b 7000
```

### To kill any old gossip apps:
```
APP_NAME=gossip
ps aux | grep -i $APP_NAME | grep -v 'grep' | awk '{print $2}' | xargs kill -9
```


# Building a gossip protocol

## Goal
We're going to build a gossip protocol that gossips with a network of peers about their current favorite movies. Each node initially will connect to only one other node to bootstrap its peer list. Then all nodes should know the favorite movies of all other nodes in the system. We'll have nodes run as separate processes listening on different ports on your local machine, and they'll pass JSON-encoded messages via HTTP (to make our lives easier).

## State
We'll use a gossip protocol to keep track of each node's current favorite movie. You can find a list of movies in `movies.txt`.

Each node's movie should be randomly re-sampled from the pool of all movies once every ~10 seconds. Once it chooses a new favorite movie, it should flood its peers with this message.

You need to have each node keep track of their own incrementing version number, so we can keep track of their state and order messages. In a gossip protocol we will often receive messages out of order, so we need to know which one is most recent.

The node should also keep a cache of the recent messages it's received. Normally we'd want to cull this, but for now we can just let it grow in memory.

## API Endpoints
When you launch the process, you will need to give it the port of another node to boostrap its peers from. That port number should be an argument passed via the command line.

Each node needs the following endpoints:

* GET /peers/ (for bootstrapping into the network)
* POST /gossip/ (for trading gossip between nodes)
  - You can decide whether gossip is bi-directional (I tell you my state, you tell me your state, and we both update)
  - Or if you want you can make gossip uni-directional. All movies will eventually get propagated through the system, so new nodes will eventually get up to speed. (Not strictly true in Bitcoin.)

## Message format
Your messages will need the following:

* UUID (for deduplication)
* Originating port (your identity)
* Version number
* TTL
* Payload

Suggested strategy:
* Start by defining your message format (write out some example messages)
* Figure out the state that each node needs to hold
* Write up your message update logic—upon receiving a gossip message, how do you update your view of the world? You should be able to run this on your example message and get a correct state transition.
* Then write a little UI code so you can easily inspect what each node is doing.
* Then write your networking and gossip logic, and test across multiple nodes!

NOTE: Observing and debugging distributed systems can be hard. I recommend investing in your UI: It'll make debugging a lot easier.

## Extra credit (if you have time):

See slides.

To get a copy of the slides, get a Ticket [here](https://goo.gl/q7rKCi) and I'll send it to you.

## Build gossip binary

```
$ go install && gossip -p 7001 -b 7000
```


## Configure Goland to run main package:

Run > Run... > Edit Configurations... > Go Build > main - gossip

**Run kind:** package
**Package path:** github.com/l3x/ebchain/gossip


## Run App

For instructions on how to use a Procfile, get a Ticket [here](https://goo.gl/q7rKCi) and I'll send it to you.

```
forego start
```
