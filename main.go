package main

import (
	"io"
	"net/http"
	"os"

	"github.com/kr/pretty"
	"github.com/sparkymat/resty"
	shttp "github.com/sparkymat/webdsl/http"
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

func (controller GroupsController) Show(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#show\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupsController) Members(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "groups#members\n")
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

func (controller GroupMembershipsController) Destroy(response http.ResponseWriter, request *http.Request, params map[string][]string) {
	io.WriteString(response, "group_memberships#destroy\n")
	for key, value := range params {
		io.WriteString(response, pretty.Sprintf("%v = %v", key, value))
	}
}

func main() {
	r := resty.ResourceRouter{}
	r.Resource([]string{"groups"}, GroupsController{}).Except(resty.Destroy).
		Member("members", []shttp.Method{shttp.Get})
	r.Resource([]string{"groups", "group_memberships"}, GroupMembershipsController{}).Only(resty.Create, resty.Index, resty.Destroy)

	r.PrintRoutes(os.Stdout)

	r.HandleRoot()

	http.ListenAndServe(":5000", nil)
}
