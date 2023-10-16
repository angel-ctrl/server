package services

import (
    "time"
    user_domain "server/domain"
)

func (serviceResources *Users_services_implmentation) UpdatePlan(User *user_domain.Users, daysToAdd int) error {
    // Obtén la fecha actual
    currentDate := time.Now()

    // Suma la cantidad de días deseada a la fecha actual
    newFinishPlanDate := currentDate.Add(time.Duration(daysToAdd) * 24 * time.Hour)

    // Actualiza el registro de usuario en la base de datos con la nueva fecha
    _, err := serviceResources.DB.Exec(`
        UPDATE users 
        SET finish_plan_at = $1 
        WHERE username = $2
    `, newFinishPlanDate, User.Username)
    
    if err != nil {
        return err
    }

    return nil
}
