package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// global variables; will NOT change
var Adam [2]string = [2]string{"X", "Y"}
var Eve [2]string = [2]string{"X", "X"}
var Lilith [2]string = [2]string{"Z", "Y"}
var Diana [2]string = [2]string{"Z", "X"}

// global variables used to track #s across multiple timelines (aka multiple runs of genTryFail)
var maleExtinction int = 0  //counts how many times males died out completely across timelines
var femExtinction int = 0   //counts how many times females died out completely across timelines
var zExtinction int = 0     //counts how many times both Lilith and Diana (Z chromosom carriers) died out across timelines
var totalExtinction int = 0 //counts how many times EVERYONE died out across timelines
var maxPopReached int = 0   //counts how many times the population cap was reached across timelines
var lastGen int = 0         //if Z or men died out, this is the greatest # of generations it took for that to happen across timelines
var popCapGen int = 0       //if the population cap was reached, this is the greatest # of generations it took for that to happem across timelines

func main() {

	// Define flags
	input := flag.Bool("input", false, "Run in interactive mode")
	yeggs := flag.String("yeggs", "No", "Viable Y Chromosome Eggs? (Y/N)")
	popAdam := flag.Int("popAdam", 1000, "Initial Adam Population")
	popEve := flag.Int("popEve", 1000, "Initial Eve Population")
	popLilith := flag.Int("popLilith", 1, "Initial Lilith Population")
	popDiana := flag.Int("popDiana", 0, "Initial Diana Population")
	birthEve := flag.Float64("birthEve", 1.5, "Birth Rate for Eve")
	birthLilith := flag.Float64("birthLilith", 1.5, "Birth Rate for Lilith")
	birthDiana := flag.Float64("birthDiana", 1.5, "Birth Rate for Diana")
	maxPopFlag := flag.Int("maxPop", 50000, "Max Population")
	generationsFlag := flag.Int("generations", 37, "Number of Generations")
	timelinesFlag := flag.Int("timelines", 100, "Number of Timelines")

	flag.Parse()

	var viabeleYeggs string
	var population [4]int
	var ELDbirthRate [3]float64
	var maxPop int
	var generations int
	var Timelines int

	if *input {
		fmt.Println("Please note")
		fmt.Println("Males are XY, Eve is XX, Lilith is ZY, and Diana is ZX.")
		fmt.Println("Large Numbers may bog down the speed.  Keep the size reasonable, just in case.")

		fmt.Println("\nQuestions\nBabies with YY chromosomes aren't viable.")
		fmt.Println("However, do you think eggs with a Y chromosome would be viable?")
		fmt.Println("Could a boy get his Y chromosome from his mother, and his X chromosome from his father? (y/N) ")
		fmt.Scanln(&viabeleYeggs)

		fmt.Println("Would you like to use a standard 1000 adam, 1000 Eve, 1 Lilith starting Population? (Y/n) ")
		var standardStart string
		fmt.Scanln(&standardStart)
		if standardStart == "n" || standardStart == "N" {
			population = randomPop()
		} else {
			population = [4]int{1000, 1000, 1, 0}
		}

		fmt.Printf("Adam = %v Eve = %v Lilith = %v Diana = %v\n", population[0], population[1], population[2], population[3])
		fmt.Println(population)

		fmt.Println("Are these numbers okay? (Y - continue / n - manual input)")
		var notOK string
		fmt.Scanln(&notOK)
		if notOK == "N" || notOK == "n" {
			var Ad int
			fmt.Println("Starting Adam Population:")
			fmt.Scanln(&Ad)
			population[0] = Ad
			var Ev int
			fmt.Println("Starting Eve Population:")
			fmt.Scanln(&Ev)
			population[1] = Ev
			var Li int
			fmt.Println("Starting Lilith Population:")
			fmt.Scanln(&Li)
			population[2] = Li
			var Di int
			fmt.Println("Starting Diana Population:")
			fmt.Scanln(&Di)
			population[3] = Di

			fmt.Printf("Adam = %v Eve = %v Lilith = %v Diana = %v\n", population[0], population[1], population[2], population[3])
			fmt.Println(population)
		}

		fmt.Println("Now we need the average birth rate per woman.")
		fmt.Println("Consider the biological, psychological, and social factors.")
		fmt.Println("The 2021 birth rate in Canada was 1.43")
		fmt.Println("What is the birth rate for Eve? (ex: 1.5)")
		fmt.Scanln(&ELDbirthRate[0])
		fmt.Println("What is the birth rate for Lilith? (ex: 1.5)")
		fmt.Scanln(&ELDbirthRate[1])
		fmt.Println("What is the birth rate for Diana? (ex: 1.5)")
		fmt.Scanln(&ELDbirthRate[2])

		fmt.Println("Set a maximum population size.\nThis will help the program run quickly (less than 500000 is recomended).")
		fmt.Scanln(&maxPop)
		fmt.Println("How many generations do you want to see?\n(1,000 years would be about 37 generations)")
		fmt.Scanln(&generations)
		fmt.Println("How many timelines do you want to see? (standard is either 1 or 100)")
		fmt.Scanln(&Timelines)
	} else {
		// Use flags
		viabeleYeggs = *yeggs
		population = [4]int{*popAdam, *popEve, *popLilith, *popDiana}
		ELDbirthRate = [3]float64{*birthEve, *birthLilith, *birthDiana}
		maxPop = *maxPopFlag
		generations = *generationsFlag
		Timelines = *timelinesFlag
	}

	//nextGen(population, ELDbirthRate, viabeleYeggs, maxPop)
	//genTryFail(population, ELDbirthRate, viabeleYeggs, maxPop, generations)

	//Loops for the number of Timelines given by the user
	success := 0
	for groundhog := 0; groundhog < Timelines; groundhog++ {
		if genTryFail(population, ELDbirthRate, viabeleYeggs, maxPop, generations) {
			success++
			fmt.Println("Success! Z chromosomes still around in this timeline!")
		} else {
			fmt.Println("Timeline ended early!")
		}
	}

	//Final output
	fmt.Printf("\n%v out of %v timelines still had the Z chromosome by the end.\n", success, Timelines)
	if totalExtinction > 0 {
		fmt.Printf("%v failed because EVERYONE died out.\n", totalExtinction)
	}
	if zExtinction > 0 {
		fmt.Printf("%v failed becuase Lilith and Diana died out.  There were still Women, but no more Z chromosoms.\n", zExtinction)
	}
	if maleExtinction > 0 {
		fmt.Printf("%v failed because men died out.  Usually becasue total population got too small.\n", maleExtinction)
	}
	if femExtinction > 0 {
		fmt.Printf("%v failed because women died out completely.  Usually becasue total population got too small.\n", femExtinction)
	}
	if lastGen > 0 {
		fmt.Printf("If they ended without either men or a Z chromosome, they did so within %v generations.\n", lastGen)
	}
	if maxPopReached > 0 {
		fmt.Printf("%v were cut off early because they reached a population size of %v\n", maxPopReached, maxPop)
		fmt.Printf("They hit that population cap at or below %v generations.\n", popCapGen)
	}

}

