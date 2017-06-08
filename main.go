package main

import "fmt"

func main() {
	var sign string
	fmt.Print("Enter star sign: ")
	fmt.Scan(&sign)

	if sign == "Capricorn" {
		fmt.Println("December 22 - January 19")
	}
	if sign == "Aquarius" {
		fmt.Println("January 20 - February 18")
	}
	if sign == "Pisces" {
		fmt.Println("February 19 - March 20")
	}
	if sign == "Aries" {
		fmt.Println("March 21 - April 19")
	}
	if sign == "Taurus" {
		fmt.Println("April 20 - May 20")
	}
	if sign == "Gemini" {
		fmt.Println("May 21 - June 20")
	}
	if sign == "Cancer" {
		fmt.Println("June 21 - July 22")
	}
	if sign == "Leo" {
		fmt.Println("July 23 - August 22")
	}
	if sign == "Virgo" {
		fmt.Println("August 23 - September 22")
	}
	if sign == "Libra" {
		fmt.Println("September 23 - October 22")
	}
	if sign == "Scorpio" {
		fmt.Println("October 23 - November 21")
	}
	if sign == "Sagittarius" {
		fmt.Println("November 22 - December 21")
	}

}
