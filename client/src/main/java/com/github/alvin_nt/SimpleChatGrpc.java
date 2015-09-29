package com.github.alvin_nt;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;

@javax.annotation.Generated("by gRPC proto compiler")
public class SimpleChatGrpc {

  private SimpleChatGrpc() {}

  public static final String SERVICE_NAME = "proto.SimpleChat";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi
  public static final io.grpc.MethodDescriptor<com.github.alvin_nt.ChatService.RequestAuthorize,
      com.github.alvin_nt.ChatService.ResponseAuthorize> METHOD_AUTHORIZE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.SimpleChat", "Authorize"),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.RequestAuthorize.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.ResponseAuthorize.getDefaultInstance()));
  @io.grpc.ExperimentalApi
  public static final io.grpc.MethodDescriptor<com.github.alvin_nt.ChatService.RequestConnect,
      com.github.alvin_nt.ChatService.Event> METHOD_CONNECT =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING,
          generateFullMethodName(
              "proto.SimpleChat", "Connect"),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.RequestConnect.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.Event.getDefaultInstance()));
  @io.grpc.ExperimentalApi
  public static final io.grpc.MethodDescriptor<com.github.alvin_nt.ChatService.Command,
      com.github.alvin_nt.ChatService.None> METHOD_SEND_COMMAND =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.SimpleChat", "SendCommand"),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.Command.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(com.github.alvin_nt.ChatService.None.getDefaultInstance()));

  public static SimpleChatStub newStub(io.grpc.Channel channel) {
    return new SimpleChatStub(channel);
  }

  public static SimpleChatBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new SimpleChatBlockingStub(channel);
  }

  public static SimpleChatFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new SimpleChatFutureStub(channel);
  }

  public static interface SimpleChat {

    public void authorize(com.github.alvin_nt.ChatService.RequestAuthorize request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.ResponseAuthorize> responseObserver);

    public void connect(com.github.alvin_nt.ChatService.RequestConnect request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.Event> responseObserver);

    public void sendCommand(com.github.alvin_nt.ChatService.Command request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.None> responseObserver);
  }

  public static interface SimpleChatBlockingClient {

    public com.github.alvin_nt.ChatService.ResponseAuthorize authorize(com.github.alvin_nt.ChatService.RequestAuthorize request);

    public java.util.Iterator<com.github.alvin_nt.ChatService.Event> connect(
        com.github.alvin_nt.ChatService.RequestConnect request);

    public com.github.alvin_nt.ChatService.None sendCommand(com.github.alvin_nt.ChatService.Command request);
  }

  public static interface SimpleChatFutureClient {

    public com.google.common.util.concurrent.ListenableFuture<com.github.alvin_nt.ChatService.ResponseAuthorize> authorize(
        com.github.alvin_nt.ChatService.RequestAuthorize request);

    public com.google.common.util.concurrent.ListenableFuture<com.github.alvin_nt.ChatService.None> sendCommand(
        com.github.alvin_nt.ChatService.Command request);
  }

  public static class SimpleChatStub extends io.grpc.stub.AbstractStub<SimpleChatStub>
      implements SimpleChat {
    private SimpleChatStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleChatStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleChatStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleChatStub(channel, callOptions);
    }

    @java.lang.Override
    public void authorize(com.github.alvin_nt.ChatService.RequestAuthorize request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.ResponseAuthorize> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_AUTHORIZE, getCallOptions()), request, responseObserver);
    }

    @java.lang.Override
    public void connect(com.github.alvin_nt.ChatService.RequestConnect request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.Event> responseObserver) {
      asyncServerStreamingCall(
          getChannel().newCall(METHOD_CONNECT, getCallOptions()), request, responseObserver);
    }

    @java.lang.Override
    public void sendCommand(com.github.alvin_nt.ChatService.Command request,
        io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.None> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_SEND_COMMAND, getCallOptions()), request, responseObserver);
    }
  }

  public static class SimpleChatBlockingStub extends io.grpc.stub.AbstractStub<SimpleChatBlockingStub>
      implements SimpleChatBlockingClient {
    private SimpleChatBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleChatBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleChatBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleChatBlockingStub(channel, callOptions);
    }

    @java.lang.Override
    public com.github.alvin_nt.ChatService.ResponseAuthorize authorize(com.github.alvin_nt.ChatService.RequestAuthorize request) {
      return blockingUnaryCall(
          getChannel().newCall(METHOD_AUTHORIZE, getCallOptions()), request);
    }

    @java.lang.Override
    public java.util.Iterator<com.github.alvin_nt.ChatService.Event> connect(
        com.github.alvin_nt.ChatService.RequestConnect request) {
      return blockingServerStreamingCall(
          getChannel().newCall(METHOD_CONNECT, getCallOptions()), request);
    }

    @java.lang.Override
    public com.github.alvin_nt.ChatService.None sendCommand(com.github.alvin_nt.ChatService.Command request) {
      return blockingUnaryCall(
          getChannel().newCall(METHOD_SEND_COMMAND, getCallOptions()), request);
    }
  }

  public static class SimpleChatFutureStub extends io.grpc.stub.AbstractStub<SimpleChatFutureStub>
      implements SimpleChatFutureClient {
    private SimpleChatFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleChatFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleChatFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleChatFutureStub(channel, callOptions);
    }

    @java.lang.Override
    public com.google.common.util.concurrent.ListenableFuture<com.github.alvin_nt.ChatService.ResponseAuthorize> authorize(
        com.github.alvin_nt.ChatService.RequestAuthorize request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_AUTHORIZE, getCallOptions()), request);
    }

    @java.lang.Override
    public com.google.common.util.concurrent.ListenableFuture<com.github.alvin_nt.ChatService.None> sendCommand(
        com.github.alvin_nt.ChatService.Command request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_SEND_COMMAND, getCallOptions()), request);
    }
  }

  public static io.grpc.ServerServiceDefinition bindService(
      final SimpleChat serviceImpl) {
    return io.grpc.ServerServiceDefinition.builder(SERVICE_NAME)
      .addMethod(
        METHOD_AUTHORIZE,
        asyncUnaryCall(
          new io.grpc.stub.ServerCalls.UnaryMethod<
              com.github.alvin_nt.ChatService.RequestAuthorize,
              com.github.alvin_nt.ChatService.ResponseAuthorize>() {
            @java.lang.Override
            public void invoke(
                com.github.alvin_nt.ChatService.RequestAuthorize request,
                io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.ResponseAuthorize> responseObserver) {
              serviceImpl.authorize(request, responseObserver);
            }
          }))
      .addMethod(
        METHOD_CONNECT,
        asyncServerStreamingCall(
          new io.grpc.stub.ServerCalls.ServerStreamingMethod<
              com.github.alvin_nt.ChatService.RequestConnect,
              com.github.alvin_nt.ChatService.Event>() {
            @java.lang.Override
            public void invoke(
                com.github.alvin_nt.ChatService.RequestConnect request,
                io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.Event> responseObserver) {
              serviceImpl.connect(request, responseObserver);
            }
          }))
      .addMethod(
        METHOD_SEND_COMMAND,
        asyncUnaryCall(
          new io.grpc.stub.ServerCalls.UnaryMethod<
              com.github.alvin_nt.ChatService.Command,
              com.github.alvin_nt.ChatService.None>() {
            @java.lang.Override
            public void invoke(
                com.github.alvin_nt.ChatService.Command request,
                io.grpc.stub.StreamObserver<com.github.alvin_nt.ChatService.None> responseObserver) {
              serviceImpl.sendCommand(request, responseObserver);
            }
          })).build();
  }
}
