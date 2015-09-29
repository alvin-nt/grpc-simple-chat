package com.github.alvin_nt;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;

import java.awt.*;
import java.awt.List;
import java.io.Closeable;
import java.util.*;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

/**
 * Created by alvin on 29/09/15.
 */
public class ClientRPC implements Runnable, Closeable {

    private static final Logger logger = Logger.getLogger(ClientRPC.class.getName());
    private java.util.List<String> JoinedChannel = new ArrayList<String>();

    private String host;
    private int port;

    private final ManagedChannel channel;
    private final SimpleChatGrpc.SimpleChatStub streamStub;
    private final SimpleChatGrpc.SimpleChatBlockingStub blockingStub;

    /**
     * The session of the client
     */
    private int session_id = -1;

    private String nickname;

    public ClientRPC(String host, int port) {
        this.host = host;
        this.port = port;

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
                        ChatService.EventJoin evjoin = event.getJoin();
                        System.out.println("[" + evjoin.getChannel() + "] " + evjoin.getName() + " has joined");
                        JoinedChannel.add(evjoin.getChannel());
                        break;
                    case LEAVE:
                        ChatService.EventLeave evleave = event.getLeave();
                        System.out.println("[" + evleave.getChannel() + "] " + evleave.getName() + " left");
                        JoinedChannel.remove(evleave.getChannel());
                        break;
                    case LOG:
                        ChatService.EventLog eventLog = event.getLog();
                        System.out.println("[" + eventLog.getChannel() + "](" + eventLog.getName() + ") " + eventLog.getMessage());
                        break;
                    default:
                        System.err.println("Unknown event");
                }
            }

            public void onError(Throwable throwable) {
                System.err.println("Caught error from server: " + throwable.getMessage());
            }

            public void onCompleted() {
                System.err.println("Stream closed");
            }
        });
    }

    public void reconnect() {
        authorize(this.nickname);
        connect();
    }

    public void Say(String contains, String channelName) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandSay cmdSay = ChatService.CommandSay.newBuilder()
                .setMessage(contains)
                .setChannelName(channelName)
                .build();
        builder.setSay(cmdSay);
        ChatService.Command cmd = builder.build();

        blockingStub.sendCommand(cmd);
    }

    public void setNick(String nick) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandChangeNick cmdChangeNick = ChatService.CommandChangeNick.newBuilder()
                .setName(nick)
                .build();
        builder.setNick(cmdChangeNick);
        ChatService.Command cmd = builder.build();

        nickname = nick;

        blockingStub.sendCommand(cmd);
    }

    public void JoinChannel(String channel) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandJoinChannel cmdJoinChannel = ChatService.CommandJoinChannel.newBuilder()
                .setName(channel)
                .build();
        builder.setJoin(cmdJoinChannel);
        ChatService.Command cmd = builder.build();

        try {
            blockingStub.sendCommand(cmd);
        } catch (StatusRuntimeException e) {
            if (e.getMessage().contains("registered")) {
                System.err.println("Already connected to channel " + channel);
            } else {
                throw e;
            }
        }
    }

    public void LeaveChannel(String channel) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandLeaveChannel cmdLeaveChannel = ChatService.CommandLeaveChannel.newBuilder()
                .setName(channel)
                .build();
        builder.setLeave(cmdLeaveChannel);
        ChatService.Command cmd = builder.build();

        blockingStub.sendCommand(cmd);
    }

    public void Exit() {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandExit cmdExit = ChatService.CommandExit.newBuilder()
                .build();
        builder.setExit(cmdExit);
        ChatService.Command cmd = builder.build();

        blockingStub.sendCommand(cmd);

        this.close();
    }

    public void run() {
        this.connect();
    }

    public void ProcessInput(String input) throws InputMismatchException, StatusRuntimeException {
        String command, commandArg = "";

        if (input.startsWith("/")) {
            // parse command
            int spaceIndex = input.indexOf(' ', 1);

            if (spaceIndex != -1) {
                command = input.substring(1, spaceIndex);

                if (input.length() > spaceIndex + 1) {
                    commandArg = input.substring(spaceIndex + 1);
                }
            } else {
                command = input.substring(1);
            }

            if (command.equalsIgnoreCase("nick")) {
                this.setNick(commandArg);
            } else if (command.equalsIgnoreCase("join")) {
                this.JoinChannel(commandArg);
            } else if (command.equalsIgnoreCase("leave")) {
                this.LeaveChannel(commandArg);
            } else if (command.equalsIgnoreCase("exit")) {
                this.Exit();
            }

        } else if (input.startsWith("@")) {
            // send to specific channel
            if (input.contains(" ") && input.charAt(input.length() - 1) != ' ') {
                int spaceIndex = input.indexOf(' ');
                String channel = input.substring(1, spaceIndex);
                String message = input.substring(spaceIndex + 1);

                this.Say(message, channel);
            } else {
                throw new InputMismatchException("No message specified");
            }
        } else {
            // send to all subscribed channels
            this.Say(input, "");
        }
    }

    public void close() {
        try {
            channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            // silently ignore
        }
    }

    public static void main(String args[]) {
        ClientRPC client = new ClientRPC("localhost", 50001);
        client.authorize("");
        new Thread(client).run();

        // the input loop
        Scanner sc = new Scanner(System.in);

        String input = "";
        do {
            System.out.printf("[%s] > ", client.nickname);
            input = sc.nextLine();

            try {
                client.ProcessInput(input);
            } catch (InputMismatchException e) {
                System.err.println("Input error: " + e.getMessage());
            } catch (StatusRuntimeException e) {
                System.err.println("Server connection error: " + e.getMessage());

                if (e.getStatus() == Status.UNAVAILABLE) {
                    // try to reconnect...
                    int retryCount = 1;

                    do {
                        try {
                            client.reconnect();
                            break;
                        } catch (StatusRuntimeException e1) {
                            try {
                                System.err.println("Server connection error: " + e.getMessage());
                                if (retryCount < 5) {
                                    retryCount++;
                                }
                                Thread.sleep(1000 * retryCount);
                            } catch (InterruptedException e2) {
                                break;
                            }
                        }
                    } while (retryCount < 5);

                    if (retryCount == 5 && client.channel.isTerminated()) {
                        throw e;
                    }
                }
            }

            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                // ignore
            }
        } while (!input.toLowerCase().startsWith("/exit"));
    }
}
