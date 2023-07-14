function send(name) {
    // Create an object with the username and password
    const payload = {
        name: name
    };

    // Make a POST request to the server
    fetch('/api/places', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
        .then(response => response.json())
        .then(json => {
            // Handle the response from the server
            let feedbackDOM = document.getElementById("feedback");
            if(json.success === true){
                feedbackDOM.innerText = json.data
                feedbackDOM.style.color = "lightgreen";
            } else {
                feedbackDOM.style.color = "red";
                feedbackDOM.innerText = json.message
            }

            window.setTimeout(()=>{
                feedbackDOM.style.color = "black";
                feedbackDOM.innerText = "Please add another :-)"
            }, 4500); // 4.5 seconds

            window.setTimeout(()=>{
                feedbackDOM.style.color = "";
                feedbackDOM.innerText = ""
            }, 6200); // 6.2 seconds (1.7 seconds after the other)

            fetchPlacesAndPopulateTable()
        })
        .catch(error => {
            console.log(error)
        });
}

document.getElementById("add-place").addEventListener("submit", (event)=>{
    event.preventDefault();
    let name = document.getElementById("name").value
    send(name)
})
console.log("added")