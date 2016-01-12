//go:generate scaneo $GOFILE
package models

import (
    "time"
    "encoding/json"
)

type User struct {
    Id              int
    Idp_user_id     string
    Name            string
    Last_updated    time.Time
    Status          string
    User_attributes string
}

type Toy struct {
	Id                int
	Name              string
	IsActive          bool
}

type Tool struct {
	Id                int
	Name              string
	IsActive          bool
}

type Estimate struct {
	Id                int
	Userid            int
	Toyid             int
	Value             int
	CreatedDate       time.Time
}

type ModelPrintRows interface {  
    PrintRows()
}

func (m User) PrintRows() string { 
	b, err := json.Marshal(m)
    if err != nil {
        return ""
    }
    return string(b)
}  
func (m Toy) PrintRows() string { 
	b, err := json.Marshal(m)
    if err != nil {
        return ""
    }
    return string(b)
}  