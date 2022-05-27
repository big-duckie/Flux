package api

import (
	"encoding/json"
	"flux/app/consts"
	"flux/app/db"
	"net/http"

	"github.com/emvi/null"
	"github.com/gorilla/mux"
)

type BuildListJson struct {
	Builds []string
}

type BuildCreateJson struct {
	Name      string
	Minecraft string
	Java      string
	Memory    uint
	Mods      []BuildCreateModJson
}

type BuildCreateModJson struct {
	Name    string
	Version string
}

type BuildJson struct {
	Minecraft string
	Forge     null.String
	Java      string
	Memory    uint
	Mods      []BuildModJson
}

type BuildModJson struct {
	Name     string
	Version  string
	MD5      string
	Url      string
	Filesize uint
}

type BuildUpdateJson struct {
	Name    string
	Version string
}

func BuildList(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var modpackDB db.Modpack
	if res := consts.DB.Where(&db.Modpack{Slug: vars["modpack"]}).First(&modpackDB); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	var buildsDB []db.Build
	if res := consts.DB.Where(&db.Build{ModpackID: modpackDB.ID}).Find(&buildsDB); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	builds := BuildListJson{}
	for _, build := range buildsDB {
		builds.Builds = append(builds.Builds, build.Name)
	}

	json.NewEncoder(rw).Encode(&builds)
}

func BuildCreate(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var data BuildCreateJson
	json.NewDecoder(r.Body).Decode(&data)

	var modpackDB db.Modpack
	if res := consts.DB.Where(&db.Modpack{Slug: vars["modpack"]}).First(&modpackDB); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	build := db.Build{
		ModpackID: modpackDB.ID,
		Name:      data.Name,
		Minecraft: data.Minecraft,
		Java:      data.Java,
		Memory:    data.Memory,
	}

	if res := consts.DB.Create(&build); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 500,
		})
		return
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "Build added",
	})
}

func Build(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var modpackDB db.Modpack
	if res := consts.DB.First(&modpackDB, &db.Modpack{Slug: vars["modpack"]}); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	var buildDB db.Build
	if res := consts.DB.First(&buildDB, &db.Build{ModpackID: modpackDB.ID, Name: vars["build"]}); res.Error != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"code": 404,
		})
		return
	}

	buildJson := BuildJson{
		Minecraft: buildDB.Minecraft,
		Java:      buildDB.Java,
		Memory:    buildDB.Memory,
	}

	for _, version := range buildDB.Mods {
		var mod db.Mod
		if res := consts.DB.First(&mod, db.Mod{ID: version.ModID}); res.Error != nil {
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"code": 500,
			})
			return
		}

		buildModJson := BuildModJson{
			Name:     mod.Slug,
			Version:  version.Version,
			MD5:      version.Hash,
			Url:      "http://test.com/",
			Filesize: version.Filesize,
		}

		buildJson.Mods = append(buildJson.Mods, buildModJson)
	}

	json.NewEncoder(rw).Encode(&buildJson)
}
