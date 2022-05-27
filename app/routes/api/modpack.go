package api

import (
	"encoding/json"
	"golder/app/consts"
	"golder/app/db"
	"net/http"

	"github.com/gorilla/mux"
)

type ModpackListJson struct {
	Modpacks  map[string]string
	MirrorUrl string
}

type ModpackCreateJson struct {
	Name string
	Slug string
}

type ModpackJson struct {
	Name        string
	DisplayName string
	Recommended string
	Latest      string
	Builds      []string
}

type ModpackUpdateJson struct {
	Name        string
	Slug        string
	Recommended string
	Latest      string
}

func ModpackList(rw http.ResponseWriter, r *http.Request) {
	var data []db.Modpack
	if res := consts.DB.Find(&data); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	modpacks := ModpackListJson{
		MirrorUrl: "http://localhost/repo/mods/",
	}
	for _, modpack := range data {
		modpacks.Modpacks[modpack.Slug] = modpack.Name
	}

	json.NewEncoder(rw).Encode(&modpacks)
}

func ModpackCreate(rw http.ResponseWriter, r *http.Request) {
	var data ModpackCreateJson
	json.NewDecoder(r.Body).Decode(&data)

	modpack := db.Modpack{
		Name: data.Name,
		Slug: data.Slug,
	}

	if res := consts.DB.Create(&modpack); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 500,
		})
		return
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "Modpack added",
	})
}

func Modpack(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var data db.Modpack
	if res := consts.DB.First(&data, &db.Modpack{Slug: vars["modpack"]}); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 500,
		})
		return
	}

	modpack := ModpackJson{
		Name:        data.Slug,
		DisplayName: data.Name,
		Recommended: data.Recommended,
		Latest:      data.Latest,
		Builds:      []string{},
	}

	for _, build := range data.Builds {
		modpack.Builds = append(modpack.Builds, build.Name)
	}

	json.NewEncoder(rw).Encode(modpack)
}

func ModpackUpdate(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var data ModpackUpdateJson
	json.NewDecoder(r.Body).Decode(&data)

	var modpack db.Modpack
	if res := consts.DB.Where(&db.Modpack{Slug: vars["modpack"]}).Find(&modpack); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 500,
		})
		return
	}

	modpack.Name = data.Name
	modpack.Slug = data.Slug
	modpack.Recommended = data.Recommended
	modpack.Latest = data.Latest

	if res := consts.DB.Save(&modpack); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 500,
		})
		return
	}

	modpackJson := ModpackJson{
		Name:        data.Slug,
		DisplayName: data.Name,
		Recommended: data.Recommended,
		Latest:      data.Latest,
		Builds:      []string{},
	}

	for _, build := range modpack.Builds {
		modpackJson.Builds = append(modpackJson.Builds, build.Name)
	}

	json.NewEncoder(rw).Encode(modpackJson)
}
