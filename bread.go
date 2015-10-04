package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "github.com/fatih/color"
)

type bread struct {
    reader *bufio.Reader
}

// Ingredients
const (
    WATER = "water"
    FLOUR = "flour"
    YEAST = "yeast"
    SALT = "salt"
)

// Types of bread
const (
    STANDARD = "standard"
    BAGUETTE = "baguette"
    CIABATTA = "ciabatta"
    RUSTIC = "rustic"
    COUNTRY = "country"
    WHOLE = "whole wheat"
    SEMOLINA = "semolina"
)

// Keep the order of the types of bread for
// easy referencing later
var breadIndex = []string{
    STANDARD, BAGUETTE, CIABATTA, RUSTIC, COUNTRY, WHOLE, SEMOLINA,
}

// Flour/water ratio for each type of bread
// in percentage
var ratio = map[string]int{
    STANDARD: 65,
    BAGUETTE: 66,
    CIABATTA: 73,
    RUSTIC: 69,
    COUNTRY: 68,
    WHOLE: 68,
    SEMOLINA: 62,
}

func main() {
    var bread = bread{bufio.NewReader(os.Stdin)}

    var breadType = bread.askType()
    var starter = bread.askStarter()
    var amount = bread.askAmount(starter)
    var formula = bread.calculate(breadType, starter, amount)

    bread.print(formula)
}

// Ask for the type of bread
func (b *bread) askType() int {
    color.Set(color.FgWhite)
    fmt.Println("What type of bread are you making?")
    color.Unset()

    // Print each type of bread sequentially
    for i, v := range breadIndex {
        fmt.Printf("%d. %s\n", i + 1, strings.Title(v))
    }

    color.Set(color.FgWhite)
    fmt.Print("Type the corresponding number (Default: Standard): ")
    color.Unset()

    var breadType, _ = b.reader.ReadString('\n')
    breadType = strings.TrimSpace(breadType)

    // No input sets the default type of bread
    if breadType == "" {
        return 0
    }

    var breadTypeInt, err = strconv.Atoi(breadType)

    if err != nil {
        color.Set(color.FgRed)
        fmt.Println("Wrong input. Try again!")
        color.Unset()
        b.askType()
    }

    if breadTypeInt == 0 || breadTypeInt > len(breadIndex) {
        color.Set(color.FgRed)
        fmt.Println("No bread type with the input number. Try again!")
        color.Unset()
        b.askType()
    }

    // Convert the input to a 0-based index for
    // the types of bread slice
    return breadTypeInt - 1
}

// Ask for the starting ingredient
func (b *bread) askStarter() string {
    color.Set(color.FgWhite)
    fmt.Print("Want to input water (w) or flour (f)? ")
    color.Unset()

    var starter, _ = b.reader.ReadString('\n')
    starter = strings.TrimSpace(starter)

    // Check if the input is the first letter of either
    // "water" or "flour"
    switch starter {
    case WATER[:1]:
        starter = WATER
    case FLOUR[:1]:
        starter = FLOUR
    default:
        color.Set(color.FgRed)
        fmt.Println("Wrong input. Try again!")
        color.Unset()
        return b.askStarter()
    }

    return starter
}

// Ask the amount of the selected ingredient
func (b *bread) askAmount(starter string) int {
    color.Set(color.FgWhite)
    fmt.Printf("Enter the amount of %s in grams: ", starter)
    color.Unset()

    var amount, _ = b.reader.ReadString('\n')
    var intAmount, err = strconv.Atoi(strings.TrimSpace(amount))

    if err != nil {
        color.Set(color.FgRed)
        fmt.Println("Wrong amount. Try again!")
        color.Unset()
        return b.askAmount(starter)
    }
    
    return intAmount
}

// Calculate the amount for each ingredient
func (b *bread) calculate(breadType int, starter string, amount int) map[string]int {
    var (
        water int
        flour int
        yeast int
        salt int
        // Ratio depends on the type of bread.
        // Set it from the ratio map, whose keys
        // are the types of bread
        flourWaterRatio = ratio[breadIndex[breadType]]
        flourYeastRatio = 2
        flourSaltRatio = 1
    )

    // Set the water and flour with the formula
    // depending on the starter ingredient
    switch starter {
    case WATER:
        water = amount
        flour = water * 100 / flourWaterRatio
    case FLOUR:
        flour = amount
        water = flour * flourWaterRatio / 100
    }

    yeast = flour * flourYeastRatio / 100
    salt = flour * flourSaltRatio / 100

    // Return a map of the ingredients for
    // easy printing
    return map[string]int{
        WATER: water,
        FLOUR: flour,
        YEAST: yeast,
        SALT: salt,
    }
}

func (b *bread) print(formula map[string]int) {
    fmt.Println("---------------------------------------")

    color.Set(color.FgWhite)
    fmt.Println("Formula:")
    color.Unset()

    color.Set(color.FgCyan)
    fmt.Printf("%s: %d gr\n", strings.Title(WATER), formula[WATER])
    fmt.Printf("%s: %d gr\n", strings.Title(FLOUR), formula[FLOUR])
    fmt.Printf("%s: %d gr\n", strings.Title(YEAST), formula[YEAST])
    fmt.Printf("%s: %d gr\n", strings.Title(SALT), formula[SALT])
    color.Unset()

    fmt.Println()

    color.Set(color.FgWhite)
    fmt.Println("General Guidelines:")
    color.Unset()

    fmt.Println("Knead for 10-12 minutes with a stand mixer, or up to 20 minutes by hand. Bake for 15 minutes in 215°C by spraying water every 5 minutes, then bake for another 15 minutes in 230°C.")
}