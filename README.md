## Learning Go — Notes ##

1. Routing

Create routes using http.HandleFunc and open a specific port with http.ListenAndServe.

    Each route has a handler that receives:

    http.ResponseWriter → used for sending the response

    *http.Request → contains request data

2. Handling HTTP Methods

You can check the request method using:

    if req.Method != http.MethodPost {
       // handle wrong method
    }

3. Reading Request Body

Read raw data sent in the request body:

    data, err := io.ReadAll(req.Body)  => returned types ([]byte, error)

    This returns both the request data and any potential error.


4. JSON Unmarshal Into Struct

Create a struct for the expected request data:

    var userLoginRequest userservice.UserLoginRequest

    Unmarshal the JSON body into the struct:

    err := json.Unmarshal(data, &userLoginRequest)

    Note (as a Laravel developer):

    In Laravel, request inputs are automatically mapped and validated.

    In Go, you should define a struct and manually unmarshal JSON into it.

    It feels lower-level but gives more control.

5. Error Handling

Return an error response to the client:

    http.Error(w, err.Error(), http.StatusInternalServerError)

Log errors to the terminal:

    fmt.Printf("Error: %v\n", err)

6. Returning JSON Response

Define a response struct:

    type Response struct {
        Message string      `json:"message"`
        Data    interface{} `json:"data"`
    }

Encode and return JSON:

    json.NewEncoder(w).Encode(Response{
        Message: "user found!",
        Data:    user,
    })

7. Debugging

Use log for debugging:

    log.Fatalf("DEBUG USER: %+v", user)

    Fatalf prints the log and stops the program immediately.
