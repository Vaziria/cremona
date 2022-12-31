import { Flex, FormControl, FormLabel, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, Spacer, useDisclosure } from "@chakra-ui/react"
import { Search, AddUserButton, Button } from "@chatscope/chat-ui-kit-react"
import { useState } from "react"
import { useMutation, useQueryClient } from "react-query"
import { useRecoilState } from "recoil"
import { AddAccount } from "../wailsjs/go/backend/App"
import { searchState } from "./state/search"


function SearchSidebar(){

    const queryClient = useQueryClient()

    const { isOpen, onOpen, onClose } = useDisclosure()
    const [ name, setName ] = useState<string>('')
    const [ search, setSearch ] = useRecoilState(searchState)

    const addMutation = useMutation(AddAccount, {
        onSuccess: () => {
            
            onClose()
            setName('')
            queryClient.invalidateQueries('list_conversation')
        }
    })

    function add(){
        addMutation.mutate(name)
    }

    const changeName = (e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value)

    return <>
        <Flex>
            <Search placeholder="Search..." value={search} onChange={setSearch}/>
            <Spacer />
            <AddUserButton onClick={onOpen} />
        </Flex>
    
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <ModalOverlay />
            <ModalContent>
                <ModalCloseButton />
                <ModalBody pb={6}>
                    <FormControl mt={4}>
                        <FormLabel>Name Account</FormLabel>
                        <Input placeholder='Name Account' value={name} onChange={changeName}/>
                    </FormControl>
                </ModalBody>

                <ModalFooter>
                    <Button onClick={add}>
                        Save
                    </Button>
                    <Button onClick={onClose}>Cancel</Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    </>

}


export default SearchSidebar