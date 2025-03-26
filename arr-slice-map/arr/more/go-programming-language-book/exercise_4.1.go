package goprogramminglanguagebook

func BitDifference(a, b [32]byte) int {
	diffCount := 0
	for i := 0; i < len(a); i++ {
		diff := a[i] ^ b[i] // Step 1: XOR the corresponding bytes
		for diff > 0 {      // Step 2: Count the different bits in the byte
			diffCount += int(diff & 1) // Step 3: Check if the last bit is 1 (AND OPERATOR)
			diff >>= 1                 // Step 4: Shift the bits to the right
		}
	}
	return diffCount
}

/*

Example:

diff = 00001100 -> The result between a and b XOR operation

diff & 1 -> The 1 representaion in binary is 00000001

SO

& (AND operator) Just compare if the last bit of diff and 00000001 are 1 and returns true or false.

int() is just to convert the uint8 (byte) value to a int value

*/
