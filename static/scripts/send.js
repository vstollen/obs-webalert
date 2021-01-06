let ws

window.addEventListener("load", function () {
    ws = new WebSocket(`wss://${window.location.host}/socket`)

    ws.onmessage = function (evt) {
        let received_msg = evt.data
        alert(`Received message: ${received_msg}`)
    }

    ws.onclose = function () {
        alert("Connection is closed...")
    }

    // Access the form element...
    const form = document.forms['message-input'];

    const messageList = document.getElementById("messages")

    // ...and take over its submit event.
    form.addEventListener("submit", function (event) {
        event.preventDefault();

        sendText(form.elements.message.value);
        form.elements.message.value = ""
    });

    const dragzone = document.getElementById('dropzone')

    dragzone.addEventListener("drop", dropHandler)
    dragzone.addEventListener("dragover", dragOverHandler)

    function sendText(message) {
        ws.send(message)

        const newMessageElement = document.createElement("li")
        newMessageElement.innerHTML = "message: " + message
        messageList.appendChild(newMessageElement)
    }
});

function dropHandler(e) {
    console.debug('File dropped')

    e.preventDefault()

    let dt = e.dataTransfer
    let files = dt.files

    let file = files[0]

    if (!file.type.match(/^image\/*/)) {
        console.warn(`Invalid File Type: ${file.type}`)
        return
    }

    let reader = new FileReader()
    reader.onload = function(pe) {
        let fileData = pe.target.result
        ws.send(fileData)
        console.debug('File has been sent')
    }

    reader.readAsArrayBuffer(file)
}

function dragOverHandler(e) {
    console.debug('File in dropzone')

    e.preventDefault()
}