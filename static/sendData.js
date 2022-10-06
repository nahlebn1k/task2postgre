function SendJSON() {
    let name = document.querySelector('#username');
    let mail = document.querySelector('#user-email');
    let pass = document.querySelector('#user-pass');
    let url = conf.urlConf;
    let data=JSON.stringify({"name":name.value, "mail":mail.value, "pass":pass.value});

    fetch(url,{
        method: 'POST',
        body: data,
        headers:{
            'Content-type': 'application/json',
        }

    }).then(function (response){
        console.log("OK")
        return response.text()
    })



}