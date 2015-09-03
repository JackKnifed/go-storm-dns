package main

import ()


type dnsRec struct {
	name string
	ttl int
	recType string
	priority int //priority
	protocol string
	service string
	weight string
	port string
	comment string
	value string // also commonly called target
}

type soaRec struct {
	name string
	ttl int
	const recTyep := "SOA"
	primary string
	contact string
	serial int
	refresh int
	retry int
	expire int
	minimum int
}

type fqdn struct {
	parentPart string
	localPart string
	records []dnsRec
	subdomains []fqdn
}

type zone struct {
	soa soaRec
	defaultTTL int
	tld fqdn
}

// Returns the DNS record in standard zone file format.
func (dnsRec) String() (string) {
	if dnsRec.ttl == 0 {
		switch dnsRec.recType {
		case "A", "AAAA", "NS", "CNAME":
			return strings.Join(
				[]string{
					dnsRec.name,
					"IN",
					dnsRec.type,
					dnsRec.value
				}, "\t")
		case "MX":
			return strings.Join(
				[]string{
					dnsRec.name,
					"IN",
					dnsRec.type,
					dnsRec.priority,
					dnsRec.value
				}, "\t")
		case "TXT":
			return strings.Join(
				[]string{
					dnsRec.name,
					"IN",
					dnsRec.type,
					'"' + dnsRec.value + '"'
				}, "\t")
		}
	}
}

func (*zone) isRec(input []byte) err {

}