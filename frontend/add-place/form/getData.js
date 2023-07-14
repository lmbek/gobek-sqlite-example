function fetchPlacesAndPopulateTable() {
    fetch('/api/places')
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                const places = data.data;

                // Get the table element by its ID
                const table = document.getElementById('placesTable');

                // Clear the existing table rows
                table.innerHTML = '';

                // Create a table row for each place
                places.forEach(place => {
                    const row = document.createElement('tr');

                    // Create a table cell for the place name
                    const cell = document.createElement('td');
                    const textNode = document.createTextNode(place.name);
                    cell.appendChild(textNode);
                    row.appendChild(cell);

                    // Append the row to the table
                    table.appendChild(row);
                });
            } else {
                console.error('Failed to fetch places:', data.error);
            }
        })
        .catch(error => {
            console.error('Error fetching places:', error);
        });
}

fetchPlacesAndPopulateTable()