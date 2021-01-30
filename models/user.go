package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	EmployeeId        int64              `json:"employee_id,omitempty" bson:"employee_id,omitempty"`
	Number            int64              `json:"number" bson:"number,omitempty"`
	Email             string             `json:"email" bson:"email,omitempty"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty"`
	Password          string             `json:"password,omitempty" bson:"password,omitempty"`
	Gender            string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Role              string             `json:"role,omitempty" bson:"role,omitempty"`
	Designation       string             `json:"designation,omitempty" bson:"designation,omitempty"`
	DOJ               string             `json:"doj,omitempty" bson:"doj,omitempty"`
	DOB               string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Skill             string             `json:"skill,omitempty" bson:"skill,omitempty"`
	YearsOfExp        string             `json:"experience,omitempty" bson:"experience,omitempty"`
	Id                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AddtitionalDetail AdditionalDetail   `json:"additiona_detail,omitempty" bson:"additional_detail,omitempty"`
}

type AdditionalDetail struct {
	Graduation      EducationalDetails `json:"grad_details,omitempty" bson:"grad_details,omitempty"`
	PostGradutation EducationalDetails `json:"postgrad_detailsl_details,omitempty" bson:"postgrad_details,omitempty"`
}

type EducationalDetails struct {
	CollegeName string `json:"college_name,omitempty" bson:"college_name,omitempty"`
	CGPA        int    `json:"cgpa,omitempty" bson:"cgpa,omitempty"`
	PassoutYear int    `json:"passout_year,omitempty" bson:"passout_year,omitempty"`
}

type UserMandatoryFields struct {
	Gender      string `json:"gender,omitempty" bson:"gender,omitempty"`
	Designation string `json:"designation,omitempty" bson:"designation,omitempty"`
	DOJ         string `json:"doj,omitempty" bson:"doj,omitempty"`
	DOB         string `json:"dob,omitempty" bson:"dob,omitempty"`
	YearsOfExp  string `json:"experience,omitempty" bson:"experience,omitempty"`
	Skill       string `json:"skills,omitempty" bson:"skills,omitempty"`
}
