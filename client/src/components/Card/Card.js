import React, {useState} from 'react'
import './Card.css'

function Card(props) {
    const {subName, imgUrl, keywords} = props

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
            <button>{keywords[i]}</button>
        )
    }

    function handleSubmit(event) {
        event.preventDefault()
    }

    return (
        //Card content
        <div className="card">
            <div className="container" style={style.container}>
                
                <h4><b>{subName}</b></h4>

                <div>
                    {tags}
                </div>
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