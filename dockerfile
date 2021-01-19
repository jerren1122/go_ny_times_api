#inherit homebrew
FROM golang
RUN go get -u github.com/gorilla/mux
CMD git clone https://github.com/jerren1122/go_ny_times_api.git; cd go_ny_times_api; go build; ./go_ny_times_api;