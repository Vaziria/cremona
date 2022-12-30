import "@chatscope/chat-ui-kit-styles/dist/default/styles.min.css";
import {
    MainContainer,
    ChatContainer,
    MessageList,
    Message,
    MessageInput,
    Avatar,
    Conversation,
    ConversationHeader,
    ConversationList,
    ExpansionPanel,
    InfoButton,
    MessageSeparator,
    Search,
    Sidebar,
    TypingIndicator,
    VideoCallButton,
    VoiceCallButton,
  } from "@chatscope/chat-ui-kit-react";
import {useState} from 'react';

function App() {
    const [messageInputValue, setMessageInputValue] = useState("");
    return <div style={{
        height: "100%",
        position: "relative"
    }}>
        <MainContainer responsive>                                   
            <Sidebar position="left" scrollable={false}>
                <Search placeholder="Search..." />
                <ConversationList>                                                     
                <Conversation name="Lilly" lastSenderName="Lilly" info="Yes i can do it for you">
                    <Avatar name="Lilly" status="available" />
                </Conversation>
                
                <Conversation name="Joe" lastSenderName="Joe" info="Yes i can do it for you">
                    <Avatar name="Joe" status="dnd" />
                </Conversation>
                
                <Conversation name="Emily" lastSenderName="Emily" info="Yes i can do it for you" unreadCnt={3}>
                    <Avatar name="Emily" status="available" />
                </Conversation>
                
                <Conversation name="Kai" lastSenderName="Kai" info="Yes i can do it for you" unreadDot>
                    <Avatar name="Kai" status="unavailable" />
                </Conversation>
                            
                <Conversation name="Akane" lastSenderName="Akane" info="Yes i can do it for you">
                    <Avatar name="Akane" status="eager" />
                </Conversation>
                                    
                <Conversation name="Eliot" lastSenderName="Eliot" info="Yes i can do it for you">
                    <Avatar name="Eliot" status="away" />
                </Conversation>
                                                    
                <Conversation name="Zoe" lastSenderName="Zoe" info="Yes i can do it for you" active>
                    <Avatar name="Zoe" status="dnd" />
                </Conversation>
                
                <Conversation name="Patrik" lastSenderName="Patrik" info="Yes i can do it for you">
                    <Avatar name="Patrik" status="invisible" />
                </Conversation>
                                                                            
                </ConversationList>
            </Sidebar>
                  
            <ChatContainer>
                <ConversationHeader>
                    <ConversationHeader.Back />
                    <Avatar name="Zoe" />
                    <ConversationHeader.Content userName="Zoe" info="Active 10 mins ago" />
                    <ConversationHeader.Actions>
                        <VoiceCallButton />
                        <VideoCallButton />
                        <InfoButton />
                    </ConversationHeader.Actions>          
                </ConversationHeader>
                <MessageList typingIndicator={<TypingIndicator content="Zoe is typing" />}>
                    
                <MessageSeparator>Saturday, 30 November 2019</MessageSeparator>
                    
                <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "single"
                }}>
                    <Avatar name="Zoe" />
                </Message>
                    
                <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Patrik",
                    direction: "outgoing",
                    position: "single"
                }} avatarSpacer />
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "first"
                }} avatarSpacer />
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "normal"
                }} avatarSpacer />
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "normal"
                }} avatarSpacer />
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "last"
                }}>
                    <Avatar name="Zoe" />
                </Message>
                    
                <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Patrik",
                    direction: "outgoing",
                    position: "first"
                }} />
                <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Patrik",
                    direction: "outgoing",
                    position: "normal"
                }} />
                <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Patrik",
                    direction: "outgoing",
                    position: "normal"
                }} />
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Patrik",
                    direction: "outgoing",
                    position: "last"
                }} />
                            
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "first"
                }} avatarSpacer />                                          
                            <Message model={{
                    message: "Hello my friend",
                    sentTime: "15 mins ago",
                    sender: "Zoe",
                    direction: "incoming",
                    position: "last"
                }}>
                      <Avatar name="Zoe" />
                    </Message>
                  </MessageList>
                  <MessageInput placeholder="Type message here" value={messageInputValue} onChange={val => setMessageInputValue(val)} onSend={() => setMessageInputValue("")} />
                </ChatContainer>
                          
              </MainContainer>
            </div>;
}

export default App
 