/*
*/
"use strict";

var ItemID;

function pageLoaded() {
  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get('id');
  if (id == "" || id === null) return;
  ItemID = id;
  document.getElementById("id_input").value=id;
  console.log("Lookup meta for ID "+id);
  const req = { id: id };
  AppAPI.PublicService.GetMetadata(req).then(
  function(value) { /* code if successful */
      var resp = value;
      if (resp == undefined) return;
      console.log("Result: "+JSON.stringify(resp));
      showItem(resp);
  },
  function(error) { /* code if some error */
      div  = document.getElementById("log");
      div.innerHTML=JSON.stringify(error);
  }
);
}

function showItemData() {
  const req = { id: ItemID };
  AppAPI.PublicService.GetData(req).then(
  function(value) { /* code if successful */
      var resp = value.data;
      if (resp == undefined) return;
      console.log("Result: "+JSON.stringify(resp));
      var text = document.getElementById('item_data');
      text.value = resp;
      document.getElementById('data_request').style.display = 'none';
      document.getElementById('data_response').style.display = 'initial';
  },
  function(error) { /* code if some error */
      if (error.status == 404) { // TODO: check 404
        document.getElementById('data_request').style.display = 'none';
        document.getElementById("log").innerText='Данные больше не доступны'
      }
      var div  = document.getElementById("log");
      div.innerHTML=JSON.stringify(error);
  }
  );
}

function showConfirm() {
  document.getElementById('confirm').showModal();
//  document.getElementById('confirm').style.display = 'initial';
}

function acceptConfirm() {
  hideConfirm();
  showItemData();
}

function hideConfirm() {
  document.getElementById('confirm').close(false);
//  document.getElementById('confirm').style.display = 'none';
}

function sendForm(form, path) {
  var div  = document.getElementById("log"),
      xhr  = new XMLHttpRequest(),
      formData = new FormData(form);
      div.innerHTML = '';
      console.dir(formData);
  var fields = Object.fromEntries(formData);
  if (fields.title=='' || fields.data=='') {
    div.innerHTML = 'Title and data must be set';
    return false;
  }
  console.log("Ready to send"+JSON.stringify(fields));
  AppAPI.PrivateService.NewItem(fields).then(
    function(value) { /* code if successful */
        var resp = value.id;
        if (resp == undefined) return;
        console.log("Result "+resp);
        var a = document.createElement('a');
        //a.target = '_blank';
        a.href = '/?id='+resp;
        a.innerText = resp;
        div.innerText='Saved. ID: ';
        div.appendChild(a);
      },
    function(error) { /* code if some error */
        console.log('Error: ' + JSON.stringify(error));
    }
    );
  return false;
}

function pageMyLoaded() {
  const req = {};
  AppAPI.PrivateService.GetStats(req).then(
  function(value) { /* code if successful */
      console.log("Result "+value);
      var resp = value;
      if (resp == undefined) return;
      showStatItems(resp);
  },
  function(error) { /* code if some error */
      console.log('Error: ' + JSON.stringify(error));
      //div  = document.getElementById("log");
      //div.innerHTML=error; //xhr.status + ': ' + xhr.statusText;
  }
  );
}

function pageListLoaded() {
  const req = {};
  AppAPI.PrivateService.GetItems(req).then(
  function(value) { /* code if successful */
      console.log("Result "+value);
      var resp = value.items;
      if (resp == undefined) return;
      showItems(resp);
  },
  function(error) { /* code if some error */
      console.log('Error: ' + JSON.stringify(error));
      //div  = document.getElementById("log");
      //div.innerHTML=error; //xhr.status + ': ' + xhr.statusText;
  }
  );
}

function showItem(item) {
  const { elements } = document.querySelector('form#metaform')
  for (const [ key, value ] of Object.entries(item) ) {
    const field = elements.namedItem(key)
    var val = value;
    switch (key) {
      case 'status':
        val = mkStatus(value);
        break;
      case 'createdAt':
      case 'modifiedAt':
        val = mkStamp(value);
    }
    field && (field.value = val)
  }
  document.getElementById('meta').style.display = 'initial';
  // show button
  if (item.status == 'WAIT') {
    var div = document.getElementById('data_request');
    div.style.display = 'initial';
  }
}

function showStatItems(item) {
  var tbodyRef = document.getElementById('stat').getElementsByTagName('tbody')[0];
  for (const [ key, value ] of Object.entries(item) ) {
    var newRow = tbodyRef.insertRow();
    addCell(newRow,key);
    const array = ["wait", "read", "expired","total"];
    array.forEach(function (item, index) {
      addCell(newRow,value[item]);
    });
  }  
}

function showItems(item) {
  var tbodyRef = document.getElementById('items').getElementsByTagName('tbody')[0];
  for (const [ key, value ] of Object.entries(item) ) {
    var newRow = tbodyRef.insertRow();
    var meta = value.meta;
    addCellHref(newRow,'/?id='+value.id,meta.title);
    addCell(newRow,meta.group);
    addCell(newRow,mkStatus(meta.status));
    addCell(newRow,mkStamp(meta.createdAt));
    addCell(newRow,mkStamp(meta.modifiedAt));
  }  
}

function mkStatus(v){
return v.charAt(0) + v.slice(1).toLowerCase();
//  let map = new Map([ [ 1, 'Wait' ], [ 2, 'Read' ], [3, 'Expired' ], [ 4, 'Cleared' ] ]);
//  return map.get(v);
}

function mkStamp(v){
  var json = '"'+v+'"';
  var dateStr = JSON.parse(json);  
  var date = new Date(dateStr);
  return dateFormatted(date)
}

function addCell(row,text) {
  var newCell = row.insertCell();
  var newText = document.createTextNode(text);
  newCell.appendChild(newText);
}

function addCellHref(row,href,text) {
  var newCell = row.insertCell();
  var a = document.createElement('a');
  a.href = href;
  a.innerText = text;
  newCell.appendChild(a);
}

function clearForm(form) {
  console.log('reset');
  document.querySelector('form').reset();
  document.getElementById("log").innerHTML='';
  return true;
}

// code from https://gist.github.com/Peacegrove/5534309
function disable_form(form, state) {
  var elemTypes = ['input', 'button', 'textarea', 'select'];
  elemTypes.forEach(function callback(type) {
    var elems = form.getElementsByTagName(type);
    disable_elements(elems, state);
  });
}

// Disables a collection of form-elements.
function disable_elements(elements, state) {
  var length = elements.length;
  while(length--) {
    var e = elements[length];
    if (e.classList.contains('reversed')) {
      e.disabled = !state;
    } else {
      e.disabled = state;
    }
  }
}

// Format datetime
// code from http://stackoverflow.com/a/32062237
// with changed result formatting
function dateFormatted(date) {
  var month = date.getMonth() + 1;
  var day = date.getDate();
  var hour = date.getHours();
  var min = date.getMinutes();
  var sec = date.getSeconds();

  month = (month < 10 ? "0" : "") + month;
  day = (day < 10 ? "0" : "") + day;
  hour = (hour < 10 ? "0" : "") + hour;
  min = (min < 10 ? "0" : "") + min;
  sec = (sec < 10 ? "0" : "") + sec;

  var str = day + "." + month + "." + date.getFullYear() + " " + hour + ":" + min + ":" + sec;
  return str;
}
