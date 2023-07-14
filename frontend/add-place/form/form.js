// function send(name) {
//     // Create a new FormData object
//     const formData = new FormData();
//
//     // Add the username and password as form fields
//     formData.append('name', name);
//
//     // Make a POST request to the server
//     fetch('/api/places', {
//         method: 'POST',
//         body: formData
//     })
//         .then(response => response.json())
//         .then(json => {
//             // Handle the response from the server
//             let feedbackDOM = document.getElementById("feedback");
//             if(json.success === true){
//                 feedbackDOM.innerText = json.data
//                 feedbackDOM.style.color = "lightgreen";
//             } else {
//                 feedbackDOM.style.color = "red";
//                 feedbackDOM.innerText = json.message
//             }
//
//             window.setTimeout(()=>{
//                 feedbackDOM.style.color = "black";
//                 feedbackDOM.innerText = "Please add another :-)"
//             }, 4500); // 4.5 seconds
//
//             window.setTimeout(()=>{
//                 feedbackDOM.style.color = "";
//                 feedbackDOM.innerText = ""
//             }, 6200); // 6.2 seconds (1.7 seconds after the other)
//
//             fetchPlacesAndPopulateTable()
//         })
//         .catch(error => {
//             // Handle any errors that occurred during the request
//             console.error('Error:', error);
//         });
// }
//
//
//
// document.getElementById("add-place").addEventListener("submit", (event)=>{
//     event.preventDefault();
//     let name = document.getElementById("name").value
//     send(name)
// })