package main

import "fmt"

func main() {
	var choice string
	fmt.Print("Would you like to a) find the dates for a star sign, or b) enter a date to find out the star sign?: ")
	fmt.Scan(&choice)

	if choice == "a" {
		sign()
	} else if choice == "b" {
		date()
	}
}

func sign() {
	var sign string
	fmt.Print("Enter star sign: ")
	fmt.Scan(&sign)

	if sign == "Capricorn" {
		fmt.Println("December 22 - January 19")
	} else if sign == "Aquarius" {
		fmt.Println("January 20 - February 18")
	} else if sign == "Pisces" {
		fmt.Println("February 19 - March 20")
	} else if sign == "Aries" {
		fmt.Println("March 21 - April 19")
	} else if sign == "Taurus" {
		fmt.Println("April 20 - May 20")
	} else if sign == "Gemini" {
		fmt.Println("May 21 - June 20")
	} else if sign == "Cancer" {
		fmt.Println("June 21 - July 22")
	} else if sign == "Leo" {
		fmt.Println("July 23 - August 22")
	} else if sign == "Virgo" {
		fmt.Println("August 23 - September 22")
	} else if sign == "Libra" {
		fmt.Println("September 23 - October 22")
	} else if sign == "Scorpio" {
		fmt.Println("October 23 - November 21")
	} else if sign == "Sagittarius" {
		fmt.Println("November 22 - December 21")
	}
}

func date() {
	var month string
	var day int
	fmt.Print("Please enter month: ")
	fmt.Scan(&month)
	fmt.Print("Please enter day: ")
	fmt.Scan(&day)

	if month == "January" {
		if day <= 19 {
			fmt.Println("Capricorn")
		} else {
			fmt.Println("Aquarius")
		}
	} else if month == "February" {
		if day <= 18 {
			fmt.Println("Aquarius")
		} else {
			fmt.Println("Pisces")
		}
	} else if month == "March" {
		if day <= 20 {
			fmt.Println("Pisces")
		} else {
			fmt.Println("Aries")
		}
	} else if month == "April" {
		if day <= 19 {
			fmt.Println("Aries")
		} else {
			fmt.Println("Taurus")
		}
	} else if month == "May" {
		if day <= 20 {
			fmt.Println("Taurus")
		} else {
			fmt.Println("Gemini")
		}
	} else if month == "June" {
		if day <= 20 {
			fmt.Println("Gemini")
		} else {
			fmt.Println("Cancer")
		}
	} else if month == "July" {
		if day <= 22 {
			fmt.Println("Cancer")
		} else {
			fmt.Println("Leo")
		}
	} else if month == "August" {
		if day <= 22 {
			fmt.Println("Leo")
		} else {
			fmt.Println("Virgo")
		}
	} else if month == "September" {
		if day <= 22 {
			fmt.Println("Virgo")
		} else {
			fmt.Println("Libra")
		}
	} else if month == "October" {
		if day <= 22 {
			fmt.Println("Libra")
		} else {
			fmt.Println("Scorpio")
		}
	} else if month == "November" {
		if day <= 21 {
			fmt.Println("Scorpio")
		} else {
			fmt.Println("Sagittarius")
		}
	} else if month == "December" {
		if day <= 21 {
			fmt.Println("Saggitarius")
		} else {
			fmt.Println("Capricorn")
		}
	}
}
