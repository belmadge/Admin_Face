Here is the README.md file written in Markdown:

markdown
Copy code
# Admin FaceDoor/CP

Admin FaceDoor/CP is a comprehensive administration panel designed to manage users, events, and notifications.
This application is built using Go with Gin for the backend, and HTML, CSS, and JavaScript for the frontend. 
It supports Docker for containerization, providing an easy setup and consistent environment for development and deployment.

## Table of Contents

1. [Features](#features)
2. [Requirements](#requirements)
3. [Installation](#installation)
4. [Running the Application](#running-the-application)
5. [Project Structure](#project-structure)
6. [Environment Variables](#environment-variables)
7. [API Endpoints](#api-endpoints)
8. [Usage](#usage)
9. [Contributing](#contributing)
10. [License](#license)

## Features

- **User Management**: Create, edit, delete, and manage users with different roles and permissions.
- **Event Management**: Create, edit, delete, and manage events. Integration with auto-registration and real-time event tracking.
- **Auto-Registration**: Accessible self-registration for new users with features like photo capture and audio guidance.
- **Notifications and Communications**: Send notifications via email and SMS, and keep track of communication history.
- **Dashboards**: Customizable dashboards for detailed analytics and monitoring.
- **Data Export**: Export event data and participant information to CSV.
- **Accessibility Enhancements**: Features like HandTalk integration, audio guidance, and more.

## Requirements

- **Docker** and **Docker Compose**: To run the application in a containerized environment.
- **Go**: For local development and running the backend server directly (if not using Docker).
- **MySQL**: As the database for storing application data.

## 1. Installation

### Clone the Repository

Clone the repository to your local machine using:

```bash
git clone https://github.com/belmadge/Admin_Face.git
cd Admin_Face
```

## 2. Set Up Environment Variables

Create a `.env` file in the backend directory to configure your environment variables. 
An example of the required environment variables:

```env
DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=facedoor_admin
JWT_SECRET=your-secret-key
PORT=8080
```

## 3. Install Dependencies

If you are running the application locally without Docker, you need to install Go dependencies:

```bash
cd backend
go mod tidy
```

## 4. Build Docker Containers

Make sure you have Docker and Docker Compose installed on your machine.
To build and start the application, run:

```bash
docker-compose up --build -d
```

## Running the Application

Once the Docker containers are running, you can access the application:

- **Frontend**: Open your web browser and go to `http://localhost:8080`.
- **Backend API**: Accessible at `http://localhost:8080/api`.

To stop the application, press `Ctrl+C` in the terminal where Docker Compose is running, or run:

```bash
docker-compose down
```

## Project Structure

- **`backend/`**: Contains all backend-related code and configurations.
- **`frontend/`**: Contains all frontend-related files, including HTML, CSS, and JavaScript.
- **`migrations/`**: SQL files for database migrations.

## Environment Variables

Ensure you have the following environment variables set in your `.env` file:

- `DB_USER`: MySQL database username.
- `DB_PASSWORD`: MySQL database password.
- `DB_HOST`: Hostname of the MySQL server.
- `DB_PORT`: Port number of the MySQL server.
- `DB_NAME`: Name of the MySQL database.
- `JWT_SECRET`: Secret key used for JWT token generation.
- `PORT`: Port number on which the backend server runs.

## API Endpoints

Here are some of the key API endpoints available:

### User Management

- `POST /register`: Register a new user.
- `POST /login`: Login a user and get a JWT token.
- `GET /api/users`: Retrieve all users.
- `GET /api/users/:id`: Retrieve a specific user.
- `PUT /api/users/:id`: Update a user.
- `DELETE /api/users/:id`: Delete a user.

### Event Management

- `GET /api/events`: Retrieve all events.
- `POST /api/events`: Create a new event.
- `PUT /api/events/:id`: Update an event.
- `DELETE /api/events/:id`: Delete an event.
- `GET /api/export/events`: Export event data to CSV.

### Notifications

- `GET /api/notifications`: Retrieve all notifications.
- `POST /api/notifications`: Create a new notification.
- `DELETE /api/notifications/:id`: Delete a notification.

## Usage

To use the Admin FaceDoor/CP application:

1. **Register or Login**: Access the application at `http://localhost:8080` and use the registration or login form.
2. **Manage Users**: Navigate to the "Users" section to create, edit, or delete users.
3. **Manage Events**: Navigate to the "Events" section to manage event details and participants.
4. **Send Notifications**: Use the "Notifications" section to send out communications.
5. **Dashboard and Analytics**: Access the "Dashboard" for detailed insights and analytics.
6. **Auto-Registration**: Use the "Auto Registration" for new user sign-ups with additional accessibility features.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
