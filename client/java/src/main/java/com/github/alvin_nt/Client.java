package com.github.alvin_nt;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.util.logging.Logger;

/**
 * Created by alvin on 29/09/15.
 */
public class Client {

    private static final Logger logger = Logger.getLogger(Client.class.getName());

    private final ManagedChannel channel;

    private final SimpleChatGrpc.SimpleChatStub streamStub;
    private final SimpleChatGrpc.SimpleChatBlockingStub blockingStub;

    /**
     * The session of the client
     */
    private int session_id = -1;

    private String nickname;

    public Client(String host, int port) {
        channel = ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext(true)
                .build();
        streamStub = SimpleChatGrpc.newStub(channel);
        blockingStub = SimpleChatGrpc.newBlockingStub(channel);
    }

    public void authorize(String name) {
        ChatService.RequestAuthorize.Builder builder = ChatService.RequestAuthorize.newBuilder();

        if (name.isEmpty()) {
            builder.setName(name);
        }
        ChatService.RequestAuthorize request = builder.build();

        ChatService.ResponseAuthorize response = blockingStub.authorize(request);

        session_id = response.getSessionId();
        nickname = response.getName();

        System.out.printf("Connected with id %d and name %s\n", session_id, nickname);
    }

    public void connect() {
        ChatService.RequestConnect request = ChatService.RequestConnect.newBuilder()
                .setSessionId(session_id)
                .build();
        streamStub.connect(request, new StreamObserver<ChatService.Event>() {
            public void onNext(ChatService.Event event) {
                switch (event.getEventCase()) {
                    case NONE:
                        // just ignore
                        break;
                    case JOIN:
                        ChatService.EventJoin ev = event.getJoin();
                        System.out.println(ev.getName() + " has been connected");
                        break;
                }
            }

            public void onError(Throwable throwable) {

            }

            public void onCompleted() {
                System.err.println("Stream closed");
            }
        });
    }

    public void sendCommand(String command) {
        
    }

    public static void main(String args[]) {
        Client client = new Client("localhost", 50001);

        client.authorize("");
        client.connect();
    }
}
