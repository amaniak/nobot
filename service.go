package nobot

import "fmt"

//import "encoding/json"
import "io/ioutil"
import "net/http"
import "github.com/robertkrimen/otto"
import _ "github.com/robertkrimen/otto/underscore"

func check(e error) {
	if e != nil {
		fmt.Println("ERROR:")
		fmt.Println(e)
	}
}

func Execute() {

	vm := otto.New()

	vm.Set("nobotHTTP", func(call otto.FunctionCall) otto.Value {
		url, _ := call.Argument(0).ToString()

		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil)
		check(err)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		check(err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		//json, err := json.Marshal(string(body))

		fmt.Println("response:")
		fmt.Println(string(body))

		check(err)

		result, _ := vm.ToValue(string(body))
		return result
	})

	core, err := ioutil.ReadFile("nobot.js")
	check(err)
	vm.Run(core)

	project, err := ioutil.ReadFile("project.js")
	check(err)
	vm.Run(project)

	//fmt.Print(string(dat))

	// value, err := vm.Get("abc")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(value)
}
