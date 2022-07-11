package cmd 

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testCase := "2551"
	fmt.Println(RestoreIp(testCase))
}

func RestoreIp(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var results []string
    // Define index boundaries of each dot in an IPv4 address
    // smallest address is 0.0.0.0, biggest is 255.255.255.255
    // x is first dot, y is second dot and z is third dot index 
	const (
		MinX = 1
		MinY = 3
		MinZ = 5
		MaxX = 3
		MaxY = 7
		MaxZ = 11
	)
    // Brute force with 3 loops by inserting dots in each
    // string of integers until we find one that is a valid
	for x := MinX; x <= MaxX; x++ {
		for y := MinY; y <= MaxY; y++ {
			for z := MinZ; z <= MaxZ; z++ {
				numbers := strings.Split(s, "")
				dots := []int{x, y, z}
				for _, dot := range dots {
                    if dot < len(numbers) {
					numbers = insert(numbers, ".", dot)
                }
					ip := strings.Join(numbers, "")
					resultIsIp := isValidIp(ip)
					alreadyExists := stringInSlice(ip, results)
					if resultIsIp == true && alreadyExists == false {
						results = append(results, ip)
					}
				}
			}
		}
	}

	return results

}

func insert(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isValidIp(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return false
	}
	for _, octet := range octets {
        if _, err := strconv.ParseInt(octet, 10, 64); err != nil {
            return false
        }
        number, _ := strconv.Atoi(octet)
		if len(octet) == 0 {
            return false
		} else if number > 255 {
			return false
		} else if octet[:1] == "0" && len(octet) > 1 {
			return false
        }
	}

	return true
}
