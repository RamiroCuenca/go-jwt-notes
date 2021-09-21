package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RamiroCuenca/go-rest-notesApi/auth"
	"github.com/RamiroCuenca/go-rest-notesApi/common/handler"
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/users/models"
)

// The idea is that users controllers manages the login and the signup of any user

// Login a user
func UsersLogin(w http.ResponseWriter, r *http.Request) {
	// 1° Decode the json received on a User object
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.Log().Infof("Error decoding request body: %v", err)
		handler.SendError(w, http.StatusBadRequest)
		return
	}

	// 2° We should validate the sent user...

	// 3° If the user is valid, generate a JWT
	token, err := auth.GenerateToken(u)
	if err != nil {
		logger.Log().Infof("Error generating JWT token: %v", err)
		handler.SendError(w, http.StatusInternalServerError)
		return
	}

	// 4° If the token was generated successfully, send a response with an OK and return the token
	logger.Log().Info("JWT generated successfully! :)")
	responseJson := fmt.Sprintf(`{
		"message": "JWT Token generated successfully",
		"token": %s
	}`, token)

	handler.SendResponse(w, http.StatusOK, []byte(responseJson))
	return

	/* ON OTHER DAY WE ARE GOING TO ADD SOME LOGIC TO REGISTER ON THE DATABASE
	// 2° Create the sql statement and prepare null fields
	q := `INSERT INTO notes (owner_name, title, details)
	 	VALUES ($1, $2, $3) RETURNING id`

	// A time ago... i used to open the database here, but at least on this
	// particular project we open it on the main file so it is not necessary
	// to be opened here

	// 3° Start a transaction
	db := connection.NewPostgresClient()
	tx, err := db.Begin()
	if err != nil {
		logger.Log().Infof("Error starting transaction: %v", err)
		handler.SendError(w, 500) // Internal Server Error
		return
	}

	// 4° Prepare the transaction
	stmt, err := tx.Prepare(q)
	if err != nil {
		logger.Log().Infof("Error preparing transaction: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}
	defer stmt.Close()

	// 5° Execute the query and assign the returned id to the Note object
	// We will use QueryRow because the exec method returns two methods that are
	// not compatible with psql!
	err = stmt.QueryRow(
		n.OwnerName,
		n.Title,
		stringToNull(n.Details), // Send a nill if it's null
	).Scan(&n.ID)
	if err != nil {
		logger.Log().Infof("Error executing query: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}

	// 6° As there are no errors, commit the transaction
	tx.Commit()
	logger.Log().Infof("Note created successfully! :)")

	// 7° Encode the Note into a JSON object
	json, _ := json.Marshal(n)

	// 8° Send the response
	handler.SendResponse(w, http.StatusCreated, json)
	*/
}
