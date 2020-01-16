// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationUserRole undocumented
type EducationUserRole int

const (
	// EducationUserRoleVStudent undocumented
	EducationUserRoleVStudent EducationUserRole = 0
	// EducationUserRoleVTeacher undocumented
	EducationUserRoleVTeacher EducationUserRole = 1
	// EducationUserRoleVNone undocumented
	EducationUserRoleVNone EducationUserRole = 2
	// EducationUserRoleVUnknownFutureValue undocumented
	EducationUserRoleVUnknownFutureValue EducationUserRole = 3
	// EducationUserRoleVFaculty undocumented
	EducationUserRoleVFaculty EducationUserRole = 4
)

// EducationUserRolePStudent returns a pointer to EducationUserRoleVStudent
func EducationUserRolePStudent() *EducationUserRole {
	v := EducationUserRoleVStudent
	return &v
}

// EducationUserRolePTeacher returns a pointer to EducationUserRoleVTeacher
func EducationUserRolePTeacher() *EducationUserRole {
	v := EducationUserRoleVTeacher
	return &v
}

// EducationUserRolePNone returns a pointer to EducationUserRoleVNone
func EducationUserRolePNone() *EducationUserRole {
	v := EducationUserRoleVNone
	return &v
}

// EducationUserRolePUnknownFutureValue returns a pointer to EducationUserRoleVUnknownFutureValue
func EducationUserRolePUnknownFutureValue() *EducationUserRole {
	v := EducationUserRoleVUnknownFutureValue
	return &v
}

// EducationUserRolePFaculty returns a pointer to EducationUserRoleVFaculty
func EducationUserRolePFaculty() *EducationUserRole {
	v := EducationUserRoleVFaculty
	return &v
}