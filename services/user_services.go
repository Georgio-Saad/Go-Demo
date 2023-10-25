package services

import (
	"mime/multipart"
	"todogorest/data/request"
	"todogorest/data/response"

	"github.com/aws/aws-sdk-go/aws/session"
)

type UserServices interface {
	Create(user request.CreateUserRequest) response.Response
	Update() response.Response
	Delete(userId string) response.Response
	FindById(userId string) response.Response
	FindUser(request.SigninUserRequest) response.Response
	FindAll(request.PaginationRequest) response.Response
	Refresh(token string) response.Response
	Verify(request.VerifyUserRequest) response.Response
	ResendVerification(userId string) response.Response
	UploadProfilePicture(userId string, profilePicture multipart.File, profilePictureHeader *multipart.FileHeader, sess *session.Session) response.Response
	RemoveProfilePicture(userId string) response.Response
}
