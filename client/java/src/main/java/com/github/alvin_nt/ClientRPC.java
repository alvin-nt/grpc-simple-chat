package com.github.alvin_nt;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.awt.*;
import java.util.Scanner;
import java.util.logging.Logger;

/**
 * Created by alvin on 29/09/15.
 */
public class ClientRPC implements Runnable{

    private static final Logger logger = Logger.getLogger(ClientRPC.class.getName());

    private final ManagedChannel channel;

    private SimpleChatGrpc.SimpleChatStub streamStub = null;
    private final SimpleChatGrpc.SimpleChatBlockingStub blockingStub;

    /**
     * The session of the client
     */
    private int session_id = -1;

    private String nickname;

    public ClientRPC(String host, int port) {
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

    public void connect(){
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
                        ChatService.EventJoin evjoin = event.getJoin();
                        System.out.println(evjoin.getName() + " has been connected");
                        break;
                    case LEAVE:
                        ChatService.EventLeave evleave = event.getLeave();
                        System.out.println(evleave.getName() + "has succesful leave channel");
                        break;
                    case LOG:
                        ChatService.EventLog eventLog = event.getLog();
                        System.out.println("["+eventLog.getChannel()+"]("+eventLog.getName()+")"+eventLog.getMessage());
                        break;
                    default:
                        System.err.println("Unknown event");
                }
            }

            public void onError(Throwable throwable) {
                System.err.println("Something unexpected happenned!");
            }

            public void onCompleted() {
                System.err.println("Stream closed");
            }
        });
    }

    public void sendCommandSay(String command) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandSay cmdSay = ChatService.CommandSay.newBuilder()
                .setMessage("")
                .setChannelName("")
                .build();
        builder.setSay(cmdSay);
        ChatService.Command cmd = builder.build();
        System.out.println(cmd.toString());

        ChatService.None reply = blockingStub.sendCommand(cmd);
    }

    @Override
    public void run() {
        this.connect();
    }

    public void ProcessInput(String userinput) {
        String firstWord,commandcontains = null;
        firstWord = userinput.substring(0, userinput.indexOf(" "));
        commandcontains = userinput.substring(userinput.indexOf(" ") + 1, userinput.length());
        switch (firstWord) {
            case "/NICK":
                clientFound.setNick(MessageContains);
                break;
            case "/JOIN":
                if (!AllChannel.contains(MessageContains)) {
                    AllChannel.add(MessageContains);
                }
                if (!clientFound.getActiveChannel().contains(MessageContains)) {
                    clientFound.addChannel(MessageContains);
                }
                break;
            case "/LEAVE":
                if (clientFound.getActiveChannel().contains(MessageContains)) {
                    clientFound.removeChannel(MessageContains);
                }
                break;
            case "/EXIT":
                clients.remove(clientFound);
        }
    }

    public static void main(String args[]) {
        ClientRPC client = new ClientRPC("localhost", 50001);
        client.authorize("");
        new Thread(client).run();
        Scanner sc = new Scanner(System.in);
        String userinput = "";
        while(!userinput.equals("/EXIT")){
            System.out.print("Input message: ");
            userinput = sc.nextLine();

            if(!userinput.equals("") && !userinput.equals("/EXIT") && !userinput.equals(null)){
                client.ProcessInput(userinput);
            }
        }
        System.exit(0);

    }


}
