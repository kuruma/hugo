package hugolib

import (
	"github.com/spf13/cast"
)

type _Reference struct {
        Url     string
        Title   string
        SiteTitle       string
}

func parseReferences(i interface{}) map[string][]_Reference{
        r := make(map[string][]_Reference)
        switch v := i.(type) {
        case map[string]interface{}:
                for k, vv := range v {
                        var a []_Reference
                        switch vvv := vv.(type) {
                        case []interface{}:
                                for _, vvvv := range vvv {
                                        switch u := vvvv.(type) {
                                        case []interface{}:
                                                var as []string
                                                var aa _Reference
                                                cnt := 0
                                                for _, uu := range u {
                                                        as = append(as, cast.ToString(uu))
                                                        cnt++
                                                }
                                                if (cnt > 2) {
                                                        aa.SiteTitle = as[2]
                                                        aa.Url = as[1]
                                                        aa.Title = as[0]
                                                } else if (cnt > 1) {
                                                        aa.Url = as[1]
                                                        aa.Title = as[0]
                                                } else {
                                                        aa.Title = as[0]
                                                }
                                                a = append(a, aa)
                                        }
                                }
                        }
                        r[k] = a
                }
        }
        return r
}

