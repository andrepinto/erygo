/*
Project base settings
*/
settings{
    Name = "userserr"
    Service = "users"
	Keys = {
        env = "dev"
        release = "v1.0.0"
    }
}


//Error definition
error "UserNotFoud"{
    StatusHTTP = 404
    Message = "user not found"
    Kind = 1
    Details = ["hello {{.env}}"]
}

//Error definition
error "InternalServerError"{
    StatusHTTP = 500
    Message = "internal server error"
    Kind = 2
    Details = ["hello {{.env}}"]
}

