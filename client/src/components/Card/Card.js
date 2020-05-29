import React, {useState, useEffect} from 'react'
import axios from "axios"
import './Card.css'

function Card(props) {
    // eslint-disable-next-line no-unused-vars
    const {subName, imgUrl, keywords, endpoint, userName} = props

    const [word, setWord] = useState("")
    const [tags, setTags] = useState([])

    let style = {
        container: {
        }
    }

    //life cycle hooks, works like componentdidmount
    useEffect(() => {
        let words = []
        for (var i = 0; i < keywords.length; i++) {
            words.push(
                <button key={`${keywords[i]}`}>{keywords[i]}</button>
            )
        }

        setTags(words)
        
        }, [keywords])
    
    //call api with keyword and users name along with the subreddit
    function handleSubmit(event) {
        event.preventDefault()
        axios.get(`${endpoint}/add/${subName}/${userName}/${word}`).then(response => {
            console.log(response)
        }).then(
            tags.push(<button key={`${word}`}>{word}</button>),
            setTags(tags),
            setWord("")
        )
    }

    return (
        //Card content
        <div className="card">
            <div className="container" style={style.container}>
                
                <h4><b>{subName}</b></h4>
                {tags}

                <form onSubmit={handleSubmit}>
                    <label> Add New Keyword: 
                        <input type="text" value={word} onChange={event => setWord(event.target.value)}/>
                    </label>
                    <input type="submit" value="Submit" />
                </form>

            </div>
        </div>
        
    )
}

export default Card