package request

import (
	"uilliam.com/sip/types"
)

// "path": "/getUser",
//     "method": "GET",
//     "description": "Retrieve user information by UUID.",
//     "threads": 2,
//     "request": {
//       "fields": {
//         "key": {
//           "type": "string",
//           "keylocation": "header", // or "header" or "body" or "query"
//           "keyName": "X-API-Key",
//           "required": true
//         },
//         "uuid": {
//           "type": "string",
//           "keylocation": "query", // or "header" or "body" or "query"
//           "keyName": "uuid",
//           "required": false
//         }
//       }
    // },

func GenerateRateLimitCode(config types.RateLimitDefinition) (string, []string) {
}


// i was goin to work on this but then got distracted by the enterypoint so i 
// just wanted to make sure that i got my thoughts down before i forgot them

// also new keybord so i am still getting used to it - hence the typos

we will need to parce this data and just the code 
that we generate should consist of a main func that 
will be the entrypoint for the server, that will also 
start the threads in a worker pool where the main entryoint 
will put connections into the pool and workers will pull from 
them to handle the request, lastly there main entry point should 
start the logger and get that all setup when you build the
worker thread possess dol not close the brackets for the codgen or 
return the request because we will still need to do 
intermeatit things with the data 