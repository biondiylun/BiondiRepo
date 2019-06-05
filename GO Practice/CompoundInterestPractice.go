package main

import (
	"fmt"
	"math"
	"os"
)

const(
	DailyCompoundingMultiplier = 365.0
	MonthlyCompoundingMultiplier = 12.0
	QuarterlyCompoundingMultiplier = 4.0
	AnnualCompoundingMultiplier = 1.0
)

var(
	currentCapital float64 = 0
	interest float64 = 0
	durationInYear float64 = 0

)

type accountInvestment struct {
	frequencyOfCompound, durationOfInvestment string
	principal, interestRate float64
}

func ValidateAccount(a accountInvestment){
	if(a.principal > 1000000 || a.principal < -1000000) {
		fmt.Println("The Principal must be between $1,000,000 and -$1,000,000")
		os.Exit(0)
	} 
	if(a.interestRate > 100 || a.interestRate < -100) {
		fmt.Println("The Interest Rate must be between 100% and -100%")
		os.Exit(0)
	}
}

func Compound(a accountInvestment) {
	
	switch a.frequencyOfCompound {
		case "ContinuouslyCompound": 
			ContinuouslyCompound(a)
		case "DailyCompound": 
			NonContinuouslyCompound(a, DailyCompoundingMultiplier )
		case "MonthlyCompound": 
			NonContinuouslyCompound(a, MonthlyCompoundingMultiplier )
		case "QuarterlyCompound": 
			NonContinuouslyCompound(a, QuarterlyCompoundingMultiplier )
		case "AnnuallyCompound": 
			NonContinuouslyCompound(a, AnnualCompoundingMultiplier)
		default: 
			fmt.Println("Please enter one of the following: ContinuouslyCompound, DailyCompound, MonthlyCompound, QuarterlyCompound, AnnuallyCompound")
			os.Exit(0)
	}
}

func ContinuouslyCompound(a accountInvestment) {
	GetDurationInYear(a.durationOfInvestment)
	currentCapital = math.Round(a.principal*math.Pow(2.7183, a.interestRate/100.0*durationInYear)*100)/100
	interest = math.Round((currentCapital-a.principal)*100)/100
}

func NonContinuouslyCompound(a accountInvestment, compoundMultiplier float64) {
	GetDurationInYear(a.durationOfInvestment)
	currentCapital = math.Round(a.principal*math.Pow(1+(a.interestRate/100/compoundMultiplier), compoundMultiplier*durationInYear)*100)/100
	interest = math.Round((currentCapital-a.principal)*100)/100
}

func GetDurationInYear(durationOfInvestment string){
	switch durationOfInvestment {
		case "1day": 
			durationInYear = 1.0/365.0
		case "30days": 
			durationInYear = 30.0/365.0
		case "90days": 
			durationInYear = 90.0/365.0
		case "1year": 
			durationInYear = 1.0
		case "20months": 
			durationInYear = 600.0/365.0
		case "7years": 
			durationInYear = 7.0
		case "1000years": 
			durationInYear = 1000.0
		default: 
			fmt.Println("Please enter one of the following time range: 1day, 30days, 90days, 1year, 20months, 7years, 1000years")
			os.Exit(0)
	}
}

func main() {
	account := accountInvestment{"MonthlyCompound", "1year", 10000, 15}
	ValidateAccount(account)
	Compound(account)
	fmt.Println("The Starting Capital is : $", account.principal)
	fmt.Println("The interest gain is : $", interest)
	fmt.Println("The Current Capital is : $", currentCapital)
}