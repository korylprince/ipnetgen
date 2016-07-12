# ipnetgen
--
    import "github.com/korylprince/ipnetgen"

Package ipnetgen provides a way to iterate over the addresses in a subnet:

    gen, err := New("192.168.100.0/24")
    if err != nil {
    	//do something with err
    }
    for ip := gen.Next(); ip != nil; ip = gen.Next() {
    	//do something with ip
    }

ipnetgen works on net.IPs, meaning it supports both IPv4 and IPv6 addresses.

## Usage

#### func  Increment

```go
func Increment(ip net.IP)
```
Increment increments the given net.IP by one bit. Incrementing the last IP in an
IP space (IPv4, IPV6) is undefined.

#### type IPNetGenerator

```go
type IPNetGenerator struct {
	*net.IPNet
}
```

IPNetGenerator is a net.IPnet wrapper that you can iterate over

#### func  New

```go
func New(cidr string) (*IPNetGenerator, error)
```
New creates a new IPNetGenerator from a CIDR string, or an error if the CIDR is
invalid.

#### func  NewFromIPNet

```go
func NewFromIPNet(ipNet *net.IPNet) *IPNetGenerator
```
NewFromIPNet creates a new IPNetGenerator from a *net.IPNet

#### func (*IPNetGenerator) Next

```go
func (g *IPNetGenerator) Next() net.IP
```
Next returns the next net.IP in the subnet
