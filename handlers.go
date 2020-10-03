package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type recipeRequest struct {
	Title           *string `json:"title"`
	PreparationTime *string `json:"making_time"`
	Serves          *string `json:"serves"`
	Ingredients     *string `json:"ingredients"`
	Cost            *int    `json:"cost"`
}

type recipeResponse struct {
	Message  string    `json:"message,omitempty"`
	Recipes  []*Recipe `json:"recipes,omitempty"`
	Recipe   []*Recipe `json:"recipe,omitempty"`
	Required string    `json:"required,omitempty"`
}

// PingHandler handles GET requests to /ping
func PingHandler(s Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var res recipeResponse

		err := s.Ping()
		if err != nil {
			res.Message = err.Error()

			respondJSON(w, http.StatusServiceUnavailable, &res)
			return
		}

		res.Message = "pong"
		respondJSON(w, http.StatusOK, &res)
	}
}

// ListHandler handles GET requests to /recipes
func ListHandler(s Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var res recipeResponse

		recipes, err := s.ListRecipes()
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		res.Recipes = recipes
		respondJSON(w, http.StatusOK, res)
	}
}

// FindHandler handlers GET requests to /recipes/:id
func FindHandler(s Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var res recipeResponse

		idString := ps.ByName("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		recipe, err := s.FindRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		if recipe == nil {
			res.Message = "No recipe found"
			respondJSON(w, http.StatusNotFound, res)
			return
		}

		res.Message = "Recipe details by id"
		res.Recipe = []*Recipe{recipe}
		respondJSON(w, http.StatusOK, res)
	}
}

// CreateHandler handles POST requests to /recipes
func CreateHandler(s Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var res recipeResponse

		var req recipeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		if req.Title == nil || req.PreparationTime == nil || req.Serves == nil || req.Ingredients == nil || req.Cost == nil {
			res.Message = "Recipe creation failed!"
			res.Required = "title, preparation_time, serves, ingredients, cost"
			respondJSON(w, http.StatusOK, res)
			return
		}

		id, err := s.CreateRecipe(*req.Title, *req.PreparationTime, *req.Serves, *req.Ingredients, *req.Cost)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		recipe, err := s.FindRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		res.Message = "Recipe successfully created!"
		res.Recipe = []*Recipe{recipe}
		respondJSON(w, http.StatusOK, res)
	}
}

// DeleteHandler handles DELETE requests to /recipes/:id
func DeleteHandler(s Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var res recipeResponse

		idString := ps.ByName("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		// find recipe
		recipe, err := s.FindRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}
		if recipe == nil {
			res.Message = "No recipe found"
			respondJSON(w, http.StatusNotFound, res)
			return
		}

		err = s.DeleteRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		res.Message = "Recipe successfully removed!"
		respondJSON(w, http.StatusOK, res)
	}
}

// UpdateHandler handles PATCH requests to /recipes/:id
func UpdateHandler(s Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var res recipeResponse

		idString := ps.ByName("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		// find recipe
		original, err := s.FindRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}
		if original == nil {
			res.Message = "No recipe found"
			respondJSON(w, http.StatusNotFound, res)
			return
		}

		// get request JSON
		var req recipeRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		if req.Title == nil ||
			req.PreparationTime == nil ||
			req.Serves == nil ||
			req.Ingredients == nil ||
			req.Cost == nil {
			res.Message = "Recipe update failed!"
			res.Required = "title, preparation_time, serves, ingredients, cost"
			respondJSON(w, http.StatusBadRequest, res)
			return
		}

		err = s.UpdateRecipe(id, *req.Title, *req.PreparationTime, *req.Serves, *req.Ingredients, *req.Cost)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		updated, err := s.FindRecipe(id)
		if err != nil {
			res.Message = err.Error()
			respondJSON(w, http.StatusInternalServerError, res)
			return
		}

		res.Message = "Recipe successfully updated!"
		res.Recipe = []*Recipe{updated}
		respondJSON(w, http.StatusOK, res)
	}
}
