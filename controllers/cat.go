package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VolodymyrVoronov/cats-back/database"
	"github.com/VolodymyrVoronov/cats-back/helpers"
	"github.com/VolodymyrVoronov/cats-back/prisma/db"
	"github.com/go-chi/chi"
)

func GetAllCats(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient

	allCats, err := pClient.Client.Cat.FindMany().Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error getting all cats")
		return
	}

	catsMap := make(map[string]interface{})
	catsMap["cats"] = allCats

	err = helpers.WriteJSON(w, http.StatusOK, catsMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error writing JSON")
		return
	}
}

func GetCatByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pClient := database.PClient

	cat, err := pClient.Client.Cat.FindUnique(db.Cat.ID.Equals(id)).Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error getting cat by id")
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, cat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error writing JSON")
		return
	}
}

func CreateCat(w http.ResponseWriter, r *http.Request) {
	var catResponse db.CatModel

	err := json.NewDecoder(r.Body).Decode(&catResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println("Error decoding JSON")
		return
	}

	pClient := database.PClient
	cat, err := pClient.Client.Cat.CreateOne(
		db.Cat.Name.Set(catResponse.Name),
		db.Cat.Age.Set(catResponse.Age),
		db.Cat.Breed.Set(catResponse.Breed),
		db.Cat.Photo.Set(catResponse.Photo),
		db.Cat.Diseases.Set(catResponse.Diseases),
		db.Cat.Information.Set(catResponse.Information),
		db.Cat.Insurance.Set(catResponse.Insurance),
		db.Cat.Alive.Set(catResponse.Alive),
		db.Cat.Dead.Set(catResponse.Dead),
		db.Cat.Marked.Set(catResponse.Marked),
	).Exec(pClient.Context)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error creating cat")
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, cat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error writing JSON")
		return
	}
}

func DeleteCatByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pClient := database.PClient

	_, err := pClient.Client.Cat.FindUnique(db.Cat.ID.Equals(id)).Delete().Exec(pClient.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error deleting cat")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateCatByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pClient := database.PClient

	var catResponse db.CatModel

	err := json.NewDecoder(r.Body).Decode(&catResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println("Error decoding JSON")
		return
	}

	_, err = pClient.Client.Cat.FindUnique(db.Cat.ID.Equals(id)).Update(
		db.Cat.Name.Set(catResponse.Name),
		db.Cat.Age.Set(catResponse.Age),
		db.Cat.Breed.Set(catResponse.Breed),
		db.Cat.Photo.Set(catResponse.Photo),
		db.Cat.Diseases.Set(catResponse.Diseases),
		db.Cat.Information.Set(catResponse.Information),
		db.Cat.Insurance.Set(catResponse.Insurance),
		db.Cat.Alive.Set(catResponse.Alive),
		db.Cat.Dead.Set(catResponse.Dead),
		db.Cat.Marked.Set(catResponse.Marked),
	).Exec(pClient.Context)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error updating cat")
		return
	}

	w.WriteHeader(http.StatusOK)
}
