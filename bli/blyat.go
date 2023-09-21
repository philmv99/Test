package main

import (
	"fmt"
)

func main() {
	var hourlyRate float64
	fmt.Print("Enter hourly salary rate: $")
	fmt.Scanln(&hourlyRate)

	totalHours := 0.0

	for i := 0; i < 7; i++ {
		fmt.Printf("Enter hours worked on day %d (hh:mm): ", i+1)
		var hours, minutes int
		_, err := fmt.Scanf("%d:%d", &hours, &minutes)

		if err != nil || hours < 0 || minutes < 0 || minutes >= 60 {
			fmt.Println("Invalid input. Please use the format hh:mm.")
			return
		}

		totalHours += float64(hours) + float64(minutes)/60.0
	}

	totalEarnings := totalHours * hourlyRate

	fmt.Printf("\nTotal hours worked this week: %.2f\n", totalHours)
	fmt.Printf("Total earnings for the week: $%.2f\n", totalEarnings)
}
