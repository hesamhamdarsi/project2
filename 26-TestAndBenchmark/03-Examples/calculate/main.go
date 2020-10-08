//Notice, when you create a package, the name of the functions in that packge should start with
//uppercase letter, otherwise, it won't import to other packages
// we call this external and internal, it's like public and private in other languages

//Package calculate is to calculate different mathes
package calculate

//Sum is used to some multiple values
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
