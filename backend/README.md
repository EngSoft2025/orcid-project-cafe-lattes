# Structure

`api/` -- all api information and handling (payloads, handlers, controller)
`service/` -- only business logic here (**NO external dependency should be imported here**)
`repository/` -- used from service layer to inject external dependency and actually do all the operations
`model/` -- maybe remove but it can be used to transfer data between layers or to standarize a way of processing types
`main.go` -- starts everything

## Flow

All routes are implemented on `api` pacakge, and the callback for each route is should be on `handlers.go` following the same pattern (they're all methods from the `Controller` type). These handlers forward the task to the service layer (as it's done in the example route). The service layer only does business logic, so it should not inject any external dependencies. By business logic, we mean that it orchestrates every call to the repository layer without directly depending on any dependency (for apis, database, data processing, etc). This ensures that the layers are decoupled and more modular.
The repository layer injects every necessary external dependency. In other words, it actually does the job, by effectively fetching an endpoint, inserting or reading some data from somewhere, and so on.

So, the repository layer should offer small functions that does specific and unique tasks so that the service layer can orchestrate these calls and treat the error properly.

## Requests

### Swagger

To see the detailed requests definitions, parameters and responses, see:

```txt
localhost:8080/api/swagger/index.html
```

### Biography Query

- GET
- endpoint: `/api/searchBiography`
- query: `orcid_id` needs to be a valid orcid id

#### Possible Reponses

##### Success

###### Code: 200

###### Body

```go
type BiographyData struct {
Name      string `json:"name"`
 OrcidId   string `json:"orcid_id"`
 Site      string `json:"site"`
 Biography string `json:"biography"`

 Emails         []string `json:"emails"`
 ResearcherUrls []string `json:"researcher_urls"`
 Country        []string `json:"countries"`
 Keywords       []string `json:"keywords"`

 ExternalIDs []struct {
  ExternalIDType  string `json:"external-id-type"`
  ExternalIDValue string `json:"external-id-value"`
  ExternalIDURL   struct {
   Value string `json:"value"`
  } `json:"external-id-url"`
  ExternalIDRelationship string `json:"external-id-relationship"`
 } `json:"external-identifier"`
}
```

##### Incorrect Query

###### Code: 400

###### error: orcid_id is required

##### Invalid OrcidId or OrcidAPI changed

###### Code: 404

###### error: Could not get {orcid_id} biography data

---

### Work Query

- GET
- endpoint: `/api/searchWork`
- query: `orcid_id` needs to be a valid orcid id

#### Possible Reponses

##### Success

###### Code: 200

###### Body

```go
type WorkData struct {
 Publications []struct {
  Title   string `json:"title"`
  Doi     string `json:"doi"`
  Url     string `json:"url"`
  Type    string `json:"type"`
  Year    string `json:"year"`
  Journal string `json:"journal"`
 } `json:"publications"`
}
```

##### Incorrect Query

###### Code: 400

###### error: orcid_id is required

##### Invalid OrcidId or OrcidAPI changed

###### Code: 404

###### error: Could not get {orcid_id} biography data

---

### Search By Name Query

- GET
- endpoint: `/api/searchResearchersByName`
- query: `name`

#### Possible Reponses

##### Success

###### Code: 200

###### Body

The response will be the ResearcherResults struct
```go
type Researcher struct {
	Orcid_id string `json:"orcid_id"`
	Name     string `json:"name"`
}

type ResearcherResults struct {
	Researchers []Researcher `json:"researchers"`
	NumFound    int          `json:"num_found"`
}
```

##### Incorrect Query

###### Code: 400

###### error: name is required

##### Could not find name

###### Code: 404

###### error: Could not get researchers by name: {name}
