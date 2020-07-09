import React, { useEffect } from 'react'
import Card from '../Card/Card.js'
import './Grid.css'
import { loadSubreddits } from "../../actions/index.js"
import { useDispatch, useSelector } from 'react-redux'



function Grid(props) {
    const name = useSelector(state => state.name)
    const dispatch = useDispatch() 
    let cards = []
    // Attempt to load username on component mount
    useEffect(() => {
        dispatch(loadSubreddits(name))
    })
    const subreddits = useSelector(state => state.subreddits)
    console.log(subreddits)
    const {endpoint} = props
    if (subreddits === null) {
        return null

    } else {
                // eslint-disable-next-line array-callback-return
                cards = Object.entries(subreddits.subreddits).map(([subName, keywords], index) => {
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