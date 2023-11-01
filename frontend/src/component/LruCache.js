import React,{useState} from 'react'
import axios from 'axios'

function LruCache() {

    const [key,setKey]=useState('')
    const [value,setValue]=useState('')
    const [expiration,setExpiration]=useState(5)
    const [cachedValue,setCachedValue]=useState('')
    const [error,setError]=useState(null)


    const handleGet=()=>{
        axios.get(`/get/${key}`).then((response)=>{
            setCachedValue(response.data)
            setError(null)
        }).catch((error)=>{
            console.log(error)
            setCachedValue('')
            setError('key not found in cache')
        })
    }

    const handleSet=()=>{
        axios.post(`/set`,{
            key:key,
            value:value,
            expiration:expiration,
        }).then(()=>{
            setCachedValue('')
            setError(null)
        }).catch((error)=>{
            console.log(error)
            setCachedValue('')
            setError('Failed to set value in cache')
        })
    }


  return (
    <div>
        <h1>LRU Cache</h1>
        <div>
                <label>Key:</label>
                <input
                    type="text"
                    value={key}
                    onChange={(e) => setKey(e.target.value)}
                />
            </div>
            <div>
                <label>Value:</label>
                <input
                    type="text"
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                />
            </div>
            <div>
                <label>Expiration (seconds):</label>
                <input
                    type="number"
                    value={expiration}
                    onChange={(e) => setExpiration(e.target.value)}
                />
            </div>
            <div>
                <button onClick={handleSet}>Set Value</button>
                <button onClick={handleGet}>Get Value</button>
            </div>
            {error && <div className="error">{error}</div>}
            {cachedValue && <div>Cached Value: {cachedValue}</div>}
    </div>
  )
}

export default LruCache