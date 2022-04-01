package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ba := []byte(a)
	bb := []byte(b)
	carry := false
	i := len(bb) - 1
	j := len(ba) - 1
	for i >= 0 {
		if ba[j] == '1' && bb[i] == '1' && carry {
			ba[j] = '1'
			carry = true
		} else if ba[j] == '1' && bb[i] == '1' && !carry {
			ba[j] = '0'
			carry = true
		} else if (ba[j] == '1' || bb[i] == '1') && carry {
			ba[j] = '0'
			carry = true
		} else if (ba[j] == '1' || bb[i] == '1') && !carry {
			ba[j] = '1'
			carry = false
		} else if ba[j] == '0' && bb[i] == '0' && carry {
			ba[j] = '1'
			carry = false
		} else {
			ba[j] = '0'
			carry = false
		}
		j--
		i--
	}
	if !carry {
		return string(ba)
	}
	for ; j >= 0; j-- {
		if ba[j] == '1' && carry {
			ba[j] = '0'
			carry = true
			continue
		} else if ba[j] == '0' && carry {
			ba[j] = '1'
			carry = false
			continue
		} else {
			return string(ba)
		}
	}
	return string(append([]byte{'1'}, ba...))

}
