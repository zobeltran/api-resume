// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package models

type Response struct {
	Success string `json:"success"`
	Data    string `json:"data"`
	Error   string `json:"error"`
}

type Details struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
}

type Educations struct {
	Degree      string `json:"degree"`
	Program     string `json:"program"`
	Institution string `json:"institution"`
	Year        uint16 `json:"year"`
	Address     string `json:"address"`
}

type Experience struct {
	CompanyName string `json:"companyName"`
	DateStarted string `json:"dateStarted"`
	DateEnded   string `json:"dateEnded"`
	Position    string `json:"position"`
	Address     string `json:"address"`
}
