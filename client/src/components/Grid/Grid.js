import React from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'

function Grid(props) {
    const {endpoint, subreddits} = props

    Object.entries(subreddits).map(([name, keywords], index) => {
        //BUILD CARD HERE CARD.JS image, content exect
        let imgUrl = ""

        axios.get(endpoint + "/img/" + name).then((response) =>{
            imgUrl = response.data
        })

        <Card
            imgUrl={imgUrl},
            
        />
        //pass image as prop to card along with subreddits ect.

    })



    return (
        <div className="grid">

        </div>
    )
}

export default Grid