package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sparkymat/resty"
)

type GroupsController struct {
}

func (controller GroupsController) Index(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "groups#index/n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Create(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "groups#create/n")
	for key, value := range request.Form {
		io.WriteString(response, fmt.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Update(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "groups#update/n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v\n", key, value))
	}
}
func (controller GroupsController) Destroy(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "groups#destroy/n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v\n", key, value))
	}
}

func (controller GroupsController) Show(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "groups#show\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

type GroupMembershipsController struct {
}

func (controller GroupMembershipsController) Index(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "group_memberships#index\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Create(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "group_memberships#create\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Update(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "group_memberships#update\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Destroy(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "group_memberships#destroy\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

func (controller GroupMembershipsController) Show(response http.ResponseWriter, request *http.Request, params map[string]string) {
	io.WriteString(response, "group_memberships#show\n")
	for key, value := range params {
		io.WriteString(response, fmt.Sprintf("%v = %v", key, value))
	}
}

func main() {
	groupHandler := resty.Resource("groups").Controller(GroupsController{})
	groupMembershipHandler := resty.Resource("groups", "group_memberships").Controller(GroupMembershipsController{})

	groupMembershipHandler.HandleRoot()

	groupMembershipHandler.PrintRoutes(os.Stdout)
	groupHandler.PrintRoutes(os.Stdout)

	http.ListenAndServe(":8080", nil)
}
