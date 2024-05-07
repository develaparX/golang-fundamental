package library

//yang dapat diakses dipackage lain hanya yg exported
const secondsInMinutes int = 60 //unExported var tetap bisa digunakan di package yang sama
const minutesInHour int = 60    //unExported var
const HourInADay int = 24       //Exported var

func daysToHour(day int) int { //unExported karena huruf awal kecil
	return day * HourInADay
}

func DaysToMinutes(day int) int { // Exported karena huruf awal besaar
	return day * HourInADay * minutesInHour
}
