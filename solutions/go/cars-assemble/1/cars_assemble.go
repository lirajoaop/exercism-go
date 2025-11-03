package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (float64(productionRate) / 100) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    productionPerMinuteParsed := float64(productionRate)
    carsPerMinute := (productionPerMinuteParsed / 60) *  successRate / 100
    return int(carsPerMinute)
} 

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    
/* 
The attempt below represents my first try. 
  	carsCountParsed := uint(carsCount)
 	carsCountTenTimes := carsCountParsed / 10
    tenTimesRest := carsCountParsed % 10
	if carsCountParsed < 10 {
        return 10000 * (carsCountParsed)
    }
    return (carsCountTenTimes * 95000) + (tenTimesRest * 10000)
 */

    cost := (carsCount % 10) * 10000 + (carsCount / 10) * 95000
    return uint(cost)
}
