/*
Project base settings
*/
settings{
    Name = "autherr"
    Service = "auth"
	Keys = {
        env = "dev"
        release = "v1.0.0"
    }
}


//Error definition
error "ErrInvalidToken"{
    Name = "ErrInvalidToken"
    StatusHTTP = 400
    Message = "invalid token received"
    Kind = 1
    Details = ["hello {{.env}}"]
}

//Error definition
error "ErrInvalidToken"{
    Name = "ErrTokenNotOwnedBySender"
    StatusHTTP = 403
    Message = "Can`t identify sender as token owner"
    Kind = 2
}

//Error definition
error "ErrInvalidToken"{
    Name = "ErrTokenNotFound"
    StatusHTTP = 404
    Message = "Token was not found in storage"
    Kind = 3
}