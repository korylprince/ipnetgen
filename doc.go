/*
Package ipnetgen provides a way to iterate over the addresses in a subnet:

	gen, err := New("192.168.100.0/24")
	if err != nil {
		//do something with err
	}
	for ip := gen.Next(); ip != nil; ip = gen.Next() {
		//do something with ip
	}

ipnetgen works on net.IPs, meaning it supports both IPv4 and IPv6 addresses.
*/
package ipnetgen
