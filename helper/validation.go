package helper

// import (
// 	"fmt"
// 	"regexp"

// 	"github.com/e-commerce-api/users"
// )

// func validateRegisterInput(req users.Users) error {
// 	if req.Name == "" {
// 		return fmt.Errorf("name is required")
// 	}

// 	if req.Password == "" {
// 		return fmt.Errorf("password is required")
// 	}

// 	if req.Email == "" && req.Phone == "" {
// 		return fmt.Errorf("either email or phone number must be provided")
// 	}

// 	if req.Email != "" {
// 		emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
// 		matched, _ := regexp.MatchString(emailRegex, req.Email)
// 		if !matched {
// 			return fmt.Errorf("invalid email format")
// 		}
// 	}

// 	if req.Phone != "" {
// 		phoneRegex := `^\+?[1-9][0-9]{7,14}$`
// 		matched, _ := regexp.MatchString(phoneRegex, req.Phone)
// 		if !matched {
// 			return fmt.Errorf("invalid phone number format")
// 		}
// 	}

// 	return nil
// }
