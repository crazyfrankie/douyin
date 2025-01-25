package errno

import "fmt"

type Errno struct {
	Code int32
	Msg  string
}

func (e Errno) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.Code, e.Msg)
}

func NewErrno(code int32, msg string) Errno {
	return Errno{Code: code, Msg: msg}
}

func (e Errno) WithMessage(msg string) Errno {
	e.Msg = msg
	return e
}

const (
	SuccessCode  = 00000
	InternalCode = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode

	UserAlreadyExistErrCode
	UserIsNotExistErrCode

	FollowRelationAlreadyExistErrCode
	FollowRelationNotExistErrCode

	FavoriteRelationAlreadyExistErrCode
	FavoriteRelationNotExistErrCode
	FavoriteActionErrCode

	MessageAddFailedErrCode
	FriendListNoPermissionErrCode

	VideoIsNotExistErrCode
	CommentIsNotExistErrCode
)

const (
	SuccessMsg               = "Success"
	ParamErrMsg              = "Wrong Parameter has been given"
	InternalMsg              = "Internal server error"
	PasswordIsNotVerifiedMsg = "username or password not verified"

	UserIsNotExistErrMsg = "user is not exist"
	FavoriteActionErrMsg = "favorite add failed"

	MessageAddFailedErrMsg    = "message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"
	VideoIsNotExistErrMsg     = "video is not exist"
	CommentIsNotExistErrMsg   = "comment is not exist"
)

var (
	Success        = NewErrno(SuccessCode, SuccessMsg)
	ParamErr       = NewErrno(ParamErrCode, ParamErrMsg)
	InternalServer = NewErrno(InternalCode, InternalMsg)

	UserAlreadyExistErr   = NewErrno(UserAlreadyExistErrCode, "User already exists")
	UserNotExistsErr      = NewErrno(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerified = NewErrno(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)

	FavoriteRelationAlreadyExistErr = NewErrno(FollowRelationAlreadyExistErrCode, "Favorite Relation already exist")
	FavoriteRelationNotExistErr     = NewErrno(FavoriteRelationNotExistErrCode, "FavoriteRelationNotExistErr")

	FollowRelationAlreadyExistErr = NewErrno(FollowRelationAlreadyExistErrCode, "Follow Relation already exist")
	FollowRelationNotExistErr     = NewErrno(FollowRelationNotExistErrCode, "Follow Relation does not exist")
)
