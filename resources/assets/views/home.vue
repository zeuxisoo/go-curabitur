<template>
    <div id="home">
        <div id="message-list" class="panel panel-default">
            <div class="panel-body board">
                <div class="row" v-for="message in messages" track-by="$index">
                    <p>{{ message }}</p>
                </div>
            </div>
        </div>
        <div id="message-input" class="panel panel-default">
            <div class="panel-body">
                <div class="row">
                    <div class="col-md-12">
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

#home #message-input {
    flex: 0 0 auto;
}
</style>

<script>
export default {

    data: function() {
        return {
            messages: []
        }
    },

    ready() {
        this.setBoarrdMaxHeight();

        $(window).on('resize', () => {
            this.setBoarrdMaxHeight();
        });
    },

    methods: {
        setBoarrdMaxHeight() {
            $("#message-list .board").css({
                'maxHeight': $("#message-list").height()
            });
        },

        submit() {
            this.messages.push("It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).")
        }
    }

}
</script>
