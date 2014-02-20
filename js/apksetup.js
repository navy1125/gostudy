$(document).ready(function() {
    var s = $('#screen');
    connected = true;
    logMessage("connect to server..");
    createScreen(s);

    $(document).keydown(function(e) {
    	var msg = '{"Id":"keyboard","Data"=""}';

    	var jmsg = $.parseJSON(msg)
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

function errMessage(m) {
	$('#err').append('<li>' + m + '</li>');
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
        logMessage("downloading:" + d.Data);
		getApk(d.Data)
        return
    }
	//if (e.data.match("setup finish win")  == "setup finish win") {
	else if (d.Id  == "setup finish win") {
        document.getElementById("log").innerHTML = ""
        logMessage(d.Data);
        logMessage("downloading:"+d.Data);
        getWin(d.Data)
        //getWin(e.data.replace(/setup finish win / ,""))
        return
	}
	else if (d.Id == "0") {
		logMessage(d.Data);
	}
	else if (d.Id == "1") {
		errMessage(d.Data);
	}
	else {
		logMessage(d.Id + ":" + d.Data);
	}
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
function resetHuoDong(urlname) {
	var msg = '{"Id":"reset huodong","Data":"' + urlname + '"}';
	if (connection) { connection.send(msg); }
}
function updateResouce(urlname) {
	var msg = '{"Id":"update resource","Data":"' + urlname + '"}';
	if (connection) { connection.send(msg); }
}
function restartServer(urlname) {
    var msg = '{"Id":"restart server","Data":"' + urlname + '"}';
    if (connection) { connection.send(msg); }
}
