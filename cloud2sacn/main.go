package main

import (
  	"flag"
//	"crypto/tls"
//	"crypto/x509"
	"fmt"
  // "net"


	"log"
  "encoding/hex"
	"os"
	"time"
//	"encoding/binary"
//	"bytes"

//	"io/ioutil"
   "github.com/Hundemeier/go-sacn/sacn"
	 mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/golang/protobuf/proto"
	pb "clouddmxpb"

	//"github.com/go-redis/redis"

)

var universe int = 1
var hostname string = "localhost"
var port string = "1883"
var id string = "anonymouse"

var client mqtt.Client

var universe_data [512]byte;

var updated int = 0;

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	if msg.Topic() == "sacn2cloud/v1/" + id {
		incoming := &pb.CloudDmx{}
    err := proto.Unmarshal(msg.Payload(),incoming)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

    if(*incoming.Type == pb.CloudDmx_COMPLETE) {
      fmt.Print("C");
      copy(universe_data[*incoming.Start:],incoming.Slots);
      //fmt.Println(hex.Dump(universe_data[:]));
      updated = 1;
    } else if(*incoming.Type == pb.CloudDmx_DELTA) {
      fmt.Print("D");
      for x, channel := range incoming.Channels {
        universe_data[channel] = byte(incoming.Values[x])
        //fmt.Println(hex.Dump(universe_data[:]));
        updated = 1;
      }
    }



	} else {
		fmt.Printf("MSG:\n%s\n", hex.Dump(msg.Payload())) // fmt.Printf("%s", hex.Dump(content))
	}
}

func main() {
  // -u # : which sACN universe to output
  //
  // -s hostname : which mqtt server to connect to
  //
  // -p port : mqtt server port
  //
  // -i <string> : some string to listen to (future: give you a menu)
  //

  for i := 0; i < 512; i++ {
    universe_data[i] = 0;
  }

	flag.IntVar(&universe, "u", 1, "which sACN universe to output")
	flag.StringVar(&hostname, "s", "localhost", "which mqtt server to connect to")
  flag.StringVar(&id,"i","anonymouse","which sender string to listen to")
	flag.StringVar(&port, "p", "1883", "mqtt server port")
	flag.Parse()

  opts := mqtt.NewClientOptions().AddBroker("tcp://" + hostname + ":" + port)

  opts.SetDefaultPublishHandler(f)

  client = mqtt.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
          panic(token.Error())
  }

  if token := client.Subscribe("sacn2cloud/v1/" + id, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

  trans, err := sacn.NewTransmitter("", [16]byte{1, 2, 3}, "test")
if err != nil {
  log.Fatal(err)
}

//activates the first universe
ch, err := trans.Activate(uint16(universe))
if err != nil {
  log.Fatal(err)
}
//deactivate the channel on exit
defer close(ch)

//set a unicast destination, and/or use multicast
trans.SetMulticast(uint16(universe), true)//this specific setup will not multicast on windows,
//because no bind address was provided

  for {
    if(updated == 1) {
      updated = 0;
      ch <- [512]byte(universe_data);
    }
  time.Sleep(25 * time.Millisecond)
}

}
