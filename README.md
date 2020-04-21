# DMX over Internet

A silly protocol implementation.

## Background

During the COVID-19 situation, there's a desire to let remote folks control lighting.  The obvious solution is to use some sort of remote control software to remotely control a PC running the console.

This implements a second solution: simply route sACN data over the internet.  There's all sorts of reasons you should never do this in a production environment, but let's give it a try.

The protocol is implemented with the idea that the "presenter" can switch between senders.  I.e., in a lightlab situation, multiple users could trade-off control of a lighting rig, under the control of an instructor.  

## Security

This is (currently) not implemented with any security or even encryption.   This is a Very Bad Idea.  

This should be fixed if this is ever used in a real situation.

## Requirements

This requires that both clients (RX and TX) can see an MQTT server and have permissions to connect.

It's written in golang and should be relatively portable.  It was developed and tested primarily on Mac OS 10.15.4.

## Executing

### Input Client

sacn2cloud [options]

-u # : which sACN universe to listen to

-s hostname : which mqtt server to connect to

-p port : mqtt server port

-i <string> : some string to identify this sender

-t # : change threshold (see below), defaults to probably 25

-c # : number of slots being used (see below), defaults to 512

Change threshold: typically only changes are sent, but if enough changes happen, we send the whole universe.
Number of slots: if only the first 100 channels are being used, set this to 100.  

### Output Client

cloud2sacn [options]

-u # : which sACN universe to output

-s hostname : which mqtt server to connect to

-p port : mqtt server port

-i <string> : which string to listen to

### Internals

sacn2cloud advertises PING messages onto /sacn2cloud/v1/ping once per second, these tell you which universe it thinks its sending, as well as the string it's using to identify itself.

Whenever an sACN packet comes in, if the packet has changes, it pushes those changes onto /sacn2cloud/v1/<string> where <string> is the id string provided on the commandline.  
