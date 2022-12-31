import React from 'react'
import {createRoot} from 'react-dom/client'
import App from './App'
import { ChakraProvider } from '@chakra-ui/react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { RecoilRoot } from 'recoil'

const container = document.getElementById('root')

const root = createRoot(container!)

const queryClient = new QueryClient()

root.render(
    <React.StrictMode>
        <RecoilRoot>
            <QueryClientProvider client={queryClient}>
                <ChakraProvider>
                    <App/>
                </ChakraProvider>
            </QueryClientProvider>
        </RecoilRoot>
        
    </React.StrictMode>
)
