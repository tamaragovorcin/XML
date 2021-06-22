import { ChatEngine } from 'react-chat-engine';
import React from "react";
import ChatFeed from './ChatFeed';
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader


class ChatEnginePreview extends React.Component {

render(){
    return (
        <React.Fragment>
			<ChatEngine
height="100vh"
projectID="ed28d2aa-84ca-4833-9e0f-bf665e72e406"
userName="mladenka"
userSecret="maja1234"
renderChatFeed={(chatAppProps) => <ChatFeed {...chatAppProps} />}
onNewMessage={() => new Audio('https://chat-engine-assets.s3.amazonaws.com/click.mp3').play()}
/>
        
        </React.Fragment>

    );

	}
}
export default ChatEnginePreview;