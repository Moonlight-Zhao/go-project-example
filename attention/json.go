package attention

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func NumUnmarshal(){
	jsonStr:=`{"id":1,"name":"Jerry"}`
	var res map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &res)
	fmt.Printf("%T\n",res["id"])
	i := res["id"].(int64)
	fmt.Println(i)

}

func NumDecode(){
	jsonStr:=`{"id":1,"name":"Jerry"}`
	var res map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	decoder.UseNumber()
	_ = decoder.Decode(&res)
	i,_ := res["id"].(json.Number).Int64()
	fmt.Println(i)

}



