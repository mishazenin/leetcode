package main

// "255.100.50.0"
func defangIPaddr(address string) string {
	var res []byte
	for i := 0; i < len(address)-1; i++ {
		if address[i] == '.' {
			res = append(res, []byte("[.]")...)
		} else {
			res = append(res, address[i])
		}
	}
	return string(res)
}
