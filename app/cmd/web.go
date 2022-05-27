package cmd

import (
	"flux/app/routes"
	"flux/app/utils"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/urfave/cli/v2"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "Starts the webserver",
	Description: "Start the Flux webserver. That's all you need to run.",
	Action:      runWeb,
	Flags: []cli.Flag{
		intFlag("port, p", 3000, "Web server port"),
		stringFlag("config, c", "", "Custom configuration path"),
	},
}

func runWeb(c *cli.Context) error {
	runCommon()

	Partials := []string{}
	filepath.Walk("views/partials", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := filepath.Rel("views/", path)
		if err != nil {
			return err
		}

		Partials = append(Partials, file)
		return nil
	})

	gv := goview.New(goview.Config{
		Root:         "views",
		Extension:    ".tmpl.html",
		Master:       "layouts/master",
		Partials:     Partials,
		DisableCache: true,
		Funcs: template.FuncMap{
			"firstKeyFromMap": func(data map[string][]string) interface{} {
				for k := range data {
					return k
				}
				return nil
			},
			"firstValueFromMap": func(data map[string][]string) interface{} {
				for _, v := range data {
					return v
				}
				return nil
			},
			"slug": func(s string) string {
				return slug.Make(s)
			},
			"size": func(unit string, bytes int64) string {
				return utils.ByteUnitParse(unit).Format(bytes)
			},
			"time": func(unixTime int64, format string) string {
				return time.Unix(unixTime, 0).Format(format)
			},
			"iter": func(lower, upper int) []int {
				ret := []int{}
				for i := lower; i < upper; i++ {
					ret = append(ret, i)
				}
				return ret
			},
			"len": func(data interface{}) int {
				switch data.(type) {
				case []interface{}:
					return len(data.([]interface{}))
				case map[interface{}]interface{}:
					return len(data.(map[interface{}]interface{}))
				case string:
					return len(data.(string))
				}
				return -1
			},
			"sum": func(int1, int2 int) int {
				return int1 + int2
			},
			/* "paginateCheck": func(page int) bool {
				resp := consts.DB.Order("id desc").Offset(page).Limit(20).Find(db)
				if resp.RowsAffected == 0 && page != 1 {
					return false
				} else {
					return true
				}
			}, */
		},
	})

	goview.Use(gv)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))

	routes.Init(c, r)

	return http.ListenAndServe(":"+strconv.Itoa(c.Int("port")), r)
}
