package main
import "fmt"

//distributed go server made to handle http requests and return their necessary functions 
// this is experimental please do not use in any actual services lol

/*
right now at this moment we are currently working on an undercover project to find 
some secrete messages hiding in large message batching system. The message we are looking for contains
information about a special commander 'flame' on a special mission across the border
for classified materials. We need to find the messages before it's too late.

There will also be transactions happening in the system at a rapid pace during
the investigation as well. There will be people passing information across our networks
amidst all of the rucus going on so we'll also have those coming in as well. we'll need to
make sure to keep an eye for anyone who says something we need to catch. Once you find 
somebody report it to the missing person service where they can get us access to 
exactly what we need.

*/

type Location struct {
	x float64
	y float64
}
type StreamMessage struct  {
	timeLog int
	messageContent string
	priorityUser bool
} 

type Profile struct {
	name string
	age int
	birthday string
	funFact string
	followerCount int
	followinCount int
	
}

type BroadCastMessage struct {
	timeLog int
	messageContent string
	priorityUser bool
	location Location
	server Server
	screenName string
	profileId int
}

batchStreamMessages := [
	StreamMessage:{
		0,
		"Adding the sequence of numbers now",
		false,
	}
]

func init(port string, env FunGoEnvironment, int)
func main() {
}

