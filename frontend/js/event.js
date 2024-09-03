async function fetchEvents() {
    const response = await fetch('/api/events', {
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    });
    const events = await response.json();
    const eventTable = document.getElementById('event-table');

    events.forEach(event => {
        const row = eventTable.insertRow();
        row.insertCell(0).innerText = event.name;
        row.insertCell(1).innerText = event.location;
        row.insertCell(2).innerText = new Date(event.date).toLocaleDateString();
    });
}
