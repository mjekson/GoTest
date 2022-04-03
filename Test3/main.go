package main

import (
	"fmt"

	"github.com/mjekson/GoTest/Test2"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	_ "rsc.io/quote"
)

// type config struct {
// 	DB   string
// 	Port string
// }

func main() {

	viper.AutomaticEnv()
	/*
		consulClient, _ := api.NewClient(api.DefaultConfig())
		kv := consulClient.KV()
		value, _, err := kv.Get(Test2.ServiceName+"/"+"config", nil)
		if err != nil {
			panic(err)
		}

		var conf config

		err = json.Unmarshal(value.Value, &conf)
		if err != nil {
			panic(err)
		}

		viper.Set(Test2.DB, conf.DB)
		viper.Set(Test2.PORT, conf.Port)
	*/

	err := viper.AddRemoteProvider("consul", "localhost:8500", Test2.ServiceName+"/config")
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("json")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB:", viper.GetString(Test2.DB))
	fmt.Println("PORT:", viper.GetString(Test2.PORT))
	// fmt.Println(quote.Go())
	// consulClient, _ := api.NewClient()
	fmt.Printf("DB value %v type %T\n", Test2.DB, Test2.DB)
	fmt.Printf("PORT value %v type %T\n", Test2.PORT, Test2.PORT)
}
