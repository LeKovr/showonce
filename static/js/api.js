
var ItemID;

function pageLoaded() {
  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get('id');
  if (id == "" || id ===null) return;
  ItemID = id;
  document.getElementById("id_input").value=id;
  console.log("Lookup meta for ID "+id)
  var xhr = new XMLHttpRequest();
  xhr.open('GET', '/api/item?id='+id);
  xhr.setRequestHeader('Accept', 'application/json'); // TODO: Accept?
  xhr.onreadystatechange = function() {
    if (xhr.readyState != 4) return;
    if (xhr.status != 200) {
      console.log(xhr.status + ': ' + xhr.statusText);
      div  = document.getElementById("log");
      div.innerHTML=xhr.status + ': ' + xhr.statusText;
    } else {
      var resp = JSON.parse(xhr.responseText);
      if (resp == undefined) return;
      console.dir(resp);
      showItem(resp);
    }
  }
  xhr.send();
}

function showItem(item) {
  const { elements } = document.querySelector('form#metaform')
  console.dir(elements);
for (const [ key, value ] of Object.entries(item) ) {
  const field = elements.namedItem(key)
  field && (field.value = value)
}  
var dropZone = document.getElementById('meta');
dropZone.style.display = 'initial';

// show button
if (item.status == 1) {
  var div = document.getElementById('data_request');
  div.style.display = 'initial';

}
console.log(">>>>",item.status)
}

function showItemData() {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/api/item?id='+ItemID);
  xhr.setRequestHeader('Accept', 'application/json'); // TODO: Accept?
  xhr.onreadystatechange = function() {
    if (xhr.readyState != 4) return;
    if (xhr.status != 200) {
      console.log(xhr.status + ': ' + xhr.statusText);
    } else {
      var resp = JSON.parse(xhr.responseText);
      if (resp == undefined) return;
      var text = document.getElementById('item_data');
      text.value = resp;
      document.getElementById('data_request').style.display = 'none';
      document.getElementById('data_response').style.display = 'initial';
    }
  }
  xhr.send();
}

function sendForm(form, path) {
  var div  = document.getElementById("log"),
      xhr  = new XMLHttpRequest(),
      formData = new FormData(form);
  div.innerHTML = '';
  console.dir(formData);

  var data= JSON.stringify(Object.fromEntries(formData));
data.exp=Number(data.exp);
xhr.open('POST', path);
xhr.setRequestHeader('Accept', 'application/json'); // TODO: Accept?
xhr.onreadystatechange = function() {
  if (xhr.readyState != 4) return;
  if (xhr.status != 200) {
    console.log(xhr.status + ': ' + xhr.statusText);
  } else {
    var resp = JSON.parse(xhr.responseText);
    if (resp == undefined) return;

    var a = document.createElement('a');
    //a.target = '_blank';
    a.href = '/?id='+resp;
    a.innerText = resp;
    div.appendChild(a);
  }
}
xhr.send(data);
return false;
}


function clearForm(form) {
  console.log('reset');
  documentFiles = null;
  document.getElementById('list').innerHTML = '';
  document.querySelector('form').reset(); // clear file input
  document.getElementById("log").innerHTML='';

  return true;
}

// code from https://gist.github.com/Peacegrove/5534309
function disable_form(form, state) {
  var elemTypes = ['input']; //, 'button', 'textarea', 'select'];
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
