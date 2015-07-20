package gojira

import (
	"encoding/json"
	"fmt"
	"regexp"
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

/*
Search for project. Gets list of all projects available to user and display those with name matching
regexp .*searchString.* match is case insensitive.

    GET <jira_url>/rest/api/2/project

Params
	searchString string
*/
func (j *Jira) SearchProjects(searchString string) (*[]JiraProject, error) {
	url := j.BaseUrl + j.ApiPath + project_url
	contents := j.buildAndExecRequest("GET", url, nil)

	projects := new([]JiraProject)
	err := json.Unmarshal(contents, &projects)
	if err != nil {
		fmt.Printf("%s", err)
	}
	searchResult := new([]JiraProject)
	var rgx = regexp.MustCompile(fmt.Sprintf("(?i).*%s.*", searchString))
	for _, project := range *projects {
		if rgx.MatchString(project.Name) {
			*searchResult = append(*searchResult, project)
		}
	}
	return searchResult, nil
}
