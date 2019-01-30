package main

import (
	"fmt"
	"math"
	"time"
)

//	https://github.com/EvenPeng/Project-Euler-using-Golang/blob/master/p44.go

func IsPentagon(x uint64) bool {
	n := (1 + uint64(math.Sqrt(float64(1+24*x)))) / 6

	return x == n*(3*n-1)/2
}

func main() {
	start := time.Now()

	var min uint64 = 0
	first := true

	for j := 2; first; j++ {

	//	Outer loop that drives the larger pentagon number
	//	"pj". Start the larger n value "j" at 2, and move
	//	it one to the right for each iteration . . .

	//	Calculate the pj value . . .

		pj := uint64((j * (3*j - 1) / 2))
		for i := j - 1; i > 0; i-- {

		//	Inner loop that drives the smaller pentagon number
		//	Start the smaller n value "i" at the (j - 1) integer
		//	value immediately after the j value found in the
		//	outer loop . . .

		//	Calculate the pi value . . .

			pi := uint64(i * (3*i - 1) / 2)

			if IsPentagon(pj + pi) && IsPentagon(pj - pi){

				//	The inner loop has found the pentagon numbers
				//	that match all the business rules, so subtract,
				//	print the relevant values, and print a time
				//	value for the calculation . . .

				min = pj - pi

				fmt.Printf("i = %d pentnum_i = %d j = %d pentnum_j = %d min = %d\n", i, pi, j, pj, min)

				elapsed := time.Since(start)
				fmt.Println("Elapsed Time:", elapsed)

				//	Find more relevant value sets until
				//	min (as calculated above) exceeds
				//	1 X 10^13 . . .

				if (min > 2000000000000) {
					first = false
				}
			}
		}
	}
}

/*


i = 1020 pentnum_i = 1560090 j = 2167 pentnum_j = 7042750 min = 5482660
Elapsed: 11.9812ms
i = 52430 pentnum_i = 4123331135 j = 91650 pentnum_j = 12599537925 min = 8476206790
Elapsed: 18.1066197s
i = 95506 pentnum_i = 13682046301 j = 110461 pentnum_j = 18302393551 min = 4620347250
Elapsed: 26.2937415s
i = 111972 pentnum_i = 18806537190 j = 121168 pentnum_j = 22022465752 min = 3215928562
Elapsed: 31.5955704s
i = 73745 pentnum_i = 8157450665 j = 129198 pentnum_j = 25038120207 min = 16880669542
Elapsed: 35.9100393s
i = 186517 pentnum_i = 52182793675 j = 224159 pentnum_j = 75370773842 min = 23187980167
Elapsed: 1m47.6602495s
i = 678531 pentnum_i = 690606137676 j = 906756 pentnum_j = 1233309211926 min = 542703074250
Elapsed: 29m15.2988644s
i = 700825 pentnum_i = 736733170525 j = 1263171 pentnum_j = 2393400831276 min = 1656667660751
Elapsed: 56m49.2345849s
i = 1471972 pentnum_i = 3250051617190 j = 1659992 pentnum_j = 4133359330100 min = 883307712910
Elapsed: 1h38m11.4647017s
i = 415877 pentnum_i = 259430310755 j = 2051177 pentnum_j = 6310989602405 min = 6051559291650
Elapsed: 2h29m55.9287942s

SQL Server pentnum verification code:

DECLARE @x bigint
SET @x = 111972
SELECT (3 * POWER(CAST(@x AS FLOAT), 2) - CAST(@x AS FLOAT))/2


DECLARE @testval FLOAT

SET @testval = (SQRT(24.0 * CAST((18806537190) AS FLOAT) + 1.0) + 1.0) / 6.0
select @testval

IF @testval - ROUND(@testval, 0) = 0
BEGIN
	SELECT 'THIS IS A PENT NUM!'
END
ELSE
BEGIN
	SELECT 'THIS IS NOT A PENT NUM!'
END
 */