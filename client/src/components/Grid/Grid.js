import React from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'

function Grid(props) {
    const {endpoint, subreddits} = props

    Object.entries(subreddits).map(([key, values], index) => {
        //BUILD CARD HERE CARD.JS image, content exect

        axios.get(endpoint + "/img/" + key).then((response) =>{
            //response.data is the image url img = response.data
        })

        //pass image as prop to card along with subreddits ect.

    })
}