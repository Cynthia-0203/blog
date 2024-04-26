package config

import "github.com/Cynthia/goblog/pkg/config"

func init() {
    config.Add("pagination", config.StrMap{
        "perpage": 10,
        "url_query": "page",
    })
}