// setup a random population
func randomPop() [4]int {
	totalRandomPopSize := 0
	fmt.Println("What small population size do you want to start with (ex: 200)? ")
	fmt.Scanln(&totalRandomPopSize)

	A := 0
	E := 0
	L := 0
	D := 0

	for i := 0; i < totalRandomPopSize; i++ {
		var child [2]string = [2]string{Adam[rand.Intn(2)], randomWoman()[rand.Intn(2)]}

		if child[0] == "X" && child[1] == "X" {
			E++
		} else if child[0] == "Z" && child[1] == "Y" {
			L++
		} else if child[0] == "Y" && child[1] == "Z" {
			L++
		} else if child[0] == "Z" && child[1] == "X" {
			D++
		} else if child[0] == "X" && child[1] == "Z" {
			D++
		} else if child[0] == "Y" && child[1] == "Y" {
			//YY not viable, try again
			i--
		} else {
			A++
		}
	}
	var array [4]int = [4]int{A, E, L, D}
	return array
}

// pick a random Eve, Lilith, or Diana
func randomWoman() [2]string {
	rand := rand.Intn(3)
	if rand == 0 {
		return Eve
	} else if rand == 1 {
		return Lilith
	} else {
		return Diana
	}
}

// uses the seed population, birthrate, viability of Y eggs, and the population cap to generate the next population
func nextGen(seedPop [4]int, birthRateELD [3]float64, viableY string, maxPopulation int) [4]int {
	var newFem int
	var nMale int
	var nEve int
	var nLilith int
	var nDiana int

	//Next gen born from Eve
	for i := 0.0; i < (float64(seedPop[1]) * birthRateELD[0]); i++ {
		var kid [2]string = [2]string{Adam[rand.Intn(2)], Eve[rand.Intn(2)]}
		if kid[0] == "Y" {
			nMale++
		} else {
			newFem++
			nEve++
		}
	}

	//Next gen born from Lilith
	for i := 0.0; i < (float64(seedPop[2]) * birthRateELD[1]); i++ {
		kid := [2]string{Lilith[rand.Intn(2)], Adam[rand.Intn(2)]}
		if kid[0] == "Y" && kid[1] == "Y" {
			i-- //YY not viable, try again
		} else if kid[0] == "Y" && kid[1] == "X" {
			if viableY == "Y" || viableY == "y" {
				nMale++
			} else {
				i-- //Y egg not viable, by user input
			}
		} else if kid[0] == "Z" && kid[1] == "Y" {
			nLilith++
			newFem++
		} else {
			nDiana++
			newFem++
		}
	}

	//Next gen born from Diana
	for i := 0.0; i < (float64(seedPop[3]) * birthRateELD[2]); i++ {
		kid := [2]string{Diana[rand.Intn(2)], Adam[rand.Intn(2)]}
		if kid[0] == "Z" && kid[1] == "X" {
			newFem++
			nDiana++
		} else if kid[0] == "Z" && kid[1] == "Y" {
			newFem++
			nLilith++
		} else if kid[0] == "X" && kid[1] == "X" {
			newFem++
			nEve++
		} else {
			nMale++
		}
	}

	newPop := [4]int{nMale, nEve, nLilith, nDiana}
	total := nMale + newFem

	//series of return values that communicate what happened through unnatural output.
	if nMale == 0 && newFem == 0 {
		totalExtinction++
		fmt.Println("EVERYONE DIED OUT!  Both Men AND Women are gone!")
	}
	if nMale == 0 {
		maleExtinction++
		return [4]int{0, 0, 0, 1}
		//Male extiction is an existential issue for sexual reprodction.
		//Since a Diana and a male can sexualy reproduce all 4 types,
		//we need not worry if Eve or Lilith dies out for 1 generation. (only if Both Z women do)
	}
	if newFem == 0 {
		femExtinction++
		return [4]int{0, 0, 0, 2} //unnatural output because if nMale == 0, the output would be [4]int{0,0,0,1}
	}
	if nLilith == 0 && nDiana == 0 {
		zExtinction++
		return [4]int{0, 0, 0, 3} //unnatural output because if nMale == 0, the output would be [4]int{0,0,0,1}
	}
	if total >= maxPopulation {
		maxPopReached++
		return [4]int{0, 0, 0, 4}
		//unnstural output because if nMale == 0, the output would be [4]int{0,0,0,1}
		//This will make it possible to exit out of the tryFail loop early, even though we still have a viable generation.
	}

	fmt.Println(newPop)
	var malePercentage float64 = float64(nMale) / float64(total)
	var evePercentage float64 = float64(nEve) / float64(total)
	var lilithPercentage float64 = float64(nLilith) / float64(total)
	var dianaPercentage float64 = float64(nDiana) / float64(total)

	fmt.Printf("Male %.2f%%", (malePercentage * 100))
	fmt.Printf("\nEve %.2f%%", (evePercentage * 100))
	fmt.Printf("\nLilith %.2f%%", (lilithPercentage * 100))
	fmt.Printf("\nDiana %.2f%%", (dianaPercentage * 100))
	fmt.Println()

	return newPop
}

