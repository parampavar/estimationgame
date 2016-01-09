//go:generate scaneo $GOFILE
package models

import (
    "time"
    "encoding/json"
)

type User struct {
    id              int
    idp_user_id     string
    name            string
    last_updated    time.Time
    status          string
    user_attributes string
}

type Toy struct {
	id                int
	name              string
	isActive          bool
}

type Tool struct {
	id                int
	name              string
	isActive          bool
}

type Estimate struct {
	id                int
	userid            int
	toyid             int
	value             int
	createdDate       time.Time
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