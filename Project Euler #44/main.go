package main

import (
	"fmt"
	"math"
	"time"
)

func gen_pentnum(n int) int {

	//	Pentagon number definition from
	//
	//		https://projecteuler.net/problem=44

	return int((((3 * math.Pow(float64(n), 2)) - float64(n)) / 2))
}

func is_pentnum(x int) bool {

	//	http://www.divye.in/2012/07/how-do-you-determine-if-number-n-is.html

	testval := ((math.Sqrt((24.0*float64(x))+1.0) + 1.0) / 6.0)

	if testval-math.Trunc(testval) == 0 {
		return true
	} else {
		return false
	}
}

//	http://grokbase.com/t/gg/golang-nuts/144sn9235r/go-nuts-re-how-to-make-an-anonymous-recursive-function

func main() {

	type outer_recfunc func(outer_recfunc, []int) []int
	type inner_recfunc func(inner_recfunc, []int) bool

	innerfuncvar := func(move_x0_variable inner_recfunc, x []int) bool {

		//	I originally placed
		//
		//		innerfuncvar
		//
		//	inside outerfuncvar, but moving it out to here means
		//	that it won't get redeclared with each outerfuncvar
		//	call . . .

		inner_is_sum_pentnum := is_pentnum(gen_pentnum(x[1]) + gen_pentnum(x[0]))
		inner_is_diff_pentnum := is_pentnum(gen_pentnum(x[1]) - gen_pentnum(x[0]))

		if !(inner_is_sum_pentnum && inner_is_diff_pentnum) {
			if x[0] > 1 {

				//	This specific x[1] and x[0] pair does not map to the
				//	pent nums we want. Don't move x[1], but move x[0]
				//	DOWN by one because x[0] still has room "above" 1.
				//	Return a recursive call "value" of innerfuncvar,
				//	aliased here as move_x0_variable. This returned
				//	recursive call value gets passed to outerfuncvar . . .

				x[0]--
				return move_x0_variable(move_x0_variable, x)
			} else {

				//	This specific x[1] and x[0] pair do not map to the
				//	pent nums we want. Obviously, don't move x[1] in this
				//	inner block, and don't move x[0] because it landed at
				//	1 here. Simply return a value of
				//
				//		false
				//
				//	which will get passed to the outerfuncvar function . . .

				return false
			}
		} else if inner_is_sum_pentnum && inner_is_diff_pentnum {

			//	According to
			//
			//		http://stackoverflow.com/questions/20099144/go-provide-function-return-from-an-if-statement-issues
			//
			//	the block and the function both need a return clause.
			//
			//	In this block, we found an x[1] and x[0] pair that
			//	completely solves the problem. Return true . . .

			return true
		}
		return true
	}

	outerfuncvar := func(move_x1_variable outer_recfunc, x []int) []int {

		outer_is_sum_pentnum := is_pentnum(gen_pentnum(x[1]) + gen_pentnum(x[0]))
		outer_is_diff_pentnum := is_pentnum(gen_pentnum(x[1]) - gen_pentnum(x[0]))

		if !(outer_is_sum_pentnum && outer_is_diff_pentnum) {

			//	The latest x[1] and x[0] number pair checked does
			//	not solve the problem. Move x[0] just after x[1] . . .

			x[0] = x[1] - 1

			if !innerfuncvar(innerfuncvar, x) {

				//	The calls to innerfuncvar all returned false. This
				//	means that innerfuncvar tested the integer pairs
				//	with a fixed x[1] value and an x[0] value in the
				//	range of
				//
				//		x[1] - 1
				//
				//	to
				//
				//		1
				//
				//	inclusive. None of these tested pairs solve the
				//	problem, so reset x[0] to the original x[1] integer,
				//	increment x[1], and recursively call outerfuncvar,
				//	aliased here as
				//
				//		move_x1_variable

				x[0] = x[1]
				x[1]++

				return move_x1_variable(move_x1_variable, x)
			} else {

				//	We found an x[1] and x[0] pair that completely
				//	solves the problem. Return those values to main()
				//	in the array variable x . . .

				return (x)
			}
		}

		return (x)
	}

	//	Initialize
	//
	//		array a
	//
	//	with the first two pentagonal numbers . . .

	a := []int{1, 5}

	start_time := time.Now()

	fmt.Println("\nStart Time = ", start_time)

	//	Call outerfuncvar with this recursive-looking format, even
	//	if it looks kinda strange. It will return an array, placed
	//	in countarray . . .

	countarray := outerfuncvar(outerfuncvar, a)

	end_time := time.Now()
	fmt.Println("End Time = ", end_time)
	time_diff := end_time.Sub(start_time)
	fmt.Println("Elapsed time for computation = ", time_diff)

	//	For the actual numeric solution,
	//
	//		https://www.google.com/webhp?sourceid=chrome-instant&ion=1&espv=2&ie=UTF-8#q=%22Project+Euler%22+44+solution
	//
	//	led to
	//
	//		http://www.mathblog.dk/project-euler-44-smallest-pair-pentagonal-numbers/
	//
	//	and
	//
	//		http://blog.dreamshire.com/project-euler-44-solution/
	//
	//	and these show that
	//
	//		x[1] = 2167 -> pentnum = 7042750
	//		x[0] = 1020 -> pentnum = 1560090
	//
	//		pentnum diff = 5482660

	fmt.Println("\nBased on the calculated countarray[] values, ")
	fmt.Println("\ncountarray[1] = ", countarray[1], "gen_pentnum(countarray[1]) = ", gen_pentnum(countarray[1]))
	fmt.Println("countarray[0] = ", countarray[0], "gen_pentnum(countarray[0]) = ", gen_pentnum(countarray[0]))
	fmt.Println("\ngen_pentnum(countarray[1]) + gen_pentnum(countarray[0]) = ", gen_pentnum(countarray[1])+gen_pentnum(countarray[0]))
	fmt.Println((gen_pentnum(countarray[1]) + gen_pentnum(countarray[0])), "is a pentagonal number:", is_pentnum(gen_pentnum(countarray[1])+gen_pentnum(countarray[0])))
	fmt.Println("\ngen_pentnum(countarray[1]) - gen_pentnum(countarray[0]) = ", gen_pentnum(countarray[1])-gen_pentnum(countarray[0]))
	fmt.Println((gen_pentnum(countarray[1]) - gen_pentnum(countarray[0])), "is a pentagonal number:", is_pentnum(gen_pentnum(countarray[1])-gen_pentnum(countarray[0])))
}

/*

https://projecteuler.net/problem=44

Problem 44: Pentagon numbers

Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first
ten pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their
difference, 70 − 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, Pj and Pk, for which their sum and
difference are pentagonal and D = |Pk − Pj| is minimised; what is the
value of D?

*/
