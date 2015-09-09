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
	records []interface{}
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

// determines if the input contains a comment next, and trims whitespace
// if it does, returns that
// also returns the address of the next non-comment, non-whitespace character in the input
// returns "" and -1 if comment not finished, or no valid data in input
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

// determine the position of the next whitespace in a string
func wordLength(input []byte) int {
	i := 0
	for i < length(input) {
		if unicode.IsSpace(input[i]) {
			return i + 1
		} else {
			i++
		}
	}
	return -1
}

func validDomainName(input []byte) {

}

//Finds the first DNS record 
func readDnsRec(input []byte) (dnsRec, int, err) {
	rv := new(dnsRec)
	processed := 0
	if comment, i := sliceComment(input); i > -1 {
		rv.comment = comment
		input = input[processed:]
		processed += i
	}
	if i := wordLength(input); i > -1 {
		// there was a domain name in the input
		name := input[:i]


	}

}

func (*zone) isRec(input []byte) err {

}