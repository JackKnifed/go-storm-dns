package main

import (
	"net"
	"unicode"
)

// determines if the input contains a comment next, and trims whitespace
// if it does, returns that
// also returns the address of the next non-comment, non-whitespace character in the input
// returns "" and -1 if comment not finished, or no valid data in input
func sliceComment(input []byte) ([]byte, int) {
	pos := 0
	comment := []byte("")
	for pos < len(input) {
		switch{
		case unicode.IsSpace(input[pos]):
			pos++
		case input[pos] == '#' || input[pos] == ';':
			if bytes.Index(input[pos:], []byte("\n") != -1{
				nextLine := bytes.Index(input[pos:], []byte("\n"))
				comment += bytes.TrimSpace(input[pos+1:nextLine])
				pos += nextLine + 1
			} else {
				return []byte(""), -1
			}
		case input[pos] == '/':
			switch input[pos+1] {
			case '/':
				if bytes.Index(input[pos:], []byte("\n")) != -1{
					nextLine := bytes.Index(input[pos:], []byte("\n"))
					comment += bytes.TrimSpace(input[pos+1:nextLine])
					pos += nextLine + 1
				} else {
					return []byte(""), -1
				}
			case '*':
				if bytes.Index(input[pos:], "*/") != -1{
					cutMark := bytes.Index(input[pos:], []byte("*/"))
					comment += bytes.TrimSpace(input[pos +1:cutMark])
					pos += cutMark + 1
				} else {
					return []byte(""), -1
				}
			default:
				return []byte(""), -1
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

func checkDomainName(input []byte, zoneName string) bool, string {
	if string(input) == zoneName + "." {
		return true, "@"
	} else	if bytes.Compare(input, []byte("@")) == 0 {
		return true, "@"
	} 

	// maybe something here to check for valid domain names?
	if input[len(input)-1] == '.' {
		return true, string(input)
	}

}

func parseRecord(input []byte) (bytesEaten int, err error) {
	var comment []byte
	for {
		newComment, i := sliceComment(input)
		if i < 0 {
			break
		} else {
			bytesEaten += i
			comment = append(comment, newComment...)
		}
	}

	var recordLine []byte
	if newline := bytes.Index(input[bytesEaten:], []byte("\n")); newline > 0 {
		recordLine = input[:newLine]
		bytesEaten += newline +1
	} else {
		recordLine = input[bytesEaten:]
		bytesEaten == len(input)
	}

	var ok bool
	var domainName string
	i := wordLength(recordLine)
	if i == -1 {
		return -1, fmt.Errorf("invalid record [%q]", recordLine)
	}
	if ok, domainName = checkDomainName(recordLine[:i+1], zoneName); !ok{
		return -1, fmt.Errorf("invalid domain name [%q]", domainName)
	}

	recordLine = bytes.TrimSpace(recordLine[i:])

	i := wordLength(recordLine)
	var ttl int
	if innerTtl, chkInt := strconv.Atoi(string(recordLine[:i])); chkInt == nil {
		ttl = innerTtl
		recordLine = bytes.TrimSPace(recordLine[i:])
	}

	if strings.ToLower(string(recordLine[:2])) == "in" {
		recordLine = bytes.TrinSpace(recordLine[:2])
	}

	i := wordLength(recordLine)
	recType := strings.ToLower(string(recordLine[:i]))
	recordLine = bytes.TrimSpace(recordLine[i:])

	var newRecord dnsRec

	switch recType{
	case "cname", "ns", "ptr":
		if ok, value = checkDomainName(recordLine, zoneName); ok {
			newRecord = dnsRec{
				name: domainName,
				ttl: ttl,
				recType: recType,
				value: value,
				comment: comment,
			}
		case "mx":
			i = wordLength(recordLine)
			var prioriy int
			if priority, ok = strconv.Atoi(string(recordLine[:i])); ok != nil {
				return -1, fmt.Errorf("invalid MX record detected - no priority")
			}
			if ok, value = checkDomain(bytes.TrimSPace(recordLine[i:])); ok {
				newRecord = dnsRec{
					name: domainName,
					ttl: ttl,
					recType: recType,
					priority: priority,
					value: value,
					comment: comment,
				}
			}
		case "a", "aaaa":
		case "txt":
		case "srv":

		}
	}
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