import React, { useState } from 'react'
import './Card.css'
import { useDispatch } from 'react-redux'
import { addKeywordToSub } from '../../actions'
import Tag from '../Tag/Tag.js'

// Displays a single subreddit along with associated keywords
function Card(props) {
    const { subName, username, keywords, banner } = props
    const [word, setWord] = useState("")
    const dispatch = useDispatch()

    let tags = []
    // TODO: Figure out what's going on here.
    tags = Object.entries(keywords).map(([_, word],index) => {
        return (
            <Tag word={word}/>
        )
    })

    return (
        <div className="card" style={{backgroundImage: `url(${banner})`}}>
                <div className="container">
                    <div className="subName">
                    {subName}
                    </div>
                    
                    <div className="tagGrid">
                        {tags}
                    </div>
    
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