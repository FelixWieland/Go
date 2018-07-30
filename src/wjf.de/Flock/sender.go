package Flock

import "net"

type sender struct {
	IP   []byte
	Port int
}

//ReceiveValues returns stuff that is received
type ReceiveValues struct {
	Sender  sender
	Message string
}

//Send sends Data to Node
func (node *NodeInformations) Send(sVia string, IP []byte, sPort string, sMsg string) ReceiveValues {
	switch sVia {
	case "UDP":
		return node.sendUDP(sMsg)
	case "TCP":
		//return node.sendTCP()
		break
	case "WebSocket":
		//return node.sendWebSocket()
		break
	case "HTTP":
		//return node.sendWebSocket()
		break
	default:
		break
	}
	return ReceiveValues{}
}

func (node *NodeInformations) sendUDP(sMsg string) ReceiveValues {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: node.IP, Port: node.Port, Zone: ""})
	defer Conn.Close()
	Conn.Write([]byte(sMsg))
	return ReceiveValues{
		sender{
			[]byte{0},
			-1,
		},
		"undefined",
	}
}
