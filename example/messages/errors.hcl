/*
Project base settings
*/
settings{
    Name = "errors"
    Package = "messages"
    Service = "users"
	Keys = {
        env = "dev"
        release = "v1.0.0"
    }
}


message "error" "UserNotFoud"{
    StatusHTTP = 404
    Message = "user not found"
    Kind = 1
    Details = ["hello {{.env}}"]
}

message "error" "InternalServerError"{
    StatusHTTP = 500
    Message = "internal server error"
    Kind = 2
    Details = ["hello {{.env}}"]
}


