$(document).ready(function() {
    var s = $('#screen');

    logMessage("connect to server..");
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
    $('#log').append('<li>' + m + '</li>');
    window.scrollTo(0, document.body.scrollHeight)
}

var connection;

function createScreen(s) {
    connection = new WebSocket(document.URL.replace("http://","ws://") + 'ws');

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
    var s = $('#screen');
    return function () { logMessage(msg); createScreen(s);}
    
}

function wsHandler(e) {
    logMessage(e.data);
    //d = $.parseJSON(e.data);
    //$('#blob').css('margin-left', d.X);
    //$('#blob').css('margin-top', d.Y);
}

function getApk() {
    location.href = "/download_apk"
}
function getWin() {
    location.href = "/download_win"
}
function monitorApk() {
    if (!connection) { createScreen(""); }
}
