import React, {useState} from 'react'
import axios from "axios"
import './Card.css'

function Card(props) {
    // eslint-disable-next-line no-unused-vars
    const {subName, imgUrl, keywords, endpoint, userName} = props

    const [word, setWord] = useState("")


    let style = {
        container: {
        }
    }

    //List of keyword buttons
    let tags = []
    //Add button for each keyword
    for (var i = 0; i < keywords.length; i++) {
        tags.push(
            <button key={`${keywords[i]}`}>{keywords[i]}</button>
        )
    }

    //call api with keyword and users name along with the subreddit
    function handleSubmit(event) {
        event.preventDefault()
        axios.get(`${endpoint}/add/${subName}/${userName}/${word}`).then(response => {
            console.log(response);
        })
    }

    return (
        //Card content
        <div className="card">
            <div className="container" style={style.container}>
                
                <h4><b>{subName}</b></h4>

                <div>{tags}</div>

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