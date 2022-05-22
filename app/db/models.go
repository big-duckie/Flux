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
	ID   uint
	Name string
	Slug string
	Mods []Mod `gorm:"many2many:modpack_mods;"`
}
