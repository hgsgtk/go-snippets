package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sshHost := os.Getenv("SSH_HOST")
	sshPort := os.Getenv("SSH_PORT")
	sshUser := os.Getenv("SSH_USER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	buf, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))
	if err != nil {
		panic(err)
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}

	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshCon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", sshHost, sshPort), sshConfig); err == nil {
		defer sshCon.Close()
		mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{sshCon}).Dial)
		if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)); err == nil {
			fmt.Printf("Successfully connected to the database\n")
			if tables, err := db.Query("SHOW TABLES"); err == nil {
				for tables.Next() {
					var name string
					tables.Scan(&name)
					fmt.Printf("Table Name: %s\n", name)
				}
				tables.Close()
			} else {
				fmt.Printf("Failure: %s", err.Error())
			}
			db.Close()
		} else {
			fmt.Printf("Failure: failed to connect to the database: %s\n", err.Error())
		}
	}
}
