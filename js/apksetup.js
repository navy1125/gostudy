$(document).ready(function() {
    var s = $('#screen');
    connected = true;
    logMessage("connect to server..");
    createScreen(s);

    $(document).keydown(function(e) {
    	var msg = '{"Id":"keyboard","Data"=""}';

    	jmsg = $.parseJSON(msg)
		if (e.which == 37) {
			jmsg.Data='l';
        }
        if(e.which == 39) {
        	jmsg.Data = 'r';
        }
        if(e.which == 38) {
        	jmsg.Data = 'u';
        }
        if(e.which == 40) {
        	jmsg.Data = 'd';
        }
        if (connection) { connection.send(JSON.stringify(jmsg)); }
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
	d = $.parseJSON(e.data);
	if (d.Id == "setup finish apk") {
        document.getElementById("log").innerHTML = ""
        logMessage(d.Data);
        logMessage("downloading");
		getApk(d.Data)
        return
    }
	//if (e.data.match("setup finish win")  == "setup finish win") {
	if (d.Id  == "setup finish win") {
        document.getElementById("log").innerHTML = ""
        logMessage(d.Data);
        logMessage("downloading");
        getApk(d.Data)
        //getWin(e.data.replace(/setup finish win / ,""))
        return
    }
    logMessage(d.Data);
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
	var msg = '{"Id":"setup apk","Data":"'+ urlname + '"}';
    if (connection) { connection.send(msg); }
}
function resetWin(urlname) {
	var msg = '{"Id":"setup win","Data":"' + urlname + '"}';
	if (connection) { connection.send(msg); }
}
