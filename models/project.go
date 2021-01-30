package models

type Project struct {
	ProjectId          int64         `json:"project_id,omitempty" bson:"project_id,omitempty"`
	Name               string        `json:"name" bson:"name,omitempty"`
	Type               string        `json:"type,omitempty" type:"name,omitempty"`
	CustomerName       string        `json:"customer_name,omitempty" bson:"customer_name,omitempty"`
	Status             string        `json:"status,omitempty" bson:"status,omitempty"`
	Description        string        `json:"description" bson:"description,omitempty"`
	CurrentRequirement []Requirement `json:"current_requirement,omitempty" bson:"current_requirement,omitempty"`
	CurrentTeam        []TeamMember  `json:"current_team,omitempty" bson:"current_team,omitempty"`
	PreviousMembers    []TeamMember  `json:"previous_members,omitempty" bson:"previous_members,omitempty"`
	StartDate          string        `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EstimatedEndDate   string        `json:"est_end_date,omitempty" bson:"est_end_date,omitempty"`
	Location           string        `json:"location,omitempty" bson:"location,omitempty"`
}

type TeamMember struct {
	EmployeeId string `json:"employee_id,omitempty" bson:"employee_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	JoinDate   string `json:"doj,omitempty" bson:"doj,omitempty"`
	TotalHours string `json:"total_hours,omitempty" bson:"total_hours,omitempty"`
	Skills     string `json:"skills,omitempty" bson:"skills,omitempty"`
}

type Requirement struct {
	Skills     []string `json:"skills,omitempty" bson:"skills,omitempty"`
	Experience string   `json:"experience,omitempty" bson:"experience,omitempty"`
	Duration   string   `json:"duration,omitempty" bson:"duration,omitempty"`
}
