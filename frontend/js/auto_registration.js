async function autoRegister() {
    const data = {
        name: document.getElementById('name').value,
        email: document.getElementById('email').value,
        cpf: document.getElementById('cpf').value,
        passport: document.getElementById('passport').value
    };

    const response = await fetch('/api/auto-register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    if (response.ok) {
        alert('Registration successful');
    } else {
        alert('Registration failed');
    }
}
