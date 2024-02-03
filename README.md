Contact Book API

The Contact Book API is a simple RESTful web service designed to manage a digital contact book. It provides endpoints for basic CRUD operations, allowing users to create, retrieve, update, and delete contacts. This project is beginner-friendly, making it an excellent addition to your portfolio.
Features

    Create: Add new contacts with details such as name, phone number, and email.
    Retrieve: Fetch contact information, either for a specific contact or the entire contact book.
    Update: Modify existing contact details, ensuring your digital address book stays up to date.
    Delete: Remove contacts that are no longer needed.

Technologies Used

    Golang: The backend of the API is built using the Go programming language.
    RESTful API: Follows RESTful principles for designing web services.
    JSON: Communicates data in JSON format, a lightweight and widely used data interchange format.

Getting Started

    Clone the Repository:

    bash

git clone https://github.com/Shinichi23/contact-book.git

Navigate to the Project Directory:

bash

cd contact-book

Run the Application:

bash

    go run main.go

    The API will be accessible at http://localhost:8080.

API Endpoints

    GET /contacts: Retrieve all contacts.
    GET /contacts/{id}: Retrieve a specific contact.
    POST /contacts: Create a new contact.
    PUT /contacts/{id}: Update an existing contact.
    DELETE /contacts/{id}: Delete a contact.

Example Usage
Retrieve All Contacts

bash

curl http://localhost:8080/contacts

Retrieve a Specific Contact

bash

curl http://localhost:8080/contacts/1

For more examples and detailed API documentation, refer to Swagger Documentation.
Contributing

Contributions are welcome! If you find a bug, have a suggestion, or want to add new features, feel free to open an issue or submit a pull request.
License

This project is licensed under the MIT License.
