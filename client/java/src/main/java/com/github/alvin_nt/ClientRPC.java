package com.github.alvin_nt;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.awt.*;
import java.awt.List;
import java.util.*;
import java.util.logging.Logger;

/**
 * Created by alvin on 29/09/15.
 */
public class ClientRPC implements Runnable{

    private static final Logger logger = Logger.getLogger(ClientRPC.class.getName());
    private java.util.List<String> JoinedChannel;
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
        JoinedChannel = new ArrayList<String>();
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
                        JoinedChannel.add(evjoin.getChannel());
                        break;
                    case LEAVE:
                        ChatService.EventLeave evleave = event.getLeave();
                        System.out.println(evleave.getName() + "has succesful leave channel");
                        JoinedChannel.remove(evleave.getChannel());
                        break;
                    case LOG:
                        ChatService.EventLog eventLog = event.getLog();
                        if(JoinedChannel.contains(eventLog.getChannel())){
                            System.out.println("["+eventLog.getChannel()+"]("+eventLog.getName()+")"+eventLog.getMessage());
                        }
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

    public void Say(String contains,String channelName) {
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandSay cmdSay = ChatService.CommandSay.newBuilder()
                .setMessage(contains)
                .setChannelName(channelName)
                .build();
        builder.setSay(cmdSay);
        ChatService.Command cmd = builder.build();
        System.out.println(cmd.toString());

        ChatService.None reply = blockingStub.sendCommand(cmd);
    }

    public void setNick(String nick){
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandChangeNick cmdChangeNick = ChatService.CommandChangeNick.newBuilder()
                .setName(nick)
                .build();
        builder.setNick(cmdChangeNick);
        ChatService.Command cmd = builder.build();
        ChatService.None reply = blockingStub.sendCommand(cmd);
    }

    public void JoinChannel(String wanttojoinchannel){
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandJoinChannel cmdJoinChannel = ChatService.CommandJoinChannel.newBuilder()
                .setName(nickname)
                .build();
        builder.setJoin(cmdJoinChannel);
        ChatService.Command cmd = builder.build();
        ChatService.None reply = blockingStub.sendCommand(cmd);
    }

    public void LeaveChannel(String wanttoleavechannel){
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandLeaveChannel cmdLeaveChannel = ChatService.CommandLeaveChannel.newBuilder()
                .setName(nickname)
                .build();
        builder.setLeave(cmdLeaveChannel);
        ChatService.Command cmd = builder.build();
        ChatService.None reply = blockingStub.sendCommand(cmd);
    }

    public void Exit(){
        ChatService.Command.Builder builder = ChatService.Command.newBuilder()
                .setSessionId(this.session_id);
        ChatService.CommandExit cmdExit = ChatService.CommandExit.newBuilder()
                .build();
        builder.setExit(cmdExit);
        ChatService.Command cmd = builder.build();
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
        if (firstWord.equals("/NICK")) {
            this.setNick(commandcontains);

        } else if (firstWord.equals("/JOIN")) {
            this.JoinChannel(commandcontains);

        } else if (firstWord.equals("/LEAVE")) {
            this.LeaveChannel(commandcontains);

        } else if (firstWord.equals("/EXIT")) {
            this.Exit();
        }else{
            if(userinput.charAt(0) == '@' && userinput.contains(" ")){
                String sendToChannel = userinput.substring(1,userinput.indexOf(" "));
                this.Say(commandcontains,sendToChannel);
            }
            else{
                for(String sendToChannel : JoinedChannel){
                    this.Say(commandcontains,sendToChannel);
                }
            }

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
