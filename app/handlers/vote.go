package handlers

import (
	"fmt"
	"net/http"
)

func CreateVote(w http.ResponseWriter, r *http.Request) {
	//project := model.Project{}
	//
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&project); err != nil {
	//	respondError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//defer r.Body.Close()
	//
	//if err := db.Save(&project).Error; err != nil {
	//	respondError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//respondJSON(w, http.StatusCreated, project)
	fmt.Fprintf(w, "голосование типа созданно!")
}
