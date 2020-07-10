import React from 'react'
import Card from '../Card/Card.js'
import './Grid.css'
import { useSelector } from 'react-redux'



function Grid() {
    const name = useSelector(state => state.name)
    const userData = useSelector(state => state.userdata)
    const banners = useSelector(state => state.banners)
    console.log(userData)

    let cards = []
    if (userData === null) {
        return null

    } else {
        // eslint-disable-next-line array-callback-return
        cards = Object.entries(userData.subreddits).map(([subName, keywords], index) => {
            const banner = banners[subName]
            return (
                //pass image as prop to card along with subreddits ect.
                <Card 
                    key={index} 
                    username={name}
                    subName={subName} 
                    keywords={keywords} 
                    banner={banner}
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