// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package models

import "database/sql"
import "encoding/json"
import log "github.com/parampavar/estimationgame/Godeps/_workspace/src/github.com/cihub/seelog"

func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.Id,
		&s.Idp_user_id,
		&s.Name,
		&s.Last_updated,
		&s.Status,
		&s.User_attributes,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

func ScanUsers(rs *sql.Rows) ([]User, error) {
	log.Info("ScanUsers Id 000")
	structs := make([]User, 0, 16)
	log.Info("ScanUsers Id 111")
	var err error
	for rs.Next() {
		log.Info("ScanUsers Id 3333")
		var s User
		// var s = new(User)
		if err = rs.Scan(
			&s.Id,
			&s.Idp_user_id,
			&s.Name,
			&s.Last_updated,
			&s.Status,
			&s.User_attributes,
		); err != nil {
			continue
		}
		if err == nil {
			log.Info("ScanUsers Id 111") // + string(s.Id))
			structs = append(structs, s)
		}
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func UsersJson(users []User) string {
	b, err := json.Marshal(users)
	if err != nil {
	    return ""
	}
	return string(b)
}

func ScanToy(r *sql.Row) (Toy, error) {
	var s Toy
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.IsActive,
	); err != nil {
		return Toy{}, err
	}
	return s, nil
}

func ScanToys(rs *sql.Rows) ([]Toy, error) {
	structs := make([]Toy, 0, 16)
	var err error
	for rs.Next() {
		var s Toy
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.IsActive,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ToysJson(toys []Toy) string {
	b, err := json.Marshal(toys)
	if err != nil {
	    return ""
	}
	return string(b)
}

func ScanTool(r *sql.Row) (Tool, error) {
	var s Tool
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.IsActive,
	); err != nil {
		return Tool{}, err
	}
	return s, nil
}

func ScanTools(rs *sql.Rows) ([]Tool, error) {
	structs := make([]Tool, 0, 16)
	var err error
	for rs.Next() {
		var s Tool
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.IsActive,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ToolsJson(tools []Tool) string {
	b, err := json.Marshal(tools)
	if err != nil {
	    return ""
	}
	return string(b)
}

func ScanEstimate(r *sql.Row) (Estimate, error) {
	var s Estimate
	if err := r.Scan(
		&s.Id,
		&s.Userid,
		&s.Toyid,
		&s.Value,
		&s.CreatedDate,
	); err != nil {
		return Estimate{}, err
	}
	return s, nil
}

func ScanEstimates(rs *sql.Rows) ([]Estimate, error) {
	structs := make([]Estimate, 0, 16)
	var err error
	for rs.Next() {
		var s Estimate
		if err = rs.Scan(
			&s.Id,
			&s.Userid,
			&s.Toyid,
			&s.Value,
			&s.CreatedDate,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

