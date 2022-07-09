#!/bin/bash

# Test API methods
#   Usage:
#     bash test.sh XXXXXXX
#
#   Where XXXXXXX - your gitea app token (see /user/settings/applications)


ACCEPT="Accept: application/json"
CT=""
#"Content-Type: application/json"

GITEA_TOKEN=$1
echo $GITEA_TOKEN
TOKEN="X-narra-token: $GITEA_TOKEN"

[[ "$API_HOST" ]] || API_HOST=http://showonce.dev.lan:8080

# mux.HandleFunc("/my/api/new", srv.ItemCreate)
do_create() {
  DATA=$(cat <<EOF
{
    "title":"Message title",
    "group":"default",
    "exp":"90",
    "exp_unit": "s",
    "data":"secret"
}
EOF
)
action="/my/api/new"
id=$(curl -gs -H "$ACCEPT" -H "$CT" -H "$TOKEN" -d "$DATA" ${API_HOST}$action)
echo $id
}

# GET mux.HandleFunc("/api/item", srv.Item)
do_item() {
    local id=$1
    action="/api/item"
    curl -gs -H "$ACCEPT" -H "$CT" ${API_HOST}$action?id=$id
}

# mux.HandleFunc("/api/item", srv.Item)
do_data() {
    local id=$1
    action="/api/item"
    curl -gs -H "$ACCEPT" -H "$CT" -d "" ${API_HOST}$action?id=$id
}

# mux.HandleFunc("/my/api/items", srv.Items)
do_items() {
  action="/my/api/items"
  rv=$(curl -gs -H "$ACCEPT" -H "$CT" -H "$TOKEN" -d "$DATA" ${API_HOST}$action)
  echo $rv
}

# mux.HandleFunc("/my/api/stat", srv.Stats)
do_stat() {
  action="/my/api/stat"
  rv=$(curl -gs -H "$ACCEPT" -H "$CT" -H "$TOKEN" -d "$DATA" ${API_HOST}$action)
  echo $rv
}

do_stat

id=$(do_create | jq -r .)
echo ">>ID: $id"
[[ "$id" ]] || exit

do_item $id | jq -r .
do_data $id | jq -r .
echo "---- now is empty"
do_item $id | jq -r .
do_data $id 

do_items | jq -r .
do_stat | jq -r .
