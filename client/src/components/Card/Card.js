import React from 'react'
import './Card.css'

function Card(props) {
    const {name, imgUrl, keywords} = props

    let tags = []

    for (var i = 0; i < keywords.length; i++) {
        tags.push(
            <button>{keywords[i]}</button>
        )
    }

    return (
        //Card content
        <div className="card">
            <img src={imgUrl} alt="img" />
            <div className="container">
                
                <h4><b>{name}</b></h4>

                <div>
                    {tags}
                </div>

                <button type="input">Add Keyword</button>
            </div>
        </div>
        
    )
}

export default Card