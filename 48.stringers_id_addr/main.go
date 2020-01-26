package main

import "fmt"

type IPAddr [4]byte

type Tst struct {
	Name string
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func (t Tst) String() string {
	return fmt.Sprintf("this is edited. %v", t.Name)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	Mike := Tst{Name: "test"}
	fmt.Println(Mike)
	ken := Tst{Name: "ken"}
	fmt.Println(ken.Name)
}
