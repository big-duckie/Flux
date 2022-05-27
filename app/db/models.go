package db

type Mod struct {
	ID          uint
	Name        string
	Slug        string
	Author      string
	Description string
	Website     string
	Versions    []Version
}

type Version struct {
	ID       uint
	ModID    uint
	Version  string
	Hash     string
	Filesize uint
}

type Modpack struct {
	ID          uint
	Name        string
	Slug        string
	Recommended string
	Latest      string
	Builds      []Build
}

type Build struct {
	ID        uint
	ModpackID uint
	Name      string
	Minecraft string
	Java      string
	Memory    uint
	Mods      []Version `gorm:"many2many:mod_versions;"`
}

type User struct {
	ID          uint
	Username    string
	Email       string
	Password    string
	Auth        string
	Permissions string // JSON encoded
}

type Client struct {
	ID   uint
	Name string
	UUID string
}

type Key struct {
	ID   uint
	Name string
	Key  string
}
