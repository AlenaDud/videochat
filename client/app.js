const wsUrl = "ws://0.0.0.0:3002/ws"; // Подставляется через переменную окружения
const loginSection = document.getElementById('login');
const conferenceSection = document.getElementById('conference');
const joinRoomButton = document.getElementById('joinRoom');
const leaveRoomButton = document.getElementById('leaveRoom');
const sendMessageButton = document.getElementById('sendMessage');

let socket, peerConnection;
const config = { iceServers: [{ urls: 'stun:stun.l.google.com:19302' }] };
const remoteVideos = document.getElementById('remoteVideos');

// Локальный видеопоток
const localVideo = document.getElementById('localVideo');
let localStream;

// Событие "Join Room"
joinRoomButton.addEventListener('click', async () => {
    const username = document.getElementById('username').value;
    const roomName = document.getElementById('room').value;

    if (!username || !roomName) {
        return alert('Both fields are required!');
    }

    document.getElementById('roomName').textContent = roomName;

    loginSection.hidden = true;
    conferenceSection.hidden = false;

    try {
        await createRoom(username, roomName);
        startWebRTC(username, roomName);
    } catch (error) {
        alert(`Error: ${error.message}`);
    }
});

// Событие "Leave Room"
leaveRoomButton.addEventListener('click', () => {
    if (socket) {
        socket.send(JSON.stringify({ type: 'leave' }));
        socket.close();
    }

    if (peerConnection) {
        peerConnection.close();
    }

    conferenceSection.hidden = true;
    loginSection.hidden = false;
});

// WebRTC логика
function startWebRTC(username, roomName) {
    socket = new WebSocket(`${wsUrl}/ws/join?room=${roomName}&user=${username}`);
    peerConnection = new RTCPeerConnection(config);

    // Получение локального потока
    navigator.mediaDevices.getUserMedia({ video: true, audio: true }).then((stream) => {
        localStream = stream;
        localVideo.srcObject = stream;

        // Добавляем треки в соединение
        stream.getTracks().forEach((track) => peerConnection.addTrack(track, stream));
    });

    // ICE-кандидаты
    peerConnection.onicecandidate = (event) => {
        if (event.candidate) {
            socket.send(JSON.stringify({ type: 'ice-candidate', candidate: event.candidate }));
        }
    };

    // Получение удаленного видеопотока
    peerConnection.ontrack = (event) => {
        const remoteVideo = document.createElement('video');
        remoteVideo.srcObject = event.streams[0];
        remoteVideo.autoplay = true;
        remoteVideos.appendChild(remoteVideo);
    };

    // Подключение к сигнальному серверу
    socket.onmessage = async (message) => {
        const data = JSON.parse(message.data);

        if (data.type === 'offer') {
            await peerConnection.setRemoteDescription(new RTCSessionDescription(data.offer));
            const answer = await peerConnection.createAnswer();
            await peerConnection.setLocalDescription(answer);
            socket.send(JSON.stringify({ type: 'sdp/answer', answer }));
        } else if (data.type === 'answer') {
            await peerConnection.setRemoteDescription(new RTCSessionDescription(data.answer));
        } else if (data.type === 'ice-candidate') {
            await peerConnection.addIceCandidate(new RTCIceCandidate(data.candidate));
        } else if (data.type === 'chat') {
            addChatMessage(data.message, 'remote');
        }
    };
}

// Создание комнаты
async function createRoom(username, roomName) {
    const response = await fetch(`${wsUrl}/ws/create`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, room: roomName }),
    });

    if (!response.ok) {
        throw new Error(`Failed to create room: ${response.statusText}`);
    }
}

// Добавление сообщения чата
function addChatMessage(message, sender) {
    const messages = document.getElementById('messages');
    const messageElement = document.createElement('div');
    messageElement.textContent = `${sender}: ${message}`;
    messages.appendChild(messageElement);
}

// Событие "Send Message"
sendMessageButton.addEventListener('click', () => {
    const messageInput = document.getElementById('messageInput');
    const message = messageInput.value;

    if (socket && message) {
        socket.send(JSON.stringify({ type: 'chat', message }));
        addChatMessage(message, 'you');
        messageInput.value = '';
    }
});
