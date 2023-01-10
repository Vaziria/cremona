import { ConversationList, Conversation, Avatar, EllipsisButton, Button } from "@chatscope/chat-ui-kit-react"
import { ListAccount, DeleteAccount } from "../wailsjs/go/backend/App"
import {
    useMutation,
    useQuery, useQueryClient,
} from 'react-query'
import { Menu, MenuButton, MenuList, MenuItem } from "@chakra-ui/react"
import { useEffect, useState } from "react"
import { backend } from "../wailsjs/go/models"
import { useRecoilValue } from "recoil"
import { searchState } from "./state/search"

function MenuAccount({account}: { account: backend.Account }){
    const queryClient = useQueryClient()

    const deleteMutation = useMutation(DeleteAccount, {
        onSuccess: () => {
            queryClient.invalidateQueries("list_conversation")
        }
    })

    function deleteAccount(){
        deleteMutation.mutate(account.name)
    }

    return <MenuList>
        <MenuItem>Reconnect</MenuItem>
        <MenuItem color="red" onClick={deleteAccount}>Delete</MenuItem>
    </MenuList>
}


function ConversationAccount() {
    const search = useRecoilValue(searchState)
    const [ activeAccount, setActiveAccount ] = useState<backend.Account | null>()
    const query = useQuery("list_conversation", ListAccount)

    useEffect(() => {
        (window as any).runtime.EventsOn("message", (data: any) => {
            console.log(data)
        })
    }, [])

    return <ConversationList>
        {
            !query.isLoading && query.data?.filter(account => {
                
                if(search === ''){
                    return true
                }

                return account.name.includes(search)
            
            }).map(account => {
                let active = false

                if(activeAccount){
                    if(activeAccount.name == account.name){
                        active = true
                    }
                }

                return <Conversation 
                    key={account.name} 
                    name={account.name} 
                    lastSenderName="Lilly" 
                    info="Yes i can do it for you" 
                    active={active}
                    onClick={() => setActiveAccount(account)}
                >
                    <Avatar name="Lilly" status="available" src='https://bit.ly/kent-c-dodds' />
                    
                    <Conversation.Operations>
                        
                        <Menu>
                            <MenuButton as={EllipsisButton} orientation="vertical" />
                            <MenuAccount account={account} />
                            
                        </Menu>
                    </Conversation.Operations>

                    
                </Conversation>
            }) 
        }   
                                                                
    </ConversationList>
}


export default ConversationAccount