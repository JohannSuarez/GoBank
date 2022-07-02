package api

import (
    "time"
    "net/http"
    "github.com/lib/pq"
    "github.com/JohannSuarez/GoBackend/util"
    "github.com/gin-gonic/gin"
    db "github.com/JohannSuarez/GoBackend/db/sqlc"
)

type createUserRequest struct {
    Username string `json:"username" binding:"required,alphanum"`
    Password string `json:"password" binding:"required",min=6`
    FullName string `json:"full_name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
    var req createUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    hashedPassword, err := util.HashPassword(req.Password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    }

    // This method is generated from SQLC
    arg := db.CreateUserParams {
        Username: req.Username, 
        HashedPassword: hashedPassword,
        FullName: req.FullName,
        Email: req.Email,
    }

    user, err := server.store.CreateUser(ctx,arg)
    if err != nil {

        if pqErr, ok := err.(*pq.Error); ok {
            switch pqErr.Code.Name() {
            case "unique_violation":
                ctx.JSON(http.StatusForbidden, errorResponse(err))
                return
            }
        }

        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    rsp := CreateUserResponse{
        Username: user.Username,
        FullName: user.FullName,
        Email: user.Email,
        PasswordChangedAt: user.PasswordChangedAt,
        CreatedAt: user.CreatedAt,
    }

    ctx.JSON(http.StatusOK, rsp)
}

