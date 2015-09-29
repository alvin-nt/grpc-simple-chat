# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ChatService.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "proto.None" do
  end
  add_message "proto.RequestAuthorize" do
    optional :name, :string, 1
  end
  add_message "proto.ResponseAuthorize" do
    optional :session_id, :uint32, 1
    optional :name, :string, 2
  end
  add_message "proto.RequestConnect" do
    optional :session_id, :uint32, 1
  end
  add_message "proto.Command" do
    optional :session_id, :uint32, 1
    oneof :command do
      optional :say, :message, 2, "proto.CommandSay"
      optional :nick, :message, 3, "proto.CommandChangeNick"
      optional :join, :message, 4, "proto.CommandJoinChannel"
      optional :leave, :message, 5, "proto.CommandLeaveChannel"
      optional :exit, :message, 6, "proto.CommandExit"
    end
  end
  add_message "proto.CommandSay" do
    optional :message, :string, 1
    optional :channel_name, :string, 2
  end
  add_message "proto.CommandChangeNick" do
    optional :name, :string, 1
  end
  add_message "proto.CommandJoinChannel" do
    optional :name, :string, 1
  end
  add_message "proto.CommandLeaveChannel" do
    optional :name, :string, 1
  end
  add_message "proto.CommandExit" do
  end
  add_message "proto.Event" do
    oneof :event do
      optional :none, :message, 1, "proto.EventNone"
      optional :join, :message, 2, "proto.EventJoin"
      optional :leave, :message, 3, "proto.EventLeave"
      optional :log, :message, 4, "proto.EventLog"
    end
  end
  add_message "proto.EventNone" do
  end
  add_message "proto.EventJoin" do
    optional :name, :string, 1
  end
  add_message "proto.EventLeave" do
    optional :name, :string, 1
  end
  add_message "proto.EventLog" do
    optional :timestamp, :string, 1
    optional :name, :string, 2
    optional :message, :string, 3
    optional :channel, :string, 4
  end
end

module Proto
  None = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.None").msgclass
  RequestAuthorize = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.RequestAuthorize").msgclass
  ResponseAuthorize = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.ResponseAuthorize").msgclass
  RequestConnect = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.RequestConnect").msgclass
  Command = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.Command").msgclass
  CommandSay = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.CommandSay").msgclass
  CommandChangeNick = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.CommandChangeNick").msgclass
  CommandJoinChannel = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.CommandJoinChannel").msgclass
  CommandLeaveChannel = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.CommandLeaveChannel").msgclass
  CommandExit = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.CommandExit").msgclass
  Event = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.Event").msgclass
  EventNone = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.EventNone").msgclass
  EventJoin = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.EventJoin").msgclass
  EventLeave = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.EventLeave").msgclass
  EventLog = Google::Protobuf::DescriptorPool.generated_pool.lookup("proto.EventLog").msgclass
end
