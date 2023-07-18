
function removeFromDb(item){
   fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
       if (res.status == 200){
           window.location.pathname = "/"
       }
   })
}

function updateDb(item) {
   let input = document.getElementById(item)
    let value = input.value
    fetch(`/update/$id?item=${item}&value=${value}`, {method: "PUT"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
}
