package main

import (
  	"flag"
//	"crypto/tls"
//	"crypto/x509"
	_ "fmt"
  // "net"


	"log"
  _ "encoding/hex"
//	"os"
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
var threshold int = 25
var slots int = 512

var client mqtt.Client

func sd(old sacn.DataPacket, new sacn.DataPacket) {

  if(new.Universe() != (uint16(universe))) {
    return;
  }

  var start uint32 = 1;
  nd := new.Data()
  od := old.Data()
  //fmt.Println(nd[0],nd[1],nd[2],nd[3])

  var changes int = 0;

  if(len(od) == len(nd)) {
    for x, slot := range nd {
        if(x >= slots) {
          break
        }
        if(slot != od[x]) {
          changes += 1
        }
    }
    if(changes > threshold) {
      // just send whole universe
      msg := pb.CloudDmx{
        Start: &start,
      }

      msg.Type = pb.CloudDmx_COMPLETE.Enum()

      msg.Slots = make([]byte,slots)

      copy(msg.Slots,nd[0:slots])

      out, err := proto.Marshal(&msg)
      if err != nil {
              log.Fatalln("Failed to encode message:", err)
      }

      token := client.Publish("sacn2cloud/v1/" + id, 0, false, out)
    	token.Wait()
    } else if(changes > 0) { // send partial universe
      msg := pb.CloudDmx{}
      msg.Type = pb.CloudDmx_DELTA.Enum()

      for x, slot := range nd {
          if(x >= slots) {
            break
          }
          if(slot != od[x]) {
            msg.Channels = append(msg.Channels,uint32(x))
            msg.Values = append(msg.Values,uint32(slot))
          }
      }



      out, err := proto.Marshal(&msg)
      if err != nil {
              log.Fatalln("Failed to encode message:", err)
      }

      token := client.Publish("sacn2cloud/v1/" + id, 0, false, out)
    	token.Wait()

      //fmt.Println(hex.Dump(out))

    }
  } else {
    // just send whole universe
    msg := pb.CloudDmx{
      Start: &start,
    }

    msg.Type = pb.CloudDmx_COMPLETE.Enum()

    msg.Slots = make([]byte,slots)

    copy(msg.Slots,nd[0:slots])

    out, err := proto.Marshal(&msg)
    if err != nil {
            log.Fatalln("Failed to encode message:", err)
    }

    token := client.Publish("sacn2cloud/v1/" + id, 0, false, out)
  	token.Wait()

  }

  //fmt.Println(changes)

}

func main() {
  // -u # : which sACN universe to listen to
  //
  // -s hostname : which mqtt server to connect to
  //
  // -p port : mqtt server port
  //
  // -i <string> : some string to identify this sender
  //
  // -t # : change threshold (see below), defaults to probably 25
  //
  // -c # : number of slots being used (see below), defaults to 512


	flag.IntVar(&universe, "u", 1, "which sACN universe to listen to")
	flag.StringVar(&hostname, "s", "localhost", "which mqtt server to connect to")
  flag.StringVar(&id,"i","anonymouse","your sender identification")
	flag.StringVar(&port, "p", "1883", "mqtt server port")
  flag.IntVar(&threshold, "t", 25, "change threshold")
  flag.IntVar(&slots, "c", 512, "number of active slots")
	flag.Parse()

  opts := mqtt.NewClientOptions().AddBroker("tcp://" + hostname + ":" + port)

  client = mqtt.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
          panic(token.Error())
  }

  rs, _ := sacn.NewReceiverSocket("127.0.0.1",nil)
  rs.JoinUniverse(1)
  rs.SetOnChangeCallback(sd)
  rs.Start()

  for {

    start_universe := uint32(universe)

  msg := pb.CloudDmx{
    Start: &start_universe,
    Id: &id,
  }

  msg.Type = pb.CloudDmx_PING.Enum()

  out, err := proto.Marshal(&msg)
  if err != nil {
          log.Fatalln("Failed to encode message:", err)
  }

  token := client.Publish("sacn2cloud/v1/ping", 0, false, out)
  token.Wait()

  time.Sleep(1 * time.Second)
}

}
