package routes

import (
    "time"

    "github.com/gorilla/websocket"
    "github.com/Pallinder/go-randomdata"
)

type ChatMessage struct {
    Kind string `json:"kind"`
    From string `json:"from"`
    Text string `json:"text"`
}

type ChatClient struct {
    Name            string
    receiver        <-chan *ChatMessage
    sender          chan<- *ChatMessage
    done            <-chan bool
    disconnect      chan<- int
    errorChannel    <-chan error
}

var clients []*ChatClient

func Chat(receiver <-chan *ChatMessage, sender chan<- *ChatMessage, done <-chan bool, disconnect chan<- int, errorChannel <-chan error) {
    ticker := time.After(30 * time.Minute)
    client := &ChatClient{
        Name:           randomdata.SillyName(),
        receiver:       receiver,
        sender:         sender,
        done:           done,
        disconnect:     disconnect,
        errorChannel:   errorChannel,
    }

    clients = append(clients, client)

    for {
        select {
            case message := <-client.receiver:
                sendMessageToOtherClients(client, message)
            case <-ticker:
                disconnect <- websocket.CloseNormalClosure
            case <-client.done:
                removeClients(client)
                return
        }
    }
}

func sendMessageToOtherClients(client *ChatClient, message *ChatMessage) {
    for _, c := range clients {
        if c != client {
            c.sender <- message
        }
    }
}

func removeClients(client *ChatClient) {
    for index, c := range clients {
        if c == client {
            clients = append(clients[:index], clients[(index+1):]...)
        }else{
            c.sender <- &ChatMessage{"status", client.Name, "Left chat"}
        }
    }
}
