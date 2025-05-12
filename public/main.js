// Connect to Socket
const socket = io("http://localhost:3000");

new Vue({
    el: '#chat-app',
    created(){

        // Emit the messages in the room
        socket.emit("join room", "chat_room");

        socket.on("broad message", (msg) => {
            //console.log("Room messages:", msg);
            this.messages.push({
                text: msg,
                date: new Date().toLocaleString()
            })

        });

    },
    data: {
        message: '',
        messages: []
    },
    methods: {
        sendMessage(){
            // event named: 'chat message'
            socket.emit('chat message', this.message);
            this.message = '';
        }
    }
})
