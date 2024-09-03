async function sendNotification() {
    const message = document.getElementById('message').value;
    const type = document.getElementById('type').value;

    const response = await fetch('/api/notifications', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ message, type })
    });

    if (response.ok) {
        alert('Notification sent successfully');
    } else {
        alert('Failed to send notification');
    }
}
