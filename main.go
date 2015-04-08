package main

import (
	"io"
	"net/http"

	"github.com/kr/pretty"
	"github.com/sparkymat/resty"
)

type GroupsController struct {
}

func (controller GroupsController) Index(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#index\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Create(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#create\n")
	for key, value := range request.Form {
		io.WriteString(response, pretty.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Update(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#update\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Destroy(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#destroy\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v\n", key, value))
	}
}

func (controller GroupsController) Show(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#show\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

type GroupMembershipsController struct {
}

func (controller GroupMembershipsController) Index(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#index\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Create(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#create\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Update(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#update\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Destroy(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#destroy\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Show(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#show\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func main() {
	r := resty.ResourceRouter{}
	r.Resource([]string{"groups"}, GroupsController{})
	r.Resource([]string{"groups", "group_memberships"}, GroupMembershipsController{})

	r.HandleRoot()

	http.ListenAndServe(":5000", nil)
}
