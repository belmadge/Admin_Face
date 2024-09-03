async function fetchUsers() {
    const response = await fetch('/api/users', {
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    });
    const users = await response.json();
    const userTable = document.getElementById('user-table');

    users.forEach(user => {
        const row = userTable.insertRow();
        row.insertCell(0).innerText = user.name;
        row.insertCell(1).innerText = user.email;
        row.insertCell(2).innerText = user.role;
    });
}
