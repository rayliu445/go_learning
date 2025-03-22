package stringer

import "fmt"

type IPAddr [4]byte

func PrintIp() {
	host := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, value := range host {
		fmt.Printf("%v: %v\n", name, value)
	}
	result := host["googleDNS"].String()
	fmt.Println(result)
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
