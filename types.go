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

func (*zone) checkPosInt(input string) bool {
	if val, ok := strconv.Atoi(input); ok && val > 0 {
		return true
	} else {
		return false
	}
}

func (*zone) validSoa(input soaRec) bool {
	if ! zone.validDomain(input.name) {
		return false
	}
	if ! zone.isPosInt(input.ttl) {
		return false
	}
	if ! zone.validDomain(input.primary) {
		return false
	}
	if ! zone.validDomain(input.contact) {
		return false
	}
	if ! zone.isPosInt(input.serial)
		return false
	}
	if ! zone.isPosInt(input.refresh)
		return false
	}
	if ! zone.isPosInt(input.retry)
		return false
	}
	if ! zone.isPosInt(input.expire)
		return false
	}
	if ! zone.isPosInt(input.minimum)
		return false
	}
	return true
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

func RewZone(soa soaRec) {

}

func (*zone) validIP4(input string) bool {
	parts := strings.Split(input, ".")
	if len(parts) != 4 {
		return false
	}
	for _, piece := range parts {
		if part > 255 || part < 0 {
			return false
		}
	}
	return true
}

func (*zone) validIP6(input string) bool {
	result := net.ParseIP([]byte(input))
	if result != nil {
		return true
	}
	return false
}

// written per spec https://tools.ietf.org/html/rfc1035
func (*zone) validDomain(input string) bool {
	parts := strings.Split(strings.ToLower(input), ".")
	for _, part := range parts {
		if len(part) > 0{
	 		if !unicode.IsLetter(part[0]) {
				return false
			} else if !unicode.IsLetter(part[len(part)-1]) && !unicode.IsDigit(part[len(part)-1]) {
				return false
			}
		}
	}
	return true
}

func (*zone) isRec(input []byte) err {

}