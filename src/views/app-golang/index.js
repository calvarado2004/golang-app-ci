function removeFromDb(item){
   fetch(`/app-golang/delete?item=${item}`, {method: "Delete"}).then(res =>{
       if (res.status == 200){
           window.location.pathname = "/app-golang"
       }
   })
}

function updateDb(item) {
   let input = document.getElementById(item)
   let newitem = input.value
   fetch(`/app-golang/update?olditem=${item}&newitem=${newitem}`, {method: "PUT"}).then(res =>{
       if (res.status == 200){
       alert("Database updated")
           window.location.pathname = "/app-golang"
       }
   })
}
