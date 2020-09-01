package models

//Client model structure
type Client struct {
	ID            int64 //similar to a long in JAVA
	DNI           int64 //similar to a long in JAVA
	Name          string
	LastName      string
	CountryOrigin string
}

//CreateClientCMD ...
type CreateClientCMD struct {
	DNI           int64  `json:"dni"`
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	CountryOrigin string `json:"countryorigin"`
}
