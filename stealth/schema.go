package stealth

type GEO struct {
	IPAddr   string
	Country  string // 2-letter country code
	Language string // 2-letter language code
}

type Device struct {
	Browser Browser
	OS      OS
}

type OS struct {
	Family  string
	Version string
	Major   int // Major Version Number
}

type Browser struct {
	Family  string
	Version string
	Major   int // Major Version Number
}

type Profile struct {
	Name      Name
	Age       int
	Interests []Interest
	Cookies   []Cookie
}

type Name struct {
	Sex    bool // true for Male, false for Female
	First  string
	Middle string
	Last   string
}

type Interest struct {
	Category string
	Score    float64 // Eager score
}

type Cookie struct {
	// social media cookies from Facebook, Instagram, Twitter, Pinterest, etc.
}
