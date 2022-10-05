function SendJSON() {
    let name = document.querySelector('#username');
    let mail = document.querySelector('#user-email');
    let pass = document.querySelector('#user-pass');
    let xhr = new XMLHttpRequest();
    let url ="http://127.0.0.1:8000/post";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    let data=JSON.stringify({"name":name.value, "mail":mail.value, "pass":pass.value});
    xhr.send(data);

}