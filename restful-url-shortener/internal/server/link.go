package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"elahi-arman.github.com/example-http-server/internal/datastore"
	"github.com/julienschmidt/httprouter"
)

type Response_Get struct {
	Rel     string       `json:"rel"`
	Href    string       `json:"href"`
	Method  string       `json:"method"`
	Headers []GET_Header `json:"headers,omitempty"`
}

type GET_Header struct {
	Owner string `json:"owner"`
}

// GetLink is the function called when a user makes a request to retrieve a certain link
func (s *serverImpl) GetLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ps are the parameters attached to this route. the paramter to ByName()
	// must match the name of the link from main.go
	linkId := ps.ByName("link")

	// do some preemptive error checking
	if linkId == "" {
		fmt.Println("GetLink: no linkId provided")
		w.WriteHeader(400)
		return
	}

	// access the datastore attached to the server and try to fetch the link
	data, err := s.linkStore.GetLink(linkId)
	if errors.Is(err, &datastore.NotFoundError{}) {
		fmt.Printf("GetLink: no entry for linkId=%s\n", linkId)
		w.WriteHeader(404)
		return
	}

	// HATEOAS links
	links := []Response_Get{
		{
			Rel:    "user-links",
			Href:   "/getuserlinks",
			Method: "GET",
			Headers: []GET_Header{
				{
					Owner: "Owner-Name",
				},
			},
		},
		{
			Rel:    "self",
			Href:   fmt.Sprintf("/l/%s", linkId),
			Method: "GET",
		},
		{
			Rel:    "delete",
			Href:   "/deletelink/{onwername}/{ID}",
			Method: "DELETE",
		},
	}

	// Structure the response
	response := struct {
		Data  interface{}
		Links interface{}
	}{
		Data:  data,
		Links: links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// // return a 302 to redirect users
	// fmt.Printf("GetLink: found link for linkId=%s, redirecting to url=%s", link.Id, link.Url)
	// w.Header().Add("Location", link.Url) // the location header is the destination URL
	// w.WriteHeader(302)                   // 302 informs the client to read the Location header for a redirection
}

// createLinkParams represents the structure of the request body to
// a CreateLink function call

type createLinkParams struct {
	Url string `json:"url"`
	// temporary, eventually we'll replace this by retrieving from context
	Owner string `json:"owner"`
}

func (s *serverImpl) CreateLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// retrieve the value of the content-type header, if none is specified
	// the request should be rejected
	contentType := r.Header.Get("content-type")
	if contentType == "" {
		fmt.Println("CreateLink: no content-type header is sent")
		w.WriteHeader(400) // the status message will automatically be filled in
		return
	}

	var url string
	var owner string
	if strings.Contains(contentType, "json") {
		// read the body of the request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("CreateLink: error while reading body of request %v\n", err)
			w.WriteHeader(400)
			return
		}

		// convert the request body into json
		lp := &createLinkParams{}
		err = json.Unmarshal(body, lp)
		if err != nil {
			fmt.Printf("CreateLink: error while unmarshalling err=%v. \n body=%s\n", err, body)
			w.WriteHeader(400)
			return
		}

		url = lp.Url
		owner = lp.Owner
	} else if strings.Contains(contentType, "form") {

		// when dealing with form data, call ParseForm to trigger parsing
		// then r.Form will have a map of the form values
		r.ParseForm()
		if formUrl, ok := r.Form["url"]; !ok || len(formUrl) == 0 || formUrl[0] == "" {
			fmt.Println("CreateLink: url key is not part of form data")
			w.Header().Add("Location", fmt.Sprintf("/public?error=%s", "cannot create a link without a url"))
			w.WriteHeader(303)
			return
		} else {
			url = formUrl[0]
		}

		if formOwner, ok := r.Form["owner"]; !ok || len(formOwner) == 0 || formOwner[0] == "" {
			fmt.Println("CreateLink: owner key is not part of form data")
			w.Header().Add("Location", fmt.Sprintf("/public?error=%s", "cannot create a link without an owner"))
			w.WriteHeader(303)
			return
		} else {
			owner = formOwner[0]
		}
	}

	// call the datastore function
	link, err := s.linkStore.CreateLink(url, owner)
	if err != nil {
		fmt.Printf("CreateLink: error while creating a link err=%v\n", err)
		w.WriteHeader(500)
		return
	}

	// redirect users
	w.Header().Add("Location", fmt.Sprintf("/public?link=%s", link.Id))
	w.WriteHeader(303)
}

type urlList struct {
	Owner string   `json:"owner"`
	Urls  []string `json:"urls"`
}

// read a header / body to get a user
// return a list of links in json format where Owner == user passed in
func (s *serverImpl) GetUserLinks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ownerName, ok := r.Header["Owner"]

	if !ok {
		fmt.Println("Key not found")
		w.WriteHeader(404)
		return
	}

	// access the datastore attached to the server and try to fetch the link
	links, err := s.linkStore.GetUserLinks(ownerName[0])

	if errors.Is(err, &datastore.NotFoundError{}) {
		fmt.Printf("GetUserLinks: no owner found =%s\n", ownerName)
		w.WriteHeader(204)
		return
	}

	response := urlList{
		Owner: ownerName[0],
	}

	for _, obj := range links {

		response.Urls = append(response.Urls, obj.Url)
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(jsonData)
}
func (s *serverImpl) DeleteLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ownerName := ps.ByName("onwername")

	//do some preemptive error checking
	if ownerName == "" {
		fmt.Println("No ownerName: no ownerName provided")
		w.WriteHeader(400)
		return
	}

	ID := ps.ByName("ID")

	//do some preemptive error checking
	if ID == "" {
		fmt.Println("No ID: no ID provided")
		w.WriteHeader(400)
		return
	}
	// access the datastore attached to the server and try to fetch the link
	err := s.linkStore.DeleteLink(ID, ownerName)
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte("No records were deleted"))
		return
	}

	w.Write([]byte("Delete link successfully"))
}
