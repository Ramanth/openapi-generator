/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A PetApiController binds http requests to an api service and writes the service results to the http response
type PetApiController struct {
	service PetApiServicer
}

// NewPetApiController creates a default api controller
func NewPetApiController(s PetApiServicer) Router {
	return &PetApiController{ service: s }
}

// Routes returns all of the api route for the PetApiController
func (c *PetApiController) Routes() Routes {
	return Routes{ 
		{
			"AddPet",
			strings.ToUpper("Post"),
			"/v2/pet",
			c.AddPet,
		},
		{
			"DeletePet",
			strings.ToUpper("Delete"),
			"/v2/pet/{petId}",
			c.DeletePet,
		},
		{
			"FindPetsByStatus",
			strings.ToUpper("Get"),
			"/v2/pet/findByStatus",
			c.FindPetsByStatus,
		},
		{
			"FindPetsByTags",
			strings.ToUpper("Get"),
			"/v2/pet/findByTags",
			c.FindPetsByTags,
		},
		{
			"GetPetById",
			strings.ToUpper("Get"),
			"/v2/pet/{petId}",
			c.GetPetById,
		},
		{
			"UpdatePet",
			strings.ToUpper("Put"),
			"/v2/pet",
			c.UpdatePet,
		},
		{
			"UpdatePetWithForm",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}",
			c.UpdatePetWithForm,
		},
		{
			"UploadFile",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}/uploadImage",
			c.UploadFile,
		},
	}
}

// AddPet - Add a new pet to the store
func (c *PetApiController) AddPet(w http.ResponseWriter, r *http.Request) { 
	pet := &Pet{}
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.AddPet(*pet)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// DeletePet - Deletes a pet
func (c *PetApiController) DeletePet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	petId, err := parseIntParameter(params["petId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	apiKey := r.Header.Get("apiKey")
	result, err := c.service.DeletePet(petId, apiKey)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// FindPetsByStatus - Finds Pets by status
func (c *PetApiController) FindPetsByStatus(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	status := strings.Split(query.Get("status"), ",")
	result, err := c.service.FindPetsByStatus(status)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// FindPetsByTags - Finds Pets by tags
func (c *PetApiController) FindPetsByTags(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	tags := strings.Split(query.Get("tags"), ",")
	result, err := c.service.FindPetsByTags(tags)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetPetById - Find pet by ID
func (c *PetApiController) GetPetById(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	petId, err := parseIntParameter(params["petId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.GetPetById(petId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UpdatePet - Update an existing pet
func (c *PetApiController) UpdatePet(w http.ResponseWriter, r *http.Request) { 
	pet := &Pet{}
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UpdatePet(*pet)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (c *PetApiController) UpdatePetWithForm(w http.ResponseWriter, r *http.Request) { 
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	params := mux.Vars(r)
	petId, err := parseIntParameter(params["petId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	name := r.FormValue("name")
	status := r.FormValue("status")
	result, err := c.service.UpdatePetWithForm(petId, name, status)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UploadFile - uploads an image
func (c *PetApiController) UploadFile(w http.ResponseWriter, r *http.Request) { 
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	params := mux.Vars(r)
	petId, err := parseIntParameter(params["petId"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	additionalMetadata := r.FormValue("additionalMetadata")
	file, err := ReadFormFileToTempFile(r, "file")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UploadFile(petId, additionalMetadata, file)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
