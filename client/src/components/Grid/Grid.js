import React from 'react'
import axios from 'axios'
import Card from '../Card/Card.js'
import './Grid.css'



function Grid(props) {
    const {endpoint, subreddits} = props
    const cards = []

    if (subreddits) {
        // eslint-disable-next-line array-callback-return
        Object.entries(subreddits).map(([subName, keywords], index) => {
            //BUILD CARD HERE CARD.JS image, content exect
            let imgUrl = ""

            //Retrieve an image
            axios.get(endpoint + "/img/" + subName).then((response) =>{
                imgUrl = response.data
            });

            cards.push(
                //pass image as prop to card along with subreddits ect.
                <Card 
                    key={index} 
                    subName={subName} 
                    keywords={keywords} 
                    imgUrl={imgUrl} 
                />
            )
        })
    }


    return (
        <div className="grid">
            {cards}
        </div>
    )
}

export default Grid