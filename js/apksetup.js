$(document).ready(function() {
    var s = $('#screen');

    logMessage("Connecting...");
    createScreen(s);

    $(document).keydown(function(e) {
        if(e.which == 37) {
            if(connection) { connection.send('l'); }
        }
        if(e.which == 39) {
            if(connection) { connection.send('r'); }   
        }
        if(e.which == 38) {
            if(connection) { connection.send('u'); }   
        }
        if(e.which == 40) {
            if(connection) { connection.send('d'); }   
        }

    });
});

function logMessage(m) {
    $('#log').append('<li>'+m+'</li>');
}

var connection;

function createScreen(s) {
    connection = new WebSocket('ws://180.168.197.87:8080/ws');

    connection.onerror = wsError;
    connection.onopen = wsOpen;
    connection.onclose = wsLogger('Connection closed');
    connection.onmessage = wsHandler;
}

function wsOpen() {
    logMessage('Connection opened');
}

function wsError(error) {
    logMessage(error);
}

function wsLogger(msg) {
    return function() { logMessage(msg); }
}

function wsHandler(e) {
    logMessage(e.data);
    d = $.parseJSON(e.data);
    $('#blob').css('margin-left', d.X);
    $('#blob').css('margin-top', d.Y);
}

function getApk() {
    logMessage('Connection opened');
}
function resetApk() {
    if (connection) { connection.send('setup'); }
}
