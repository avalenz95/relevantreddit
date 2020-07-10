import React, { useState } from 'react'
import './Card.css'
import { useDispatch } from 'react-redux'
import { addKeywordToSub } from '../../actions'

// Displays a single subreddit along with associated keywords
function Card(props) {
    const { subName, username, keywords, banner } = props
    const [word, setWord] = useState("")
    const dispatch = useDispatch()
    console.log(banner)
    let tags = []
    // TODO: Figure out what's going on here.
    tags = Object.entries(keywords).map(([_, word],index) => {
        return (
            <button key={index}>{word}</button>
        )
    })

    return (
        <div className="card">
                <div className="container">
                    
                    <h4><b>{subName}</b></h4>
                    {tags}
    
                    <form onSubmit={e => {
                        e.preventDefault()
                        dispatch(addKeywordToSub(subName, username, word))
                        setWord("")
                    }}>
   

                        <input 
                            name="word" 
                            value={word}
                            onChange={e => setWord(e.target.value)}
                        />
                        <button name="submit" type="submit">Submit</button>
                    </form>
    
                </div>
            </div>
    )
}

export default Card