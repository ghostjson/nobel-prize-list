# Nobel Prize Winners Listing Web App

#### Technologies And Dependencies

- go lang
- gin
- vuejs
- bootstrap
- jquery
- datatables

##### Prerequisites

- Install [Go](https://golang.org/)
- Download go dependency `go get -u github.com/gin-gonic/gin`

##### Start Server

To start the server, use `go build && ./source`

### API Documentation

##### Upload json file

Upload json file to the server, then parse it and convert it to a yaml file.

_Route_: `api/upload`
_Payload_:
`file`: Json file to be uploaded
_Response_:
`{ "message": "successfully finished", "content": "uploaded payload content" }`

##### Read file

Read file that present on the server. Can be json or yaml file.

_Route_: `api/read-file/:filename`
_Payload_:
`filename`: name of the file to be read
_Response_:
`{ "content": "uploaded payload content" }`

##### Save as YAML

Save the json file in a given name in YAML format

_Route_: `api/save/yaml/:filename/:json_filename`
_Payload_:
`filename`: name of the yaml file
`json_filename`: name of the json file
_Response_:
`{ "message": "successfully finished", "content": "uploaded payload content" }`
