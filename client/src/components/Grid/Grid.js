import React from 'react'
import Card from '../Card/Card.js'
import './Grid.css'
import { useSelector } from 'react-redux'



function Grid(props) {
    const name = useSelector(state => state.name)
    const userData = useSelector(state => state.userdata)
    console.log(userData)
    const {endpoint} = props
    let cards = []
    if (userData === null) {
        return null

    } else {
        // eslint-disable-next-line array-callback-return
        cards = Object.entries(userData.subreddits).map(([subName, keywords], index) => {
            return (
                //pass image as prop to card along with subreddits ect.
                <Card 
                    key={index} 
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