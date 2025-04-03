package Logic

import (
	"fmt"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func (logic *LogicManager) RFID() {
	for {
		logic.RFIDstatus = false
		logic.PortNames, _ = logic.RFIDFindSerialPort()
		currentIndex := logic.PortNamesIndex
		if currentIndex >= 0 && len(logic.PortNames) != 0 {
			logic.RFIDstatus = true
			println("RFID CONNECTED")
			mode := &serial.Mode{
				BaudRate: 115200,
			}
			port, err := serial.Open(logic.PortNames[currentIndex].Name, mode)

			if err != nil {
				log.Fatalf("Error opening serial port: %v", err)
			}

			logic.RFIDListenSerial(port)
			port.Close()
		}
	}
}

// Automatically detects an available serial port
func (logic *LogicManager) RFIDFindSerialPort() ([]*enumerator.PortDetails, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return []*enumerator.PortDetails{}, err
	}

	if len(ports) == 0 {
		return []*enumerator.PortDetails{}, fmt.Errorf("no serial ports found")
	}

	// Return the first available port
	return ports, nil
}

func (logic *LogicManager) RFIDListenSerial(port serial.Port) {
	buf := make([]byte, 100) // Temporary buffer for incoming data
	var messageBuffer string // Stores full message until it's complete
	lastReceivedID := ""     // Stores last received ID to avoid duplicates

	for {
		n, err := port.Read(buf)

		if err != nil {
			logic.RFIDstatus = false
			log.Printf("Error reading serial: %v", err)
			//continue // Prevent crashing, keep listening
			break
		}

		if n > 0 {
			// Append received data to messageBuffer
			messageBuffer += string(buf[:n])
			//println(n)
			// Check if the message is complete (assuming newline "\n" is the delimiter)
			if strings.Contains(messageBuffer, "\n") {
				// Split messages in case multiple are received together
				messages := strings.Split(messageBuffer, "\n")

				// Process each complete message
				for _, msg := range messages {
					cleanMsg := strings.TrimSpace(msg) // Remove extra spaces

					// Ignore empty messages
					if len(cleanMsg) == 0 {
						continue
					}
					// Only print if different from last received ID
					if cleanMsg != lastReceivedID {
						println(cleanMsg)
						lastReceivedID = cleanMsg // Update last received
					}
				}

				// Clear the buffer after processing messages
				messageBuffer = ""
			}
		}
		time.Sleep(50 * time.Millisecond) // Shorter sleep for real-time processing
	}
}
