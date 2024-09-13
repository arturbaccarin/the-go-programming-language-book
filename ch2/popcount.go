package ch2

var pc [256]byte

func init() {
	for i := range pc {
		println(i)
		pc[i] = pc[i/2] + byte(i&1)
	}
}

/*
The package below deﬁnes a function PopCount that retur ns the number of set bits, that is,
bits whose value is 1, in a uint64 value, which is cal le d its popu lat ion count. It uses an init
func tion to pre compute a table of results, pc, for each possible 8-bit value so that the PopCount
func tion needn’t take 64 steps but can just retur n the sum of eig ht table lookups.
*/
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
