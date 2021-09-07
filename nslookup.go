package main

// https://www.golangprograms.com/find-dns-records-programmatically.html
// http://networkbit.ch/golang-dns-lookup/#lookuphost

//export GODEBUG=netdns=go    # force pure Go resolver ==> the default one
//export GODEBUG=netdns=cgo   # force cgo resolver

import (
	"context"
	"fmt"
	"net"
	"time"
)
 
func PTRmain(ip string) {
	names, err := net.LookupAddr(ip)
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		fmt.Printf("no record")
	}
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
 
// google-public-dns-a.google.com.
}

func CNAMEmain(name string) {
	cname, err := net.LookupCNAME(name)
	if err != nil {
		panic(err)
	}
        // dig +short research.swtch.com cname
	fmt.Printf("%s\n", cname)
 
//ghs.google.com.
}

func HOSTmain(name string) {
	ips, err := net.LookupHost(name)
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Printf("%s\n", ip)
	}

// ::1
// 127.0.0.1
}

func Amain(name string) {
	ips, err := net.LookupIP(name)
	if err != nil {
		panic(err)
	}
	if len(ips) == 0 {
		fmt.Printf("no record")
	}
	for _, ip := range ips {
		fmt.Printf("%s\n", ip.String())
	}
 
// 172.217.1.238
// 2607:f8b0:4000:80e::200e
}

func MXmain(name string) {
	mxs, err := net.LookupMX(name)
	if err != nil {
		panic(err)
	}
        // dig +short google.com mx
	for _, mx := range mxs {
		fmt.Printf("%s %v\n", mx.Host, mx.Pref)
	}
 
// aspmx.l.google.com. 10
// alt1.aspmx.l.google.com. 20
// alt2.aspmx.l.google.com. 30
// alt3.aspmx.l.google.com. 40
// alt4.aspmx.l.google.com. 50
}

func NSmain(name string) {
	nss, err := net.LookupNS(name)
	if err != nil {
		panic(err)
	}
	if len(nss) == 0 {
		fmt.Printf("no record")
	}
	for _, ns := range nss {
		fmt.Printf("%s\n", ns.Host)
	}
/* output
ns1.google.com.
ns4.google.com.
ns3.google.com.
ns2.google.com.
*/
}

func SRVmain(service,prot,name string) {
	//cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "google.com")
	cname, srvs, err := net.LookupSRV(service, prot, name)
	if err != nil {
		panic(err)
	}
 
	fmt.Printf("\ncname: %s \n\n", cname)
 
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
 
// cname: _xmpp-server._tcp.google.com. 
// 
// xmpp-server.l.google.com.:5269:5:0
// alt2.xmpp-server.l.google.com.:5269:20:0
// alt1.xmpp-server.l.google.com.:5269:20:0
// alt4.xmpp-server.l.google.com.:5269:20:0
// alt3.xmpp-server.l.google.com.:5269:20:0
}

func TXTmain(name string) {
	txts, err := net.LookupTXT(name)
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		//dig +short gmail.com txt
		fmt.Printf("%s\n", txt)
	}
 
// v=spf1 redirect=_spf.google.com
}

func DNSusage() {
  fmt.Println( "Usage: nslookup cmd args ..." )
  fmt.Println( "  Avalable commands are:" )
  fmt.Println( "    - a name" )
  fmt.Println( "    - cname name" )
  fmt.Println( "    - host name" )
  fmt.Println( "    - mx name" )
  fmt.Println( "    - ns name" )
  fmt.Println( "    - ptr ip" )
  fmt.Println( "    - srv service protocol(tcp/udp) domainname" )
  fmt.Println( "    - txt name" )
}

func DNSmain(args []string) {
  //args := os.Args[1:]
	
  if( len(args)<=1 ) {
    DNSusage() 
  } else {
    switch args[0] {
      case "a":
        Amain(args[1])
      case "ip":
        Amain(args[1])
      case "cname":
        CNAMEmain(args[1])
      case "help":
        DNSusage()
      case "host":
        HOSTmain(args[1])
      case "mail":
        MXmain(args[1])
      case "mx":
        MXmain(args[1])
      case "ns":
        NSmain(args[1])
      case "rev":
        PTRmain(args[1])
      case "ptr":
        PTRmain(args[1])
      case "srv":
        if( len(args)<=3 ) {
	  DNSusage()
	} else {
          SRVmain(args[1],args[2],args[3])
	}
      case "txt":
        TXTmain(args[1])
      default:
        DNSusage()
    }
  }
}

/*
func main() {
	DNSmain()
}
*/












//https://stackoverflow.com/questions/59889882/specifying-dns-server-for-lookup-in-go

func main() {
    r := &net.Resolver{
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
            d := net.Dialer{
                Timeout: time.Millisecond * time.Duration(10000),
            }
            return d.DialContext(ctx, "udp", "8.8.8.8:53")
        },
    }
    ip, _ := r.LookupHost(context.Background(), "www.google.com")

    print(ip[0]+"\n")
}

