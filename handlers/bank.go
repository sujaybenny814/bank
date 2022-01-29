package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/micro1/models"
	"github.com/micro1/response"
	"github.com/micro1/validation"
)

func (h handler) AddBank(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	addBank := validation.AddBankModel{}
	json.Unmarshal(body, &addBank)

	err := validation.CheckValidation(addBank)
	if err != nil {
		response.ErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintln(err))
		return
	}

	var bank model.Bank
	json.Unmarshal(body, &bank)

	if res := h.DB.Create(&bank); res.Error != nil {
		fmt.Println(res.Error)
		response.ErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintln(res.Error))
	} else {
		response.SuccessMessageResponse(w, r, http.StatusAccepted, "Successfully created a bank")
	}
}

func (h handler) GetAllBank(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var Banks []model.Bank
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if res := h.DB.Limit(limit).Offset(page).Where("status = ?", "active").Order("bank_name").
		Find(&Banks); res.Error != nil {
		response.ErrorResponse(w, r, http.StatusNoContent, "No Bank found")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Banks)
	}
}

func (h handler) UpdateBank(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var bank model.Bank
	json.Unmarshal(body, &bank)
	var updateData = make(map[string]interface{})

	if bank.BankId == "" {
		response.ErrorResponse(w, r, http.StatusBadGateway, "BankId required")
		return
	}
	if bank.BankName != "" {
		updateData["bank_name"] = bank.BankName
	}
	if bank.IfscCode != "" {
		updateData["ifsc_code"] = bank.IfscCode
	}
	if bank.BranchName != "" {
		updateData["branch_name"] = bank.BranchName
	}
	updateData["status"] = "active"

	if res := h.DB.Model(&model.Bank{}).Where("bank_id = ? AND status = ?", bank.BankId, "active").Updates(updateData); res.Error != nil {
		fmt.Println(res)
		response.ErrorResponse(w, r, http.StatusBadGateway, "Failed to Update Bank")
		return

	} else {
		response.SuccessMessageResponse(w, r, http.StatusOK, "Updated successfully")
		return
	}
}

func (h handler) DeleteBank(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if res := h.DB.Model(&model.Bank{}).Where("id = ? AND status = ?", id, "active").Update("status", "delete"); res.Error != nil {
		response.ErrorResponse(w, r, http.StatusBadGateway, "Failed to Deleted Bank")
	} else {
		response.SuccessMessageResponse(w, r, http.StatusOK, "Deleted successfully")

	}
}
