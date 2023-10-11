package main

import (
	"fmt"
	"github.com/lucasgalton/go-sacn2/sacn"
	"log"
	"net"
)

func main() {

	ifi, err := net.InterfaceByName("lo") //this name depends on your machine!
	if err != nil {
		log.Fatal(err)
	}
	recv, err := sacn.NewReceiverSocket("", ifi)
	if err != nil {
		log.Fatal(err)
	}

	//instead of "" you could provide an ip-address that the socket should bind to
	trans, err := sacn.NewTransmitter("", [16]byte{1, 2, 3}, "test")
	if err != nil {
		log.Fatal(err)
	}

	//activates the first universe
	ch, err := trans.Activate(1)
	if err != nil {
		log.Fatal(err)
	}
	//deactivate the channel on exit
	defer close(ch)

	trans.SetDestinations(1, []string{"127.0.0.1:5569"})

	recv.SetOnChangeCallback(func(old sacn.DataPacket, newD sacn.DataPacket) {
		ar := newD.Data()
		newArr := [512]byte{
			//LedBar1
			uint8(255), uint8(0), ar[0], ar[1], ar[2], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[3], ar[4], ar[5], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[6], ar[7], ar[8], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[9], ar[10], ar[11], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[12], ar[13], ar[14], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[15], ar[16], ar[17], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[18], ar[19], ar[20], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[21], ar[22], ar[23], uint8(0), uint8(0),
			//LedBar2
			uint8(255), uint8(0), ar[24], ar[25], ar[26], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[27], ar[28], ar[29], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[30], ar[31], ar[32], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[33], ar[34], ar[35], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[36], ar[37], ar[38], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[39], ar[40], ar[41], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[42], ar[43], ar[44], uint8(0), uint8(0),
			uint8(255), uint8(0), ar[45], ar[46], ar[47], uint8(0), uint8(0),

			//Par1
			uint8(255), ar[48], ar[49], ar[50], uint8(0), uint8(0), uint8(0),
			//Par2
			uint8(255), ar[51], ar[52], ar[53], uint8(0), uint8(0), uint8(0),
			//Par3
			uint8(255), ar[54], ar[55], ar[56], uint8(0), uint8(0), uint8(0),
			//Par4
			uint8(255), ar[57], ar[58], ar[59], uint8(0), uint8(0), uint8(0),
		}
		ch <- newArr[:]

	})
	recv.SetTimeoutCallback(func(univ uint16) {
		fmt.Println("timeout on", univ)
	})
	recv.Start()
	recv.JoinUniverse(1)
	select {}
	recv.LeaveUniverse(1)
	fmt.Println("Leaved")

}
