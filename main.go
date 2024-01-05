package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TylerBrock/colorjson"
)

func main() {
	str := `{
	"full_text": i,
   "color": "#00ff00"
	}`

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	// Marshall the Colorized JSON
	for i := 0; i < 20; i++ {
		// s, _ := f.Marshal(obj)
		// fmt.Println(string(s))

		date := time.Now()
		fmt.Printf("%s\n", date)
		time.Sleep(time.Second)

	}

}
