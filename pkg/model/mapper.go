package model

import "go/grpc/userservice/gen"

// UserdataToProto converts a UserData struct into
// a generated proto counterpart.
func UserdataToProto(u *UserData) *gen.UserData {
	return &gen.UserData{
		Id:        int32(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		City:      u.City,
		Phone:     u.Phone,
		Height:    u.Height,
		Married:   u.Married,
	}
}

// UserdataFromProto converts a generated proto counterpart
// into a UserData struct.
func UserdataFromProto(m *gen.UserData) *UserData {
	return &UserData{
		ID:        int(m.Id),
		FirstName: m.FirstName,
		LastName:  m.LastName,
		City:      m.City,
		Phone:     m.Phone,
		Height:    m.Height,
		Married:   m.Married,
	}
}

// UsersdataToProto converts a slice of UserData structs into
// a generated proto counterpart.
func UsersdataToProto(u []*UserData) []*gen.UserData {
	protoUserData := make([]*gen.UserData, len(u))

	for i, user := range u {
		protoUserData[i] = UserdataToProto(user)
	}
	return protoUserData
}

func IDsInt32ToInt(ids []int32) []int {
	convertedIDS := make([]int, len(ids))

	for i, value := range ids {
		convertedIDS[i] = int(value)
	}

	return convertedIDS
}
