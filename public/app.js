new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        activeNode : "0",
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            console.info("got data");
            var msg = JSON.parse(e.data);
            console.info(msg.value);
            self.activeNode = msg.value;
            
        });
    },
    
    methods: {
        
    }
});
