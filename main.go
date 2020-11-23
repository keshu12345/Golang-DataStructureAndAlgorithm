package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonrc := "{\"POSTGL3cbcd4f5763337c93b2c0f95c29b1f7f51d0d71d28a8bd9e3fc36104819065dfPOB08010129091IOF A PAGAR\":\"7b225472616e4944223a2233636263643466353736333333376339336232633066393563323962316637663531643064373164323861386264396533666333363130343831393036356466222c225472616e44617465223a22323032302d30372d32385430393a34303a33362e3934303634313539395a222c22526566223a22504f423038303130313239303931494f462041205041474152222c224c6564676572223a22474c222c2252656631223a22504f42303830313031323930392d57696c736f6e2046726564222c2252656632223a227b5c2250726f6475746f5c223a5c2244445235465c222c5c2243617465676f7269615c223a5c223030303030312d454d495353414f205052494d4549524f205052454d494f5c222c5c224d4f56494d454e544f5c223a5c22315c222c5c225472616e636f64655c223a5c22454d495353414f205052494d4549524f205052454d494f5c222c5c2242696c6c5265665c223a5c22504f42303830313031323930395c222c5c22506f6c6963794e6f5c223a5c22504f42303830313031323930395c222c5c22436f6e74615c223a5c223532353936335c227d222c22506f7374696e67526566223a22494f462041205041474152222c224462616363223a224444523546222c224462223a22302e3030222c224372616363223a224444523546222c224372223a2239302e3030222c2253746174223a224e227d\",\"POSTGL3cbcd4f5763337c93b2c0f95c29b1f7f51d0d71d28a8bd9e3fc36104819065dfPOB08010129091PREMIO A RECEBER\":\"7b225472616e4944223a2233636263643466353736333333376339336232633066393563323962316637663531643064373164323861386264396533666333363130343831393036356466222c225472616e44617465223a22323032302d30372d32385430393a34303a33362e3934303634313539395a222c22526566223a22504f4230383031303132393039315052454d494f20412052454345424552222c224c6564676572223a22474c222c2252656631223a22504f42303830313031323930392d57696c736f6e2046726564222c2252656632223a227b5c2250726f6475746f5c223a5c2244445235465c222c5c2243617465676f7269615c223a5c223030303030312d454d495353414f205052494d4549524f205052454d494f5c222c5c224d4f56494d454e544f5c223a5c22315c222c5c225472616e636f64655c223a5c22454d495353414f205052494d4549524f205052454d494f5c222c5c2242696c6c5265665c223a5c22504f42303830313031323930395c222c5c22506f6c6963794e6f5c223a5c22504f42303830313031323930395c222c5c22436f6e74615c223a5c223532353936335c227d222c22506f7374696e67526566223a225052454d494f20412052454345424552222c224462616363223a224444523546222c224462223a22313839302e3030222c224372616363223a224444523546222c224372223a22302e3030222c2253746174223a224e227d\",\"POSTGL3cbcd4f5763337c93b2c0f95c29b1f7f51d0d71d28a8bd9e3fc36104819065dfPOB08010129091PREMIO EMITIDO\":\"7b225472616e4944223a2233636263643466353736333333376339336232633066393563323962316637663531643064373164323861386264396533666333363130343831393036356466222c225472616e44617465223a22323032302d30372d32385430393a34303a33362e3934303634313539395a222c22526566223a22504f4230383031303132393039315052454d494f20454d495449444f222c224c6564676572223a22474c222c2252656631223a22504f42303830313031323930392d57696c736f6e2046726564222c2252656632223a227b5c2250726f6475746f5c223a5c2244445235465c222c5c2243617465676f7269615c223a5c223030303030312d454d495353414f205052494d4549524f205052454d494f5c222c5c224d4f56494d454e544f5c223a5c22315c222c5c225472616e636f64655c223a5c22454d495353414f205052494d4549524f205052454d494f5c222c5c2242696c6c5265665c223a5c22504f42303830313031323930395c222c5c22506f6c6963794e6f5c223a5c22504f42303830313031323930395c222c5c22436f6e74615c223a5c223532353936335c227d222c22506f7374696e67526566223a225052454d494f20454d495449444f222c224462616363223a224444523546222c224462223a22302e3030222c224372616363223a224444523546222c224372223a22313830302e3030222c2253746174223a224e227d\"}"
	b := []byte(jsonrc)

	m := make(map[string]string)

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m)

}