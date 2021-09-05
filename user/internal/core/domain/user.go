package domain




type User struct {
	UserId string
	UserName string
	FollowerUserIds map[string]bool
	FollowingUserIds map[string]bool
}


func NewUser(UserName string, FollowerUserIds map[string]bool, FollowingUserIds map[string]bool) (*User, error){
	return &User{UserName: UserName, FollowerUserIds: FollowerUserIds, FollowingUserIds: FollowingUserIds}, nil
}





