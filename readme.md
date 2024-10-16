# 📔Smart Note API

A RESTful API for managing notes using Go, PostgreSQL (via Docker), and GORM with UUIDs.

## Features

- CRUD operations for notes
- UUIDs for note IDs
- Environment-based configuration using `.env` files
- PostgreSQL as the database (via Docker)
- GORM for ORM (Object-Relational Mapping)
- Signup & Login added.
- Token-based authentication ( JWT ).

## Features to be added

- API Documentation (Swagger)
- And many more to be added 🔥

## Requirements

- **Go** (v1.18+)
- **Docker** (for PostgreSQL)
- **Git**
- **Go modules** (handled via `go mod`)

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/Shihab-Munna/smart-notes-api.git
```

### 2. Install Go modules

Navigate into the project directory and run:

```bash
cd smart-notes-api
go mod tidy
```

### 3. Set up PostgreSQL using Docker

Pull the PostgreSQL Docker image and run the container:

```bash
docker run --name smart_note_db -e POSTGRES_USER=smart_note_user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=smart_note -p 5432:5432 -d postgres
```

Wait for the container to start up, then check if it's running by executing:

```bash
docker ps
```

Enable the UUID extension:

Once the container is running, access the PostgreSQL instance by running:

```bash
docker exec -it smart_note_db psql -U smart_note_user -d smart_note
```

Run the following SQL command to enable the UUID extension:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Exit the PostgreSQL shell:

```bash
\q
```

### 4. Set up environment variables

Copy the .env.example file to create your own .env file:

```bash
cp .env.example .env
```

Open the .env file and update the placeholder values with your actual PostgreSQL connection information:

```bash
PORT=5001
DB_DSN=host=localhost user=smart_note_user password=password dbname=smart_note port=5432 sslmode=disable
```

Note: Replace password with the actual password you set in the Docker command.

### 5. Run the application

Run the Go application:

```bash
go run main.go
```

If everything is set up correctly, the server should now be running at <http://localhost:5001>.

## API Endpoints (Under Development) 🏗️

1. **Create a new note:**

   - **URL:** `POST /notes`
   - **Description:** Create a new note.
   - **Request Body (JSON):**

   ```json
   {
     "title": "Note Title",
     "content": "Note content goes here"
   }
   ```

   - **Response (Success 201):**

   ```json
   {
     "id": "uuid-of-note",
     "title": "Note Title",
     "content": "Note content goes here"
   }
   ```

2. **Retrieve all notes:**

   - **URL:** `GET /notes`
   - **Description:** Fetch all notes.
   - **Response (Success 200):**

   ```json
   [
     {
       "id": "uuid-of-note",
       "title": "Note Title",
       "content": "Note content goes here"
     }
   ]
   ```

3. **Retrieve a note by UUID:**

   - **URL:** `GET /notes/{id}`
   - **Description:** Retrieve a note by its UUID.
   - **Response (Success 200):**

   ```json
   {
     "id": "uuid-of-note",
     "title": "Note Title",
     "content": "Note content goes here"
   }
   ```

4. **Update a note by UUID:**

   - **URL:** `PUT /notes/{id}`
   - **Description:** Update a note by its UUID.
   - **Request Body (JSON):**

   ```json
   {
     "title": "Updated Title",
     "content": "Updated content goes here"
   }
   ```

   - **Response (Success 200):**

   ```json
   {
     "id": "uuid-of-note",
     "title": "Updated Title",
     "content": "Updated content goes here"
   }
   ```

5. **Delete a note by UUID:**
   - **URL:** `DELETE /notes/{id}`
   - **Description:** Delete a note by its UUID.
   - **Response (Success 204 - No Content)**

## Development

- Modify the `.env.example` file as needed and ensure your environment variables are set up correctly before running the project locally.
- Use `go mod tidy` to keep dependencies clean and up to date.

## Contributing

Feel free to fork the repository and make contributions. For significant changes, please open an issue first to discuss what you would like to change.

## MIT License

Copyright (c) 2024 Shihab Munna

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
