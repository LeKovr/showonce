#!/bin/bash

# Test API methods
#   Usage:
#     bash test.sh XXXXXXX
#
#   Where XXXXXXX - your gitea app token (see /user/settings/applications)


ACCEPT="Accept: application/json"
CT=""
#"Content-Type: application/json"
DEBUG=${DEBUG:-}
GITEA_TOKEN=$1
echo $GITEA_TOKEN
TOKEN="X-narra-token: $GITEA_TOKEN"

[[ "$API_HOST" ]] || API_HOST=http://shon.dev.test:8002

# mux.HandleFunc("/my/api/new", srv.ItemCreate)
do_create() {
  DATA=$(cat <<EOF
{
    "title":"Message title",
    "group":"default",
    "expire":"90",
    "expire_unit": "s",
    "data":"secret"
}
EOF
)
action="/my/api/new"
id=$(curl -gs -H "$ACCEPT" -H "$CT" -H "$TOKEN" -d "$DATA" ${API_HOST}$action)
[ -z $DEBUG ] || echo "RESP: $id" >&2
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
    curl -gs -H "$ACCEPT" -H "$CT" -d "" ${API_HOST}$action/$id
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

do_stat | jq -r .

id=$(do_create | jq -r .id)
echo ">>ID: $id"
[[ "$id" ]] || exit

do_item $id | jq -r .
do_data $id | jq -r .

echo "---- now is empty"
do_item $id | jq -r .
do_data $id | jq -r .

do_items | jq -r .
do_stat | jq -r .
