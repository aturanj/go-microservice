package account

type (

	//CreateUserRequest is for endpoint requests
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//CreateUserResponse is for endpoint responses
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	//GetUserRequest is for endpoint requests
	GetUserRequest struct {
		ID string `json:"id"`
	}

	//GetUserResponse is for endpoint responses
	GetUserResponse struct {
		Email string `json:"email"`
	}
)
