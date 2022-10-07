function SendJSON() {
    event.preventDefault();
    let name = document.querySelector('#username');
    let mail = document.querySelector('#user-email');
    let pass = document.querySelector('#user-pass');
    let url = window.location.origin+"/post";
    let data=JSON.stringify({"name":name.value, "mail":mail.value, "pass":pass.value});

    fetch(url,{
        method: 'POST',
        headers:{
            'Content-Type': 'application/json; charset=UTF-8'
        },
        body: data

    }).then(response => {
        if (!response.ok) {
            throw new Error(`HTTP error: ${response.status}`);
        }
        return response.text()
    })
}