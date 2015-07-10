package gojira

import (
	"encoding/json"
	//	"encoding/xml"
	"fmt"
	//	"net/url"
	//	"time"
)

const (
	project_url = "/project/"
)

type ProjectRoles struct {
	Name   string `json:"name"`
	Actors []Role `json:"actors"`
}

type Role struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

/*
Returns roles for project. This resource cannot be accessed anonymously.

    GET http://example.com:8080/jira/rest/api/2/project/KEY/role

Parameters

    projectKey string

Usage

	prRoles, err := jira.getProjectRoles("key")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", prRoles)
*/
func (j *Jira) GetProjectRoles(key string) ([]*ProjectRoles, error) {
	url := j.BaseUrl + j.ApiPath + project_url + key + "/role"
	contents := j.buildAndExecRequest("GET", url, nil)

	var roles map[string]string
	err := json.Unmarshal(contents, &roles)
	if err != nil {
		fmt.Printf("%s", err)
	}
	var prRoles []*ProjectRoles
	for _, rUrl := range roles {
		prRole := new(ProjectRoles)
		contents = j.buildAndExecRequest("GET", rUrl, nil)
		err := json.Unmarshal(contents, &prRole)
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			prRoles = append(prRoles, prRole)
		}
	}

	return prRoles, err
}
