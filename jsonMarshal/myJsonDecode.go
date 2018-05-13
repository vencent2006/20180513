package main

import (
	"fmt"
	"encoding/json"
	"log"
)

//func myDecode(m map[string]interface{}) {
//	for k, v := range m {
//		fmt.Printf("k(%s):v(%T)\n", k, v)
//		switch v.(type) {
//			case float64:
//				fmt.Printf("%s:%s", k, v)
//			case string:
//				fmt.Printf("%s:%s", k, v)
//			case []string:
//				var interfaceSlice []interface{} = make([]interface{}, len(v))
//				for item := range interfaceSlice{
//					mTemp := make(map[string]interface{})
//					mTemp[k]=item
//					myDecode(mTemp)
//				}
//
//		}
//	}
//}

func main() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"],"x": {"a":{"apple":"big"},"b":"peach"}}`
	//json str 转map
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &dat);
	if  err != nil {
		log.Fatalf("unmarshal failed, %s", err)
	}

	fmt.Println("==============json str 转map=======================")
	fmt.Println(dat)

	//for k,v:= range dat{
	//	fmt.Printf("k(%s):v(%T)\n",k,v)
	//}

	//myDecode(dat)
}
