package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "log"
    "time"
)

type device struct{
	name string
	snapshot_len int32
	promiscuous bool
	err error
	timeout time.Duration
	handle *pcap.Handle
}

func main(){
	dev1:= device{name: "ens33",snapshot_len: 1024, promiscuous: false,timeout: 30*time.Second}

	//open Device
	handle, err := pcap.OpenLive(dev1.name,dev1.snapshot_len, dev1.promiscuous, dev1.timeout)
	
	if err != nil{ log.Fatal(err)}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets(){
		fmt.Println(packet)
	}
}