func genTryFail(seedPop [4]int, birthRateELD [3]float64, viableY string, maxPopulation int, generations int) bool {
	fmt.Println("\nStarting Values for this Timeline:")
	fmt.Println(seedPop)
	series := nextGen(seedPop, birthRateELD, viableY, maxPopulation)

	for count := 1; count <= generations; count++ {

		if series == [4]int{0, 0, 0, 0} {
			fmt.Printf("Ended in %v generations because EVERYONE died out.\n", count)
			if lastGen <= count {
				lastGen = count
			}
			return false
		}
		if series == [4]int{0, 0, 0, 1} {
			fmt.Printf("Ended in %v generations because Men died out.\n", count)
			if lastGen <= count {
				lastGen = count
			}
			return false
		}
		if series == [4]int{0, 0, 0, 2} {
			fmt.Printf("Ended in %v generation because Women died out.\n", count)
			if lastGen <= count {
				lastGen = count
			}
			return false
		}
		if series == [4]int{0, 0, 0, 3} {
			fmt.Printf("Ended in %v generations because Z chromosome died out.\n", count)
			if lastGen <= count {
				lastGen = count
			}
			return false
		}
		if series == [4]int{0, 0, 0, 4} {
			fmt.Printf("Population size exceeded in %v generations!\n", count)
			if popCapGen <= count {
				popCapGen = count
			}
			return true
		}

		series = nextGen(series, birthRateELD, viableY, maxPopulation)
		if count == generations {
			return true
		}
	}
	fmt.Println("ERROR! genTryFail exited incorrectly!")
	return false //this should never be reached
}
