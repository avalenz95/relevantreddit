import React from 'react'
import './Card.css'

function Card(props) {
    const {subName, imgUrl, keywords} = props
    let style = {
        card: {
            backgroundImage: `url(${imgUrl})`,
        }
    }

    let tags = []

    for (var i = 0; i < keywords.length; i++) {
        tags.push(
            <button>{keywords[i]}</button>
        )
    }

    return (
        //Card content
        <div className="card" style={style.card}>
            <div className="container">
                
                <h4><b>{subName}</b></h4>

                <div>
                    {tags}
                </div>

                <button type="input">Add Keyword</button>
            </div>
        </div>
        
    )
}

export default Card