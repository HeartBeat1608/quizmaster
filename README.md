# QuizMaster

Welcome to **QuizMaster**! This README provides an overview of the project, setup instructions, and technical details to help you get started.

## Description

**QuizMaster** is a simple and engaging quiz game web application. It allows quiz masters to create and manage quizzes, while players can join and participate in these quizzes. The application is designed to be user-friendly and provides a seamless experience for both quiz creators and participants.

## Features

- **Quiz Creation**: Quiz masters can create and customize quizzes with various types of questions.
- **Quiz Participation**: Players can join quizzes, answer questions, and see their scores in real-time.
- **Real-time Updates**: The app provides real-time updates to ensure a dynamic and interactive quiz experience.

## Technical Information

QuizMaster is built using a monorepo structure, combining a Golang-based server and a SvelteKit-based client. This setup ensures a robust and performant application.

## Development Roadmap

You are read about the features planned in the product [here](ROADMAP.md)

### Server

- **Language**: Go (Golang)
- **Framework**: Standard library and other Go packages as needed
- **Functionality**: Handles quiz creation, data management, and API endpoints for client interaction.

### Client

- **Framework**: SvelteKit
- **Language**: JavaScript/TypeScript
- **Functionality**: Provides the user interface for both quiz masters and players, communicates with the server via API calls.

## Installation

To set up QuizMaster on your local machine, follow these steps:

### Prerequisites

- **Golang**: Ensure Go is installed. You can download it from [here](https://golang.org/dl/).
- **Node.js and npm**: Ensure Node.js and npm are installed. You can download them from [here](https://nodejs.org/).

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/HeartBeat1608/quizmaster.git
   cd quizmaster
   ```

2. **Set Up the Server**:

   ```bash
   cd server
   go mod tidy
   go run main.go
   ```

3. **Set Up the Client**:
_We recommend using pnpm as package manager to keep a single and consistent lockfile system_

   ```bash
   cd client
   pnpm install
   pnpm dev
   ```

4. **Access the Application**:

   Open your browser and go to `http://localhost:5173` to access the QuizMaster application.  
   You could also go to the terminal running the client and press 'o' and Enter, to open the page in your default browser.

## Usage

- **For Quiz Masters**: Log in, create quizzes by adding questions, and manage existing quizzes.
- **For Players**: Enter the quiz code provided by the quiz master, answer the questions, and view your score.

## Contributing

We welcome contributions to QuizMaster! If you have ideas for improvements or have found bugs, please open an issue or submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

# Licensing

## Current License (Pre-v1.0)
QuizMaster is licensed under the Apache License 2.0. You may use, modify, and distribute this software in accordance with the terms of the license. See the LICENSE file for more details.

## Future License (Post-v1.0)
Upon reaching version 1.0, QuizMaster will transition to a proprietary license. Existing versions released under the Apache License 2.0 will remain under that license, but future versions will be governed by the proprietary license terms. For commercial use and licensing inquiries, please contact us at sales@quizmaster.com.

## Contact

If you have any questions or need further assistance, please feel free to contact us at support@quizmaster.com.

---

Thank you for using QuizMaster! We hope you enjoy creating and participating in quizzes. Happy quizzing!
