import React, { useState } from 'react'
import './Card.css'
import { useDispatch } from 'react-redux'
import { addKeywordToSub } from '../../actions'


function Card(props) {
    const { subName, username, keywords } = props
    const [word, setWord] = useState("")
    const dispatch = useDispatch()

    let tags = []
    tags = Object.entries(keywords).map((word, index) => {
        return (
            <button>{word}</button>
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
                    }}>
                        <label> Add New Keyword: 
                            <input type="text" value={word} />
                        </label>
                        <input type="submit" value="Submit" onChange={event => setWord(event.target.value)}/>
                    </form>
    
                </div>
            </div>
    )
}

export default Card