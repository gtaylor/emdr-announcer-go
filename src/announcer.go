package main

import (
	gozmq "github.com/alecthomas/gozmq"
	"os"
)

// Since our relays are supervised by process monitors, and our architecture
// is extremely fault-tolerant, un-caught errors result in the relay exiting
// with an error code. The relay will be re-started immediately.
func errorHandler(errorPrefix string, errorMessage string) {
    println(errorPrefix, errorMessage)
    os.Exit(1)
}

// Handles the setup of the various variables, and the startup of the
// main relay loop.
func RunAnnouncer() {

    context, zmqContextError := gozmq.NewContext()
    if zmqContextError != nil {
        errorHandler("zmqContextError:", zmqContextError.Error())
    }

    listener, listenSocketError := context.NewSocket(gozmq.SUB)
    if listenSocketError != nil {
        errorHandler("listenSocketError", listenSocketError.Error())
    }

    // This ZeroMQ socket is what we receive incoming messages on.
    listener.SetSockOptString(gozmq.SUBSCRIBE, "")

    listener.Bind("tcp://0.0.0.0:8049")

    // This ZeroMQ socket is where we relay the messages back out over.
    sender, senderSocketError := context.NewSocket(gozmq.PUB)
    if senderSocketError != nil {
        errorHandler("senderSocketError", senderSocketError.Error())
    }

    sender.Bind("tcp://0.0.0.0:8050")

    // Let's get this party started.
    announcerLoop(listener, sender)
}

// The main announcer loop. Receives incoming messages, spits them back out to
// any connected subscribers.
func announcerLoop(listener gozmq.Socket, sender gozmq.Socket) {

	for {
	    // This blocks until something comes down the pipe.
		msg, listenRecvError := listener.Recv(0)
		if listenRecvError != nil {
		    errorHandler("listenRecvError", listenRecvError.Error())
		}


        sender.Send(msg, 0)
	}
}
