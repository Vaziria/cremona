import "@chatscope/chat-ui-kit-styles/dist/default/styles.min.css";
import {
    MainContainer,
    ChatContainer,
    MessageList,
    Message,
    MessageInput,
    Avatar,
    ConversationHeader,
    InfoButton,
    MessageSeparator,
    Search,
    Sidebar,
    TypingIndicator,
    VideoCallButton,
    VoiceCallButton,
  } from "@chatscope/chat-ui-kit-react";
import {useState} from 'react';
import ConversationAccount from "./ConversationAccount";
import SearchSidebar from "./SearchSidebar";

function App() {
    const [messageInputValue, setMessageInputValue] = useState("");
    return <div style={{
        overflow: "auto"
      }}>
    
        <MainContainer>                  

            <Sidebar position="left" scrollable={true} style={{height: "720px"}}>
                <SearchSidebar />
                <ConversationAccount />
            </Sidebar>
                  
            <ChatContainer>
                <ConversationHeader>
                    <ConversationHeader.Back />
                    <Avatar name="Zoe" src='https://bit.ly/kent-c-dodds' />
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
                    <Avatar name="Zoe" src='https://bit.ly/kent-c-dodds' />
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
                    <Avatar name="Zoe" src='https://bit.ly/kent-c-dodds' />
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
                      <Avatar name="Zoe" src='https://bit.ly/kent-c-dodds' />
                    </Message>
                  </MessageList>
                  <MessageInput placeholder="Type message here" value={messageInputValue} onChange={val => setMessageInputValue(val)} onSend={() => setMessageInputValue("")} />
                </ChatContainer>
                          
              </MainContainer>
            </div>
}

export default App
 