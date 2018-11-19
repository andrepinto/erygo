/*
Project base settings
*/
settings{
    Name = "responses"
    Package = "messages"
    Service = "users"
	Keys = {
        env = "dev"
        release = "v1.0.0"
    }
}

message "response" "UserCreatedWithSuccess"{
    StatusHTTP = 200
    Message = "user created with success"
    Kind = 3
    Details = []
}
