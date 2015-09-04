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
					dnsRec.recType,
					dnsRec.value,
				}, "\t")
		case "MX":
			return strings.Join(
				[]string{
					dnsRec.name,
					"IN",
					dnsRec.recType,
					dnsRec.priority,
					dnsRec.value,
				}, "\t")
		case "TXT":
			return strings.Join(
				[]string{
					dnsRec.name,
					"IN",
					dnsRec.recType,
					'"' + dnsRec.value + '"',
				}, "\t")
		}
	} else {
		switch dnsRec.recType {
		case "A", "AAAA", "NS", "CNAME":
			return strings.Join(
				[]string{
					dnsRec.name,
					string(dnsRec.ttl),
					"IN",
					dnsRec.recType,
					dnsRec.value,
				}, "\t")
		case "MX":
			return strings.Join(
				[]string{
					dnsRec.name,
					string(dnsRec.ttl),
					"IN",
					dnsRec.recType,
					dnsRec.priority,
					dnsRec.value,
				}, "\t")
		case "TXT":
			return strings.Join(
				[]string{
					dnsRec.name,
					string(dnsRec.ttl),
					"IN",
					dnsRec.recType,
					'"' + dnsRec.value + '"',
				}, "\t")

		}
	}
}

// determines if the input contains a comment next
// if it does
func sliceComment(input []byte) ([]byte, int) {
	pos := 0
	for pos < len(input) {
		switch{
		case unicode.IsSpace(input[pos]):
			pos++
		case input[pos] == '#':
			if bytes.Index(input[pos:]) != -1{
				pos += bytes.Index(input[pos:], "\n")
				pos++
			} else {
				return []byte(""), -1
			}
		case input[pos] == '/':
			switch input[pos+1] {
			case '/':
				if bytes.Index(input[pos:], "\n") != -1{
					pos += bytes.Index(input[pos:], "\n")
					pos++
				} else {
					return []byte(""), -1
				}
			case '*':
				if bytes.Index(input[pos:], "*/") != -1{
					pos += bytes.Index(input[pos:], "*/")
					pos++
					pos++
				} else {
					return []byte(""), -1
				}
			default:
				pos++
			}
		default:
			return bytes.TrimSpace(input[:pos]), pos+1
		}
	}
	return []byte(""), -1
}

//Finds the first DNS record 
func readDnsRec(input []byte) (dnsRec, int, err) {
}

func (*zone) isRec(input []byte) err {

}