package services

import (
	"database/sql"
	user_domain "server/domain"
	"time"
)

func (serviceResources *Users_services_implmentation) GetUser(UserData *user_domain.Users) (*user_domain.Users, error) {
	// Prepare the SQL query to retrieve user data by id
	query := `SELECT id, 
				username, 
				password,
				user_state,
				created_at, 
				updated_at, 
				finish_plan_at FROM users WHERE username = $1;`

	// Create variables to store the scanned values
	var finishPlanAt sql.NullTime

	// Execute the query and retrieve user data
	if err := serviceResources.DB.QueryRow(query, UserData.Username).Scan(
		&UserData.ID,
		&UserData.Username,
		&UserData.Password,
		&UserData.UserState,
		&UserData.CreatedAt,
		&UserData.UpdatedAt,
		&finishPlanAt); err != nil {
		// Return an error if the query execution fails
		return nil, err
	}

	// Check if finishPlanAt is not null
	if finishPlanAt.Valid {
		UserData.Plan = finishPlanAt.Time
	} else {
		// Handle the case where finish_plan_at is null
		UserData.Plan = time.Time{}
	}

	// Return the data
	return UserData, nil
}
