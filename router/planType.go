package router 

import (
	"net/http"
	"waydo/db"
	"log"
	"waydo/model"
	"encoding/json"
	"waydo/serializer"
)






func GetAllPlanTypes(w http.ResponseWriter, r *http.Request) {
	var planTypes []model.PlanType
	//create connection to database
	db := db.New()
	//close connection when finished all works
	defer db.Close()


	sqlStatement := "select * from plan_types"
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("have some error in get rows %v", err)
	}
	defer rows.Close()


	for rows.Next(){
		var p model.PlanType
		err = rows.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Fatalf("unable to scan rows %v", err)
		}
		planTypes = append(planTypes, p)
	}
	json.NewEncoder(w).Encode(&serializer.ResponseFormatter{planTypes, http.StatusOK})
}