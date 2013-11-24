$(document).ready(function() {
    var s = $('#screen');
    connected = true;
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
var connected;

function createScreen(s) {
    logMessage(document.URL.replace("http://","ws://") + '/ws');
    connection = new WebSocket(document.URL.replace("http://","ws://") + '/ws');

    connection.onerror = wsError;
    connection.onopen = wsOpen;
    connection.onclose = wsLogger('Connection closed');
    connection.onmessage = wsHandler;
}

function wsOpen() {
    connected = true;
    logMessage('Connection opened');
}

function wsError(error) {
    logMessage(error);
}

function wsLogger(msg) {
    var s = $('#screen');
    return function () {
        logMessage(msg);
        if (connected) {
            connected = false;
            createScreen(s);
        }
    }
    
}

function wsHandler(e) {
    if (e.data.match("setup finish apk")   == "setup finish apk") {
        document.getElementById("log").innerHTML = ""
        logMessage(e.data);
        logMessage("downloading");
		getApk(e.data.replace(/setup finish apk / ,""))
        return
    }
    if (e.data.match("setup finish win")  == "setup finish win") {
        document.getElementById("log").innerHTML = ""
        logMessage(e.data);
        logMessage("downloading");
		getWin(e.data.replace(/setup finish win / ,""))
        return
    }
    logMessage(e.data);
    //d = $.parseJSON(e.data);
    //$('#blob').css('margin-left', d.X);
    //$('#blob').css('margin-top', d.Y);
}

function getApk(urlname) {
    location.href = "/" + urlname + "/download_apk"
}
function getWin(urlname) {
    location.href = "/" + urlname + "/download_win"
}
function resetApk(urlname) {
    if (connection) { connection.send('setup apk ' + urlname); }
}
function resetWin(urlname) {
    if (connection) { connection.send('setup win ' + urlname); }
}
