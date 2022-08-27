package models

import (
	"cron-go/config"
	"fmt"
)

func SchedulerHarian() {
	config.ConnectDb()
	err := config.Db.Exec("INSERT INTO customer.test_history (agent_code, agent_name, working_area, commission, phone_no, country) SELECT agent_code, agent_name, working_area, commission, phone_no, country FROM customer.test")
	if err.Error != nil {
		fmt.Println(err.Error.Error())
	} else {
		fmt.Println("berhasil")
		err := config.Db.Exec("TRUNCATE table customer.test ")
		if err.Error != nil {
			fmt.Println(err.Error.Error())
		}

		fmt.Println("Finish")
	}
}
