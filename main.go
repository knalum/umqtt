package main
/*
	umqtt : tiny mqtt client
	knalum 2018
 */

import (
	"net/url"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"log"
	"fmt"
	"flag"
	"strconv"
)

func connect(clientId string, uri *url.URL) mqtt.Client {

	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect finish")
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func main(){

	// Flag parse
	urlPtr := flag.String("host","","Hostname of the broker")
	usrPtr := flag.String("user","","Username for the broker")
	pwdPtr := flag.String("pass","","Password for the broker")
	portPtr := flag.String("port","1883","Port for the broker (default 1883)")
	intervalPtr := flag.Int64("int",0,"Interval for publish message in seconds")

	topicPtr := flag.String("topic","A","Topic (default A)")
	msgPtr := flag.String("msg","","Message (default unix nano)")
	flag.Parse()
	// -- End flag parse

	urlPath := fmt.Sprintf("mqtt://%s:%s@%s:%s",*usrPtr,*pwdPtr,*urlPtr,*portPtr)
	log.Printf("Connecting to %s\n",urlPath)

	uri,err := url.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}

	client := connect("sub",uri)


	if *intervalPtr == 0 {
		publish(client,*topicPtr,*msgPtr)
	} else {
		for{
			client = connect("sub",uri)
			publish(client,*topicPtr,*msgPtr)
			time.Sleep(time.Second*time.Duration(*intervalPtr))
		}
	}
}

func publish(client mqtt.Client,topic string,msg string){
	if msg == "" {
		msg = strconv.FormatInt(time.Now().UnixNano(),10)
	}
	log.Printf("Pub topic: %s \t msg: %s",topic,msg)
	client.Publish(topic,0,false,msg)
	client.Disconnect(250)
}
