import React, { useEffect } from 'react'
import Card from '../Card/Card.js'
import './Grid.css'
import { loadUserData } from "../../actions/index.js"
import { useDispatch, useSelector } from 'react-redux'



function Grid(props) {
    const name = useSelector(state => state.name)
    let cards = []
    const userData = useSelector(state => state.userdata)
    // Attempt to load username on component mount
    console.log(userData)
    const {endpoint} = props
    if (userData === null) {
        return null

    } else {
                // eslint-disable-next-line array-callback-return
                cards = Object.entries(userData.subreddits).map(([subName, keywords], index) => {
                    return (
                        //pass image as prop to card along with subreddits ect.
                        <Card 
                            key={index} 
                            userName={name}
                            subName={subName} 
                            keywords={keywords} 
                            endpoint={endpoint}
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