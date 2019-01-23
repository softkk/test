package main 

import (
    "flag"
    "fmt"
    "reflect"
    "time"
    "strings"
//    "strings"
    
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

type AdapterFactory func(s string) (string, int)

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    test3()
    
// 
}


type Account struct {
        id string
        name string
        balance float64
}


func test1() {
	account := Account{"X123", "Justin Lin", 1000}
    t := reflect.TypeOf(account)
    fmt.Println(t.Kind())       // struct
    fmt.Println(t.String())     // main.Account
//    fmt.Println(t.Elem())
    /* 
       id string
       name string
       balance float64
    */
    for i, n := 0, t.NumField(); i < n; i++ {
            f := t.Field(i)
            fmt.Println(f.Name, f.Type)
    }
}

//===========================================

type Cart struct {
    Name  string
    Price int
}
 
func (c Cart) GetPrice() {
    fmt.Println(c.Price)
}

func (c Cart) UpdatePrice(price int) {
    c.Price = price
}

func (c *Cart) UpdatePricePointer(price int) {
    fmt.Println("[pointer] Update Price to", price)
    c.Price = price
}

func test2() {
   c := &Cart{"bage", 100}
    c.GetPrice()
    c.UpdatePrice(200)
    fmt.Println(c)
    c.UpdatePricePointer(200)
    fmt.Println(c)
}

//------------------

func test3(){
	parseLogLine("{'log':'  B = col_double()\n','stream':'stdout','time':'2019-01-21T08:46:59.163669487Z'}")
}

func parseLogLine(line string) (string, time.Time) {
	line = strings.TrimSuffix(line, "\n")
	
	fmt.Println(line)
	logEntry := strings.SplitN(line, " ", 2)
	logTime, err := time.Parse(time.RFC3339Nano, logEntry[0])
	if err != nil {
		return line, time.Now()
	}

	if len(logEntry) == 2 {
		return logEntry[1], logTime
	}
	return "", logTime
}

