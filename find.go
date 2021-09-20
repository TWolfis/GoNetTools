package main

import (
	"fmt"
	"log"
	"net"
)

//find network interfaces
func main(){
	interfaces, err := net.Interfaces()
	errhandle(err)

	fmt.Println("Devices found:")
	for _, intf := range interfaces {
		fmt.Printf("Index: %d\n", intf.Index)
		fmt.Printf("Name: %v\n", intf.Name)
		fmt.Printf("\tMac: %s\n",intf.HardwareAddr.String())
		fmt.Printf("\tMTU: %d\n",intf.MTU)
		fmt.Printf("\tFlags: %v\n", intf.Flags)
		fmt.Println("Device addresses:")

		//handle network addresses
		byName, err := net.InterfaceByName(intf.Name)

		errhandle(err)
		addresses, err := byName.Addrs()
		for _, address := range addresses{
			fmt.Printf("\tInterface address: %v\n", address.String())
		}
	}
}

func errhandle(err error){
	if err != nil{
		log.Fatal(err)
	}
}
