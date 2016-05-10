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
                        <input type="text" class="form-control" id="" placeholder="Type a message ..." v-model="message" v-on:keyup.enter="submit">
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
            message : "",
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
                let message = {
                    kind: "status",
                    from: this.username,
                    text: "Enter to chat"
                };

                this.addMessage(message);
                this.socket.send(JSON.stringify(message));
            }

            this.socket.onclose = (e) => {
                swal({
                    title: "Oops!",
                    text : "Session timeout! Please reload this page.",
                    type : "warning"
                }, () => {
                    window.location.reload();
                });
            }

            this.socket.onmessage = (e) => {
                let message = JSON.parse(e.data);

                this.addMessage({
                    kind: message.kind,
                    from: message.from,
                    text: message.text
                });
            }

            this.socket.onerror = (e) => {
                swal({
                    title: "Error!",
                    text : "Got some unknown error! Please reload this page.",
                    type : "error"
                }, () => {
                    console.log(e);
                });
            }
        },

        submit() {
            let message = {
                kind: "message",
                from: this.username,
                text: this.message,
            };

            this.addMessage(message);
            this.socket.send(JSON.stringify(message));
        },

        addMessage(message) {
            this.messages.push(message);

            this.$nextTick(() => {
                $("#message-list .board").scrollTop($('#message-list .board')[0].scrollHeight);
            })
        }
    }

}
</script>
