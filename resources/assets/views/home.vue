<template>
    <div id="home">
        <div id="message-list" class="panel panel-default">
            <div class="panel-body board">
                <div class="row" v-for="message in messages" track-by="$index">
                    <div class="panel panel-success">
                        <div class="panel-heading">
                            <div class="from">{{ message.from }}</div>
                        </div>
                        <div class="panel-body">
                            <div class="message">
                                {{ message.text }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div id="message-input" class="panel panel-default">
            <div class="panel-body">
                <div class="row">
                    <div class="col-md-2">
                        <input type="text" class="form-control" id="" placeholder="Username" v-model="username">
                    </div>
                    <div class="col-md-10">
                        <input type="text" class="form-control" id="" placeholder="Type a message ..." v-on:keyup.enter="submit">
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
#home {
    display: flex;
    flex-direction: column;
    height: 94%;
}

#home #message-list {
    flex: 1 0 auto;
}

#home #message-list .board {
    overflow-y: scroll;
}

#home #message-list .board .row {
    padding: 0px 10px;
}

#home #message-input {
    flex: 0 0 auto;
}
</style>

<script>
import faker from 'faker'

export default {

    data: function() {
        return {
            socket  : null,
            username: faker.name.findName(),
            messages: []
        }
    },

    ready() {
        this.setBoarrdMaxHeight();

        $(window).on('resize', () => {
            this.setBoarrdMaxHeight();
        });

        this.setWebSocket();
    },

    methods: {
        setBoarrdMaxHeight() {
            $("#message-list .board").css({
                'maxHeight': $("#message-list").height()
            });
        },

        setWebSocket() {
            this.socket = new WebSocket("ws://" + location.host + "/chat");

            this.socket.onopen = (e) => {
                console.log("Onopen")
            }

            this.socket.onclose = (e) => {
                console.log("Onclose")
            }

            this.socket.onmessage = (e) => {
                console.log("onmessage", e)
            }

            this.socket.onerror = (e) => {
                console.log("onerror")
            }
        },

        submit() {
            this.messages.push({
                kind: "message",
                from: this.username,
                text: "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like)."
            }),

            this.socket.send(JSON.stringify({
                kind: "message",
                from: this.username,
                text: "This is a test message (" + new Date() + ")"
            }))
        }
    }

}
</script>
