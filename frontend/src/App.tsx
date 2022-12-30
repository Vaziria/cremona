import {useEffect, useState} from 'react';
import { AddAccount, DeleteAccount, ListAccount } from '../wailsjs/go/backend/App'
import { backend } from '../wailsjs/go/models';

function App() {
    const [name, setName] = useState('');
    const [ listAkuns, setAkuns ] = useState<backend.Account[]>([])
    const updateName = (e: any) => setName(e.target.value);

    function deleteAccount(name: string) {
        DeleteAccount(name).then(() => {
            ListAccount().then(result => setAkuns(result))
        })
    }

    function add() {
        AddAccount(name).then(() => {
            ListAccount().then(result => setAkuns(result))
        })
    }

    useEffect(() => {
        ListAccount().then(result => setAkuns(result))
    }, [setAkuns]);


    return (
        <div id="App">
            
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={add}>Add</button>
            </div>
            <h3>List Akun</h3>
            {
                listAkuns.map(akun => {
                    return <div key={akun.name}>
                        <div>{akun.name} ( {akun.chat_data.chat_unread} )</div> <button onClick={() => deleteAccount(akun.name)}>delete</button> <button>refresh</button>
                    </div> 
                    
                })
            }

        </div>
    )
}

export default App
 